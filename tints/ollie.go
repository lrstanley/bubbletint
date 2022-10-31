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

// TintOllie (Ollie) is a collection of lipgloss styles.
type TintOllie struct{}

// DisplayName returns the display name of the tint.
func (t *TintOllie) DisplayName() string {
	return "Ollie"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintOllie) ID() string {
	return "ollie"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintOllie) About() string {
	return `Tint: Ollie`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintOllie) Fg() lipgloss.Color {
	return lipgloss.Color("#8a8dae")
}

// Bg returns the recommended default background color for this tint.
func (t *TintOllie) Bg() lipgloss.Color {
	return lipgloss.Color("#222125")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintOllie) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintOllie) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintOllie) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#5b3725")
}

func (t *TintOllie) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#4488ff")
}

func (t *TintOllie) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#1ffaff")
}

func (t *TintOllie) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#3bff99")
}

func (t *TintOllie) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#ffc21d")
}

func (t *TintOllie) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ff3d48")
}

func (t *TintOllie) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#5b6ea7")
}

func (t *TintOllie) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#ff5e1e")
}

func (t *TintOllie) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintOllie) Blue() lipgloss.Color {
	return lipgloss.Color("#2d57ac")
}

func (t *TintOllie) Cyan() lipgloss.Color {
	return lipgloss.Color("#1fa6ac")
}

func (t *TintOllie) Green() lipgloss.Color {
	return lipgloss.Color("#31ac61")
}

func (t *TintOllie) Purple() lipgloss.Color {
	return lipgloss.Color("#b08528")
}

func (t *TintOllie) Red() lipgloss.Color {
	return lipgloss.Color("#ac2e31")
}

func (t *TintOllie) White() lipgloss.Color {
	return lipgloss.Color("#8a8eac")
}

func (t *TintOllie) Yellow() lipgloss.Color {
	return lipgloss.Color("#ac4300")
}
