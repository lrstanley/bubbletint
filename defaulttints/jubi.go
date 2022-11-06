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

// TintJubi (jubi) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Jubi
type TintJubi struct{}

// DisplayName returns the display name of the tint.
func (t *TintJubi) DisplayName() string {
	return "jubi"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintJubi) ID() string {
	return "jubi"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintJubi) About() string {
	return `Tint: jubi`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintJubi) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c3d3de")
}

// Bg returns the recommended default background color for this tint.
func (t *TintJubi) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#262b33")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintJubi) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintJubi) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintJubi) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#a874ce")
}

func (t *TintJubi) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#8c9fcd")
}

func (t *TintJubi) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#b7c9ef")
}

func (t *TintJubi) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#bcdd61")
}

func (t *TintJubi) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e16c87")
}

func (t *TintJubi) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#de90ab")
}

func (t *TintJubi) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#d5e5f1")
}

func (t *TintJubi) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#87e9ea")
}

func (t *TintJubi) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#3b3750")
}

func (t *TintJubi) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#576ea6")
}

func (t *TintJubi) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#75a7d2")
}

func (t *TintJubi) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#90a94b")
}

func (t *TintJubi) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#bc4f68")
}

func (t *TintJubi) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#cf7b98")
}

func (t *TintJubi) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c3d3de")
}

func (t *TintJubi) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#6ebfc0")
}
