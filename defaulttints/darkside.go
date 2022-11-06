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

// TintDarkside (Darkside) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Darkside
type TintDarkside struct{}

// DisplayName returns the display name of the tint.
func (t *TintDarkside) DisplayName() string {
	return "Darkside"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDarkside) ID() string {
	return "darkside"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDarkside) About() string {
	return `Tint: Darkside`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDarkside) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#bababa")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDarkside) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#222324")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDarkside) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDarkside) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDarkside) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintDarkside) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#387cd3")
}

func (t *TintDarkside) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#3d97e2")
}

func (t *TintDarkside) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#77b869")
}

func (t *TintDarkside) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#957bbe")
}

func (t *TintDarkside) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e05a4f")
}

func (t *TintDarkside) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#bababa")
}

func (t *TintDarkside) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#efd64b")
}

func (t *TintDarkside) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintDarkside) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#1c98e8")
}

func (t *TintDarkside) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#1c98e8")
}

func (t *TintDarkside) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#68c256")
}

func (t *TintDarkside) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#8e69c9")
}

func (t *TintDarkside) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#e8341c")
}

func (t *TintDarkside) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bababa")
}

func (t *TintDarkside) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f2d42c")
}
