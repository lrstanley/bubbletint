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

// TintTomorrowNight (Tomorrow Night) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#TomorrowNight
type TintTomorrowNight struct{}

// DisplayName returns the display name of the tint.
func (t *TintTomorrowNight) DisplayName() string {
	return "Tomorrow Night"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintTomorrowNight) ID() string {
	return "tomorrow_night"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintTomorrowNight) About() string {
	return `Tint: Tomorrow Night`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintTomorrowNight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c5c8c6")
}

// Bg returns the recommended default background color for this tint.
func (t *TintTomorrowNight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1d1f21")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintTomorrowNight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintTomorrowNight) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintTomorrowNight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintTomorrowNight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#81a2be")
}

func (t *TintTomorrowNight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8abeb7")
}

func (t *TintTomorrowNight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#b5bd68")
}

func (t *TintTomorrowNight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#b294bb")
}

func (t *TintTomorrowNight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#cc6666")
}

func (t *TintTomorrowNight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintTomorrowNight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f0c674")
}

func (t *TintTomorrowNight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintTomorrowNight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#81a2be")
}

func (t *TintTomorrowNight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8abeb7")
}

func (t *TintTomorrowNight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#b5bd68")
}

func (t *TintTomorrowNight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#b294bb")
}

func (t *TintTomorrowNight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#cc6666")
}

func (t *TintTomorrowNight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintTomorrowNight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f0c674")
}
