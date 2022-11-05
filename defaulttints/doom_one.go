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

// TintDoomOne (DoomOne) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=DoomOne
type TintDoomOne struct{}

// DisplayName returns the display name of the tint.
func (t *TintDoomOne) DisplayName() string {
	return "DoomOne"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDoomOne) ID() string {
	return "doom_one"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDoomOne) About() string {
	return `Tint: DoomOne`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDoomOne) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#bbc2cf")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDoomOne) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#282c34")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDoomOne) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDoomOne) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDoomOne) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintDoomOne) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#a9a1e1")
}

func (t *TintDoomOne) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#51afef")
}

func (t *TintDoomOne) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#99bb66")
}

func (t *TintDoomOne) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c678dd")
}

func (t *TintDoomOne) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6655")
}

func (t *TintDoomOne) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#bfbfbf")
}

func (t *TintDoomOne) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ecbe7b")
}

func (t *TintDoomOne) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintDoomOne) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#a9a1e1")
}

func (t *TintDoomOne) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#51afef")
}

func (t *TintDoomOne) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#98be65")
}

func (t *TintDoomOne) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c678dd")
}

func (t *TintDoomOne) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6c6b")
}

func (t *TintDoomOne) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bbc2cf")
}

func (t *TintDoomOne) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ecbe7b")
}