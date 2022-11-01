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

// TintPurpleRain (Purple Rain) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Purple+Rain
type TintPurpleRain struct{}

// DisplayName returns the display name of the tint.
func (t *TintPurpleRain) DisplayName() string {
	return "Purple Rain"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintPurpleRain) ID() string {
	return "purple_rain"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintPurpleRain) About() string {
	return `Tint: Purple Rain`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintPurpleRain) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#fffbf6")
}

// Bg returns the recommended default background color for this tint.
func (t *TintPurpleRain) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#21084a")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintPurpleRain) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintPurpleRain) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintPurpleRain) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#565656")
}

func (t *TintPurpleRain) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#00a6ff")
}

func (t *TintPurpleRain) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#74fdf3")
}

func (t *TintPurpleRain) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#b8e36e")
}

func (t *TintPurpleRain) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ac7bf0")
}

func (t *TintPurpleRain) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff4250")
}

func (t *TintPurpleRain) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintPurpleRain) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd852")
}

func (t *TintPurpleRain) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintPurpleRain) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#00a2fa")
}

func (t *TintPurpleRain) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00deef")
}

func (t *TintPurpleRain) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#9be205")
}

func (t *TintPurpleRain) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#815bb5")
}

func (t *TintPurpleRain) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff260e")
}

func (t *TintPurpleRain) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintPurpleRain) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffc400")
}
