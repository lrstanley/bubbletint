// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"encoding/json"
	"errors"
	"image/color"
	"slices"
	"testing"
)

func TestTintComplete(t *testing.T) {
	t.Parallel()
	red, brightBlue := FromHex("#ff0000"), FromHex("#8888ff")
	tn := &Tint{ID: "mine", Red: red, Blue: FromHex("#0000ff"), BrightBlue: brightBlue}
	tn.Complete()
	if tn.BrightRed != red || tn.BrightBlue != brightBlue || tn.BrightGreen != nil {
		t.Errorf("brights = %v, %v, %v", tn.BrightRed, tn.BrightBlue, tn.BrightGreen)
	}
	if tn.DisplayName != "mine" {
		t.Errorf("DisplayName = %q", tn.DisplayName)
	}
	named := &Tint{ID: "mine", DisplayName: "Mine"}
	named.Complete()
	if named.DisplayName != "Mine" {
		t.Errorf("Complete replaced DisplayName with %q", named.DisplayName)
	}
}

func TestTintMissing(t *testing.T) {
	t.Parallel()
	if got := TintTokyoNight.Missing(); len(got) != 0 {
		t.Errorf("a built-in tint is missing %v", got)
	}
	tn := &Tint{Fg: FromHex("#ffffff"), Red: FromHex("#ff0000"), Purple: FromHex("#ff00ff")}
	want := []string{"bg", "black", "green", "yellow", "blue", "cyan", "white"}
	if got := tn.Missing(); !slices.Equal(got, want) {
		t.Errorf("Missing() = %v, want %v", got, want)
	}
}

func TestParseJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, in string
		want     bool
	}{
		{name: "dark bg", in: `{"bg": "#101010"}`, want: true},
		{name: "light bg", in: `{"bg": "#f0f0f0"}`},
		{name: "set", in: `{"bg": "#101010", "dark": false}`},
		{name: "set light", in: `{"bg": "#f0f0f0", "dark": true}`, want: true},
		{name: "no bg", in: `{"id": "x"}`},
	}
	for _, tt := range tests {
		tn, err := ParseJSON([]byte(tt.in))
		if err != nil {
			t.Fatalf("%s: %v", tt.name, err)
		}
		if tn.Dark != tt.want {
			t.Errorf("%s: Dark = %v, want %v", tt.name, tn.Dark, tt.want)
		}
	}
	// Everything else decodes, and it's completed.
	tn, err := ParseJSON([]byte(`{"id": "x", "red": {"r": 255, "g": 0, "b": 0}}`))
	if err != nil {
		t.Fatal(err)
	}
	if tn.DisplayName != "x" || tn.Red == nil || *tn.Red != (Color{R: 0xff, A: 0xff}) || tn.BrightRed != tn.Red {
		t.Errorf("decoded %+v", tn)
	}
	if _, err = ParseJSON([]byte(`{"red": 3}`)); err == nil {
		t.Error("a bad color isn't an error")
	}
	// A tint round-trips, light ones included.
	b, err := json.Marshal(&Tint{ID: "x", Bg: FromHex("#f0f0f0"), Dark: false})
	if err != nil {
		t.Fatal(err)
	}
	if back, perr := ParseJSON(b); perr != nil || back.Dark {
		t.Errorf("round trip: %s -> %+v, %v", b, back, perr)
	}
}

// Decoding into a tint, or a struct that embeds one, is the standard
// decoder's: it keeps what the JSON leaves out, and decodes the struct's own
// fields.
func TestTintStandardJSON(t *testing.T) {
	t.Parallel()
	tn := *TintTokyoNight
	if err := json.Unmarshal([]byte(`{"id": "mine"}`), &tn); err != nil {
		t.Fatal(err)
	}
	if tn.ID != "mine" || tn.Red != TintTokyoNight.Red {
		t.Errorf("decoding into a tint dropped its colors: %+v", tn)
	}
	var outer struct {
		Tint

		Guards bool `json:"guards"`
	}
	if err := json.Unmarshal([]byte(`{"id": "x", "guards": true}`), &outer); err != nil {
		t.Fatal(err)
	}
	if !outer.Guards || outer.ID != "x" {
		t.Errorf("a struct embedding a tint lost its own fields: %+v", outer)
	}
}

func TestPair(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, out string
		start   Pair
		want    Pair
		err     bool
	}{
		{in: `"nord"`, out: `"nord"`, want: Pair{Dark: "nord", Light: "nord"}},
		{in: `{"dark": "a", "light": "b"}`, out: `{"dark":"a","light":"b"}`, want: Pair{Dark: "a", Light: "b"}},
		{in: `{"light": "b"}`, out: `{"dark":"a","light":"b"}`, start: Pair{Dark: "a", Light: "a"}, want: Pair{Dark: "a", Light: "b"}},
		{in: `{"dark": "a", "other": "b"}`, err: true},
		{in: `3`, err: true},
	}
	for _, tt := range tests {
		p := tt.start
		err := json.Unmarshal([]byte(tt.in), &p)
		if tt.err {
			if !errors.Is(err, ErrInvalidPair) {
				t.Errorf("%s: error = %v, want ErrInvalidPair", tt.in, err)
			}
			continue
		}
		if err != nil || p != tt.want {
			t.Errorf("%s: = %+v, %v, want %+v", tt.in, p, err, tt.want)
			continue
		}
		if b, merr := json.Marshal(p); merr != nil || string(b) != tt.out {
			t.Errorf("%s: encoded %s, %v, want %s", tt.in, b, merr, tt.out)
		}
	}
	p := Pair{Dark: "a", Light: "b"}
	if p.For(true) != "a" || p.For(false) != "b" {
		t.Error("For()")
	}
}

func TestSchemas(t *testing.T) {
	t.Parallel()
	for _, s := range []string{ColorSchema, PairSchema} {
		if !json.Valid([]byte(s)) {
			t.Errorf("schema isn't valid JSON: %s", s)
		}
	}
}

func TestColorUnmarshalAlpha(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   string
		want Color
		err  bool
	}{
		{in: `{"r": 1, "g": 2, "b": 3}`, want: Color{R: 1, G: 2, B: 3, A: 0xff}},
		{in: `{"r": 1, "g": 2, "b": 3, "a": 4}`, want: Color{R: 1, G: 2, B: 3, A: 4}},
		{in: `{"r": 1, "g": 2}`, err: true},
		{in: `{"r": 1, "g": 2, "b": 3, "a": 300}`, err: true},
		{in: `"#f80"`, want: Color{R: 0xff, G: 0x88, A: 0xff}},
	}
	for _, tt := range tests {
		var c Color
		err := json.Unmarshal([]byte(tt.in), &c)
		if (err != nil) != tt.err || (!tt.err && c != tt.want) {
			t.Errorf("%s: = %+v, %v, want %+v", tt.in, c, err, tt.want)
		}
	}
}

func TestSameRGB(t *testing.T) {
	t.Parallel()
	var none *Color
	tests := []struct {
		a, b color.Color
		want bool
	}{
		{a: Color{R: 1, G: 2, B: 3, A: 0xff}, b: color.NRGBA{R: 1, G: 2, B: 3, A: 0xff}, want: true},
		{a: Color{R: 1, G: 2, B: 3, A: 0}, b: &Color{R: 1, G: 2, B: 3, A: 0xff}, want: true},
		{a: color.NRGBA{R: 0xff, A: 0x80}, b: Color{R: 0xff, A: 0xff}, want: true},
		{a: &color.NRGBA{R: 0xff}, b: color.NRGBA{R: 0xff}, want: true},
		{a: &color.NRGBA64{R: 0xffff}, b: Color{R: 0xff, A: 0xff}, want: true},
		{a: Color{R: 1, G: 2, B: 3, A: 0xff}, b: Color{R: 1, G: 2, B: 4, A: 0xff}},
		{a: nil, b: Color{A: 0xff}},
		{a: none, b: nil, want: true},
		{a: (*color.NRGBA)(nil), b: nil, want: true},
		{a: (*color.NRGBA64)(nil), b: Color{A: 0xff}},
		{a: (*color.RGBA)(nil), b: nil, want: true},
		{a: nil, b: nil, want: true},
	}
	for _, tt := range tests {
		if got := SameRGB(tt.a, tt.b); got != tt.want {
			t.Errorf("SameRGB(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}
