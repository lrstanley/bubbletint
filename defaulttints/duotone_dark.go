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

// TintDuotoneDark (Duotone Dark) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#DuotoneDark
type TintDuotoneDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintDuotoneDark) DisplayName() string {
	return "Duotone Dark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDuotoneDark) ID() string {
	return "duotone_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDuotoneDark) About() string {
	return `Tint: Duotone Dark`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDuotoneDark) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#b7a1ff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDuotoneDark) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1f1d27")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDuotoneDark) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDuotoneDark) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDuotoneDark) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#353147")
}

func (t *TintDuotoneDark) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc284")
}

func (t *TintDuotoneDark) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#2488ff")
}

func (t *TintDuotoneDark) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#2dcd73")
}

func (t *TintDuotoneDark) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#de8d40")
}

func (t *TintDuotoneDark) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#d9393e")
}

func (t *TintDuotoneDark) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#eae5ff")
}

func (t *TintDuotoneDark) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#d9b76e")
}

func (t *TintDuotoneDark) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1f1d27")
}

func (t *TintDuotoneDark) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc284")
}

func (t *TintDuotoneDark) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#2488ff")
}

func (t *TintDuotoneDark) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#2dcd73")
}

func (t *TintDuotoneDark) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#de8d40")
}

func (t *TintDuotoneDark) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d9393e")
}

func (t *TintDuotoneDark) White() lipgloss.TerminalColor {
	return lipgloss.Color("#b7a1ff")
}

func (t *TintDuotoneDark) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#d9b76e")
}
