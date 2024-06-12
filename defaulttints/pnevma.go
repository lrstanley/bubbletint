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

// TintPnevma (Pnevma) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Pnevma
type TintPnevma struct{}

// DisplayName returns the display name of the tint.
func (t *TintPnevma) DisplayName() string {
	return "Pnevma"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintPnevma) ID() string {
	return "pnevma"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintPnevma) About() string {
	return `Tint: Pnevma`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintPnevma) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#d0d0d0")
}

// Bg returns the recommended default background color for this tint.
func (t *TintPnevma) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1c1c1c")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintPnevma) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintPnevma) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintPnevma) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#4a4845")
}

func (t *TintPnevma) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#a1bdce")
}

func (t *TintPnevma) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#b1e7dd")
}

func (t *TintPnevma) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#afbea2")
}

func (t *TintPnevma) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d7beda")
}

func (t *TintPnevma) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#d78787")
}

func (t *TintPnevma) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#efefef")
}

func (t *TintPnevma) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e4c9af")
}

func (t *TintPnevma) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#2f2e2d")
}

func (t *TintPnevma) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#7fa5bd")
}

func (t *TintPnevma) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8adbb4")
}

func (t *TintPnevma) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#90a57d")
}

func (t *TintPnevma) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c79ec4")
}

func (t *TintPnevma) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#a36666")
}

func (t *TintPnevma) White() lipgloss.TerminalColor {
	return lipgloss.Color("#d0d0d0")
}

func (t *TintPnevma) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#d7af87")
}
