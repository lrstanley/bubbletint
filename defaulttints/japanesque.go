// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
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

// TintJapanesque (Japanesque) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Japanesque
type TintJapanesque struct{}

// DisplayName returns the display name of the tint.
func (t *TintJapanesque) DisplayName() string {
	return "Japanesque"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintJapanesque) ID() string {
	return "japanesque"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintJapanesque) About() string {
	return `Tint: Japanesque`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintJapanesque) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#f7f6ec")
}

// Bg returns the recommended default background color for this tint.
func (t *TintJapanesque) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1e1e1e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintJapanesque) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintJapanesque) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintJapanesque) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#595b59")
}

func (t *TintJapanesque) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#135979")
}

func (t *TintJapanesque) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#76bbca")
}

func (t *TintJapanesque) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#767f2c")
}

func (t *TintJapanesque) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#604291")
}

func (t *TintJapanesque) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#d18fa6")
}

func (t *TintJapanesque) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#b2b5ae")
}

func (t *TintJapanesque) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#78592f")
}

func (t *TintJapanesque) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#343935")
}

func (t *TintJapanesque) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#4c9ad4")
}

func (t *TintJapanesque) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#389aad")
}

func (t *TintJapanesque) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#7bb75b")
}

func (t *TintJapanesque) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a57fc4")
}

func (t *TintJapanesque) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#cf3f61")
}

func (t *TintJapanesque) White() lipgloss.TerminalColor {
	return lipgloss.Color("#fafaf6")
}

func (t *TintJapanesque) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e9b32a")
}
