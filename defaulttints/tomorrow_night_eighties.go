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

// TintTomorrowNightEighties (Tomorrow Night Eighties) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Tomorrow+Night+Eighties
type TintTomorrowNightEighties struct{}

// DisplayName returns the display name of the tint.
func (t *TintTomorrowNightEighties) DisplayName() string {
	return "Tomorrow Night Eighties"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintTomorrowNightEighties) ID() string {
	return "tomorrow_night_eighties"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintTomorrowNightEighties) About() string {
	return `Tint: Tomorrow Night Eighties`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintTomorrowNightEighties) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#cccccc")
}

// Bg returns the recommended default background color for this tint.
func (t *TintTomorrowNightEighties) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#2d2d2d")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintTomorrowNightEighties) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintTomorrowNightEighties) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintTomorrowNightEighties) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintTomorrowNightEighties) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6699cc")
}

func (t *TintTomorrowNightEighties) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#66cccc")
}

func (t *TintTomorrowNightEighties) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#99cc99")
}

func (t *TintTomorrowNightEighties) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#cc99cc")
}

func (t *TintTomorrowNightEighties) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f2777a")
}

func (t *TintTomorrowNightEighties) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintTomorrowNightEighties) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcc66")
}

func (t *TintTomorrowNightEighties) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintTomorrowNightEighties) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#6699cc")
}

func (t *TintTomorrowNightEighties) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#66cccc")
}

func (t *TintTomorrowNightEighties) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#99cc99")
}

func (t *TintTomorrowNightEighties) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#cc99cc")
}

func (t *TintTomorrowNightEighties) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f2777a")
}

func (t *TintTomorrowNightEighties) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintTomorrowNightEighties) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcc66")
}