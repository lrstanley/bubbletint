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

// TintOneHalfLight (OneHalfLight) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#OneHalfLight
type TintOneHalfLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintOneHalfLight) DisplayName() string {
	return "OneHalfLight"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintOneHalfLight) ID() string {
	return "one_half_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintOneHalfLight) About() string {
	return `Tint: OneHalfLight`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintOneHalfLight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#383a42")
}

// Bg returns the recommended default background color for this tint.
func (t *TintOneHalfLight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#fafafa")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintOneHalfLight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintOneHalfLight) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintOneHalfLight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#4f525e")
}

func (t *TintOneHalfLight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#61afef")
}

func (t *TintOneHalfLight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#56b6c2")
}

func (t *TintOneHalfLight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#98c379")
}

func (t *TintOneHalfLight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c678dd")
}

func (t *TintOneHalfLight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e06c75")
}

func (t *TintOneHalfLight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintOneHalfLight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5c07b")
}

func (t *TintOneHalfLight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#383a42")
}

func (t *TintOneHalfLight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0184bc")
}

func (t *TintOneHalfLight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#0997b3")
}

func (t *TintOneHalfLight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#50a14f")
}

func (t *TintOneHalfLight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a626a4")
}

func (t *TintOneHalfLight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#e45649")
}

func (t *TintOneHalfLight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#fafafa")
}

func (t *TintOneHalfLight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#c18401")
}
