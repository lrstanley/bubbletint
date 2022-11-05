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

// TintFideloper (Fideloper) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Fideloper
type TintFideloper struct{}

// DisplayName returns the display name of the tint.
func (t *TintFideloper) DisplayName() string {
	return "Fideloper"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintFideloper) ID() string {
	return "fideloper"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintFideloper) About() string {
	return `Tint: Fideloper`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintFideloper) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#dbdae0")
}

// Bg returns the recommended default background color for this tint.
func (t *TintFideloper) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#292f33")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintFideloper) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintFideloper) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintFideloper) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#092028")
}

func (t *TintFideloper) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#7c85c4")
}

func (t *TintFideloper) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#819090")
}

func (t *TintFideloper) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#d4605a")
}

func (t *TintFideloper) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#5c5db2")
}

func (t *TintFideloper) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#d4605a")
}

func (t *TintFideloper) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fcf4df")
}

func (t *TintFideloper) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#a86671")
}

func (t *TintFideloper) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#292f33")
}

func (t *TintFideloper) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2e78c2")
}

func (t *TintFideloper) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#309186")
}

func (t *TintFideloper) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#edb8ac")
}

func (t *TintFideloper) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c0236f")
}

func (t *TintFideloper) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#cb1e2d")
}

func (t *TintFideloper) White() lipgloss.TerminalColor {
	return lipgloss.Color("#eae3ce")
}

func (t *TintFideloper) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#b7ab9b")
}