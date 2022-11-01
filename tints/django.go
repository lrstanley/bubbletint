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

// TintDjango (Django) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Django
type TintDjango struct{}

// DisplayName returns the display name of the tint.
func (t *TintDjango) DisplayName() string {
	return "Django"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDjango) ID() string {
	return "django"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDjango) About() string {
	return `Tint: Django`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDjango) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#f8f8f8")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDjango) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0b2f20")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDjango) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDjango) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDjango) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#323232")
}

func (t *TintDjango) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#568264")
}

func (t *TintDjango) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#cfffd1")
}

func (t *TintDjango) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#73da70")
}

func (t *TintDjango) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintDjango) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff943b")
}

func (t *TintDjango) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintDjango) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff94")
}

func (t *TintDjango) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintDjango) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#245032")
}

func (t *TintDjango) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#9df39f")
}

func (t *TintDjango) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#41a83e")
}

func (t *TintDjango) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#f8f8f8")
}

func (t *TintDjango) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#fd6209")
}

func (t *TintDjango) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintDjango) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffe862")
}
