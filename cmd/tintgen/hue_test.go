// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestHueBin(t *testing.T) {
	t.Parallel()

	if len(hueStops) == 0 {
		t.Fatal("no hue stops")
	}
	for i := 1; i < len(hueStops); i++ {
		if hueStops[i].Until <= hueStops[i-1].Until || hueStops[i].Until >= 360 {
			t.Fatalf("hue stop %q bound %v is out of order", hueStops[i].Name, hueStops[i].Until)
		}
	}

	for i, stop := range hueStops {
		start := 0.0
		if i > 0 {
			start = hueStops[i-1].Until
		}
		if got := hueBin(start); got != stop.Name {
			t.Errorf("hueBin(%v) = %q, want %q", start, got, stop.Name)
		}
		beforeEnd := stop.Until - 0.001
		if got := hueBin(beforeEnd); got != stop.Name {
			t.Errorf("hueBin(%v) = %q, want %q", beforeEnd, got, stop.Name)
		}
	}

	first := hueStops[0].Name
	wraps := []float64{hueStops[len(hueStops)-1].Until, 359.999, 360, -15, -360}
	for _, angle := range wraps {
		if got := hueBin(angle); got != first {
			t.Errorf("hueBin(%v) = %q, want wrap bucket %q", angle, got, first)
		}
	}
}

func TestHueFromWeights(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		weights map[string]float64
		want    string
	}{
		{name: "nothing qualifies", weights: map[string]float64{}, want: hueNeutral},
		{name: "leading half", weights: map[string]float64{"red": 0.5, "blue": 0.5}, want: "red"},
		{name: "clear leader", weights: map[string]float64{"blue": 0.6, "green": 0.4}, want: "blue"},
		{
			name:    "below share",
			weights: map[string]float64{"red": 0.4, "blue": 0.3, "green": 0.3},
			want:    hueMixed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := hueFromWeights(tt.weights); got != tt.want {
				t.Errorf("hueFromWeights() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestClassifyHue(t *testing.T) {
	t.Parallel()

	red := colorful.Hsl(0, 1, 0.5).Hex()
	green := colorful.Hsl(120, 1, 0.5).Hex()
	blue := colorful.Hsl(220, 1, 0.5).Hex()
	orange := colorful.Hsl(30, 0.9, 0.55).Hex()

	mixed := Tint{
		Red: red, BrightRed: red, Yellow: red, BrightYellow: red,
		Green: green, BrightGreen: green, Cyan: green, BrightCyan: green,
		Blue: blue, BrightBlue: blue, Purple: blue, BrightPurple: blue,
	}

	accent := filledTint("#3a3a3a")
	accent.Bg = "#1c1c1c"
	accent.Fg = "#d8d8d8"
	accent.Black = "#1c1c1c"
	accent.White = "#e8e8e8"
	accent.BrightBlack = "#4a4a4a"
	accent.BrightWhite = "#f5f5f5"
	accent.Red = orange
	accent.BrightRed = orange
	accent.Yellow = orange

	// A saturated background and foreground must not move the hue.
	ignored := filledTint("#4a4a4a")
	ignored.Red = orange
	ignored.Bg = colorful.Hsl(220, 1, 0.5).Hex()
	ignored.Fg = colorful.Hsl(120, 1, 0.5).Hex()

	tests := []struct {
		name string
		tint Tint
		want string
	}{
		{name: "grayscale", tint: filledTint("#808080"), want: hueNeutral},
		{name: "too dark", tint: filledTint("#100000"), want: hueNeutral},
		{name: "too light", tint: filledTint("#fff5f5"), want: hueNeutral},
		{name: "low saturation", tint: filledTint(colorful.Hsl(0, 0.05, 0.5).Hex()), want: hueNeutral},
		{name: "red", tint: filledTint(red), want: "red"},
		{name: "orange", tint: filledTint(orange), want: "orange"},
		{name: "yellow", tint: filledTint(colorful.Hsl(60, 1, 0.5).Hex()), want: "yellow"},
		{name: "green", tint: filledTint(green), want: "green"},
		{name: "cyan", tint: filledTint(colorful.Hsl(180, 1, 0.5).Hex()), want: "cyan"},
		{name: "blue", tint: filledTint(blue), want: "blue"},
		{name: "purple", tint: filledTint(colorful.Hsl(280, 1, 0.5).Hex()), want: "purple"},
		{name: "pink", tint: filledTint(colorful.Hsl(330, 1, 0.5).Hex()), want: "pink"},
		{name: "mixed", tint: mixed, want: hueMixed},
		{name: "orange accents on a dark gray field", tint: accent, want: "orange"},
		{name: "saturated background and foreground are ignored", tint: ignored, want: "orange"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := classifyHue(tt.tint); got != tt.want {
				t.Errorf("classifyHue() = %q, want %q", got, tt.want)
			}
		})
	}
}

func filledTint(hex string) Tint {
	return Tint{
		Fg:           hex,
		Bg:           hex,
		Black:        hex,
		Red:          hex,
		Green:        hex,
		Yellow:       hex,
		Blue:         hex,
		Purple:       hex,
		Cyan:         hex,
		White:        hex,
		BrightBlack:  hex,
		BrightRed:    hex,
		BrightGreen:  hex,
		BrightYellow: hex,
		BrightBlue:   hex,
		BrightPurple: hex,
		BrightCyan:   hex,
		BrightWhite:  hex,
	}
}
