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

// TintBuiltinSolarizedLight (Builtin Solarized Light) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Builtin+Solarized+Light
type TintBuiltinSolarizedLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintBuiltinSolarizedLight) DisplayName() string {
	return "Builtin Solarized Light"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBuiltinSolarizedLight) ID() string {
	return "builtin_solarized_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBuiltinSolarizedLight) About() string {
	return `Tint: Builtin Solarized Light`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBuiltinSolarizedLight) Fg() lipgloss.Color {
	return lipgloss.Color("#657b83")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBuiltinSolarizedLight) Bg() lipgloss.Color {
	return lipgloss.Color("#fdf6e3")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBuiltinSolarizedLight) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintBuiltinSolarizedLight) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintBuiltinSolarizedLight) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#002b36")
}

func (t *TintBuiltinSolarizedLight) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#839496")
}

func (t *TintBuiltinSolarizedLight) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#93a1a1")
}

func (t *TintBuiltinSolarizedLight) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#586e75")
}

func (t *TintBuiltinSolarizedLight) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#6c71c4")
}

func (t *TintBuiltinSolarizedLight) BrightRed() lipgloss.Color {
	return lipgloss.Color("#cb4b16")
}

func (t *TintBuiltinSolarizedLight) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#fdf6e3")
}

func (t *TintBuiltinSolarizedLight) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#657b83")
}

func (t *TintBuiltinSolarizedLight) Black() lipgloss.Color {
	return lipgloss.Color("#073642")
}

func (t *TintBuiltinSolarizedLight) Blue() lipgloss.Color {
	return lipgloss.Color("#268bd2")
}

func (t *TintBuiltinSolarizedLight) Cyan() lipgloss.Color {
	return lipgloss.Color("#2aa198")
}

func (t *TintBuiltinSolarizedLight) Green() lipgloss.Color {
	return lipgloss.Color("#859900")
}

func (t *TintBuiltinSolarizedLight) Purple() lipgloss.Color {
	return lipgloss.Color("#d33682")
}

func (t *TintBuiltinSolarizedLight) Red() lipgloss.Color {
	return lipgloss.Color("#dc322f")
}

func (t *TintBuiltinSolarizedLight) White() lipgloss.Color {
	return lipgloss.Color("#eee8d5")
}

func (t *TintBuiltinSolarizedLight) Yellow() lipgloss.Color {
	return lipgloss.Color("#b58900")
}
