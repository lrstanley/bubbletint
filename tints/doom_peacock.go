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

// TintDoomPeacock (Doom Peacock) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Doom+Peacock
type TintDoomPeacock struct{}

// DisplayName returns the display name of the tint.
func (t *TintDoomPeacock) DisplayName() string {
	return "Doom Peacock"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDoomPeacock) ID() string {
	return "doom_peacock"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDoomPeacock) About() string {
	return `Tint: Doom Peacock`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDoomPeacock) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ede0ce")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDoomPeacock) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#2b2a27")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDoomPeacock) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDoomPeacock) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDoomPeacock) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#2b2a27")
}

func (t *TintDoomPeacock) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#51afef")
}

func (t *TintDoomPeacock) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#46d9ff")
}

func (t *TintDoomPeacock) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#98be65")
}

func (t *TintDoomPeacock) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c678dd")
}

func (t *TintDoomPeacock) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5d38")
}

func (t *TintDoomPeacock) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#dfdfdf")
}

func (t *TintDoomPeacock) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e6f972")
}

func (t *TintDoomPeacock) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1c1f24")
}

func (t *TintDoomPeacock) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2a6cc6")
}

func (t *TintDoomPeacock) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5699af")
}

func (t *TintDoomPeacock) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#26a6a6")
}

func (t *TintDoomPeacock) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a9a1e1")
}

func (t *TintDoomPeacock) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#cb4b16")
}

func (t *TintDoomPeacock) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ede0ce")
}

func (t *TintDoomPeacock) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#bcd42a")
}
