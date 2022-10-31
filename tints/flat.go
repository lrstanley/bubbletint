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

// TintFlat (Flat) is a collection of lipgloss styles.
type TintFlat struct{}

// DisplayName returns the display name of the tint.
func (t *TintFlat) DisplayName() string {
	return "Flat"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintFlat) ID() string {
	return "flat"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintFlat) About() string {
	return `Tint: Flat`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintFlat) Fg() lipgloss.Color {
	return lipgloss.Color("#2cc55d")
}

// Bg returns the recommended default background color for this tint.
func (t *TintFlat) Bg() lipgloss.Color {
	return lipgloss.Color("#002240")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintFlat) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintFlat) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintFlat) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#212c3c")
}

func (t *TintFlat) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#3c7dd2")
}

func (t *TintFlat) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#35b387")
}

func (t *TintFlat) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#2d9440")
}

func (t *TintFlat) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#8230a7")
}

func (t *TintFlat) BrightRed() lipgloss.Color {
	return lipgloss.Color("#d4312e")
}

func (t *TintFlat) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#e7eced")
}

func (t *TintFlat) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#e5be0c")
}

func (t *TintFlat) Black() lipgloss.Color {
	return lipgloss.Color("#222d3f")
}

func (t *TintFlat) Blue() lipgloss.Color {
	return lipgloss.Color("#3167ac")
}

func (t *TintFlat) Cyan() lipgloss.Color {
	return lipgloss.Color("#2c9370")
}

func (t *TintFlat) Green() lipgloss.Color {
	return lipgloss.Color("#32a548")
}

func (t *TintFlat) Purple() lipgloss.Color {
	return lipgloss.Color("#781aa0")
}

func (t *TintFlat) Red() lipgloss.Color {
	return lipgloss.Color("#a82320")
}

func (t *TintFlat) White() lipgloss.Color {
	return lipgloss.Color("#b0b6ba")
}

func (t *TintFlat) Yellow() lipgloss.Color {
	return lipgloss.Color("#e58d11")
}
