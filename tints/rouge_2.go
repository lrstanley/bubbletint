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

// TintRouge2 (Rouge 2) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Rouge+2
type TintRouge2 struct{}

// DisplayName returns the display name of the tint.
func (t *TintRouge2) DisplayName() string {
	return "Rouge 2"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintRouge2) ID() string {
	return "rouge_2"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintRouge2) About() string {
	return `Tint: Rouge 2`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintRouge2) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#a2a3aa")
}

// Bg returns the recommended default background color for this tint.
func (t *TintRouge2) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#17182b")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintRouge2) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintRouge2) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintRouge2) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#616274")
}

func (t *TintRouge2) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#98b3cd")
}

func (t *TintRouge2) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#abcbd3")
}

func (t *TintRouge2) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#e6dcc4")
}

func (t *TintRouge2) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#8283a1")
}

func (t *TintRouge2) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#c6797e")
}

func (t *TintRouge2) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e8e8ea")
}

func (t *TintRouge2) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e6dcc4")
}

func (t *TintRouge2) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#5d5d6b")
}

func (t *TintRouge2) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#6e94b9")
}

func (t *TintRouge2) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8ab6c1")
}

func (t *TintRouge2) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#969e92")
}

func (t *TintRouge2) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#4c4e78")
}

func (t *TintRouge2) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#c6797e")
}

func (t *TintRouge2) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e8e8ea")
}

func (t *TintRouge2) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#dbcdab")
}
