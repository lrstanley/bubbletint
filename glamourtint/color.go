// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package glamourtint

import (
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

// step moves c by amount. Positive amounts lighten, negative amounts darken.
// A missing color returns fallback.
func step(c, fallback color.Color, amount float64) color.Color {
	if !hasColor(c) {
		if !hasColor(fallback) {
			return nil
		}
		return fallback
	}
	switch {
	case amount > 0:
		return tint.Lighten(c, amount)
	case amount < 0:
		return tint.Darken(c, -amount)
	default:
		return c
	}
}
