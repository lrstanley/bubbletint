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
//
// Additional credit to:
//  * jos3s (https://github.com/jos3s)

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintOneStar (OneStar) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#OneStar
//
// Credit to:
//   - jos3s (https://github.com/jos3s)
type TintOneStar struct{}

// DisplayName returns the display name of the tint.
func (t *TintOneStar) DisplayName() string {
	return "OneStar"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintOneStar) ID() string {
	return "one_star"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintOneStar) About() string {
	return `Tint: OneStar
Tint credits:
  * jos3s (https://github.com/jos3s)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintOneStar) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#e4e4e4")
}

// Bg returns the recommended default background color for this tint.
func (t *TintOneStar) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0e0e0e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintOneStar) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintOneStar) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintOneStar) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#666666")
}

func (t *TintOneStar) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#3b8eea")
}

func (t *TintOneStar) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#29b8db")
}

func (t *TintOneStar) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#23d18b")
}

func (t *TintOneStar) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d861d8")
}

func (t *TintOneStar) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fa4b4b")
}

func (t *TintOneStar) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fafafa")
}

func (t *TintOneStar) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fcfc5c")
}

func (t *TintOneStar) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintOneStar) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2472c8")
}

func (t *TintOneStar) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#33a0bb")
}

func (t *TintOneStar) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#0dbc79")
}

func (t *TintOneStar) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c42cc4")
}

func (t *TintOneStar) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d13b3b")
}

func (t *TintOneStar) White() lipgloss.TerminalColor {
	return lipgloss.Color("#f1f1f1")
}

func (t *TintOneStar) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#dfdf44")
}
