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

// TintWildCherry (WildCherry) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#WildCherry
type TintWildCherry struct{}

// DisplayName returns the display name of the tint.
func (t *TintWildCherry) DisplayName() string {
	return "WildCherry"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintWildCherry) ID() string {
	return "wild_cherry"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintWildCherry) About() string {
	return `Tint: WildCherry`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintWildCherry) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#dafaff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintWildCherry) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1f1726")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintWildCherry) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintWildCherry) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintWildCherry) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#009cc9")
}

func (t *TintWildCherry) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#308cba")
}

func (t *TintWildCherry) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ff919d")
}

func (t *TintWildCherry) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#f4dca5")
}

func (t *TintWildCherry) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ae636b")
}

func (t *TintWildCherry) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#da6bac")
}

func (t *TintWildCherry) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e4838d")
}

func (t *TintWildCherry) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#eac066")
}

func (t *TintWildCherry) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000507")
}

func (t *TintWildCherry) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#883cdc")
}

func (t *TintWildCherry) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#c1b8b7")
}

func (t *TintWildCherry) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#2ab250")
}

func (t *TintWildCherry) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ececec")
}

func (t *TintWildCherry) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d94085")
}

func (t *TintWildCherry) White() lipgloss.TerminalColor {
	return lipgloss.Color("#fff8de")
}

func (t *TintWildCherry) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd16f")
}
