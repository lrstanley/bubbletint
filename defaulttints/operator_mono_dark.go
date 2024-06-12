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

// TintOperatorMonoDark (Operator Mono Dark) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#OperatorMonoDark
type TintOperatorMonoDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintOperatorMonoDark) DisplayName() string {
	return "Operator Mono Dark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintOperatorMonoDark) ID() string {
	return "operator_mono_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintOperatorMonoDark) About() string {
	return `Tint: Operator Mono Dark`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintOperatorMonoDark) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c3cac2")
}

// Bg returns the recommended default background color for this tint.
func (t *TintOperatorMonoDark) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#191919")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintOperatorMonoDark) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintOperatorMonoDark) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintOperatorMonoDark) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#9a9b99")
}

func (t *TintOperatorMonoDark) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#89d3f6")
}

func (t *TintOperatorMonoDark) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#82eada")
}

func (t *TintOperatorMonoDark) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#83d0a2")
}

func (t *TintOperatorMonoDark) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff2c7a")
}

func (t *TintOperatorMonoDark) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#c37d62")
}

func (t *TintOperatorMonoDark) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fdfdf6")
}

func (t *TintOperatorMonoDark) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fdfdc5")
}

func (t *TintOperatorMonoDark) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#5a5a5a")
}

func (t *TintOperatorMonoDark) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#4387cf")
}

func (t *TintOperatorMonoDark) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#72d5c6")
}

func (t *TintOperatorMonoDark) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#4d7b3a")
}

func (t *TintOperatorMonoDark) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#b86cb4")
}

func (t *TintOperatorMonoDark) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ca372d")
}

func (t *TintOperatorMonoDark) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ced4cd")
}

func (t *TintOperatorMonoDark) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#d4d697")
}
