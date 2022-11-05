// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintGalaxy (Galaxy) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Galaxy
type TintGalaxy struct{}

// DisplayName returns the display name of the tint.
func (t *TintGalaxy) DisplayName() string {
	return "Galaxy"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintGalaxy) ID() string {
	return "galaxy"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintGalaxy) About() string {
	return `Tint: Galaxy`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintGalaxy) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintGalaxy) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1d2837")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintGalaxy) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintGalaxy) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintGalaxy) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintGalaxy) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#589df6")
}

func (t *TintGalaxy) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#3979bc")
}

func (t *TintGalaxy) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#35bb9a")
}

func (t *TintGalaxy) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#e75699")
}

func (t *TintGalaxy) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fa8c8f")
}

func (t *TintGalaxy) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintGalaxy) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff55")
}

func (t *TintGalaxy) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintGalaxy) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#589df6")
}

func (t *TintGalaxy) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#1f9ee7")
}

func (t *TintGalaxy) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#21b089")
}

func (t *TintGalaxy) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#944d95")
}

func (t *TintGalaxy) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f9555f")
}

func (t *TintGalaxy) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bbbbbb")
}

func (t *TintGalaxy) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fef02a")
}