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

// TintMolokai (Molokai) is a collection of lipgloss styles.
type TintMolokai struct{}

// DisplayName returns the display name of the tint.
func (t *TintMolokai) DisplayName() string {
	return "Molokai"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintMolokai) ID() string {
	return "molokai"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintMolokai) About() string {
	return `Tint: Molokai`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintMolokai) Fg() lipgloss.Color {
	return lipgloss.Color("#bbbbbb")
}

// Bg returns the recommended default background color for this tint.
func (t *TintMolokai) Bg() lipgloss.Color {
	return lipgloss.Color("#121212")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintMolokai) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintMolokai) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintMolokai) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#555555")
}

func (t *TintMolokai) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#00afff")
}

func (t *TintMolokai) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#51ceff")
}

func (t *TintMolokai) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#b1e05f")
}

func (t *TintMolokai) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#af87ff")
}

func (t *TintMolokai) BrightRed() lipgloss.Color {
	return lipgloss.Color("#f6669d")
}

func (t *TintMolokai) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintMolokai) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fff26d")
}

func (t *TintMolokai) Black() lipgloss.Color {
	return lipgloss.Color("#121212")
}

func (t *TintMolokai) Blue() lipgloss.Color {
	return lipgloss.Color("#1080d0")
}

func (t *TintMolokai) Cyan() lipgloss.Color {
	return lipgloss.Color("#43a8d0")
}

func (t *TintMolokai) Green() lipgloss.Color {
	return lipgloss.Color("#98e123")
}

func (t *TintMolokai) Purple() lipgloss.Color {
	return lipgloss.Color("#8700ff")
}

func (t *TintMolokai) Red() lipgloss.Color {
	return lipgloss.Color("#fa2573")
}

func (t *TintMolokai) White() lipgloss.Color {
	return lipgloss.Color("#bbbbbb")
}

func (t *TintMolokai) Yellow() lipgloss.Color {
	return lipgloss.Color("#dfd460")
}
