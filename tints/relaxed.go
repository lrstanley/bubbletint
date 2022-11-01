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

// TintRelaxed (Relaxed) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Relaxed
type TintRelaxed struct{}

// DisplayName returns the display name of the tint.
func (t *TintRelaxed) DisplayName() string {
	return "Relaxed"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintRelaxed) ID() string {
	return "relaxed"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintRelaxed) About() string {
	return `Tint: Relaxed`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintRelaxed) Fg() lipgloss.Color {
	return lipgloss.Color("#d9d9d9")
}

// Bg returns the recommended default background color for this tint.
func (t *TintRelaxed) Bg() lipgloss.Color {
	return lipgloss.Color("#353a44")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintRelaxed) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintRelaxed) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintRelaxed) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#636363")
}

func (t *TintRelaxed) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#7eaac7")
}

func (t *TintRelaxed) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#acbbd0")
}

func (t *TintRelaxed) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#a0ac77")
}

func (t *TintRelaxed) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#b06698")
}

func (t *TintRelaxed) BrightRed() lipgloss.Color {
	return lipgloss.Color("#bc5653")
}

func (t *TintRelaxed) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#f7f7f7")
}

func (t *TintRelaxed) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#ebc17a")
}

func (t *TintRelaxed) Black() lipgloss.Color {
	return lipgloss.Color("#151515")
}

func (t *TintRelaxed) Blue() lipgloss.Color {
	return lipgloss.Color("#6a8799")
}

func (t *TintRelaxed) Cyan() lipgloss.Color {
	return lipgloss.Color("#c9dfff")
}

func (t *TintRelaxed) Green() lipgloss.Color {
	return lipgloss.Color("#909d63")
}

func (t *TintRelaxed) Purple() lipgloss.Color {
	return lipgloss.Color("#b06698")
}

func (t *TintRelaxed) Red() lipgloss.Color {
	return lipgloss.Color("#bc5653")
}

func (t *TintRelaxed) White() lipgloss.Color {
	return lipgloss.Color("#d9d9d9")
}

func (t *TintRelaxed) Yellow() lipgloss.Color {
	return lipgloss.Color("#ebc17a")
}
