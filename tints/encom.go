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

// TintENCOM (ENCOM) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=ENCOM
type TintENCOM struct{}

// DisplayName returns the display name of the tint.
func (t *TintENCOM) DisplayName() string {
	return "ENCOM"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintENCOM) ID() string {
	return "encom"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintENCOM) About() string {
	return `Tint: ENCOM`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintENCOM) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#00a595")
}

// Bg returns the recommended default background color for this tint.
func (t *TintENCOM) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintENCOM) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintENCOM) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintENCOM) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintENCOM) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#0000ff")
}

func (t *TintENCOM) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00cdcd")
}

func (t *TintENCOM) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#00ee00")
}

func (t *TintENCOM) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff00ff")
}

func (t *TintENCOM) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff0000")
}

func (t *TintENCOM) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintENCOM) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff00")
}

func (t *TintENCOM) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintENCOM) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0081ff")
}

func (t *TintENCOM) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#008b8b")
}

func (t *TintENCOM) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#008b00")
}

func (t *TintENCOM) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#bc00ca")
}

func (t *TintENCOM) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#9f0000")
}

func (t *TintENCOM) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bbbbbb")
}

func (t *TintENCOM) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd000")
}
