// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
)

var _ color.Color = &Color{} // Ensure that Color implements the [color.Color] interface.

// toUint8 converts a JSON number (float64 or int) to uint8. Returns false if the
// value is missing, wrong type, or out of range.
func toUint8(v any) (uint8, bool) {
	switch n := v.(type) {
	case float64:
		if n >= 0 && n <= 255 {
			return uint8(n), true
		}
	case int:
		if n >= 0 && n <= 255 {
			return uint8(n), true
		}
	}
	return 0, false
}

// Color implements the [color.Color] interface, in addition to adding a hex field. It
// can be used anywhere a [color.Color] is expected.
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

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (c Color) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"r": int(c.R),
		"g": int(c.G),
		"b": int(c.B),
		"a": int(c.A),
	}
	return json.Marshal(data)
}

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface. It supports
// both the RGBA values as defined by the [Color] struct, as well as a "hex" string
// in place of the RGBA object, for those which do not want to write theme config
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
		var ok bool
		c.R, ok = toUint8(m["r"])
		if ok {
			c.G, ok = toUint8(m["g"])
		}
		if ok {
			c.B, ok = toUint8(m["b"])
		}
		if ok {
			c.A, ok = toUint8(m["a"])
		}
		if !ok {
			return fmt.Errorf("invalid color values: %v", m)
		}
		return nil
	}

	// Try and parse the hex value, if one exists.
	var hex string

	herr := json.Unmarshal(data, &hex)
	if herr != nil {
		return fmt.Errorf("failed to unmarshal tint color (not RGBA object or hex string): %w", errors.Join(err, herr))
	}

	hexColor := FromHex(hex)
	if hexColor == nil {
		return fmt.Errorf("invalid hex color %q (expected #RRGGBB or #RGB format)", hex)
	}
	c.R = hexColor.R
	c.G = hexColor.G
	c.B = hexColor.B
	c.A = hexColor.A
	return nil
}

// FromHex converts a hex color string to a [Color]. If the hex value is invalid,
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
