// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"slices"

	"github.com/lucasb-eyer/go-colorful"
)

var _ color.Color = &Color{} // Ensure that Color implements the color.Color interface.

// Color implements the color.Color interface, in addition to adding a hex field. It
// can be used anywhere a color.Color is expected.
type Color struct { //nolint:recvcheck
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return r, g, b, a
}

// Hex returns the hex value of the color.
func (c Color) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// MarshalJSON implements the json.Marshaler interface.
func (c Color) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"r": int(c.R),
		"g": int(c.G),
		"b": int(c.B),
		"a": int(c.A),
	}
	return json.Marshal(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It supports both
// the RGBA values as defined by the Color struct, as well as a "hex" string in
// place of the RGBA object, for those which do not want to write theme config
// files by the RGBA values.
//
// Examples:
//
//	{"r": 255, "g": 0, "b": 0, "a": 255}
//	"#ff0000"
func (c *Color) UnmarshalJSON(data []byte) error {
	var m map[string]any
	err := json.Unmarshal(data, &m)
	if err == nil {
		r, rok := m["r"].(int)
		g, gok := m["g"].(int)
		b, bok := m["b"].(int)
		a, aok := m["a"].(int)
		if rok && gok && bok && aok {
			c.R = uint8(r) //nolint:gosec
			c.G = uint8(g) //nolint:gosec
			c.B = uint8(b) //nolint:gosec
			c.A = uint8(a) //nolint:gosec
			return nil
		}
	}

	// Try and parse the hex value, if one exists.
	var hex string

	herr := json.Unmarshal(data, &hex)
	if herr != nil {
		return fmt.Errorf("failed to unmarshal tint color (not RGBA object or hex string): %w", errors.Join(err, herr))
	}

	hexColor := FromHex(hex)
	c.R = hexColor.R
	c.G = hexColor.G
	c.B = hexColor.B
	c.A = hexColor.A
	return nil
}

// FromHex converts a hex color string to a Color. If the hex value is invalid,
// it will default to rgba(0, 0, 0, 255). Supports both #RRGGBB and #RGB formats.
func FromHex(s string) *Color {
	if s == "" || s[0] != '#' {
		return nil
	}

	var hasError bool
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		hasError = true
		return 0
	}

	c := &Color{}
	c.A = 0xff

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		hasError = true
	}
	if hasError {
		return nil
	}
	return c
}

// ensureNotTransparent ensures that the alpha value of a color is not 0, and if
// it is, we will set it to 1. This is useful for when we are converting from
// RGB -> RGBA, and the alpha value is lost in the conversion for gradient purposes.
func ensureNotTransparent(c color.Color) color.Color {
	_, _, _, a := c.RGBA()
	if a == 0 {
		return Alpha(c, 1)
	}
	return c
}

// Gradient blends a series of colors together using multiple stops, into the provided
// number of steps. Uses the "CIE L*, C*, h°"/HCL color-space.
//
// Note that if any of the provided colors are completely transparent, we will
// assume that the alpha value was lost in conversion from RGB -> RGBA, and we
// will set the alpha to opaque, as it's not possible to blend something completely
// transparent.
func Gradient(steps int, stops ...color.Color) []color.Color {
	// Bound to a minimum of 2 steps. If they only provided one, it's actually invalid,
	// but will ensure that we don't panic.
	if steps < 2 {
		steps = 2
	}

	// Ensure they didn't provide any nil colors.
	stops = slices.DeleteFunc(stops, func(c color.Color) bool {
		return c == nil
	})

	if len(stops) == 0 {
		return nil // We can't safely fallback.
	}

	// If they only provided one valid color (or some nil colors), we will just return
	// an array of that color, for the amount of steps they requested.
	if len(stops) == 1 {
		singleColor := stops[0]
		result := make([]color.Color, steps)
		for i := range result {
			result[i] = singleColor
		}
		return result
	}

	blended := make([]color.Color, steps)

	// Convert stops to colorful.Color once
	cstops := make([]colorful.Color, len(stops))
	for i, k := range stops {
		cstops[i], _ = colorful.MakeColor(ensureNotTransparent(k))
	}

	numSegments := len(cstops) - 1
	defaultSize := steps / numSegments
	remainingSteps := steps % numSegments

	resultIndex := 0
	for i := range numSegments {
		from := cstops[i]
		to := cstops[i+1]

		// Calculate segment size.
		segmentSize := defaultSize
		if i < remainingSteps {
			segmentSize++
		}

		divisor := float64(segmentSize - 1)

		// Generate colors for this segment.
		for j := range segmentSize {
			var blendingFactor float64
			if segmentSize > 1 {
				blendingFactor = float64(j) / divisor
			}
			blended[resultIndex] = from.BlendLab(to, blendingFactor).Clamped()
			resultIndex++
		}
	}

	return blended
}

func clamp[T cmp.Ordered](v, low, high T) T {
	if high < low {
		high, low = low, high
	}
	return min(high, max(low, v))
}

// Darken takes a color and makes it darker by a specific percentage (0-1, clamped).
func Darken(c color.Color, percent float64) color.Color {
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

// Lighten makes a color lighter by a specific percentage (0-1, clamped).
func Lighten(c color.Color, percent float64) color.Color {
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

// Alpha adjusts the alpha value of a color using a 0-1 (clamped) float scale
// 0 = transparent, 1 = opaque.
func Alpha(c color.Color, alpha float64) color.Color {
	if c == nil {
		return nil
	}
	r, g, b, _ := c.RGBA()
	return color.RGBA{
		R: uint8(min(255, float64(r>>8))),
		G: uint8(min(255, float64(g>>8))),
		B: uint8(min(255, float64(b>>8))),
		A: uint8(clamp(alpha, 0, 1) * 255),
	}
}
