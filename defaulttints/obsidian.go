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

// TintObsidian (Obsidian) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Obsidian
type TintObsidian struct{}

// DisplayName returns the display name of the tint.
func (t *TintObsidian) DisplayName() string {
	return "Obsidian"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintObsidian) ID() string {
	return "obsidian"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintObsidian) About() string {
	return `Tint: Obsidian`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintObsidian) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#cdcdcd")
}

// Bg returns the recommended default background color for this tint.
func (t *TintObsidian) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#283033")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintObsidian) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintObsidian) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintObsidian) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintObsidian) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#a1d7ff")
}

func (t *TintObsidian) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#55ffff")
}

func (t *TintObsidian) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#93c863")
}

func (t *TintObsidian) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff55ff")
}

func (t *TintObsidian) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff0003")
}

func (t *TintObsidian) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintObsidian) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fef874")
}

func (t *TintObsidian) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintObsidian) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#3a9bdb")
}

func (t *TintObsidian) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00bbbb")
}

func (t *TintObsidian) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#00bb00")
}

func (t *TintObsidian) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#bb00bb")
}

func (t *TintObsidian) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#a60001")
}

func (t *TintObsidian) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bbbbbb")
}

func (t *TintObsidian) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fecd22")
}
