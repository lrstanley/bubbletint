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

// TintHarper (Harper) is a collection of lipgloss styles.
type TintHarper struct{}

// DisplayName returns the display name of the tint.
func (t *TintHarper) DisplayName() string {
	return "Harper"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintHarper) ID() string {
	return "harper"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintHarper) About() string {
	return `Tint: Harper`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintHarper) Fg() lipgloss.Color {
	return lipgloss.Color("#a8a49d")
}

// Bg returns the recommended default background color for this tint.
func (t *TintHarper) Bg() lipgloss.Color {
	return lipgloss.Color("#010101")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintHarper) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintHarper) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintHarper) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#726e6a")
}

func (t *TintHarper) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#489e48")
}

func (t *TintHarper) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#f5bfd7")
}

func (t *TintHarper) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#7fb5e1")
}

func (t *TintHarper) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#b296c6")
}

func (t *TintHarper) BrightRed() lipgloss.Color {
	return lipgloss.Color("#f8b63f")
}

func (t *TintHarper) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#fefbea")
}

func (t *TintHarper) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#d6da25")
}

func (t *TintHarper) Black() lipgloss.Color {
	return lipgloss.Color("#010101")
}

func (t *TintHarper) Blue() lipgloss.Color {
	return lipgloss.Color("#489e48")
}

func (t *TintHarper) Cyan() lipgloss.Color {
	return lipgloss.Color("#f5bfd7")
}

func (t *TintHarper) Green() lipgloss.Color {
	return lipgloss.Color("#7fb5e1")
}

func (t *TintHarper) Purple() lipgloss.Color {
	return lipgloss.Color("#b296c6")
}

func (t *TintHarper) Red() lipgloss.Color {
	return lipgloss.Color("#f8b63f")
}

func (t *TintHarper) White() lipgloss.Color {
	return lipgloss.Color("#a8a49d")
}

func (t *TintHarper) Yellow() lipgloss.Color {
	return lipgloss.Color("#d6da25")
}
