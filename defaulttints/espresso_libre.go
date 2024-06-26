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

// TintEspressoLibre (Espresso Libre) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#EspressoLibre
type TintEspressoLibre struct{}

// DisplayName returns the display name of the tint.
func (t *TintEspressoLibre) DisplayName() string {
	return "Espresso Libre"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintEspressoLibre) ID() string {
	return "espresso_libre"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintEspressoLibre) About() string {
	return `Tint: Espresso Libre`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintEspressoLibre) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#b8a898")
}

// Bg returns the recommended default background color for this tint.
func (t *TintEspressoLibre) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#2a211c")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintEspressoLibre) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintEspressoLibre) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintEspressoLibre) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555753")
}

func (t *TintEspressoLibre) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#43a8ed")
}

func (t *TintEspressoLibre) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#34e2e2")
}

func (t *TintEspressoLibre) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#9aff87")
}

func (t *TintEspressoLibre) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff818a")
}

func (t *TintEspressoLibre) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ef2929")
}

func (t *TintEspressoLibre) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#eeeeec")
}

func (t *TintEspressoLibre) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fffb5c")
}

func (t *TintEspressoLibre) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintEspressoLibre) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0066ff")
}

func (t *TintEspressoLibre) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#06989a")
}

func (t *TintEspressoLibre) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#1a921c")
}

func (t *TintEspressoLibre) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c5656b")
}

func (t *TintEspressoLibre) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#cc0000")
}

func (t *TintEspressoLibre) White() lipgloss.TerminalColor {
	return lipgloss.Color("#d3d7cf")
}

func (t *TintEspressoLibre) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f0e53a")
}
