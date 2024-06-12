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

// TintHighway (Highway) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Highway
type TintHighway struct{}

// DisplayName returns the display name of the tint.
func (t *TintHighway) DisplayName() string {
	return "Highway"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintHighway) ID() string {
	return "highway"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintHighway) About() string {
	return `Tint: Highway`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintHighway) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ededed")
}

// Bg returns the recommended default background color for this tint.
func (t *TintHighway) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#222225")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintHighway) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintHighway) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintHighway) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#5d504a")
}

func (t *TintHighway) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#4fc2fd")
}

func (t *TintHighway) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5d504a")
}

func (t *TintHighway) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#b1d130")
}

func (t *TintHighway) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#de0071")
}

func (t *TintHighway) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f07e18")
}

func (t *TintHighway) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintHighway) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fff120")
}

func (t *TintHighway) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintHighway) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#006bb3")
}

func (t *TintHighway) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#384564")
}

func (t *TintHighway) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#138034")
}

func (t *TintHighway) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#6b2775")
}

func (t *TintHighway) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d00e18")
}

func (t *TintHighway) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ededed")
}

func (t *TintHighway) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcb3e")
}
