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

// TintSeaShells (SeaShells) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#SeaShells
type TintSeaShells struct{}

// DisplayName returns the display name of the tint.
func (t *TintSeaShells) DisplayName() string {
	return "SeaShells"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSeaShells) ID() string {
	return "sea_shells"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSeaShells) About() string {
	return `Tint: SeaShells`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSeaShells) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#deb88d")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSeaShells) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#09141b")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSeaShells) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSeaShells) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSeaShells) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#434b53")
}

func (t *TintSeaShells) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#1bbcdd")
}

func (t *TintSeaShells) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#87acb4")
}

func (t *TintSeaShells) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#628d98")
}

func (t *TintSeaShells) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#bbe3ee")
}

func (t *TintSeaShells) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#d48678")
}

func (t *TintSeaShells) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fee4ce")
}

func (t *TintSeaShells) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fdd39f")
}

func (t *TintSeaShells) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#17384c")
}

func (t *TintSeaShells) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#1e4950")
}

func (t *TintSeaShells) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#50a3b5")
}

func (t *TintSeaShells) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#027c9b")
}

func (t *TintSeaShells) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#68d4f1")
}

func (t *TintSeaShells) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d15123")
}

func (t *TintSeaShells) White() lipgloss.TerminalColor {
	return lipgloss.Color("#deb88d")
}

func (t *TintSeaShells) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fca02f")
}
