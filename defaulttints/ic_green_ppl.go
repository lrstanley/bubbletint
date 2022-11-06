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

// TintICGreenPPL (IC_Green_PPL) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#ICGreenPPL
type TintICGreenPPL struct{}

// DisplayName returns the display name of the tint.
func (t *TintICGreenPPL) DisplayName() string {
	return "IC_Green_PPL"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintICGreenPPL) ID() string {
	return "ic_green_ppl"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintICGreenPPL) About() string {
	return `Tint: IC_Green_PPL`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintICGreenPPL) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#e0f1dc")
}

// Bg returns the recommended default background color for this tint.
func (t *TintICGreenPPL) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#2c2c2c")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintICGreenPPL) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintICGreenPPL) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintICGreenPPL) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#035c03")
}

func (t *TintICGreenPPL) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#2efaeb")
}

func (t *TintICGreenPPL) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#3cfac8")
}

func (t *TintICGreenPPL) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#aefb86")
}

func (t *TintICGreenPPL) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#50fafa")
}

func (t *TintICGreenPPL) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#b4fa5c")
}

func (t *TintICGreenPPL) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e0f1dc")
}

func (t *TintICGreenPPL) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#dafa87")
}

func (t *TintICGreenPPL) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#014401")
}

func (t *TintICGreenPPL) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2ec3b9")
}

func (t *TintICGreenPPL) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#3ca078")
}

func (t *TintICGreenPPL) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#41a638")
}

func (t *TintICGreenPPL) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#50a096")
}

func (t *TintICGreenPPL) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff2736")
}

func (t *TintICGreenPPL) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e6fef2")
}

func (t *TintICGreenPPL) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#76a831")
}
