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

// TintFunForrest (FunForrest) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#FunForrest
type TintFunForrest struct{}

// DisplayName returns the display name of the tint.
func (t *TintFunForrest) DisplayName() string {
	return "FunForrest"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintFunForrest) ID() string {
	return "fun_forrest"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintFunForrest) About() string {
	return `Tint: FunForrest`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintFunForrest) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#dec165")
}

// Bg returns the recommended default background color for this tint.
func (t *TintFunForrest) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#251200")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintFunForrest) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintFunForrest) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintFunForrest) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7f6a55")
}

func (t *TintFunForrest) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#7cc9cf")
}

func (t *TintFunForrest) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#e6a96b")
}

func (t *TintFunForrest) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#bfc65a")
}

func (t *TintFunForrest) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d26349")
}

func (t *TintFunForrest) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e55a1c")
}

func (t *TintFunForrest) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffeaa3")
}

func (t *TintFunForrest) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcb1b")
}

func (t *TintFunForrest) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintFunForrest) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#4699a3")
}

func (t *TintFunForrest) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#da8213")
}

func (t *TintFunForrest) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#919c00")
}

func (t *TintFunForrest) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#8d4331")
}

func (t *TintFunForrest) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d6262b")
}

func (t *TintFunForrest) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ddc265")
}

func (t *TintFunForrest) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#be8a13")
}
