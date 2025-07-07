// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"

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
func FromHex(hex string) *Color {
	c, err := colorful.Hex(hex)
	if err != nil {
		return &Color{A: 255}
	}

	return &Color{
		R: uint8(c.R * 255),
		G: uint8(c.G * 255),
		B: uint8(c.B * 255),
		A: 255,
	}
}

func ensureValidAlpha(c color.Color) color.Color {
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
	if len(stops) < 2 {
		panic("must provide at least two stops")
	}

	if steps < 2 {
		panic("must provide at least two steps")
	}

	cstops := make([]colorful.Color, len(stops)) // colorful-converted stops.
	for i, k := range stops {
		cstops[i], _ = colorful.MakeColor(ensureValidAlpha(k))
	}

	// Each segment will have a certain amount of colors we need to calculate and
	// distribute. We will likely have a remainder, so we need to distribute those
	// across the segments.
	segments := make([]int, len(cstops)-1)
	defaultSize := steps / len(segments)
	remainingSteps := steps % len(segments)

	for i := range segments {
		segments[i] = defaultSize

		// Distribute remaining across segments so we can meet the number of steps
		// exactly.
		if i < remainingSteps {
			segments[i]++
		}
	}

	blended := make([]color.Color, 0, steps)

	// For each segment, we will blend the colors between the start and end stop
	// we calculated for that specific segment.
	for i := range segments {
		from := cstops[i]
		to := cstops[i+1]
		size := segments[i]

		for j := range size {
			// Defaults to 0, which would result in the "from" color.
			var blendingFactor float64

			// Adjust the blending factor (angle in the HCL color space) when there
			// is more than one color in this segment.
			if size > 1 {
				blendingFactor = float64(j) / float64(size-1)
			}

			blended = append(blended, from.BlendHcl(to, blendingFactor))
		}
	}

	return blended
}

// Darken takes a color and makes it darker by a specific percentage (0-100, clamped).
func Darken(c color.Color, percent int) color.Color {
	mult := 1.0 - max(min(float64(percent), 100), 0)/100.0

	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(float64(r>>8) * mult),
		G: uint8(float64(g>>8) * mult),
		B: uint8(float64(b>>8) * mult),
		A: uint8(a >> 8), //nolint:gosec
	}
}

// Lighten makes a color lighter by a specific percentage (0-100, clamped).
func Lighten(c color.Color, percent int) color.Color {
	add := 255 * (max(min(float64(percent), 100), 0) / 100.0)

	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(min(255, float64(r>>8)+add)),
		G: uint8(min(255, float64(g>>8)+add)),
		B: uint8(min(255, float64(b>>8)+add)),
		A: uint8(a >> 8), //nolint:gosec
	}
}

// Alpha adjusts the alpha value of a color using a 0-1 float scale (0 = fully
// transparent, 1 = fully opaque).
func Alpha(c color.Color, alpha float64) color.Color {
	r, g, b, _ := c.RGBA()
	return color.RGBA{
		R: uint8(r >> 8), //nolint:gosec
		G: uint8(g >> 8), //nolint:gosec
		B: uint8(b >> 8), //nolint:gosec
		A: uint8(max(min(alpha, 1), 0) * 255),
	}
}
