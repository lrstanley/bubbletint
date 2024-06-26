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

// TintDarkPlus (Dark+) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#DarkPlus
type TintDarkPlus struct{}

// DisplayName returns the display name of the tint.
func (t *TintDarkPlus) DisplayName() string {
	return "Dark+"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDarkPlus) ID() string {
	return "dark_plus"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDarkPlus) About() string {
	return `Tint: Dark+`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDarkPlus) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#cccccc")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDarkPlus) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0e0e0e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDarkPlus) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDarkPlus) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDarkPlus) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#666666")
}

func (t *TintDarkPlus) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#3b8eea")
}

func (t *TintDarkPlus) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#29b8db")
}

func (t *TintDarkPlus) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#23d18b")
}

func (t *TintDarkPlus) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d670d6")
}

func (t *TintDarkPlus) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f14c4c")
}

func (t *TintDarkPlus) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e5e5")
}

func (t *TintDarkPlus) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f5f543")
}

func (t *TintDarkPlus) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintDarkPlus) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2472c8")
}

func (t *TintDarkPlus) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#11a8cd")
}

func (t *TintDarkPlus) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#0dbc79")
}

func (t *TintDarkPlus) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#bc3fbc")
}

func (t *TintDarkPlus) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#cd3131")
}

func (t *TintDarkPlus) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e5e5")
}

func (t *TintDarkPlus) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e510")
}
