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

// TintMaterial (Material) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Material
type TintMaterial struct{}

// DisplayName returns the display name of the tint.
func (t *TintMaterial) DisplayName() string {
	return "Material"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintMaterial) ID() string {
	return "material"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintMaterial) About() string {
	return `Tint: Material`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintMaterial) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#232322")
}

// Bg returns the recommended default background color for this tint.
func (t *TintMaterial) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#eaeaea")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintMaterial) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintMaterial) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintMaterial) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#424242")
}

func (t *TintMaterial) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#54a4f3")
}

func (t *TintMaterial) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#26bbd1")
}

func (t *TintMaterial) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#7aba3a")
}

func (t *TintMaterial) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#aa4dbc")
}

func (t *TintMaterial) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e83b3f")
}

func (t *TintMaterial) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#d9d9d9")
}

func (t *TintMaterial) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffea2e")
}

func (t *TintMaterial) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#212121")
}

func (t *TintMaterial) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#134eb2")
}

func (t *TintMaterial) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#0e717c")
}

func (t *TintMaterial) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#457b24")
}

func (t *TintMaterial) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#560088")
}

func (t *TintMaterial) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#b7141f")
}

func (t *TintMaterial) White() lipgloss.TerminalColor {
	return lipgloss.Color("#efefef")
}

func (t *TintMaterial) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f6981e")
}
