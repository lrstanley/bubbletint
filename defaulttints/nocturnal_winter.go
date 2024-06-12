// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// All tints can be previewed here:
//  * https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes
//  * https://github.com/mbadolato/iTerm2-Color-Schemes

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintNocturnalWinter (Nocturnal Winter) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#NocturnalWinter
type TintNocturnalWinter struct{}

// DisplayName returns the display name of the tint.
func (t *TintNocturnalWinter) DisplayName() string {
	return "Nocturnal Winter"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintNocturnalWinter) ID() string {
	return "nocturnal_winter"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintNocturnalWinter) About() string {
	return `Tint: Nocturnal Winter`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintNocturnalWinter) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#e6e5e5")
}

// Bg returns the recommended default background color for this tint.
func (t *TintNocturnalWinter) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0d0d17")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintNocturnalWinter) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintNocturnalWinter) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintNocturnalWinter) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#808080")
}

func (t *TintNocturnalWinter) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6096ff")
}

func (t *TintNocturnalWinter) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#0ae78d")
}

func (t *TintNocturnalWinter) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#0ae78d")
}

func (t *TintNocturnalWinter) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff78a2")
}

func (t *TintNocturnalWinter) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f16d86")
}

func (t *TintNocturnalWinter) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintNocturnalWinter) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fffc67")
}

func (t *TintNocturnalWinter) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#4d4d4d")
}

func (t *TintNocturnalWinter) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#3182e0")
}

func (t *TintNocturnalWinter) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#09c87a")
}

func (t *TintNocturnalWinter) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#09cd7e")
}

func (t *TintNocturnalWinter) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff2b6d")
}

func (t *TintNocturnalWinter) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f12d52")
}

func (t *TintNocturnalWinter) White() lipgloss.TerminalColor {
	return lipgloss.Color("#fcfcfc")
}

func (t *TintNocturnalWinter) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f5f17a")
}
