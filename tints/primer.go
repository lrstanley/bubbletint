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

// TintPrimer (Primer) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Primer
type TintPrimer struct{}

// DisplayName returns the display name of the tint.
func (t *TintPrimer) DisplayName() string {
	return "Primer"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintPrimer) ID() string {
	return "primer"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintPrimer) About() string {
	return `Tint: Primer`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintPrimer) Fg() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintPrimer) Bg() lipgloss.Color {
	return lipgloss.Color("#1a2b3c")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintPrimer) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintPrimer) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintPrimer) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#4f5861")
}

func (t *TintPrimer) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#c8e1ff")
}

func (t *TintPrimer) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#a2ecec")
}

func (t *TintPrimer) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#bef5cb")
}

func (t *TintPrimer) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#d1bcf9")
}

func (t *TintPrimer) BrightRed() lipgloss.Color {
	return lipgloss.Color("#fdaeb7")
}

func (t *TintPrimer) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintPrimer) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fff5b1")
}

func (t *TintPrimer) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintPrimer) Blue() lipgloss.Color {
	return lipgloss.Color("#2188ff")
}

func (t *TintPrimer) Cyan() lipgloss.Color {
	return lipgloss.Color("#15e2e2")
}

func (t *TintPrimer) Green() lipgloss.Color {
	return lipgloss.Color("#34d058")
}

func (t *TintPrimer) Purple() lipgloss.Color {
	return lipgloss.Color("#8a63d2")
}

func (t *TintPrimer) Red() lipgloss.Color {
	return lipgloss.Color("#ea4a5a")
}

func (t *TintPrimer) White() lipgloss.Color {
	return lipgloss.Color("#ecf0f1")
}

func (t *TintPrimer) Yellow() lipgloss.Color {
	return lipgloss.Color("#ffdf5d")
}
