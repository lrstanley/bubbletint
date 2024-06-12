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

// TintEspresso (Espresso) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Espresso
type TintEspresso struct{}

// DisplayName returns the display name of the tint.
func (t *TintEspresso) DisplayName() string {
	return "Espresso"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintEspresso) ID() string {
	return "espresso"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintEspresso) About() string {
	return `Tint: Espresso`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintEspresso) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintEspresso) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#323232")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintEspresso) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintEspresso) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintEspresso) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#535353")
}

func (t *TintEspresso) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#8ab7d9")
}

func (t *TintEspresso) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#dcf4ff")
}

func (t *TintEspresso) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#c2e075")
}

func (t *TintEspresso) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#efb5f7")
}

func (t *TintEspresso) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f00c0c")
}

func (t *TintEspresso) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintEspresso) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e1e48b")
}

func (t *TintEspresso) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#353535")
}

func (t *TintEspresso) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#6c99bb")
}

func (t *TintEspresso) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#bed6ff")
}

func (t *TintEspresso) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#a5c261")
}

func (t *TintEspresso) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#d197d9")
}

func (t *TintEspresso) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d25252")
}

func (t *TintEspresso) White() lipgloss.TerminalColor {
	return lipgloss.Color("#eeeeec")
}

func (t *TintEspresso) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc66d")
}
