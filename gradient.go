// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

// FromTo converts a color from one color space to another colorful's BlendLuv --
// which blends two colors in the CIE-L*u*v* color-space. The result is a slice of
// lipgloss.TerminalColor (compatible with lipgloss methods).
func FromTo(from, to lipgloss.TerminalColor, steps int) (gradient []lipgloss.TerminalColor) {
	gradient = make([]lipgloss.TerminalColor, steps)

	for i, hex := range FromToHex(from, to, steps) {
		gradient[i] = lipgloss.Color(hex)
	}

	return gradient
}

// FromToHex converts a color from one color space to another colorful's BlendLuv --
// which blends two colors in the CIE-L*u*v* color-space. The result is a slice of
// hex colors.
func FromToHex(from, to lipgloss.TerminalColor, steps int) (gradient []string) {
	start, ok := colorful.MakeColor(from)
	if !ok {
		panic("invalid from color specified")
	}

	end, ok := colorful.MakeColor(to)
	if !ok {
		panic("invalid to color specified")
	}

	fsteps := float64(steps)

	for i := 0.0; i < fsteps; i++ {
		gradient = append(gradient, start.BlendLuv(end, i/fsteps).Hex())
	}

	return gradient
}
