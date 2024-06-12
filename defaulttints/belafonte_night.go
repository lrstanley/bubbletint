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

// TintBelafonteNight (Belafonte Night) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#BelafonteNight
type TintBelafonteNight struct{}

// DisplayName returns the display name of the tint.
func (t *TintBelafonteNight) DisplayName() string {
	return "Belafonte Night"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBelafonteNight) ID() string {
	return "belafonte_night"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBelafonteNight) About() string {
	return `Tint: Belafonte Night`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBelafonteNight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#968c83")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBelafonteNight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#20111b")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBelafonteNight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintBelafonteNight) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintBelafonteNight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#5e5252")
}

func (t *TintBelafonteNight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#426a79")
}

func (t *TintBelafonteNight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#989a9c")
}

func (t *TintBelafonteNight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#858162")
}

func (t *TintBelafonteNight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#97522c")
}

func (t *TintBelafonteNight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#be100e")
}

func (t *TintBelafonteNight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#d5ccba")
}

func (t *TintBelafonteNight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#eaa549")
}

func (t *TintBelafonteNight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#20111b")
}

func (t *TintBelafonteNight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#426a79")
}

func (t *TintBelafonteNight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#989a9c")
}

func (t *TintBelafonteNight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#858162")
}

func (t *TintBelafonteNight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#97522c")
}

func (t *TintBelafonteNight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#be100e")
}

func (t *TintBelafonteNight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#968c83")
}

func (t *TintBelafonteNight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#eaa549")
}
