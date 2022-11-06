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

// TintHybrid (Hybrid) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Hybrid
type TintHybrid struct{}

// DisplayName returns the display name of the tint.
func (t *TintHybrid) DisplayName() string {
	return "Hybrid"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintHybrid) ID() string {
	return "hybrid"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintHybrid) About() string {
	return `Tint: Hybrid`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintHybrid) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#b7bcba")
}

// Bg returns the recommended default background color for this tint.
func (t *TintHybrid) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#161719")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintHybrid) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintHybrid) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintHybrid) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#1d1f22")
}

func (t *TintHybrid) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#4b6b88")
}

func (t *TintHybrid) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#4d7b74")
}

func (t *TintHybrid) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#798431")
}

func (t *TintHybrid) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#6e5079")
}

func (t *TintHybrid) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#8d2e32")
}

func (t *TintHybrid) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#5a626a")
}

func (t *TintHybrid) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e58a50")
}

func (t *TintHybrid) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#2a2e33")
}

func (t *TintHybrid) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#6e90b0")
}

func (t *TintHybrid) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#7fbfb4")
}

func (t *TintHybrid) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#b3bf5a")
}

func (t *TintHybrid) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a17eac")
}

func (t *TintHybrid) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#b84d51")
}

func (t *TintHybrid) White() lipgloss.TerminalColor {
	return lipgloss.Color("#b5b9b6")
}

func (t *TintHybrid) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e4b55e")
}
