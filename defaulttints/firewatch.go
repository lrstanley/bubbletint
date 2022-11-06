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

// TintFirewatch (Firewatch) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Firewatch
type TintFirewatch struct{}

// DisplayName returns the display name of the tint.
func (t *TintFirewatch) DisplayName() string {
	return "Firewatch"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintFirewatch) ID() string {
	return "firewatch"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintFirewatch) About() string {
	return `Tint: Firewatch`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintFirewatch) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#9ba2b2")
}

// Bg returns the recommended default background color for this tint.
func (t *TintFirewatch) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1e2027")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintFirewatch) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintFirewatch) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintFirewatch) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#585f6d")
}

func (t *TintFirewatch) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#4c89c5")
}

func (t *TintFirewatch) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#44a8b6")
}

func (t *TintFirewatch) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#5ab977")
}

func (t *TintFirewatch) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d55119")
}

func (t *TintFirewatch) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#d95360")
}

func (t *TintFirewatch) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e6e5ff")
}

func (t *TintFirewatch) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#dfb563")
}

func (t *TintFirewatch) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#585f6d")
}

func (t *TintFirewatch) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#4d89c4")
}

func (t *TintFirewatch) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#44a8b6")
}

func (t *TintFirewatch) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#5ab977")
}

func (t *TintFirewatch) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#d55119")
}

func (t *TintFirewatch) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d95360")
}

func (t *TintFirewatch) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e6e5ff")
}

func (t *TintFirewatch) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#dfb563")
}
