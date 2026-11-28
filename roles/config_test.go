// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"encoding/json"
	"errors"
	"slices"
	"testing"

	"github.com/charmbracelet/colorprofile"
	tint "github.com/lrstanley/bubbletint/v2"
)

func TestParseAccent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want string
		slot     tint.Slot
		err      bool
	}{
		{in: "", want: "auto"},
		{in: "auto", want: "auto"},
		{in: "AUTO", want: "auto"},
		{in: "blue", want: "blue", slot: tint.SlotBlue},
		{in: "purple", want: "purple", slot: tint.SlotMagenta},
		{in: "magenta", want: "purple", slot: tint.SlotMagenta},
		{in: "bright_magenta", want: "bright_purple", slot: tint.SlotMagenta},
		{in: "Bright_Cyan", want: "bright_cyan", slot: tint.SlotCyan},
		{in: "#f80", want: "#ff8800"},
		{in: "#FF8800", want: "#ff8800"},
		{in: "white", err: true},
		{in: "bright_black", err: true},
		{in: "#ff88", err: true},
		{in: "teal", err: true},
	}
	tn := tint.TintTokyoNight
	for _, tt := range tests {
		got, err := ParseAccent(tt.in)
		if tt.err {
			if !errors.Is(err, ErrInvalidAccent) {
				t.Errorf("ParseAccent(%q) error = %v, want ErrInvalidAccent", tt.in, err)
			}
			continue
		}
		if err != nil || got.String() != tt.want {
			t.Errorf("ParseAccent(%q) = %q, %v, want %q", tt.in, got, err, tt.want)
			continue
		}
		p := Derive(tn, got.Option())
		switch {
		case got.Auto():
			if p.accentSlot != Derive(tn).accentSlot {
				t.Errorf("%q: auto changed the accent", tt.in)
			}
		case tt.want[0] == '#':
			if p.Color(Accent).Distance(*tint.FromHex(tt.want)) > 0.05 {
				t.Errorf("%q: accent %s", tt.in, p.Color(Accent).Hex())
			}
		case p.accentSlot != tt.slot:
			t.Errorf("%q: accent slot %v, want %v", tt.in, p.accentSlot, tt.slot)
		}
	}
}

func TestConfigJSON(t *testing.T) {
	t.Parallel()
	var cfg struct {
		Accent AccentChoice `json:"accent"`
		Colors ColorLimit   `json:"colors"`
	}
	if err := json.Unmarshal([]byte(`{"accent": "purple", "colors": "256"}`), &cfg); err != nil {
		t.Fatal(err)
	}
	if cfg.Accent.String() != "purple" || cfg.Colors != Colors256 {
		t.Errorf("decoded %v, %v", cfg.Accent, cfg.Colors)
	}
	b, err := json.Marshal(cfg)
	if err != nil || string(b) != `{"accent":"purple","colors":"256"}` {
		t.Errorf("encoded %s, %v", b, err)
	}
	if bad := json.Unmarshal([]byte(`{"colors": "lots"}`), &cfg); !errors.Is(bad, ErrInvalidColorLimit) {
		t.Errorf("bad colors: %v", bad)
	}
	for _, schema := range []string{AccentSchema, ColorLimitSchema} {
		if !json.Valid([]byte(schema)) {
			t.Errorf("schema isn't valid JSON: %s", schema)
		}
	}
}

func TestColorLimitCap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		limit    ColorLimit
		detected colorprofile.Profile
		want     colorprofile.Profile
	}{
		{limit: ColorsAuto, detected: colorprofile.Unknown, want: colorprofile.TrueColor},
		{limit: ColorsAuto, detected: colorprofile.ANSI, want: colorprofile.ANSI},
		{limit: Colors256, detected: colorprofile.TrueColor, want: colorprofile.ANSI256},
		{limit: Colors256, detected: colorprofile.ANSI, want: colorprofile.ANSI},
		{limit: Colors16, detected: colorprofile.Unknown, want: colorprofile.ANSI},
		{limit: ColorsNone, detected: colorprofile.TrueColor, want: colorprofile.ASCII},
		{limit: Colors256, detected: colorprofile.NoTTY, want: colorprofile.ASCII},
	}
	for _, tt := range tests {
		if got := tt.limit.Cap(tt.detected); got != tt.want {
			t.Errorf("%v.Cap(%v) = %v, want %v", tt.limit, tt.detected, got, tt.want)
		}
	}
	for _, s := range []string{"", "auto", "256", "16", "None"} {
		if _, err := ParseColorLimit(s); err != nil {
			t.Errorf("ParseColorLimit(%q): %v", s, err)
		}
	}
}

// The choices to offer each read back as themselves, and are in the schemas.
func TestChoices(t *testing.T) {
	t.Parallel()
	var accent struct {
		AnyOf []struct {
			Enum []string `json:"enum"`
		} `json:"anyOf"`
	}
	if err := json.Unmarshal([]byte(AccentSchema), &accent); err != nil {
		t.Fatal(err)
	}
	accents := Accents()
	if len(accents) != 13 || !accents[0].Auto() {
		t.Errorf("Accents() = %v", accents)
	}
	for i, a := range accents {
		if s, ok := a.Slot(); ok != (i > 0) || ok && s.IsBright() != (i > 6) {
			t.Errorf("accent %q: Slot() = %v, %v", a, s, ok)
		}
		back, err := ParseAccent(a.String())
		if err != nil || back != a || !slices.Contains(accent.AnyOf[0].Enum, a.String()) {
			t.Errorf("accent %q reads back as %q, %v, or isn't in the schema", a, back, err)
		}
	}
	hex, _ := ParseAccent("#f80")
	if _, ok := hex.Slot(); ok {
		t.Error("a fixed color has a slot")
	}
	var limit struct {
		Enum []string `json:"enum"`
	}
	if err := json.Unmarshal([]byte(ColorLimitSchema), &limit); err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, l := range ColorLimits() {
		if back, err := ParseColorLimit(l.String()); err != nil || back != l {
			t.Errorf("limit %q reads back as %q, %v", l, back, err)
		}
		names = append(names, l.String())
	}
	if !slices.Equal(names, limit.Enum) {
		t.Errorf("ColorLimits() = %v, want the schema's %v", names, limit.Enum)
	}
}
