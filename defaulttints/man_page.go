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

// TintManPage (Man Page) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#ManPage
type TintManPage struct{}

// DisplayName returns the display name of the tint.
func (t *TintManPage) DisplayName() string {
	return "Man Page"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintManPage) ID() string {
	return "man_page"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintManPage) About() string {
	return `Tint: Man Page`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintManPage) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// Bg returns the recommended default background color for this tint.
func (t *TintManPage) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#fef49c")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintManPage) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintManPage) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintManPage) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#666666")
}

func (t *TintManPage) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#0000ff")
}

func (t *TintManPage) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00e5e5")
}

func (t *TintManPage) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#00d900")
}

func (t *TintManPage) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e500e5")
}

func (t *TintManPage) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e50000")
}

func (t *TintManPage) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e5e5")
}

func (t *TintManPage) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e500")
}

func (t *TintManPage) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintManPage) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0000b2")
}

func (t *TintManPage) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00a6b2")
}

func (t *TintManPage) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#00a600")
}

func (t *TintManPage) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#b200b2")
}

func (t *TintManPage) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#cc0000")
}

func (t *TintManPage) White() lipgloss.TerminalColor {
	return lipgloss.Color("#cccccc")
}

func (t *TintManPage) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#999900")
}
