// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// All tints can be previewed here:
//  * https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes
//  * https://github.com/mbadolato/iTerm2-Color-Schemes

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintTangoAdapted (Tango Adapted) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#TangoAdapted
type TintTangoAdapted struct{}

// DisplayName returns the display name of the tint.
func (t *TintTangoAdapted) DisplayName() string {
	return "Tango Adapted"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintTangoAdapted) ID() string {
	return "tango_adapted"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintTangoAdapted) About() string {
	return `Tint: Tango Adapted`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintTangoAdapted) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// Bg returns the recommended default background color for this tint.
func (t *TintTangoAdapted) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintTangoAdapted) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintTangoAdapted) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintTangoAdapted) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#8f928b")
}

func (t *TintTangoAdapted) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#88c9ff")
}

func (t *TintTangoAdapted) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00feff")
}

func (t *TintTangoAdapted) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#93ff00")
}

func (t *TintTangoAdapted) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e9a7e1")
}

func (t *TintTangoAdapted) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff0013")
}

func (t *TintTangoAdapted) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f6f6f4")
}

func (t *TintTangoAdapted) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fff121")
}

func (t *TintTangoAdapted) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintTangoAdapted) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#00a2ff")
}

func (t *TintTangoAdapted) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00d0d6")
}

func (t *TintTangoAdapted) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#59d600")
}

func (t *TintTangoAdapted) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c17ecc")
}

func (t *TintTangoAdapted) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff0000")
}

func (t *TintTangoAdapted) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e6ebe1")
}

func (t *TintTangoAdapted) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f0cb00")
}
