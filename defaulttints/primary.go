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

// TintPrimary (primary) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=primary
type TintPrimary struct{}

// DisplayName returns the display name of the tint.
func (t *TintPrimary) DisplayName() string {
	return "primary"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintPrimary) ID() string {
	return "primary"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintPrimary) About() string {
	return `Tint: primary`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintPrimary) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// Bg returns the recommended default background color for this tint.
func (t *TintPrimary) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintPrimary) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintPrimary) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintPrimary) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintPrimary) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#4285f4")
}

func (t *TintPrimary) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#0f9d58")
}

func (t *TintPrimary) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#0f9d58")
}

func (t *TintPrimary) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#4285f4")
}

func (t *TintPrimary) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#db4437")
}

func (t *TintPrimary) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintPrimary) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f4b400")
}

func (t *TintPrimary) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintPrimary) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#4285f4")
}

func (t *TintPrimary) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#4285f4")
}

func (t *TintPrimary) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#0f9d58")
}

func (t *TintPrimary) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#db4437")
}

func (t *TintPrimary) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#db4437")
}

func (t *TintPrimary) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintPrimary) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f4b400")
}