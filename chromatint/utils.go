// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package chromatint

import (
	"cmp"
	"image/color"
)

// darken takes a color and makes it darker by a specific percentage (0-1, clamped).
func darken(c color.Color, percent float64) color.Color {
	if c == nil {
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

// lighten makes a color lighter by a specific percentage (0-1, clamped).
func lighten(c color.Color, percent float64) color.Color {
	if c == nil {
		return nil
	}

	add := 255 * (clamp(percent, 0, 1))

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
