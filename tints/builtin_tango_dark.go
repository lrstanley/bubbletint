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

// TintBuiltinTangoDark (Builtin Tango Dark) is a collection of lipgloss styles.
type TintBuiltinTangoDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintBuiltinTangoDark) DisplayName() string {
	return "Builtin Tango Dark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBuiltinTangoDark) ID() string {
	return "builtin_tango_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBuiltinTangoDark) About() string {
	return `Tint: Builtin Tango Dark`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBuiltinTangoDark) Fg() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBuiltinTangoDark) Bg() lipgloss.Color {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBuiltinTangoDark) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintBuiltinTangoDark) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintBuiltinTangoDark) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#555753")
}

func (t *TintBuiltinTangoDark) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#729fcf")
}

func (t *TintBuiltinTangoDark) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#34e2e2")
}

func (t *TintBuiltinTangoDark) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#8ae234")
}

func (t *TintBuiltinTangoDark) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#ad7fa8")
}

func (t *TintBuiltinTangoDark) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ef2929")
}

func (t *TintBuiltinTangoDark) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#eeeeec")
}

func (t *TintBuiltinTangoDark) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fce94f")
}

func (t *TintBuiltinTangoDark) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintBuiltinTangoDark) Blue() lipgloss.Color {
	return lipgloss.Color("#3465a4")
}

func (t *TintBuiltinTangoDark) Cyan() lipgloss.Color {
	return lipgloss.Color("#06989a")
}

func (t *TintBuiltinTangoDark) Green() lipgloss.Color {
	return lipgloss.Color("#4e9a06")
}

func (t *TintBuiltinTangoDark) Purple() lipgloss.Color {
	return lipgloss.Color("#75507b")
}

func (t *TintBuiltinTangoDark) Red() lipgloss.Color {
	return lipgloss.Color("#cc0000")
}

func (t *TintBuiltinTangoDark) White() lipgloss.Color {
	return lipgloss.Color("#d3d7cf")
}

func (t *TintBuiltinTangoDark) Yellow() lipgloss.Color {
	return lipgloss.Color("#c4a000")
}
