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

// TintKolorit (Kolorit) is a collection of lipgloss styles.
type TintKolorit struct{}

// DisplayName returns the display name of the tint.
func (t *TintKolorit) DisplayName() string {
	return "Kolorit"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintKolorit) ID() string {
	return "kolorit"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintKolorit) About() string {
	return `Tint: Kolorit`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintKolorit) Fg() lipgloss.Color {
	return lipgloss.Color("#efecec")
}

// Bg returns the recommended default background color for this tint.
func (t *TintKolorit) Bg() lipgloss.Color {
	return lipgloss.Color("#1d1a1e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintKolorit) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintKolorit) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintKolorit) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#1d1a1e")
}

func (t *TintKolorit) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#5db4ee")
}

func (t *TintKolorit) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#57e9eb")
}

func (t *TintKolorit) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#47d7a1")
}

func (t *TintKolorit) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#da6cda")
}

func (t *TintKolorit) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ff5b82")
}

func (t *TintKolorit) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ededed")
}

func (t *TintKolorit) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#e8e562")
}

func (t *TintKolorit) Black() lipgloss.Color {
	return lipgloss.Color("#1d1a1e")
}

func (t *TintKolorit) Blue() lipgloss.Color {
	return lipgloss.Color("#5db4ee")
}

func (t *TintKolorit) Cyan() lipgloss.Color {
	return lipgloss.Color("#57e9eb")
}

func (t *TintKolorit) Green() lipgloss.Color {
	return lipgloss.Color("#47d7a1")
}

func (t *TintKolorit) Purple() lipgloss.Color {
	return lipgloss.Color("#da6cda")
}

func (t *TintKolorit) Red() lipgloss.Color {
	return lipgloss.Color("#ff5b82")
}

func (t *TintKolorit) White() lipgloss.Color {
	return lipgloss.Color("#ededed")
}

func (t *TintKolorit) Yellow() lipgloss.Color {
	return lipgloss.Color("#e8e562")
}
