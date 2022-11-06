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

// TintBreeze (Breeze) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Breeze
type TintBreeze struct{}

// DisplayName returns the display name of the tint.
func (t *TintBreeze) DisplayName() string {
	return "Breeze"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBreeze) ID() string {
	return "breeze"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBreeze) About() string {
	return `Tint: Breeze`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBreeze) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#eff0f1")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBreeze) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#31363b")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBreeze) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintBreeze) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintBreeze) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7f8c8d")
}

func (t *TintBreeze) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#3daee9")
}

func (t *TintBreeze) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#16a085")
}

func (t *TintBreeze) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#1cdc9a")
}

func (t *TintBreeze) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#8e44ad")
}

func (t *TintBreeze) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#c0392b")
}

func (t *TintBreeze) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fcfcfc")
}

func (t *TintBreeze) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fdbc4b")
}

func (t *TintBreeze) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#31363b")
}

func (t *TintBreeze) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#1d99f3")
}

func (t *TintBreeze) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#1abc9c")
}

func (t *TintBreeze) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#11d116")
}

func (t *TintBreeze) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#9b59b6")
}

func (t *TintBreeze) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ed1515")
}

func (t *TintBreeze) White() lipgloss.TerminalColor {
	return lipgloss.Color("#eff0f1")
}

func (t *TintBreeze) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f67400")
}
