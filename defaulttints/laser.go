// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
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

// TintLaser (Laser) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Laser
type TintLaser struct{}

// DisplayName returns the display name of the tint.
func (t *TintLaser) DisplayName() string {
	return "Laser"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintLaser) ID() string {
	return "laser"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintLaser) About() string {
	return `Tint: Laser`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintLaser) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#f106e3")
}

// Bg returns the recommended default background color for this tint.
func (t *TintLaser) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#030d18")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintLaser) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintLaser) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintLaser) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#8f8f8f")
}

func (t *TintLaser) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#f92883")
}

func (t *TintLaser) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#e6e7fe")
}

func (t *TintLaser) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#d6fcba")
}

func (t *TintLaser) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ffb2fe")
}

func (t *TintLaser) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc4be")
}

func (t *TintLaser) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintLaser) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fffed5")
}

func (t *TintLaser) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#626262")
}

func (t *TintLaser) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#fed300")
}

func (t *TintLaser) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#d1d1fe")
}

func (t *TintLaser) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#b4fb73")
}

func (t *TintLaser) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff90fe")
}

func (t *TintLaser) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff8373")
}

func (t *TintLaser) White() lipgloss.TerminalColor {
	return lipgloss.Color("#f1f1f1")
}

func (t *TintLaser) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#09b4bd")
}
