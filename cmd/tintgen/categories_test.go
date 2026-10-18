// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"encoding/json/v2"
	"maps"
	"math"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestBucketFromChoice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		selected      string
		probabilities map[string]float64
		want          string
		ok            bool
	}{
		{
			name:          "real tone at threshold",
			selected:      "vivid",
			probabilities: map[string]float64{"vivid": 0.5},
			want:          "vivid",
			ok:            true,
		},
		{
			name:          "real tone below threshold",
			selected:      "vivid",
			probabilities: map[string]float64{"vivid": 0.49},
			want:          "uncategorized",
			ok:            true,
		},
		{
			name:          "uncategorized wins",
			selected:      "uncategorized",
			probabilities: map[string]float64{"uncategorized": 0.9},
			want:          "uncategorized",
			ok:            true,
		},
		{
			name:          "uncategorized below threshold",
			selected:      "uncategorized",
			probabilities: map[string]float64{"uncategorized": 0.2},
			want:          "uncategorized",
			ok:            true,
		},
		{
			name:          "unknown option",
			selected:      "glossy",
			probabilities: map[string]float64{"glossy": 0.9},
		},
		{
			name:     "missing probability",
			selected: "muted",
		},
		{
			name:          "not a probability",
			selected:      "muted",
			probabilities: map[string]float64{"muted": math.NaN()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, ok := bucketFromChoice(tt.selected, tt.probabilities)
			if ok != tt.ok || got != tt.want {
				t.Errorf("bucketFromChoice() = %q, %v; want %q, %v", got, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestResolveTones(t *testing.T) {
	t.Parallel()

	cached := map[string]string{"Kept": "muted", "Gone": "retro"}

	tests := []struct {
		name       string
		current    []string
		reclassify bool
		asked      bool
		outcomes   map[string]toneOutcome
		want       map[string]string
		failed     []string
	}{
		{
			name:    "no token keeps cache and drops removed themes",
			current: []string{"Kept", "New"},
			want:    map[string]string{"Kept": "muted"},
		},
		{
			name:    "classifies only uncached themes",
			current: []string{"Kept", "New", "Bad"},
			asked:   true,
			outcomes: map[string]toneOutcome{
				"New": {bucket: "neon", ok: true},
				"Bad": {},
			},
			want:   map[string]string{"Kept": "muted", "New": "neon"},
			failed: []string{"Bad"},
		},
		{
			name:       "reclassify ignores cache",
			current:    []string{"Kept", "Bad"},
			reclassify: true,
			asked:      true,
			outcomes: map[string]toneOutcome{
				"Kept": {bucket: "vivid", ok: true},
				"Bad":  {},
			},
			want:   map[string]string{"Kept": "vivid"},
			failed: []string{"Bad"},
		},
		{
			name:    "uncategorized stays in the cache",
			current: []string{"New"},
			asked:   true,
			outcomes: map[string]toneOutcome{
				"New": {bucket: "uncategorized", ok: true},
			},
			want: map[string]string{"New": "uncategorized"},
		},
		{
			name:    "missing outcome is retried",
			current: []string{"New"},
			asked:   true,
			want:    map[string]string{},
			failed:  []string{"New"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, failed := resolveTones(tt.current, cached, tt.reclassify, tt.asked, tt.outcomes)
			if !maps.Equal(got, tt.want) {
				t.Errorf("kept = %#v, want %#v", got, tt.want)
			}
			if !slices.Equal(failed, tt.failed) {
				t.Errorf("failed = %#v, want %#v", failed, tt.failed)
			}
		})
	}
}

func TestGeneratedTonesOmitUncategorized(t *testing.T) {
	t.Parallel()

	got := generatedTones(map[string]string{
		"Muted": "muted",
		"None":  "uncategorized",
	})
	want := map[string]string{"Muted": "muted"}
	if !maps.Equal(got, want) {
		t.Errorf("generatedTones() = %#v, want %#v", got, want)
	}
}

func TestToneCacheRoundTrip(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), categoriesFileName)
	assignments := map[string]string{
		"Zed": "muted",
		"Ada": "muted",
		"Neo": "uncategorized",
	}
	if err := writeToneCache(path, assignments); err != nil {
		t.Fatalf("writeToneCache() error = %v", err)
	}

	got, err := loadToneCache(path)
	if err != nil {
		t.Fatalf("loadToneCache() error = %v", err)
	}
	if !maps.Equal(got, assignments) {
		t.Errorf("loadToneCache() = %#v, want %#v", got, assignments)
	}

	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read cache: %v", err)
	}
	var file toneFile
	if err = json.Unmarshal(body, &file); err != nil {
		t.Fatalf("parse cache: %v", err)
	}
	if !slices.Equal(file.Tone["muted"], []string{"Ada", "Zed"}) {
		t.Errorf("muted = %#v, want sorted names", file.Tone["muted"])
	}
	if file.Schema != categoriesSchema {
		t.Errorf("schema = %q", file.Schema)
	}
}

func TestLoadToneCacheMissing(t *testing.T) {
	t.Parallel()

	got, err := loadToneCache(filepath.Join(t.TempDir(), categoriesFileName))
	if err != nil {
		t.Fatalf("loadToneCache() error = %v", err)
	}
	if len(got) != 0 {
		t.Errorf("loadToneCache() = %#v, want empty", got)
	}
}

func TestLoadToneCacheRejectsUnknownBucket(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), categoriesFileName)
	writeFile(t, path, `{"$schema":"./categories.schema.json","tone":{"nope":[]}}`)
	if _, err := loadToneCache(path); err == nil {
		t.Fatal("loadToneCache() error = nil, want unknown field")
	}
}

func TestClassifyTonesWithoutToken(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), categoriesFileName)
	if err := writeToneCache(path, map[string]string{"Kept": "retro", "Gone": "neon"}); err != nil {
		t.Fatal(err)
	}

	got, err := classifyTones(t.Context(), []Tint{{Name: "Kept"}, {Name: "New"}}, path, "", false)
	if err != nil {
		t.Fatalf("classifyTones() error = %v", err)
	}
	want := map[string]string{"Kept": "retro"}
	if !maps.Equal(got, want) {
		t.Errorf("tones = %#v, want %#v", got, want)
	}

	cached, err := loadToneCache(path)
	if err != nil {
		t.Fatal(err)
	}
	if !maps.Equal(cached, want) {
		t.Errorf("cache = %#v, want Gone removed and New left out", cached)
	}
}

func TestClassifyTonesReclassifyRequiresToken(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), categoriesFileName)
	original := map[string]string{"Kept": "retro"}
	if err := writeToneCache(path, original); err != nil {
		t.Fatal(err)
	}

	_, err := classifyTones(t.Context(), []Tint{{Name: "Kept"}}, path, "  ", true)
	if err == nil {
		t.Fatal("classifyTones() error = nil, want missing token")
	}

	cached, loadErr := loadToneCache(path)
	if loadErr != nil {
		t.Fatal(loadErr)
	}
	if !maps.Equal(cached, original) {
		t.Errorf("cache = %#v, want original cache preserved", cached)
	}
}
