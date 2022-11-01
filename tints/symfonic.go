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

// TintSymfonic (Symfonic) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Symfonic
type TintSymfonic struct{}

// DisplayName returns the display name of the tint.
func (t *TintSymfonic) DisplayName() string {
	return "Symfonic"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSymfonic) ID() string {
	return "symfonic"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSymfonic) About() string {
	return `Tint: Symfonic`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSymfonic) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSymfonic) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSymfonic) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSymfonic) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSymfonic) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#1b1d21")
}

func (t *TintSymfonic) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#0084d4")
}

func (t *TintSymfonic) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ccccff")
}

func (t *TintSymfonic) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#56db3a")
}

func (t *TintSymfonic) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#b729d9")
}

func (t *TintSymfonic) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#dc322f")
}

func (t *TintSymfonic) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintSymfonic) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ff8400")
}

func (t *TintSymfonic) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintSymfonic) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0084d4")
}

func (t *TintSymfonic) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ccccff")
}

func (t *TintSymfonic) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#56db3a")
}

func (t *TintSymfonic) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#b729d9")
}

func (t *TintSymfonic) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#dc322f")
}

func (t *TintSymfonic) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintSymfonic) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ff8400")
}
