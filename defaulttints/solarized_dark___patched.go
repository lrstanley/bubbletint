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

// TintSolarizedDarkPatched (Solarized Dark - Patched) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#SolarizedDarkPatched
type TintSolarizedDarkPatched struct{}

// DisplayName returns the display name of the tint.
func (t *TintSolarizedDarkPatched) DisplayName() string {
	return "Solarized Dark - Patched"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSolarizedDarkPatched) ID() string {
	return "solarized_dark___patched"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSolarizedDarkPatched) About() string {
	return `Tint: Solarized Dark - Patched`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSolarizedDarkPatched) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#708284")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSolarizedDarkPatched) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#001e27")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSolarizedDarkPatched) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSolarizedDarkPatched) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSolarizedDarkPatched) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#475b62")
}

func (t *TintSolarizedDarkPatched) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#708284")
}

func (t *TintSolarizedDarkPatched) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#819090")
}

func (t *TintSolarizedDarkPatched) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#475b62")
}

func (t *TintSolarizedDarkPatched) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#5956ba")
}

func (t *TintSolarizedDarkPatched) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#bd3613")
}

func (t *TintSolarizedDarkPatched) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fcf4dc")
}

func (t *TintSolarizedDarkPatched) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#536870")
}

func (t *TintSolarizedDarkPatched) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#002831")
}

func (t *TintSolarizedDarkPatched) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2176c7")
}

func (t *TintSolarizedDarkPatched) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#259286")
}

func (t *TintSolarizedDarkPatched) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#738a05")
}

func (t *TintSolarizedDarkPatched) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c61c6f")
}

func (t *TintSolarizedDarkPatched) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d11c24")
}

func (t *TintSolarizedDarkPatched) White() lipgloss.TerminalColor {
	return lipgloss.Color("#eae3cb")
}

func (t *TintSolarizedDarkPatched) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#a57706")
}
