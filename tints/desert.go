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

// TintDesert (Desert) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Desert
type TintDesert struct{}

// DisplayName returns the display name of the tint.
func (t *TintDesert) DisplayName() string {
	return "Desert"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDesert) ID() string {
	return "desert"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDesert) About() string {
	return `Tint: Desert`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDesert) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDesert) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#333333")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDesert) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDesert) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDesert) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintDesert) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#87ceff")
}

func (t *TintDesert) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd700")
}

func (t *TintDesert) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#55ff55")
}

func (t *TintDesert) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff55ff")
}

func (t *TintDesert) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5555")
}

func (t *TintDesert) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintDesert) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff55")
}

func (t *TintDesert) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#4d4d4d")
}

func (t *TintDesert) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#cd853f")
}

func (t *TintDesert) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ffa0a0")
}

func (t *TintDesert) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#98fb98")
}

func (t *TintDesert) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ffdead")
}

func (t *TintDesert) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff2b2b")
}

func (t *TintDesert) White() lipgloss.TerminalColor {
	return lipgloss.Color("#f5deb3")
}

func (t *TintDesert) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f0e68c")
}
