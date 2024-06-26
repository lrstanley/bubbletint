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

// TintSubliminal (Subliminal) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Subliminal
type TintSubliminal struct{}

// DisplayName returns the display name of the tint.
func (t *TintSubliminal) DisplayName() string {
	return "Subliminal"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSubliminal) ID() string {
	return "subliminal"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSubliminal) About() string {
	return `Tint: Subliminal`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSubliminal) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#d4d4d4")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSubliminal) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#282c35")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSubliminal) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSubliminal) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSubliminal) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7f7f7f")
}

func (t *TintSubliminal) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6699cc")
}

func (t *TintSubliminal) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5fb3b3")
}

func (t *TintSubliminal) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#a9cfa4")
}

func (t *TintSubliminal) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#f1a5ab")
}

func (t *TintSubliminal) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e15a60")
}

func (t *TintSubliminal) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#d4d4d4")
}

func (t *TintSubliminal) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffe2a9")
}

func (t *TintSubliminal) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#7f7f7f")
}

func (t *TintSubliminal) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#6699cc")
}

func (t *TintSubliminal) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5fb3b3")
}

func (t *TintSubliminal) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#a9cfa4")
}

func (t *TintSubliminal) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#f1a5ab")
}

func (t *TintSubliminal) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#e15a60")
}

func (t *TintSubliminal) White() lipgloss.TerminalColor {
	return lipgloss.Color("#d4d4d4")
}

func (t *TintSubliminal) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffe2a9")
}
