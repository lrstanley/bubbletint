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

// TintSpring (Spring) is a collection of lipgloss styles.
type TintSpring struct{}

// DisplayName returns the display name of the tint.
func (t *TintSpring) DisplayName() string {
	return "Spring"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSpring) ID() string {
	return "spring"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSpring) About() string {
	return `Tint: Spring`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSpring) Fg() lipgloss.Color {
	return lipgloss.Color("#4d4d4c")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSpring) Bg() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSpring) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintSpring) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintSpring) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintSpring) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#15a9fd")
}

func (t *TintSpring) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#3e999f")
}

func (t *TintSpring) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#1fc231")
}

func (t *TintSpring) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#8959a8")
}

func (t *TintSpring) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ff0021")
}

func (t *TintSpring) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintSpring) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#d5b807")
}

func (t *TintSpring) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintSpring) Blue() lipgloss.Color {
	return lipgloss.Color("#1dd3ee")
}

func (t *TintSpring) Cyan() lipgloss.Color {
	return lipgloss.Color("#3e999f")
}

func (t *TintSpring) Green() lipgloss.Color {
	return lipgloss.Color("#1f8c3b")
}

func (t *TintSpring) Purple() lipgloss.Color {
	return lipgloss.Color("#8959a8")
}

func (t *TintSpring) Red() lipgloss.Color {
	return lipgloss.Color("#ff4d83")
}

func (t *TintSpring) White() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintSpring) Yellow() lipgloss.Color {
	return lipgloss.Color("#1fc95b")
}
