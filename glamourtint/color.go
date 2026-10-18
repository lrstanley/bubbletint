// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package glamourtint

import (
	"cmp"
	"fmt"
	"image/color"

	tint "github.com/lrstanley/bubbletint/v2"
)

// hasColor reports whether c is a usable color. A nil *tint.Color stored in a
// color.Color interface is not a nil interface, and calling RGBA on it panics.
func hasColor(c color.Color) bool {
	switch c := c.(type) {
	case nil:
		return false
	case *tint.Color:
		return c != nil
	default:
		return true
	}
}

func toHex(c color.Color) string {
	if !hasColor(c) {
		return ""
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}

func colorPtr(c color.Color) *string {
	if !hasColor(c) {
		return nil
	}
	s := toHex(c)
	return &s
}

// pick returns the dark color when isDark is set, and the light color otherwise.
func pick(isDark bool, light, dark *tint.Color) *tint.Color {
	if isDark {
		return dark
	}
	return light
}

func firstColor(colors ...*tint.Color) *tint.Color {
	for _, c := range colors {
		if c != nil {
			return c
		}
	}
	return nil
}

// codeChip is a background a step away from the tint background, so inline
// code reads as a chip on both dark and light themes.
func codeChip(isDark bool, bg, fallback *tint.Color) color.Color {
	if bg == nil {
		if fallback == nil {
			return nil
		}
		return fallback
	}
	if isDark {
		return lighten(bg, 0.08)
	}
	return darken(bg, 0.06)
}

func darken(c color.Color, percent float64) color.Color {
	if !hasColor(c) {
		return nil
	}

	mult := 1.0 - clamp(percent, 0, 1)

	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(float64(r>>8) * mult),
		G: uint8(float64(g>>8) * mult),
		B: uint8(float64(b>>8) * mult),
		A: uint8(min(255, float64(a>>8))),
	}
}

func lighten(c color.Color, percent float64) color.Color {
	if !hasColor(c) {
		return nil
	}

	add := 255 * clamp(percent, 0, 1)

	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(min(255, float64(r>>8)+add)),
		G: uint8(min(255, float64(g>>8)+add)),
		B: uint8(min(255, float64(b>>8)+add)),
		A: uint8(min(255, float64(a>>8))),
	}
}

func clamp[T cmp.Ordered](v, low, high T) T {
	if high < low {
		high, low = low, high
	}
	return min(high, max(low, v))
}
