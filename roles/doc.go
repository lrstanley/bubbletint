// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package roles derives semantic colors from a tint: what a color is for,
// such as [Muted] text, a [Raised] dialog or a [Danger] pill, rather than which
// of the tint's 16 colors it is. Every role is worked out from the tint's
// background, text and ANSI colors, so the same code gives a dark, a light, or
// a single-hue tint usable surfaces, readable text and distinct status colors.
//
// [Derive] works out a [Palette] in true color. Contrast guards keep text,
// marks and lines readable against what they're drawn on, moving a color's
// lightness but keeping its hue. [Palette.Fit] then fits the palette to the
// terminal's color depth: with 256 colors no two neighboring roles collapse
// into one index and the guards still hold; with 16, each role takes a slot
// from its family that reads on the terminal's palette, or the terminal's
// default plus an [Attr]; with none, only attributes are left.
//
// Fitted colors are [image/color.Color]s, so they work with anything that
// draws those. With Lip Gloss, for example:
//
//	pal := roles.Derive(tint.Current())
//	f := pal.Fit(roles.Target{Depth: roles.Colors256})
//
//	dialog := lipgloss.NewStyle().Background(f.Color(roles.Raised))
//	hint := lipgloss.NewStyle().Foreground(f.Color(roles.Subtle))
//	failed := lipgloss.NewStyle().
//		Background(f.Color(roles.Danger)).
//		Foreground(f.Color(roles.Danger.On()))
//
// An app adds roles of its own with [Define]; they're derived, guarded and
// fitted like the core roles.
//
// Role names and the API follow semantic versioning. The exact colors derived
// for a role don't: they may be tuned in a minor release, as the derivation
// improves.
package roles
