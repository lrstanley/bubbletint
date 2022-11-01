// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes
//
// Additional credit to:
//  * jmmv (https://github.com/jmmv)

package tints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintQB64SuperDarkBlue (QB64 Super Dark Blue) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=QB64+Super+Dark+Blue
//
// Credit to:
//   - jmmv (https://github.com/jmmv)
type TintQB64SuperDarkBlue struct{}

// DisplayName returns the display name of the tint.
func (t *TintQB64SuperDarkBlue) DisplayName() string {
	return "QB64 Super Dark Blue"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintQB64SuperDarkBlue) ID() string {
	return "qb_64_super_dark_blue"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintQB64SuperDarkBlue) About() string {
	return `Tint: QB64 Super Dark Blue
Tint credits:
  * jmmv (https://github.com/jmmv)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintQB64SuperDarkBlue) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#D8D8D8")
}

// Bg returns the recommended default background color for this tint.
func (t *TintQB64SuperDarkBlue) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#000027")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintQB64SuperDarkBlue) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.Color("#00586C")
}

// Cursor returns the recommended color for the cursor.
func (t *TintQB64SuperDarkBlue) Cursor() lipgloss.TerminalColor {
	return lipgloss.Color("#D8D8D8")
}

func (t *TintQB64SuperDarkBlue) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#626262")
}

func (t *TintQB64SuperDarkBlue) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#457693")
}

func (t *TintQB64SuperDarkBlue) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00586C")
}

func (t *TintQB64SuperDarkBlue) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#55CE55")
}

func (t *TintQB64SuperDarkBlue) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#934593")
}

func (t *TintQB64SuperDarkBlue) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#D8624E")
}

func (t *TintQB64SuperDarkBlue) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#D8D8D8")
}

func (t *TintQB64SuperDarkBlue) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#FFA700")
}

func (t *TintQB64SuperDarkBlue) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintQB64SuperDarkBlue) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#054663")
}

func (t *TintQB64SuperDarkBlue) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00485C")
}

func (t *TintQB64SuperDarkBlue) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#157E15")
}

func (t *TintQB64SuperDarkBlue) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#631563")
}

func (t *TintQB64SuperDarkBlue) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#98220E")
}

func (t *TintQB64SuperDarkBlue) White() lipgloss.TerminalColor {
	return lipgloss.Color("#989898")
}

func (t *TintQB64SuperDarkBlue) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#808000")
}
