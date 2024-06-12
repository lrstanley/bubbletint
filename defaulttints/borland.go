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

// TintBorland (Borland) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Borland
type TintBorland struct{}

// DisplayName returns the display name of the tint.
func (t *TintBorland) DisplayName() string {
	return "Borland"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBorland) ID() string {
	return "borland"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBorland) About() string {
	return `Tint: Borland`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBorland) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff4e")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBorland) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0000a4")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBorland) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintBorland) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintBorland) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7c7c7c")
}

func (t *TintBorland) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#b5dcff")
}

func (t *TintBorland) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#dfdffe")
}

func (t *TintBorland) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#ceffac")
}

func (t *TintBorland) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff9cfe")
}

func (t *TintBorland) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ffb6b0")
}

func (t *TintBorland) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintBorland) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffcc")
}

func (t *TintBorland) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#4f4f4f")
}

func (t *TintBorland) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#96cbfe")
}

func (t *TintBorland) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#c6c5fe")
}

func (t *TintBorland) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#a8ff60")
}

func (t *TintBorland) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff73fd")
}

func (t *TintBorland) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6c60")
}

func (t *TintBorland) White() lipgloss.TerminalColor {
	return lipgloss.Color("#eeeeee")
}

func (t *TintBorland) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffb6")
}
