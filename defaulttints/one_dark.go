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
//
// Additional credit to:
//  * azrikahar (https://github.com/azrikahar)

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintOneDark (OneDark) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#OneDark
//
// Credit to:
//   - azrikahar (https://github.com/azrikahar)
type TintOneDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintOneDark) DisplayName() string {
	return "OneDark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintOneDark) ID() string {
	return "one_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintOneDark) About() string {
	return `Tint: OneDark
Tint credits:
  * azrikahar (https://github.com/azrikahar)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintOneDark) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#abb2bf")
}

// Bg returns the recommended default background color for this tint.
func (t *TintOneDark) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1e2127")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintOneDark) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintOneDark) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintOneDark) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#5c6370")
}

func (t *TintOneDark) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#61afef")
}

func (t *TintOneDark) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#56b6c2")
}

func (t *TintOneDark) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#98c379")
}

func (t *TintOneDark) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c678dd")
}

func (t *TintOneDark) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e06c75")
}

func (t *TintOneDark) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintOneDark) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#d19a66")
}

func (t *TintOneDark) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1e2127")
}

func (t *TintOneDark) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#61afef")
}

func (t *TintOneDark) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#56b6c2")
}

func (t *TintOneDark) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#98c379")
}

func (t *TintOneDark) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c678dd")
}

func (t *TintOneDark) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#e06c75")
}

func (t *TintOneDark) White() lipgloss.TerminalColor {
	return lipgloss.Color("#abb2bf")
}

func (t *TintOneDark) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#d19a66")
}
