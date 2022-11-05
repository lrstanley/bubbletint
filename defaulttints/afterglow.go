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

// TintAfterglow (Afterglow) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Afterglow
type TintAfterglow struct{}

// DisplayName returns the display name of the tint.
func (t *TintAfterglow) DisplayName() string {
	return "Afterglow"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintAfterglow) ID() string {
	return "afterglow"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintAfterglow) About() string {
	return `Tint: Afterglow`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintAfterglow) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#d0d0d0")
}

// Bg returns the recommended default background color for this tint.
func (t *TintAfterglow) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#212121")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintAfterglow) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintAfterglow) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintAfterglow) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#505050")
}

func (t *TintAfterglow) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6c99bb")
}

func (t *TintAfterglow) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#7dd6cf")
}

func (t *TintAfterglow) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#7e8e50")
}

func (t *TintAfterglow) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#9f4e85")
}

func (t *TintAfterglow) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ac4142")
}

func (t *TintAfterglow) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f5f5f5")
}

func (t *TintAfterglow) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5b567")
}

func (t *TintAfterglow) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#151515")
}

func (t *TintAfterglow) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#6c99bb")
}

func (t *TintAfterglow) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#7dd6cf")
}

func (t *TintAfterglow) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#7e8e50")
}

func (t *TintAfterglow) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#9f4e85")
}

func (t *TintAfterglow) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ac4142")
}

func (t *TintAfterglow) White() lipgloss.TerminalColor {
	return lipgloss.Color("#d0d0d0")
}

func (t *TintAfterglow) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5b567")
}