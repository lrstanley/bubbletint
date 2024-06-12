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

// TintCyberdyne (Cyberdyne) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Cyberdyne
type TintCyberdyne struct{}

// DisplayName returns the display name of the tint.
func (t *TintCyberdyne) DisplayName() string {
	return "Cyberdyne"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintCyberdyne) ID() string {
	return "cyberdyne"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintCyberdyne) About() string {
	return `Tint: Cyberdyne`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintCyberdyne) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#00ff92")
}

// Bg returns the recommended default background color for this tint.
func (t *TintCyberdyne) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#151144")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintCyberdyne) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintCyberdyne) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintCyberdyne) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#2e2e2e")
}

func (t *TintCyberdyne) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#c2e3ff")
}

func (t *TintCyberdyne) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#e6e7fe")
}

func (t *TintCyberdyne) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#d6fcba")
}

func (t *TintCyberdyne) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ffb2fe")
}

func (t *TintCyberdyne) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc4be")
}

func (t *TintCyberdyne) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintCyberdyne) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fffed5")
}

func (t *TintCyberdyne) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#080808")
}

func (t *TintCyberdyne) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0071cf")
}

func (t *TintCyberdyne) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#6bffdd")
}

func (t *TintCyberdyne) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#00c172")
}

func (t *TintCyberdyne) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff90fe")
}

func (t *TintCyberdyne) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff8373")
}

func (t *TintCyberdyne) White() lipgloss.TerminalColor {
	return lipgloss.Color("#f1f1f1")
}

func (t *TintCyberdyne) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#d2a700")
}
