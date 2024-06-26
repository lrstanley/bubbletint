// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// All tints can be previewed here:
//  * https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes
//  * https://github.com/mbadolato/iTerm2-Color-Schemes

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintPiattoLight (Piatto Light) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#PiattoLight
type TintPiattoLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintPiattoLight) DisplayName() string {
	return "Piatto Light"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintPiattoLight) ID() string {
	return "piatto_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintPiattoLight) About() string {
	return `Tint: Piatto Light`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintPiattoLight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#414141")
}

// Bg returns the recommended default background color for this tint.
func (t *TintPiattoLight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintPiattoLight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintPiattoLight) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintPiattoLight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#3f3f3f")
}

func (t *TintPiattoLight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#3c5ea8")
}

func (t *TintPiattoLight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#829429")
}

func (t *TintPiattoLight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#829429")
}

func (t *TintPiattoLight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#a454b2")
}

func (t *TintPiattoLight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#db3365")
}

func (t *TintPiattoLight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f2f2f2")
}

func (t *TintPiattoLight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#cd6f34")
}

func (t *TintPiattoLight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#414141")
}

func (t *TintPiattoLight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#3c5ea8")
}

func (t *TintPiattoLight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#66781e")
}

func (t *TintPiattoLight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#66781e")
}

func (t *TintPiattoLight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a454b2")
}

func (t *TintPiattoLight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#b23771")
}

func (t *TintPiattoLight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintPiattoLight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#cd6f34")
}
