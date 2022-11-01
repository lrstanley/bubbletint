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

// TintBuiltinTangoLight (Builtin Tango Light) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Builtin+Tango+Light
type TintBuiltinTangoLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintBuiltinTangoLight) DisplayName() string {
	return "Builtin Tango Light"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBuiltinTangoLight) ID() string {
	return "builtin_tango_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBuiltinTangoLight) About() string {
	return `Tint: Builtin Tango Light`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBuiltinTangoLight) Fg() lipgloss.Color {
	return lipgloss.Color("#000000")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBuiltinTangoLight) Bg() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBuiltinTangoLight) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintBuiltinTangoLight) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintBuiltinTangoLight) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#555753")
}

func (t *TintBuiltinTangoLight) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#729fcf")
}

func (t *TintBuiltinTangoLight) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#34e2e2")
}

func (t *TintBuiltinTangoLight) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#8ae234")
}

func (t *TintBuiltinTangoLight) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#ad7fa8")
}

func (t *TintBuiltinTangoLight) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ef2929")
}

func (t *TintBuiltinTangoLight) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#eeeeec")
}

func (t *TintBuiltinTangoLight) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fce94f")
}

func (t *TintBuiltinTangoLight) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintBuiltinTangoLight) Blue() lipgloss.Color {
	return lipgloss.Color("#3465a4")
}

func (t *TintBuiltinTangoLight) Cyan() lipgloss.Color {
	return lipgloss.Color("#06989a")
}

func (t *TintBuiltinTangoLight) Green() lipgloss.Color {
	return lipgloss.Color("#4e9a06")
}

func (t *TintBuiltinTangoLight) Purple() lipgloss.Color {
	return lipgloss.Color("#75507b")
}

func (t *TintBuiltinTangoLight) Red() lipgloss.Color {
	return lipgloss.Color("#cc0000")
}

func (t *TintBuiltinTangoLight) White() lipgloss.Color {
	return lipgloss.Color("#d3d7cf")
}

func (t *TintBuiltinTangoLight) Yellow() lipgloss.Color {
	return lipgloss.Color("#c4a000")
}
