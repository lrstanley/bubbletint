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

// TintSolarizedDarkHigherContrast (Solarized Dark Higher Contrast) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#SolarizedDarkHigherContrast
type TintSolarizedDarkHigherContrast struct{}

// DisplayName returns the display name of the tint.
func (t *TintSolarizedDarkHigherContrast) DisplayName() string {
	return "Solarized Dark Higher Contrast"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSolarizedDarkHigherContrast) ID() string {
	return "solarized_dark_higher_contrast"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSolarizedDarkHigherContrast) About() string {
	return `Tint: Solarized Dark Higher Contrast`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSolarizedDarkHigherContrast) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#9cc2c3")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSolarizedDarkHigherContrast) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#001e27")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSolarizedDarkHigherContrast) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSolarizedDarkHigherContrast) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSolarizedDarkHigherContrast) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#006488")
}

func (t *TintSolarizedDarkHigherContrast) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#178ec8")
}

func (t *TintSolarizedDarkHigherContrast) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00b39e")
}

func (t *TintSolarizedDarkHigherContrast) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#51ef84")
}

func (t *TintSolarizedDarkHigherContrast) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e24d8e")
}

func (t *TintSolarizedDarkHigherContrast) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#f5163b")
}

func (t *TintSolarizedDarkHigherContrast) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fcf4dc")
}

func (t *TintSolarizedDarkHigherContrast) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#b27e28")
}

func (t *TintSolarizedDarkHigherContrast) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#002831")
}

func (t *TintSolarizedDarkHigherContrast) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2176c7")
}

func (t *TintSolarizedDarkHigherContrast) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#259286")
}

func (t *TintSolarizedDarkHigherContrast) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#6cbe6c")
}

func (t *TintSolarizedDarkHigherContrast) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c61c6f")
}

func (t *TintSolarizedDarkHigherContrast) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d11c24")
}

func (t *TintSolarizedDarkHigherContrast) White() lipgloss.TerminalColor {
	return lipgloss.Color("#eae3cb")
}

func (t *TintSolarizedDarkHigherContrast) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#a57706")
}
