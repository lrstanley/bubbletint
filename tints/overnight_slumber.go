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

// TintOvernightSlumber (Overnight Slumber) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Overnight+Slumber
type TintOvernightSlumber struct{}

// DisplayName returns the display name of the tint.
func (t *TintOvernightSlumber) DisplayName() string {
	return "Overnight Slumber"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintOvernightSlumber) ID() string {
	return "overnight_slumber"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintOvernightSlumber) About() string {
	return `Tint: Overnight Slumber`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintOvernightSlumber) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ced2d6")
}

// Bg returns the recommended default background color for this tint.
func (t *TintOvernightSlumber) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0e1729")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintOvernightSlumber) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintOvernightSlumber) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintOvernightSlumber) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#575656")
}

func (t *TintOvernightSlumber) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#8dabe1")
}

func (t *TintOvernightSlumber) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ffa7c4")
}

func (t *TintOvernightSlumber) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#85cc95")
}

func (t *TintOvernightSlumber) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c792eb")
}

func (t *TintOvernightSlumber) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ffa7c4")
}

func (t *TintOvernightSlumber) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintOvernightSlumber) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcb8b")
}

func (t *TintOvernightSlumber) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#0a1222")
}

func (t *TintOvernightSlumber) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#8dabe1")
}

func (t *TintOvernightSlumber) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#78ccf0")
}

func (t *TintOvernightSlumber) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#85cc95")
}

func (t *TintOvernightSlumber) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c792eb")
}

func (t *TintOvernightSlumber) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ffa7c4")
}

func (t *TintOvernightSlumber) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintOvernightSlumber) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcb8b")
}
