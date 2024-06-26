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

// TintBrightLights (Bright Lights) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#BrightLights
type TintBrightLights struct{}

// DisplayName returns the display name of the tint.
func (t *TintBrightLights) DisplayName() string {
	return "Bright Lights"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBrightLights) ID() string {
	return "bright_lights"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBrightLights) About() string {
	return `Tint: Bright Lights`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBrightLights) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#b3c9d7")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBrightLights) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#191919")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBrightLights) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintBrightLights) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintBrightLights) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#191919")
}

func (t *TintBrightLights) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#76d5ff")
}

func (t *TintBrightLights) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#6cbfb5")
}

func (t *TintBrightLights) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#b7e876")
}

func (t *TintBrightLights) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ba76e7")
}

func (t *TintBrightLights) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff355b")
}

func (t *TintBrightLights) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#c2c8d7")
}

func (t *TintBrightLights) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc251")
}

func (t *TintBrightLights) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#191919")
}

func (t *TintBrightLights) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#76d4ff")
}

func (t *TintBrightLights) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#6cbfb5")
}

func (t *TintBrightLights) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#b7e876")
}

func (t *TintBrightLights) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ba76e7")
}

func (t *TintBrightLights) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff355b")
}

func (t *TintBrightLights) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c2c8d7")
}

func (t *TintBrightLights) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc251")
}
