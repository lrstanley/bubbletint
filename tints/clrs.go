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

// TintCLRS (CLRS) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=CLRS
type TintCLRS struct{}

// DisplayName returns the display name of the tint.
func (t *TintCLRS) DisplayName() string {
	return "CLRS"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintCLRS) ID() string {
	return "clrs"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintCLRS) About() string {
	return `Tint: CLRS`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintCLRS) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#262626")
}

// Bg returns the recommended default background color for this tint.
func (t *TintCLRS) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintCLRS) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintCLRS) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintCLRS) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555753")
}

func (t *TintCLRS) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#1670ff")
}

func (t *TintCLRS) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#3ad5ce")
}

func (t *TintCLRS) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#2cc631")
}

func (t *TintCLRS) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e900b0")
}

func (t *TintCLRS) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fb0416")
}

func (t *TintCLRS) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#eeeeec")
}

func (t *TintCLRS) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fdd727")
}

func (t *TintCLRS) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintCLRS) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#135cd0")
}

func (t *TintCLRS) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#33c3c1")
}

func (t *TintCLRS) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#328a5d")
}

func (t *TintCLRS) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#9f00bd")
}

func (t *TintCLRS) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f8282a")
}

func (t *TintCLRS) White() lipgloss.TerminalColor {
	return lipgloss.Color("#b3b3b3")
}

func (t *TintCLRS) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fa701d")
}
