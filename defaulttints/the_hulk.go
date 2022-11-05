// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintTheHulk (The Hulk) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=The+Hulk
type TintTheHulk struct{}

// DisplayName returns the display name of the tint.
func (t *TintTheHulk) DisplayName() string {
	return "The Hulk"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintTheHulk) ID() string {
	return "the_hulk"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintTheHulk) About() string {
	return `Tint: The Hulk`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintTheHulk) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#b5b5b5")
}

// Bg returns the recommended default background color for this tint.
func (t *TintTheHulk) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1b1d1e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintTheHulk) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintTheHulk) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintTheHulk) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#505354")
}

func (t *TintTheHulk) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#506b95")
}

func (t *TintTheHulk) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#4085a6")
}

func (t *TintTheHulk) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#48ff77")
}

func (t *TintTheHulk) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#72589d")
}

func (t *TintTheHulk) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#8dff2a")
}

func (t *TintTheHulk) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e6e1")
}

func (t *TintTheHulk) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#3afe16")
}

func (t *TintTheHulk) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1b1d1e")
}

func (t *TintTheHulk) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2525f5")
}

func (t *TintTheHulk) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#378ca9")
}

func (t *TintTheHulk) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#13ce30")
}

func (t *TintTheHulk) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#641f74")
}

func (t *TintTheHulk) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#269d1b")
}

func (t *TintTheHulk) White() lipgloss.TerminalColor {
	return lipgloss.Color("#d9d8d1")
}

func (t *TintTheHulk) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#63e457")
}