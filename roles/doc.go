// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package roles derives semantic colors from a tint: what a color is for,
// such as [Muted] text, a [Raised] dialog or a [Danger] pill, rather than which
// of the tint's 16 colors it is. Every role is worked out from the tint's
// background, text and ANSI colors, so the same code gives a dark, a light, or
// a single-hue tint usable surfaces, readable text and distinct status colors.
//
// # Layers, grounds and states
//
// The layers ([Bg], [Surface], [Raised] and [Overlay]) are the same wherever
// they're drawn. Everything else is worked out from what it's drawn on, its
// ground: [Palette.On] and [Fitted.On] return the palette for drawing on any
// fill, whether a dialog, a selected row or a danger pill. There, body text is
// whatever reads on the fill, muted text and lines are mixed from it, tones
// stay readable, [Control] and soft fills step off the fill, and a [Sunken]
// well sinks into it.
//
// Any role can be in a [State]: Overlay.Hover() is a menu's row under the
// pointer, Bg.Selected() the page's selected row, and Accent.Pressed() an
// accent button held down. A state is a step from the role at rest, so it
// shows on any fill, and it's kept apart from the role at rest with 256
// colors, or marked with an attribute where no step can be drawn.
//
// # Fitting
//
// [Derive] works out a [Palette] in true color. Contrast guards keep text,
// marks and lines readable against what they're drawn on, moving a color's
// lightness but keeping its hue. [Palette.Fit] then fits the palette to the
// terminal's color depth: with 256 colors no two neighboring roles collapse
// into one index and the guards still hold; with 16, each role takes a slot
// from its family that reads on the terminal's palette, or the terminal's
// default plus an [Attr]; with none, only attributes are left. Where the
// tint's own hues can't tell the status tones apart, attributes mark them
// too; see [Fitted.Mono].
//
// # Styles
//
// [Fitted.Text] and [Fitted.Fill] return Lip Gloss styles for a role, with
// the attributes that stand in for its color where the terminal can't draw
// it. Text sets no background, for renderers that keep the background under
// text; [Fitted.Span] adds the ground's, for text rendered on its own and
// placed on a fill.
//
//	f := roles.Derive(tint.Current()).Fit(roles.Target{Profile: profile})
//
//	hint := f.Text(roles.Subtle)
//	failed := f.Fill(roles.Danger) // with the text that reads on it
//
//	dialog := f.Fill(roles.Raised)
//	onDialog := f.On(roles.Raised)
//	note := onDialog.Text(roles.Muted)       // muted text, kept readable on the dialog
//	cancel := onDialog.Fill(roles.Control)   // a button, a step off the dialog
//	hover := onDialog.Fill(roles.Control.Hover())
//
//	row := f.Fill(roles.Bg.Selected())
//	rowHint := f.On(roles.Bg.Selected()).Text(roles.Subtle)
//
// Fitted colors are [image/color.Color]s too, for anything else that draws
// those: [Fitted.Color] returns a role's color, [Fitted.Attr] what to add to
// it, and [Fitted.Fit] fits a color of the app's own, such as a gradient's
// step. [Fitted.Distinct] reports whether two roles look different on the
// terminal, to choose a design by, such as outlining a dialog whose surface
// wouldn't show. [Lerp] crossfades between two fitted palettes.
//
// An app adds roles of its own with [Define]; they're derived, guarded and
// fitted like the core roles, on every ground. [AccentChoice] and
// [ColorLimit] parse the settings an app reads from its config file.
//
// Role names and the API follow semantic versioning. The exact colors derived
// for a role don't: they may be tuned in a minor release, as the derivation
// improves.
package roles
