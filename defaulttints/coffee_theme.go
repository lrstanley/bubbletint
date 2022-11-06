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

// TintCoffeeTheme (coffee_theme) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#CoffeeTheme
type TintCoffeeTheme struct{}

// DisplayName returns the display name of the tint.
func (t *TintCoffeeTheme) DisplayName() string {
	return "coffee_theme"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintCoffeeTheme) ID() string {
	return "coffee_theme"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintCoffeeTheme) About() string {
	return `Tint: coffee_theme`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintCoffeeTheme) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// Bg returns the recommended default background color for this tint.
func (t *TintCoffeeTheme) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#f5deb3")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintCoffeeTheme) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintCoffeeTheme) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintCoffeeTheme) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#686868")
}

func (t *TintCoffeeTheme) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6871ff")
}

func (t *TintCoffeeTheme) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#60fdff")
}

func (t *TintCoffeeTheme) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#5ffa68")
}

func (t *TintCoffeeTheme) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff77ff")
}

func (t *TintCoffeeTheme) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6e67")
}

func (t *TintCoffeeTheme) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintCoffeeTheme) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fffc67")
}

func (t *TintCoffeeTheme) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintCoffeeTheme) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0225c7")
}

func (t *TintCoffeeTheme) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00c5c7")
}

func (t *TintCoffeeTheme) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#00c200")
}

func (t *TintCoffeeTheme) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ca30c7")
}

func (t *TintCoffeeTheme) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#c91b00")
}

func (t *TintCoffeeTheme) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c7c7c7")
}

func (t *TintCoffeeTheme) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#c7c400")
}
