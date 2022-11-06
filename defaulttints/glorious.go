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
//
// Additional credit to:
//  * alex (https://github.com/AlexMailo)

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintGlorious (Glorious) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Glorious
//
// Credit to:
//   - alex (https://github.com/AlexMailo)
type TintGlorious struct{}

// DisplayName returns the display name of the tint.
func (t *TintGlorious) DisplayName() string {
	return "Glorious"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintGlorious) ID() string {
	return "glorious"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintGlorious) About() string {
	return `Tint: Glorious
Tint credits:
  * alex (https://github.com/AlexMailo)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintGlorious) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c8c8c9")
}

// Bg returns the recommended default background color for this tint.
func (t *TintGlorious) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#201d27")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintGlorious) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintGlorious) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintGlorious) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7f7061")
}

func (t *TintGlorious) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#4e917c")
}

func (t *TintGlorious) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#7db669")
}

func (t *TintGlorious) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#d5d5d7")
}

func (t *TintGlorious) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c77089")
}

func (t *TintGlorious) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f73028")
}

func (t *TintGlorious) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e6d4a3")
}

func (t *TintGlorious) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f7b125")
}

func (t *TintGlorious) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1e1e1e")
}

func (t *TintGlorious) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#16777a")
}

func (t *TintGlorious) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#578e57")
}

func (t *TintGlorious) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#868715")
}

func (t *TintGlorious) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a04b73")
}

func (t *TintGlorious) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#be0f17")
}

func (t *TintGlorious) White() lipgloss.TerminalColor {
	return lipgloss.Color("#978771")
}

func (t *TintGlorious) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#cc881a")
}
