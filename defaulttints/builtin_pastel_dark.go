// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintBuiltinPastelDark (Builtin Pastel Dark) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Builtin+Pastel+Dark
type TintBuiltinPastelDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintBuiltinPastelDark) DisplayName() string {
	return "Builtin Pastel Dark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBuiltinPastelDark) ID() string {
	return "builtin_pastel_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBuiltinPastelDark) About() string {
	return `Tint: Builtin Pastel Dark`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBuiltinPastelDark) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#bbbbbb")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBuiltinPastelDark) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBuiltinPastelDark) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintBuiltinPastelDark) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintBuiltinPastelDark) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7c7c7c")
}

func (t *TintBuiltinPastelDark) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#b5dcff")
}

func (t *TintBuiltinPastelDark) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#dfdffe")
}

func (t *TintBuiltinPastelDark) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#ceffac")
}

func (t *TintBuiltinPastelDark) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff9cfe")
}

func (t *TintBuiltinPastelDark) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ffb6b0")
}

func (t *TintBuiltinPastelDark) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintBuiltinPastelDark) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffcc")
}

func (t *TintBuiltinPastelDark) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#4f4f4f")
}

func (t *TintBuiltinPastelDark) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#96cbfe")
}

func (t *TintBuiltinPastelDark) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#c6c5fe")
}

func (t *TintBuiltinPastelDark) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#a8ff60")
}

func (t *TintBuiltinPastelDark) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff73fd")
}

func (t *TintBuiltinPastelDark) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6c60")
}

func (t *TintBuiltinPastelDark) White() lipgloss.TerminalColor {
	return lipgloss.Color("#eeeeee")
}

func (t *TintBuiltinPastelDark) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffb6")
}