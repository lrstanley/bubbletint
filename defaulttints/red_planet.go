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

// TintRedPlanet (Red Planet) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#RedPlanet
type TintRedPlanet struct{}

// DisplayName returns the display name of the tint.
func (t *TintRedPlanet) DisplayName() string {
	return "Red Planet"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintRedPlanet) ID() string {
	return "red_planet"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintRedPlanet) About() string {
	return `Tint: Red Planet`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintRedPlanet) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c2b790")
}

// Bg returns the recommended default background color for this tint.
func (t *TintRedPlanet) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#222222")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintRedPlanet) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintRedPlanet) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintRedPlanet) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#676767")
}

func (t *TintRedPlanet) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#60827e")
}

func (t *TintRedPlanet) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#38add8")
}

func (t *TintRedPlanet) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#869985")
}

func (t *TintRedPlanet) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#de4974")
}

func (t *TintRedPlanet) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#b55242")
}

func (t *TintRedPlanet) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#d6bfb8")
}

func (t *TintRedPlanet) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ebeb91")
}

func (t *TintRedPlanet) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#202020")
}

func (t *TintRedPlanet) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#69819e")
}

func (t *TintRedPlanet) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5b8390")
}

func (t *TintRedPlanet) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#728271")
}

func (t *TintRedPlanet) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#896492")
}

func (t *TintRedPlanet) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#8c3432")
}

func (t *TintRedPlanet) White() lipgloss.TerminalColor {
	return lipgloss.Color("#b9aa99")
}

func (t *TintRedPlanet) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e8bf6a")
}
