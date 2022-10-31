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

// TintElementary (Elementary) is a collection of lipgloss styles.
type TintElementary struct{}

// DisplayName returns the display name of the tint.
func (t *TintElementary) DisplayName() string {
	return "Elementary"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintElementary) ID() string {
	return "elementary"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintElementary) About() string {
	return `Tint: Elementary`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintElementary) Fg() lipgloss.Color {
	return lipgloss.Color("#efefef")
}

// Bg returns the recommended default background color for this tint.
func (t *TintElementary) Bg() lipgloss.Color {
	return lipgloss.Color("#181818")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintElementary) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintElementary) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintElementary) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#4b4b4b")
}

func (t *TintElementary) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#0955ff")
}

func (t *TintElementary) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#3ea8fc")
}

func (t *TintElementary) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#6bc219")
}

func (t *TintElementary) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#fb0050")
}

func (t *TintElementary) BrightRed() lipgloss.Color {
	return lipgloss.Color("#fc1c18")
}

func (t *TintElementary) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#8c00ec")
}

func (t *TintElementary) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fec80e")
}

func (t *TintElementary) Black() lipgloss.Color {
	return lipgloss.Color("#242424")
}

func (t *TintElementary) Blue() lipgloss.Color {
	return lipgloss.Color("#063b8c")
}

func (t *TintElementary) Cyan() lipgloss.Color {
	return lipgloss.Color("#2595e1")
}

func (t *TintElementary) Green() lipgloss.Color {
	return lipgloss.Color("#5aa513")
}

func (t *TintElementary) Purple() lipgloss.Color {
	return lipgloss.Color("#e40038")
}

func (t *TintElementary) Red() lipgloss.Color {
	return lipgloss.Color("#d71c15")
}

func (t *TintElementary) White() lipgloss.Color {
	return lipgloss.Color("#efefef")
}

func (t *TintElementary) Yellow() lipgloss.Color {
	return lipgloss.Color("#fdb40c")
}
