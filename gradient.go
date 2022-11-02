// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bubbletint

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

// FromTo converts a color from one color space to another colorful's BlendLuv --
// which blends two colors in the CIE-L*u*v* color-space.
func FromTo(from, to lipgloss.TerminalColor, steps int) (gradient []lipgloss.TerminalColor) {
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
		gradient = append(gradient, lipgloss.Color(start.BlendLuv(end, i/fsteps).Hex()))
	}

	return gradient
}
