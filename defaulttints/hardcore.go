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

// TintHardcore (Hardcore) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Hardcore
type TintHardcore struct{}

// DisplayName returns the display name of the tint.
func (t *TintHardcore) DisplayName() string {
	return "Hardcore"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintHardcore) ID() string {
	return "hardcore"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintHardcore) About() string {
	return `Tint: Hardcore`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintHardcore) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#a0a0a0")
}

// Bg returns the recommended default background color for this tint.
func (t *TintHardcore) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#121212")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintHardcore) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintHardcore) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintHardcore) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#505354")
}

func (t *TintHardcore) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#66d9ef")
}

func (t *TintHardcore) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#a3babf")
}

func (t *TintHardcore) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#beed5f")
}

func (t *TintHardcore) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#9e6ffe")
}

func (t *TintHardcore) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff669d")
}

func (t *TintHardcore) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f8f8f2")
}

func (t *TintHardcore) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e6db74")
}

func (t *TintHardcore) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1b1d1e")
}

func (t *TintHardcore) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#66d9ef")
}

func (t *TintHardcore) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5e7175")
}

func (t *TintHardcore) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#a6e22e")
}

func (t *TintHardcore) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#9e6ffe")
}

func (t *TintHardcore) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f92672")
}

func (t *TintHardcore) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ccccc6")
}

func (t *TintHardcore) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fd971f")
}
