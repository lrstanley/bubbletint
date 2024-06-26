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
//  * atomiks (https://github.com/atomiks)

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintMoonlightIi (Moonlight II) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#MoonlightIi
//
// Credit to:
//   - atomiks (https://github.com/atomiks)
type TintMoonlightIi struct{}

// DisplayName returns the display name of the tint.
func (t *TintMoonlightIi) DisplayName() string {
	return "Moonlight II"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintMoonlightIi) ID() string {
	return "moonlight_ii"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintMoonlightIi) About() string {
	return `Tint: Moonlight II
Tint credits:
  * atomiks (https://github.com/atomiks)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintMoonlightIi) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c8d3f5")
}

// Bg returns the recommended default background color for this tint.
func (t *TintMoonlightIi) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#222436")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintMoonlightIi) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintMoonlightIi) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintMoonlightIi) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#828bb8")
}

func (t *TintMoonlightIi) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#82aaff")
}

func (t *TintMoonlightIi) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#86e1fc")
}

func (t *TintMoonlightIi) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#c3e88d")
}

func (t *TintMoonlightIi) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c099ff")
}

func (t *TintMoonlightIi) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff757f")
}

func (t *TintMoonlightIi) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#c8d3f5")
}

func (t *TintMoonlightIi) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc777")
}

func (t *TintMoonlightIi) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#191a2a")
}

func (t *TintMoonlightIi) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#82aaff")
}

func (t *TintMoonlightIi) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#86e1fc")
}

func (t *TintMoonlightIi) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#c3e88d")
}

func (t *TintMoonlightIi) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c099ff")
}

func (t *TintMoonlightIi) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff757f")
}

func (t *TintMoonlightIi) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c8d3f5")
}

func (t *TintMoonlightIi) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc777")
}
