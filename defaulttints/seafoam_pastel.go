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

// TintSeafoamPastel (Seafoam Pastel) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#SeafoamPastel
type TintSeafoamPastel struct{}

// DisplayName returns the display name of the tint.
func (t *TintSeafoamPastel) DisplayName() string {
	return "Seafoam Pastel"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSeafoamPastel) ID() string {
	return "seafoam_pastel"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSeafoamPastel) About() string {
	return `Tint: Seafoam Pastel`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSeafoamPastel) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#d4e7d4")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSeafoamPastel) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#243435")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSeafoamPastel) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSeafoamPastel) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSeafoamPastel) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#8a8a8a")
}

func (t *TintSeafoamPastel) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#7ac3cf")
}

func (t *TintSeafoamPastel) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ade0e0")
}

func (t *TintSeafoamPastel) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#98d9aa")
}

func (t *TintSeafoamPastel) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d6b2a1")
}

func (t *TintSeafoamPastel) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#cf937a")
}

func (t *TintSeafoamPastel) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e0e0e0")
}

func (t *TintSeafoamPastel) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fae79d")
}

func (t *TintSeafoamPastel) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#757575")
}

func (t *TintSeafoamPastel) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#4d7b82")
}

func (t *TintSeafoamPastel) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#729494")
}

func (t *TintSeafoamPastel) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#728c62")
}

func (t *TintSeafoamPastel) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#8a7267")
}

func (t *TintSeafoamPastel) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#825d4d")
}

func (t *TintSeafoamPastel) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e0e0e0")
}

func (t *TintSeafoamPastel) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ada16d")
}
