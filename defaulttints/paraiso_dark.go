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

// TintParaisoDark (Paraiso Dark) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#ParaisoDark
type TintParaisoDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintParaisoDark) DisplayName() string {
	return "Paraiso Dark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintParaisoDark) ID() string {
	return "paraiso_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintParaisoDark) About() string {
	return `Tint: Paraiso Dark`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintParaisoDark) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#a39e9b")
}

// Bg returns the recommended default background color for this tint.
func (t *TintParaisoDark) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#2f1e2e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintParaisoDark) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintParaisoDark) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintParaisoDark) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#776e71")
}

func (t *TintParaisoDark) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#06b6ef")
}

func (t *TintParaisoDark) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5bc4bf")
}

func (t *TintParaisoDark) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#48b685")
}

func (t *TintParaisoDark) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#815ba4")
}

func (t *TintParaisoDark) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ef6155")
}

func (t *TintParaisoDark) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e7e9db")
}

func (t *TintParaisoDark) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fec418")
}

func (t *TintParaisoDark) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#2f1e2e")
}

func (t *TintParaisoDark) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#06b6ef")
}

func (t *TintParaisoDark) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5bc4bf")
}

func (t *TintParaisoDark) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#48b685")
}

func (t *TintParaisoDark) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#815ba4")
}

func (t *TintParaisoDark) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ef6155")
}

func (t *TintParaisoDark) White() lipgloss.TerminalColor {
	return lipgloss.Color("#a39e9b")
}

func (t *TintParaisoDark) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fec418")
}
