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

// TintJackieBrown (Jackie Brown) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#JackieBrown
type TintJackieBrown struct{}

// DisplayName returns the display name of the tint.
func (t *TintJackieBrown) DisplayName() string {
	return "Jackie Brown"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintJackieBrown) ID() string {
	return "jackie_brown"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintJackieBrown) About() string {
	return `Tint: Jackie Brown`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintJackieBrown) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcc2f")
}

// Bg returns the recommended default background color for this tint.
func (t *TintJackieBrown) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#2c1d16")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintJackieBrown) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintJackieBrown) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintJackieBrown) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#666666")
}

func (t *TintJackieBrown) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#0000ff")
}

func (t *TintJackieBrown) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00e5e5")
}

func (t *TintJackieBrown) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#86a93e")
}

func (t *TintJackieBrown) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e500e5")
}

func (t *TintJackieBrown) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e50000")
}

func (t *TintJackieBrown) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e5e5")
}

func (t *TintJackieBrown) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e500")
}

func (t *TintJackieBrown) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#2c1d16")
}

func (t *TintJackieBrown) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#246eb2")
}

func (t *TintJackieBrown) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00acee")
}

func (t *TintJackieBrown) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#2baf2b")
}

func (t *TintJackieBrown) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#d05ec1")
}

func (t *TintJackieBrown) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ef5734")
}

func (t *TintJackieBrown) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bfbfbf")
}

func (t *TintJackieBrown) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#bebf00")
}
