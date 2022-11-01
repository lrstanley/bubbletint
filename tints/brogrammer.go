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

// TintBrogrammer (Brogrammer) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Brogrammer
type TintBrogrammer struct{}

// DisplayName returns the display name of the tint.
func (t *TintBrogrammer) DisplayName() string {
	return "Brogrammer"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBrogrammer) ID() string {
	return "brogrammer"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBrogrammer) About() string {
	return `Tint: Brogrammer`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBrogrammer) Fg() lipgloss.Color {
	return lipgloss.Color("#d6dbe5")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBrogrammer) Bg() lipgloss.Color {
	return lipgloss.Color("#131313")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBrogrammer) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintBrogrammer) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintBrogrammer) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#d6dbe5")
}

func (t *TintBrogrammer) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#1081d6")
}

func (t *TintBrogrammer) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#0f7ddb")
}

func (t *TintBrogrammer) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#1dd361")
}

func (t *TintBrogrammer) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#5350b9")
}

func (t *TintBrogrammer) BrightRed() lipgloss.Color {
	return lipgloss.Color("#de352e")
}

func (t *TintBrogrammer) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintBrogrammer) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#f3bd09")
}

func (t *TintBrogrammer) Black() lipgloss.Color {
	return lipgloss.Color("#1f1f1f")
}

func (t *TintBrogrammer) Blue() lipgloss.Color {
	return lipgloss.Color("#2a84d2")
}

func (t *TintBrogrammer) Cyan() lipgloss.Color {
	return lipgloss.Color("#1081d6")
}

func (t *TintBrogrammer) Green() lipgloss.Color {
	return lipgloss.Color("#2dc55e")
}

func (t *TintBrogrammer) Purple() lipgloss.Color {
	return lipgloss.Color("#4e5ab7")
}

func (t *TintBrogrammer) Red() lipgloss.Color {
	return lipgloss.Color("#f81118")
}

func (t *TintBrogrammer) White() lipgloss.Color {
	return lipgloss.Color("#d6dbe5")
}

func (t *TintBrogrammer) Yellow() lipgloss.Color {
	return lipgloss.Color("#ecba0f")
}
