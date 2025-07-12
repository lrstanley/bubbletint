// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func mustHex(t *testing.T, hex string) color.Color {
	t.Helper()

	c, err := colorful.Hex(hex)
	if err != nil {
		panic(err)
	}

	return c
}

func colorMatches(t *testing.T, got, want color.Color) {
	t.Helper()

	gr, gg, gb, ga := got.RGBA()
	wr, wg, wb, wa := want.RGBA()

	gru, ggu, gbu, gau := uint8(gr>>8), uint8(gg>>8), uint8(gb>>8), uint8(ga>>8)
	wru, wgu, wbu, wau := uint8(wr>>8), uint8(wg>>8), uint8(wb>>8), uint8(wa>>8)

	if gru != wru || ggu != wgu || gbu != wbu || gau != wau {
		t.Errorf("colorMatches() = %s, want %s", rgbaString(t, got), rgbaString(t, want))
	}
}

func rgbaString(t *testing.T, c color.Color) string {
	t.Helper()

	r, g, b, a := c.RGBA()
	ur, ug, ub, ua := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)
	return fmt.Sprintf("rgba(%d,%d,%d,%d)", ur, ug, ub, ua)
}

func TestDarken(t *testing.T) {
	tests := []struct {
		name     string
		color    color.Color
		percent  int
		expected color.Color
	}{
		{
			name:     "darken-white-50-percent",
			color:    mustHex(t, "#ffffff"),
			percent:  50,
			expected: color.RGBA{R: 127, G: 127, B: 127, A: 255}, // #7f7f7f
		},
		{
			name:     "darken-red-25-percent",
			color:    mustHex(t, "#ff0000"),
			percent:  25,
			expected: color.RGBA{R: 191, G: 0, B: 0, A: 255}, // #bf0000
		},
		{
			name:     "darken-blue-75-percent",
			color:    mustHex(t, "#0000ff"),
			percent:  75,
			expected: color.RGBA{R: 0, G: 0, B: 63, A: 255}, // #00003f
		},
		{
			name:     "darken-black-10-percent",
			color:    mustHex(t, "#000000"),
			percent:  10,
			expected: color.RGBA{R: 0, G: 0, B: 0, A: 255}, // #000000
		},
		{
			name:     "darken-with-clamp-min",
			color:    mustHex(t, "#ffffff"),
			percent:  0,
			expected: color.RGBA{R: 255, G: 255, B: 255, A: 255}, // #fcfcfc
		},
		{
			name:     "darken-with-clamp-max",
			color:    mustHex(t, "#ffffff"),
			percent:  100,
			expected: color.RGBA{R: 0, G: 0, B: 0, A: 255}, // #020202
		},
		{
			name:     "darken-nil-color",
			color:    nil,
			percent:  50,
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Darken(tt.color, tt.percent)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Darken() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLighten(t *testing.T) {
	tests := []struct {
		name     string
		color    color.Color
		percent  int
		expected color.Color
	}{
		{
			name:     "lighten-black-50-percent",
			color:    mustHex(t, "#000000"),
			percent:  50,
			expected: color.RGBA{R: 127, G: 127, B: 127, A: 255}, // #7f7f7f
		},
		{
			name:     "lighten-red-25-percent",
			color:    mustHex(t, "#800000"),
			percent:  25,
			expected: color.RGBA{R: 191, G: 63, B: 63, A: 255}, // #bf3f3f
		},
		{
			name:     "lighten-blue-75-percent",
			color:    mustHex(t, "#000080"),
			percent:  75,
			expected: color.RGBA{R: 191, G: 191, B: 255, A: 255}, // #bfbfff
		},
		{
			name:     "lighten-white-10-percent",
			color:    mustHex(t, "#ffffff"),
			percent:  10,
			expected: color.RGBA{R: 255, G: 255, B: 255, A: 255}, // #ffffff
		},
		{
			name:     "lighten-with-clamp-min",
			color:    mustHex(t, "#000000"),
			percent:  0,
			expected: color.RGBA{R: 0, G: 0, B: 0, A: 255}, // #020202
		},
		{
			name:     "lighten-with-clamp-max",
			color:    mustHex(t, "#000000"),
			percent:  100,
			expected: color.RGBA{R: 255, G: 255, B: 255, A: 255}, // #fcfcfc
		},
		{
			name:     "lighten-nil-color",
			color:    nil,
			percent:  50,
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Lighten(tt.color, tt.percent)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Lighten() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAlpha(t *testing.T) {
	tests := []struct {
		name     string
		color    color.Color
		alpha    float64
		expected color.Color
	}{
		{
			name:     "alpha-full-opacity",
			color:    color.RGBA{R: 255, G: 0, B: 0, A: 255},
			alpha:    1.0,
			expected: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			name:     "alpha-half-opacity",
			color:    color.RGBA{R: 0, G: 255, B: 0, A: 255},
			alpha:    0.5,
			expected: color.RGBA{R: 0, G: 255, B: 0, A: 127},
		},
		{
			name:     "alpha-quarter-opacity",
			color:    color.RGBA{R: 0, G: 0, B: 255, A: 255},
			alpha:    0.25,
			expected: color.RGBA{R: 0, G: 0, B: 255, A: 63},
		},
		{
			name:     "alpha-zero-opacity",
			color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
			alpha:    0.0,
			expected: color.RGBA{R: 255, G: 255, B: 255, A: 0},
		},
		{
			name:     "alpha-clamp-above-max",
			color:    color.RGBA{R: 255, G: 0, B: 255, A: 255},
			alpha:    1.5,
			expected: color.RGBA{R: 255, G: 0, B: 255, A: 255},
		},
		{
			name:     "alpha-clamp-below-min",
			color:    color.RGBA{R: 255, G: 255, B: 0, A: 255},
			alpha:    -0.5,
			expected: color.RGBA{R: 255, G: 255, B: 0, A: 0},
		},
		{
			name:     "alpha-complex-color",
			color:    color.RGBA{R: 18, G: 52, B: 86, A: 255},
			alpha:    0.75,
			expected: color.RGBA{R: 18, G: 52, B: 86, A: 191},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Alpha(tt.color, tt.alpha)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Alpha() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGradient(t *testing.T) {
	tests := []struct {
		name     string
		steps    int
		stops    []color.Color
		expected []color.Color
	}{
		{
			name:  "2-colors-10-steps",
			steps: 10,
			stops: []color.Color{
				color.RGBA{R: 255, G: 0, B: 0, A: 255},
				color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
			expected: []color.Color{
				&color.RGBA{R: 255, G: 0, B: 0, A: 255},
				&color.RGBA{R: 255, G: 0, B: 37, A: 255},
				&color.RGBA{R: 255, G: 0, B: 62, A: 255},
				&color.RGBA{R: 255, G: 0, B: 87, A: 255},
				&color.RGBA{R: 255, G: 0, B: 114, A: 255},
				&color.RGBA{R: 243, G: 0, B: 143, A: 255},
				&color.RGBA{R: 221, G: 0, B: 173, A: 255},
				&color.RGBA{R: 188, G: 0, B: 202, A: 255},
				&color.RGBA{R: 138, G: 0, B: 230, A: 255},
				&color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
		},
		{
			name:  "3-colors-4-steps",
			steps: 4,
			stops: []color.Color{
				color.RGBA{R: 255, G: 0, B: 0, A: 255},
				color.RGBA{R: 0, G: 255, B: 0, A: 255},
				color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
			expected: []color.Color{
				&color.RGBA{R: 255, G: 0, B: 0, A: 255},
				&color.RGBA{R: 0, G: 255, B: 0, A: 255},
				&color.RGBA{R: 0, G: 255, B: 0, A: 255},
				&color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
		},
		{
			name:  "black-to-white-5-steps",
			steps: 5,
			stops: []color.Color{
				color.RGBA{R: 0, G: 0, B: 0, A: 255},
				color.RGBA{R: 255, G: 255, B: 255, A: 255},
			},
			expected: []color.Color{
				&color.RGBA{R: 0, G: 0, B: 0, A: 255},
				&color.RGBA{R: 59, G: 59, B: 59, A: 255},
				&color.RGBA{R: 119, G: 119, B: 119, A: 255},
				&color.RGBA{R: 185, G: 185, B: 185, A: 255},
				&color.RGBA{R: 255, G: 255, B: 255, A: 255},
			},
		},
		{
			name:  "4-colors-6-steps",
			steps: 6,
			stops: []color.Color{
				color.RGBA{R: 255, G: 0, B: 0, A: 255},
				color.RGBA{R: 255, G: 255, B: 0, A: 255},
				color.RGBA{R: 0, G: 255, B: 0, A: 255},
				color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
			expected: []color.Color{
				&color.RGBA{R: 255, G: 0, B: 0, A: 255},
				&color.RGBA{R: 255, G: 255, B: 0, A: 255},
				&color.RGBA{R: 255, G: 255, B: 0, A: 255},
				&color.RGBA{R: 0, G: 255, B: 0, A: 255},
				&color.RGBA{R: 0, G: 255, B: 0, A: 255},
				&color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
		},
		{
			name:  "2-steps-5-stops",
			steps: 2,
			stops: []color.Color{
				color.RGBA{R: 255, G: 0, B: 0, A: 255},
				color.RGBA{R: 0, G: 255, B: 0, A: 255},
				color.RGBA{R: 0, G: 0, B: 255, A: 255},
				color.RGBA{R: 255, G: 255, B: 0, A: 255},
				color.RGBA{R: 0, G: 0, B: 0, A: 255},
			},
			expected: []color.Color{
				&color.RGBA{R: 255, G: 0, B: 0, A: 255},
				&color.RGBA{R: 0, G: 255, B: 0, A: 255},
			},
		},
		{
			name:  "insufficient-stops",
			steps: 3,
			stops: []color.Color{
				color.RGBA{R: 255, G: 0, B: 0, A: 255},
			},
			expected: []color.Color{
				&color.RGBA{R: 255, G: 0, B: 0, A: 255},
				&color.RGBA{R: 255, G: 0, B: 0, A: 255},
				&color.RGBA{R: 255, G: 0, B: 0, A: 255},
			},
		},
		{
			name:  "insufficient-steps",
			steps: 1,
			stops: []color.Color{
				color.RGBA{R: 255, G: 0, B: 0, A: 255},
				color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
			expected: []color.Color{
				&color.RGBA{R: 255, G: 0, B: 0, A: 255},
				&color.RGBA{R: 0, G: 0, B: 255, A: 255},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Gradient(tt.steps, tt.stops...)

			if len(got) != len(tt.expected) {
				t.Errorf("Gradient() = %v length, want %v length", len(got), len(tt.expected))
			}

			for i := range tt.expected {
				colorMatches(t, got[i], tt.expected[i])
			}
		})
	}
}
