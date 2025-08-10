// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"fmt"
	"image/color"
	"testing"
)

// hex converts a color to a hex string or panics if invalid.
func hex(hex string) color.Color {
	cf := FromHex(hex)
	if cf == nil {
		panic(fmt.Errorf("invalid hex: %s", hex))
	}
	return cf
}

func expectColorMatches(t *testing.T, got, want color.Color) {
	t.Helper()

	if (got == nil) != (want == nil) {
		t.Errorf("expectColorMatches() = %s, want %s", rgbaString(t, got), rgbaString(t, want))
	}

	if got == nil {
		return
	}

	gr, gg, gb, ga := got.RGBA()
	wr, wg, wb, wa := want.RGBA()

	gru, ggu, gbu, gau := uint8(gr>>8), uint8(gg>>8), uint8(gb>>8), uint8(ga>>8)
	wru, wgu, wbu, wau := uint8(wr>>8), uint8(wg>>8), uint8(wb>>8), uint8(wa>>8)

	if gru != wru || ggu != wgu || gbu != wbu || gau != wau {
		t.Errorf("expectColorMatches() = %s, want %s", rgbaString(t, got), rgbaString(t, want))
	}
}

func rgbaString(t *testing.T, c color.Color) string {
	t.Helper()

	if c == nil {
		return "nil"
	}

	r, g, b, a := c.RGBA()
	return fmt.Sprintf("rgba(%d,%d,%d,%d)", uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
}

func TestParseHex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *Color
	}{
		{name: "valid-6-red", input: "#FF0000", expected: &Color{R: 0xff, G: 0x00, B: 0x00, A: 0xff}},
		{name: "valid-6-green", input: "#00FF00", expected: &Color{R: 0x00, G: 0xff, B: 0x00, A: 0xff}},
		{name: "valid-6-blue", input: "#0000FF", expected: &Color{R: 0x00, G: 0x00, B: 0xff, A: 0xff}},
		{name: "valid-6-white", input: "#FFFFFF", expected: &Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}},
		{name: "valid-6-black", input: "#000000", expected: &Color{R: 0x00, G: 0x00, B: 0x00, A: 0xff}},
		{name: "valid-6-gray", input: "#808080", expected: &Color{R: 0x80, G: 0x80, B: 0x80, A: 0xff}},
		{name: "valid-3-red", input: "#F00", expected: &Color{R: 0xff, G: 0x00, B: 0x00, A: 0xff}},
		{name: "valid-3-green", input: "#0F0", expected: &Color{R: 0x00, G: 0xff, B: 0x00, A: 0xff}},
		{name: "valid-3-blue", input: "#00F", expected: &Color{R: 0x00, G: 0x00, B: 0xff, A: 0xff}},
		{name: "valid-3-white", input: "#FFF", expected: &Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}},
		{name: "valid-3-black", input: "#000", expected: &Color{R: 0x00, G: 0x00, B: 0x00, A: 0xff}},
		{name: "valid-6-lowercase", input: "#ff0000", expected: &Color{R: 0xff, G: 0x00, B: 0x00, A: 0xff}},
		{name: "valid-6-mixed-case", input: "#Ff0000", expected: &Color{R: 0xff, G: 0x00, B: 0x00, A: 0xff}},
		{name: "valid-3-lowercase", input: "#f00", expected: &Color{R: 0xff, G: 0x00, B: 0x00, A: 0xff}},
		{name: "missing-hash-prefix", input: "FF0000", expected: nil},
		{name: "empty-string", input: "", expected: nil},
		{name: "only-hash", input: "#", expected: nil},
		{name: "too-short-3", input: "#F0", expected: nil},
		{name: "too-long-6", input: "#FF00000", expected: nil},
		{name: "invalid-char", input: "#FG0000", expected: nil},
		{name: "invalid-char-3", input: "#FG0", expected: nil},
		{name: "invalid-char-lowercase", input: "#fg0000", expected: nil},
		{name: "invalid-char-mixed", input: "#Fg0000", expected: nil},
		{name: "wrong-len-5", input: "#FF000", expected: nil},
		{name: "wrong-len-8", input: "#FF000000", expected: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FromHex(tt.input)

			if tt.expected == nil && result != nil {
				t.Errorf("FromHex() expected nil but got %s for input %q", rgbaString(t, result), tt.input)
			}

			if tt.expected != nil && result == nil {
				t.Errorf("FromHex() expected %s but got nil for input %q", rgbaString(t, tt.expected), tt.input)
				return
			}

			if tt.expected != nil && result != nil {
				expectColorMatches(t, result, tt.expected)
			}
		})
	}
}
