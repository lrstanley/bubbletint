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

// TintLiquidCarbon (LiquidCarbon) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#LiquidCarbon
type TintLiquidCarbon struct{}

// DisplayName returns the display name of the tint.
func (t *TintLiquidCarbon) DisplayName() string {
	return "LiquidCarbon"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintLiquidCarbon) ID() string {
	return "liquid_carbon"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintLiquidCarbon) About() string {
	return `Tint: LiquidCarbon`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintLiquidCarbon) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#afc2c2")
}

// Bg returns the recommended default background color for this tint.
func (t *TintLiquidCarbon) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#303030")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintLiquidCarbon) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintLiquidCarbon) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintLiquidCarbon) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintLiquidCarbon) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#0099cc")
}

func (t *TintLiquidCarbon) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#7ac4cc")
}

func (t *TintLiquidCarbon) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#559a70")
}

func (t *TintLiquidCarbon) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#cc69c8")
}

func (t *TintLiquidCarbon) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff3030")
}

func (t *TintLiquidCarbon) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#bccccc")
}

func (t *TintLiquidCarbon) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ccac00")
}

func (t *TintLiquidCarbon) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintLiquidCarbon) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0099cc")
}

func (t *TintLiquidCarbon) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#7ac4cc")
}

func (t *TintLiquidCarbon) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#559a70")
}

func (t *TintLiquidCarbon) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#cc69c8")
}

func (t *TintLiquidCarbon) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff3030")
}

func (t *TintLiquidCarbon) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bccccc")
}

func (t *TintLiquidCarbon) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ccac00")
}
