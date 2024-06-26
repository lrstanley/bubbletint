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

// TintWombat (Wombat) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Wombat
type TintWombat struct{}

// DisplayName returns the display name of the tint.
func (t *TintWombat) DisplayName() string {
	return "Wombat"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintWombat) ID() string {
	return "wombat"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintWombat) About() string {
	return `Tint: Wombat`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintWombat) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#dedacf")
}

// Bg returns the recommended default background color for this tint.
func (t *TintWombat) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#171717")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintWombat) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintWombat) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintWombat) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#313131")
}

func (t *TintWombat) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#a5c7ff")
}

func (t *TintWombat) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#b7fff9")
}

func (t *TintWombat) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#ddf88f")
}

func (t *TintWombat) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ddaaff")
}

func (t *TintWombat) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f58c80")
}

func (t *TintWombat) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintWombat) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#eee5b2")
}

func (t *TintWombat) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintWombat) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#5da9f6")
}

func (t *TintWombat) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#82fff7")
}

func (t *TintWombat) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#b1e969")
}

func (t *TintWombat) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#e86aff")
}

func (t *TintWombat) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff615a")
}

func (t *TintWombat) White() lipgloss.TerminalColor {
	return lipgloss.Color("#dedacf")
}

func (t *TintWombat) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ebd99c")
}
