// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes

package tints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintMonokaiSoda (Monokai Soda) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Monokai+Soda
type TintMonokaiSoda struct{}

// DisplayName returns the display name of the tint.
func (t *TintMonokaiSoda) DisplayName() string {
	return "Monokai Soda"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintMonokaiSoda) ID() string {
	return "monokai_soda"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintMonokaiSoda) About() string {
	return `Tint: Monokai Soda`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintMonokaiSoda) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c4c5b5")
}

// Bg returns the recommended default background color for this tint.
func (t *TintMonokaiSoda) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1a1a1a")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintMonokaiSoda) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintMonokaiSoda) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintMonokaiSoda) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#625e4c")
}

func (t *TintMonokaiSoda) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#9d65ff")
}

func (t *TintMonokaiSoda) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#58d1eb")
}

func (t *TintMonokaiSoda) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#98e024")
}

func (t *TintMonokaiSoda) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#f4005f")
}

func (t *TintMonokaiSoda) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f4005f")
}

func (t *TintMonokaiSoda) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f6f6ef")
}

func (t *TintMonokaiSoda) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e0d561")
}

func (t *TintMonokaiSoda) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1a1a1a")
}

func (t *TintMonokaiSoda) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#9d65ff")
}

func (t *TintMonokaiSoda) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#58d1eb")
}

func (t *TintMonokaiSoda) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#98e024")
}

func (t *TintMonokaiSoda) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#f4005f")
}

func (t *TintMonokaiSoda) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f4005f")
}

func (t *TintMonokaiSoda) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c4c5b5")
}

func (t *TintMonokaiSoda) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fa8419")
}
