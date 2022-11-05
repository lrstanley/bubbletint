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

// TintHipsterGreen (Hipster Green) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Hipster+Green
type TintHipsterGreen struct{}

// DisplayName returns the display name of the tint.
func (t *TintHipsterGreen) DisplayName() string {
	return "Hipster Green"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintHipsterGreen) ID() string {
	return "hipster_green"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintHipsterGreen) About() string {
	return `Tint: Hipster Green`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintHipsterGreen) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#84c138")
}

// Bg returns the recommended default background color for this tint.
func (t *TintHipsterGreen) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#100b05")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintHipsterGreen) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintHipsterGreen) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintHipsterGreen) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#666666")
}

func (t *TintHipsterGreen) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#0000ff")
}

func (t *TintHipsterGreen) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00e5e5")
}

func (t *TintHipsterGreen) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#86a93e")
}

func (t *TintHipsterGreen) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e500e5")
}

func (t *TintHipsterGreen) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e50000")
}

func (t *TintHipsterGreen) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e5e5")
}

func (t *TintHipsterGreen) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5e500")
}

func (t *TintHipsterGreen) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintHipsterGreen) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#246eb2")
}

func (t *TintHipsterGreen) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00a6b2")
}

func (t *TintHipsterGreen) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#00a600")
}

func (t *TintHipsterGreen) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#b200b2")
}

func (t *TintHipsterGreen) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#b6214a")
}

func (t *TintHipsterGreen) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bfbfbf")
}

func (t *TintHipsterGreen) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#bfbf00")
}