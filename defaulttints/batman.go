// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
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

// TintBatman (Batman) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Batman
type TintBatman struct{}

// DisplayName returns the display name of the tint.
func (t *TintBatman) DisplayName() string {
	return "Batman"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBatman) ID() string {
	return "batman"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBatman) About() string {
	return `Tint: Batman`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBatman) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#6f6f6f")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBatman) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1b1d1e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBatman) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintBatman) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintBatman) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#505354")
}

func (t *TintBatman) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#919495")
}

func (t *TintBatman) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#a3a3a6")
}

func (t *TintBatman) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#fff27d")
}

func (t *TintBatman) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#9a9a9d")
}

func (t *TintBatman) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fff78e")
}

func (t *TintBatman) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#dadbd6")
}

func (t *TintBatman) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#feed6c")
}

func (t *TintBatman) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1b1d1e")
}

func (t *TintBatman) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#737174")
}

func (t *TintBatman) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#62605f")
}

func (t *TintBatman) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#c8be46")
}

func (t *TintBatman) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#747271")
}

func (t *TintBatman) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#e6dc44")
}

func (t *TintBatman) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c6c5bf")
}

func (t *TintBatman) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f4fd22")
}
