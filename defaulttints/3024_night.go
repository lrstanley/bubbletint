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

// Tint3024Night (3024 Night) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#3024Night
type Tint3024Night struct{}

// DisplayName returns the display name of the tint.
func (t *Tint3024Night) DisplayName() string {
	return "3024 Night"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *Tint3024Night) ID() string {
	return "3024_night"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *Tint3024Night) About() string {
	return `Tint: 3024 Night`
}

// Fg returns the recommended default foreground color for this tint.
func (t *Tint3024Night) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#a5a2a2")
}

// Bg returns the recommended default background color for this tint.
func (t *Tint3024Night) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#090300")
}

// SelectionBg returns the recommended background color for selected text.
func (t *Tint3024Night) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *Tint3024Night) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *Tint3024Night) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#5c5855")
}

func (t *Tint3024Night) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#807d7c")
}

func (t *Tint3024Night) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#cdab53")
}

func (t *Tint3024Night) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#3a3432")
}

func (t *Tint3024Night) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d6d5d4")
}

func (t *Tint3024Night) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e8bbd0")
}

func (t *Tint3024Night) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f7f7f7")
}

func (t *Tint3024Night) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#4a4543")
}

func (t *Tint3024Night) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#090300")
}

func (t *Tint3024Night) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#01a0e4")
}

func (t *Tint3024Night) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#b5e4f4")
}

func (t *Tint3024Night) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#01a252")
}

func (t *Tint3024Night) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a16a94")
}

func (t *Tint3024Night) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#db2d20")
}

func (t *Tint3024Night) White() lipgloss.TerminalColor {
	return lipgloss.Color("#a5a2a2")
}

func (t *Tint3024Night) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fded02")
}
