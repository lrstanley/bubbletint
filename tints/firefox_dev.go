// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes

package tints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintFirefoxDev (FirefoxDev) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=FirefoxDev
type TintFirefoxDev struct{}

// DisplayName returns the display name of the tint.
func (t *TintFirefoxDev) DisplayName() string {
	return "FirefoxDev"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintFirefoxDev) ID() string {
	return "firefox_dev"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintFirefoxDev) About() string {
	return `Tint: FirefoxDev`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintFirefoxDev) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#7c8fa4")
}

// Bg returns the recommended default background color for this tint.
func (t *TintFirefoxDev) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0e1011")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintFirefoxDev) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintFirefoxDev) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintFirefoxDev) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#001e27")
}

func (t *TintFirefoxDev) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#006fc0")
}

func (t *TintFirefoxDev) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#005794")
}

func (t *TintFirefoxDev) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#1d9000")
}

func (t *TintFirefoxDev) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#a200da")
}

func (t *TintFirefoxDev) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e1003f")
}

func (t *TintFirefoxDev) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e2e2e2")
}

func (t *TintFirefoxDev) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#cd9409")
}

func (t *TintFirefoxDev) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#002831")
}

func (t *TintFirefoxDev) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#359ddf")
}

func (t *TintFirefoxDev) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#4b73a2")
}

func (t *TintFirefoxDev) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#5eb83c")
}

func (t *TintFirefoxDev) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#d75cff")
}

func (t *TintFirefoxDev) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#e63853")
}

func (t *TintFirefoxDev) White() lipgloss.TerminalColor {
	return lipgloss.Color("#dcdcdc")
}

func (t *TintFirefoxDev) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#a57706")
}
