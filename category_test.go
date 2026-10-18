// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"reflect"
	"slices"
	"testing"
)

func TestRegistryCategoryFilters(t *testing.T) {
	t.Parallel()

	blueMuted := &Tint{ID: "m", Hue: HueBlue, Tone: ToneMuted}
	blue := &Tint{ID: "a", Hue: HueBlue}
	redRetro := &Tint{ID: "z", Hue: HueRed, Tone: ToneRetro}
	unset := &Tint{ID: "b"}

	reg := NewRegistry(nil, blueMuted, unset, redRetro, blue)

	tests := []struct {
		name string
		got  []*Tint
		want []*Tint
	}{
		{
			name: "hue",
			got:  reg.TintsByHue(HueBlue),
			want: []*Tint{blue, blueMuted},
		},
		{
			name: "empty hue",
			got:  reg.TintsByHue(""),
			want: []*Tint{unset},
		},
		{
			name: "tone",
			got:  reg.TintsByTone(ToneRetro),
			want: []*Tint{redRetro},
		},
		{
			name: "empty tone",
			got:  reg.TintsByTone(""),
			want: []*Tint{blue, unset},
		},
		{
			name: "hue and tone",
			got:  reg.TintsByHueAndTone(HueBlue, ToneMuted),
			want: []*Tint{blueMuted},
		},
		{
			name: "hue with empty tone",
			got:  reg.TintsByHueAndTone(HueBlue, ""),
			want: []*Tint{blue},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("got %#v, want %#v", ids(tt.got), ids(tt.want))
			}
		})
	}
}

func TestDefaultTintsFiltersCoverCatalog(t *testing.T) {
	t.Parallel()

	all := DefaultTints()
	byHue := 0
	for _, hue := range append(AllHues(), "") {
		for _, tint := range TintsByHue(hue) {
			if tint.Hue != hue {
				t.Errorf("TintsByHue(%q) included %s with hue %q", hue, tint.ID, tint.Hue)
			}
			byHue++
		}
	}
	if byHue != len(all) {
		t.Errorf("hue filters returned %d tints, want %d", byHue, len(all))
	}

	byTone := 0
	for _, tone := range append(AllTones(), "") {
		for _, tint := range TintsByTone(tone) {
			if tint.Tone != tone {
				t.Errorf("TintsByTone(%q) included %s with tone %q", tone, tint.ID, tint.Tone)
			}
			byTone++
		}
	}
	if byTone != len(all) {
		t.Errorf("tone filters returned %d tints, want %d", byTone, len(all))
	}
}

func TestBuiltinCategories(t *testing.T) {
	t.Parallel()

	for _, tint := range DefaultTints() {
		if !slices.Contains(AllHues(), tint.Hue) {
			t.Errorf("%s hue %q is not a known hue", tint.ID, tint.Hue)
		}
		if tint.Tone != "" && !slices.Contains(AllTones(), tint.Tone) {
			t.Errorf("%s tone %q is not a known tone", tint.ID, tint.Tone)
		}
	}
}

func TestTintsByHueIgnoresDefaultRegistry(t *testing.T) {
	prev := DefaultRegistry
	t.Cleanup(func() { DefaultRegistry = prev })
	DefaultRegistry = nil

	got := TintsByHue(HueBlue)
	if len(got) == 0 {
		t.Fatal("TintsByHue returned nothing")
	}
	for _, tint := range got {
		if tint.Hue != HueBlue {
			t.Errorf("TintsByHue() included %s with hue %q", tint.ID, tint.Hue)
		}
	}
	for _, tint := range TintsByTone(ToneMuted) {
		if tint.Tone != ToneMuted {
			t.Errorf("TintsByTone() included %s with tone %q", tint.ID, tint.Tone)
		}
	}
	for _, tint := range TintsByHueAndTone(HueBlue, ToneMuted) {
		if tint.Hue != HueBlue || tint.Tone != ToneMuted {
			t.Errorf("TintsByHueAndTone() included %s hue %q tone %q", tint.ID, tint.Hue, tint.Tone)
		}
	}
}

func ids(tints []*Tint) []string {
	out := make([]string, len(tints))
	for i, tint := range tints {
		out[i] = tint.ID
	}
	return out
}
