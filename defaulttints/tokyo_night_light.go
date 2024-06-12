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

// TintTokyoNightLight (TokyoNightLight) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#TokyoNightLight
type TintTokyoNightLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintTokyoNightLight) DisplayName() string {
	return "TokyoNightLight"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintTokyoNightLight) ID() string {
	return "tokyo_night_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintTokyoNightLight) About() string {
	return `Tint: TokyoNightLight`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintTokyoNightLight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#4c505e")
}

// Bg returns the recommended default background color for this tint.
func (t *TintTokyoNightLight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#cbccd1")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintTokyoNightLight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintTokyoNightLight) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintTokyoNightLight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#0f0f14")
}

func (t *TintTokyoNightLight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#34548a")
}

func (t *TintTokyoNightLight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#0f4b6e")
}

func (t *TintTokyoNightLight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#33635c")
}

func (t *TintTokyoNightLight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#5a4a78")
}

func (t *TintTokyoNightLight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#8c4351")
}

func (t *TintTokyoNightLight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#828594")
}

func (t *TintTokyoNightLight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#8f5e15")
}

func (t *TintTokyoNightLight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#0f0f14")
}

func (t *TintTokyoNightLight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#34548a")
}

func (t *TintTokyoNightLight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#0f4b6e")
}

func (t *TintTokyoNightLight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#33635c")
}

func (t *TintTokyoNightLight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#5a4a78")
}

func (t *TintTokyoNightLight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#8c4351")
}

func (t *TintTokyoNightLight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#828594")
}

func (t *TintTokyoNightLight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#8f5e15")
}
