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

// TintUltraViolent (UltraViolent) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#UltraViolent
type TintUltraViolent struct{}

// DisplayName returns the display name of the tint.
func (t *TintUltraViolent) DisplayName() string {
	return "UltraViolent"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintUltraViolent) ID() string {
	return "ultra_violent"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintUltraViolent) About() string {
	return `Tint: UltraViolent`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintUltraViolent) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c1c1c1")
}

// Bg returns the recommended default background color for this tint.
func (t *TintUltraViolent) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#242728")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintUltraViolent) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintUltraViolent) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintUltraViolent) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#636667")
}

func (t *TintUltraViolent) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#7fecff")
}

func (t *TintUltraViolent) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#69fcd3")
}

func (t *TintUltraViolent) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#deff8c")
}

func (t *TintUltraViolent) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e681ff")
}

func (t *TintUltraViolent) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fb58b4")
}

func (t *TintUltraViolent) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f9f9f5")
}

func (t *TintUltraViolent) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ebe087")
}

func (t *TintUltraViolent) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#242728")
}

func (t *TintUltraViolent) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#47e0fb")
}

func (t *TintUltraViolent) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#0effbb")
}

func (t *TintUltraViolent) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#b6ff00")
}

func (t *TintUltraViolent) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#d731ff")
}

func (t *TintUltraViolent) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff0090")
}

func (t *TintUltraViolent) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e1e1e1")
}

func (t *TintUltraViolent) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fff727")
}
