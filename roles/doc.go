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
// [Fitted.Text] and [Fitted.Fill] return Lip Gloss styles for a role, with
// the attributes that stand in for its color where the terminal can't draw
// it. [Fitted.On] gives the roles for drawing on a surface, such as a dialog
// or the selected row: their text stays readable there, and is drawn on the
// surface's color.
//
//	f := roles.Derive(tint.Current()).Fit(roles.Target{Profile: profile})
//
//	hint := f.Text(roles.Subtle)
//	failed := f.Fill(roles.Danger) // with its on-danger text
//
//	dialog := f.Fill(roles.Raised)
//	note := f.On(roles.Raised).Text(roles.Muted) // muted text, on the dialog
//
// Fitted colors are [image/color.Color]s too, for anything else that draws
// those: [Fitted.Color] returns a role's color, and [Fitted.Attr] what to add
// to it.
//
// An app adds roles of its own with [Define]; they're derived, guarded and
// fitted like the core roles.
//
// Role names and the API follow semantic versioning. The exact colors derived
// for a role don't: they may be tuned in a minor release, as the derivation
// improves.
package roles
