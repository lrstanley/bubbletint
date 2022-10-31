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

// TintChester (Chester) is a collection of lipgloss styles.
type TintChester struct{}

// DisplayName returns the display name of the tint.
func (t *TintChester) DisplayName() string {
	return "Chester"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintChester) ID() string {
	return "chester"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintChester) About() string {
	return `Tint: Chester`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintChester) Fg() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintChester) Bg() lipgloss.Color {
	return lipgloss.Color("#2c3643")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintChester) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintChester) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintChester) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#6f6b68")
}

func (t *TintChester) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#278ad6")
}

func (t *TintChester) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#27dede")
}

func (t *TintChester) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#16c98d")
}

func (t *TintChester) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#d34590")
}

func (t *TintChester) BrightRed() lipgloss.Color {
	return lipgloss.Color("#fa5e5b")
}

func (t *TintChester) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintChester) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#feef6d")
}

func (t *TintChester) Black() lipgloss.Color {
	return lipgloss.Color("#080200")
}

func (t *TintChester) Blue() lipgloss.Color {
	return lipgloss.Color("#288ad6")
}

func (t *TintChester) Cyan() lipgloss.Color {
	return lipgloss.Color("#28ddde")
}

func (t *TintChester) Green() lipgloss.Color {
	return lipgloss.Color("#16c98d")
}

func (t *TintChester) Purple() lipgloss.Color {
	return lipgloss.Color("#d34590")
}

func (t *TintChester) Red() lipgloss.Color {
	return lipgloss.Color("#fa5e5b")
}

func (t *TintChester) White() lipgloss.Color {
	return lipgloss.Color("#e7e7e7")
}

func (t *TintChester) Yellow() lipgloss.Color {
	return lipgloss.Color("#ffc83f")
}
