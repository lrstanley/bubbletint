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

// TintAfterglow (Afterglow) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Afterglow
type TintAfterglow struct{}

// DisplayName returns the display name of the tint.
func (t *TintAfterglow) DisplayName() string {
	return "Afterglow"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintAfterglow) ID() string {
	return "afterglow"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintAfterglow) About() string {
	return `Tint: Afterglow`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintAfterglow) Fg() lipgloss.Color {
	return lipgloss.Color("#d0d0d0")
}

// Bg returns the recommended default background color for this tint.
func (t *TintAfterglow) Bg() lipgloss.Color {
	return lipgloss.Color("#212121")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintAfterglow) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintAfterglow) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintAfterglow) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#505050")
}

func (t *TintAfterglow) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#6c99bb")
}

func (t *TintAfterglow) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#7dd6cf")
}

func (t *TintAfterglow) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#7e8e50")
}

func (t *TintAfterglow) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#9f4e85")
}

func (t *TintAfterglow) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ac4142")
}

func (t *TintAfterglow) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#f5f5f5")
}

func (t *TintAfterglow) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#e5b567")
}

func (t *TintAfterglow) Black() lipgloss.Color {
	return lipgloss.Color("#151515")
}

func (t *TintAfterglow) Blue() lipgloss.Color {
	return lipgloss.Color("#6c99bb")
}

func (t *TintAfterglow) Cyan() lipgloss.Color {
	return lipgloss.Color("#7dd6cf")
}

func (t *TintAfterglow) Green() lipgloss.Color {
	return lipgloss.Color("#7e8e50")
}

func (t *TintAfterglow) Purple() lipgloss.Color {
	return lipgloss.Color("#9f4e85")
}

func (t *TintAfterglow) Red() lipgloss.Color {
	return lipgloss.Color("#ac4142")
}

func (t *TintAfterglow) White() lipgloss.Color {
	return lipgloss.Color("#d0d0d0")
}

func (t *TintAfterglow) Yellow() lipgloss.Color {
	return lipgloss.Color("#e5b567")
}
