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

// TintNightOwlishLight (Night Owlish Light) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#NightOwlishLight
type TintNightOwlishLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintNightOwlishLight) DisplayName() string {
	return "Night Owlish Light"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintNightOwlishLight) ID() string {
	return "night_owlish_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintNightOwlishLight) About() string {
	return `Tint: Night Owlish Light`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintNightOwlishLight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#403f53")
}

// Bg returns the recommended default background color for this tint.
func (t *TintNightOwlishLight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintNightOwlishLight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintNightOwlishLight) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintNightOwlishLight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7a8181")
}

func (t *TintNightOwlishLight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#5ca7e4")
}

func (t *TintNightOwlishLight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00c990")
}

func (t *TintNightOwlishLight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#49d0c5")
}

func (t *TintNightOwlishLight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#697098")
}

func (t *TintNightOwlishLight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f76e6e")
}

func (t *TintNightOwlishLight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#989fb1")
}

func (t *TintNightOwlishLight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#dac26b")
}

func (t *TintNightOwlishLight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#011627")
}

func (t *TintNightOwlishLight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#4876d6")
}

func (t *TintNightOwlishLight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#08916a")
}

func (t *TintNightOwlishLight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#2aa298")
}

func (t *TintNightOwlishLight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#403f53")
}

func (t *TintNightOwlishLight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d3423e")
}

func (t *TintNightOwlishLight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#7a8181")
}

func (t *TintNightOwlishLight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#daaa01")
}
