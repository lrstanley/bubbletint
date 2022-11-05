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

// TintDarkPastel (Dark Pastel) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Dark+Pastel
type TintDarkPastel struct{}

// DisplayName returns the display name of the tint.
func (t *TintDarkPastel) DisplayName() string {
	return "Dark Pastel"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDarkPastel) ID() string {
	return "dark_pastel"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDarkPastel) About() string {
	return `Tint: Dark Pastel`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDarkPastel) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDarkPastel) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDarkPastel) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDarkPastel) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDarkPastel) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintDarkPastel) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#5555ff")
}

func (t *TintDarkPastel) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#55ffff")
}

func (t *TintDarkPastel) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#55ff55")
}

func (t *TintDarkPastel) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff55ff")
}

func (t *TintDarkPastel) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5555")
}

func (t *TintDarkPastel) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintDarkPastel) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff55")
}

func (t *TintDarkPastel) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintDarkPastel) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#5555ff")
}

func (t *TintDarkPastel) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#55ffff")
}

func (t *TintDarkPastel) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#55ff55")
}

func (t *TintDarkPastel) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff55ff")
}

func (t *TintDarkPastel) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5555")
}

func (t *TintDarkPastel) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bbbbbb")
}

func (t *TintDarkPastel) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff55")
}