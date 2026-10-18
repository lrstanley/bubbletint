// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRegistryTemplateOmitsToneHelpers(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	err := registryTmpl.Execute(&buf, []TintTemplate{{
		StructName:     "Example",
		NameNormalized: "example",
	}})
	if err != nil {
		t.Fatal(err)
	}
	body := buf.String()
	if !strings.Contains(body, "TintExample") {
		t.Fatal("registry template did not emit the tint")
	}
	if strings.Contains(body, "DefaultMutedTints") {
		t.Fatal("registry template emitted a per-tone helper")
	}
}

func TestToneCacheMatchesRepo(t *testing.T) {
	t.Parallel()

	repo := filepath.Join("..", "..", categoriesFileName)
	want, err := os.ReadFile(repo)
	if err != nil {
		t.Fatal(err)
	}
	assignments, err := loadToneCache(repo)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), categoriesFileName)
	if err = writeToneCache(path, assignments); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Fatal("rewriting categories.json changed bytes")
	}
}

func TestHTMLTemplateCategoryFilters(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	err := tintHTMLTmpl.Execute(&buf, map[string]any{
		"Tints": []TintTemplate{{
			StructName:     "Example",
			NameNormalized: "example",
			Hue:            "mixed",
			Tone:           "high-contrast",
			Tint:           Tint{Name: "Example"},
		}},
		"Hues":  hueNames(),
		"Tones": append(toneNames(), toneUncategorized),
	})
	if err != nil {
		t.Fatal(err)
	}
	body := buf.String()
	for _, want := range []string{
		`id="filter-hue-mixed"`,
		`id="filter-tone-high-contrast"`,
		`id="filter-tone-uncategorized"`,
		`data-hue="mixed"`,
		`data-tone="high-contrast"`,
		`>High contrast<`,
		`<details class="nav-toc">`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("index html missing %s", want)
		}
	}
}
