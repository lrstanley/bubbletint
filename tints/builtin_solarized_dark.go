// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes

package tints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintBuiltinSolarizedDark (Builtin Solarized Dark) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Builtin+Solarized+Dark
type TintBuiltinSolarizedDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintBuiltinSolarizedDark) DisplayName() string {
	return "Builtin Solarized Dark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBuiltinSolarizedDark) ID() string {
	return "builtin_solarized_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBuiltinSolarizedDark) About() string {
	return `Tint: Builtin Solarized Dark`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBuiltinSolarizedDark) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#839496")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBuiltinSolarizedDark) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#002b36")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBuiltinSolarizedDark) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintBuiltinSolarizedDark) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintBuiltinSolarizedDark) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#002b36")
}

func (t *TintBuiltinSolarizedDark) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#839496")
}

func (t *TintBuiltinSolarizedDark) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#93a1a1")
}

func (t *TintBuiltinSolarizedDark) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#586e75")
}

func (t *TintBuiltinSolarizedDark) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#6c71c4")
}

func (t *TintBuiltinSolarizedDark) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#cb4b16")
}

func (t *TintBuiltinSolarizedDark) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fdf6e3")
}

func (t *TintBuiltinSolarizedDark) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#657b83")
}

func (t *TintBuiltinSolarizedDark) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#073642")
}

func (t *TintBuiltinSolarizedDark) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#268bd2")
}

func (t *TintBuiltinSolarizedDark) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#2aa198")
}

func (t *TintBuiltinSolarizedDark) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#859900")
}

func (t *TintBuiltinSolarizedDark) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#d33682")
}

func (t *TintBuiltinSolarizedDark) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#dc322f")
}

func (t *TintBuiltinSolarizedDark) White() lipgloss.TerminalColor {
	return lipgloss.Color("#eee8d5")
}

func (t *TintBuiltinSolarizedDark) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#b58900")
}
