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

// TintZenburn (Zenburn) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Zenburn
type TintZenburn struct{}

// DisplayName returns the display name of the tint.
func (t *TintZenburn) DisplayName() string {
	return "Zenburn"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintZenburn) ID() string {
	return "zenburn"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintZenburn) About() string {
	return `Tint: Zenburn`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintZenburn) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#dcdccc")
}

// Bg returns the recommended default background color for this tint.
func (t *TintZenburn) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#3f3f3f")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintZenburn) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintZenburn) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintZenburn) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#709080")
}

func (t *TintZenburn) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#94bff3")
}

func (t *TintZenburn) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#93e0e3")
}

func (t *TintZenburn) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#c3bf9f")
}

func (t *TintZenburn) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ec93d3")
}

func (t *TintZenburn) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#dca3a3")
}

func (t *TintZenburn) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintZenburn) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e0cf9f")
}

func (t *TintZenburn) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#4d4d4d")
}

func (t *TintZenburn) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#506070")
}

func (t *TintZenburn) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8cd0d3")
}

func (t *TintZenburn) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#60b48a")
}

func (t *TintZenburn) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#dc8cc3")
}

func (t *TintZenburn) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#705050")
}

func (t *TintZenburn) White() lipgloss.TerminalColor {
	return lipgloss.Color("#dcdccc")
}

func (t *TintZenburn) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f0dfaf")
}
