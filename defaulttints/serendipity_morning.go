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
//
// Additional credit to:
//  * Michael Andreuzza (https://github.com/michael-andreuzza)
//  * Chaphasilor (https://github.com/Chaphasilor)

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintSerendipityMorning (Serendipity Morning) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#SerendipityMorning
//
// Credit to:
//   - Michael Andreuzza (https://github.com/michael-andreuzza)
//   - Chaphasilor (https://github.com/Chaphasilor)
type TintSerendipityMorning struct{}

// DisplayName returns the display name of the tint.
func (t *TintSerendipityMorning) DisplayName() string {
	return "Serendipity Morning"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSerendipityMorning) ID() string {
	return "serendipity_morning"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSerendipityMorning) About() string {
	return `Tint: Serendipity Morning
Tint credits:
  * Michael Andreuzza (https://github.com/michael-andreuzza)
  * Chaphasilor (https://github.com/Chaphasilor)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSerendipityMorning) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#575279")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSerendipityMorning) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#FFFAF3")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSerendipityMorning) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.Color("#F4EFEA")
}

// Cursor returns the recommended color for the cursor.
func (t *TintSerendipityMorning) Cursor() lipgloss.TerminalColor {
	return lipgloss.Color("#9893A5")
}

func (t *TintSerendipityMorning) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#6E6A86")
}

func (t *TintSerendipityMorning) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#3788BE")
}

func (t *TintSerendipityMorning) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#7397DE")
}

func (t *TintSerendipityMorning) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#77AAB3")
}

func (t *TintSerendipityMorning) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#886CDB")
}

func (t *TintSerendipityMorning) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#D26A5D")
}

func (t *TintSerendipityMorning) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#575279")
}

func (t *TintSerendipityMorning) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#C8A299")
}

func (t *TintSerendipityMorning) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#F2E9DE")
}

func (t *TintSerendipityMorning) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#3788BE")
}

func (t *TintSerendipityMorning) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#7397DE")
}

func (t *TintSerendipityMorning) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#77AAB3")
}

func (t *TintSerendipityMorning) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#886CDB")
}

func (t *TintSerendipityMorning) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#D26A5D")
}

func (t *TintSerendipityMorning) White() lipgloss.TerminalColor {
	return lipgloss.Color("#575279")
}

func (t *TintSerendipityMorning) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#C8A299")
}
