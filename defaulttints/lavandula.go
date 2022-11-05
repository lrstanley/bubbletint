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

// TintLavandula (Lavandula) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Lavandula
type TintLavandula struct{}

// DisplayName returns the display name of the tint.
func (t *TintLavandula) DisplayName() string {
	return "Lavandula"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintLavandula) ID() string {
	return "lavandula"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintLavandula) About() string {
	return `Tint: Lavandula`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintLavandula) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#736e7d")
}

// Bg returns the recommended default background color for this tint.
func (t *TintLavandula) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#050014")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintLavandula) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintLavandula) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintLavandula) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#372d46")
}

func (t *TintLavandula) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#8e87e0")
}

func (t *TintLavandula) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#9ad4e0")
}

func (t *TintLavandula) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#52e0c4")
}

func (t *TintLavandula) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#a776e0")
}

func (t *TintLavandula) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#e05167")
}

func (t *TintLavandula) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#8c91fa")
}

func (t *TintLavandula) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e0c386")
}

func (t *TintLavandula) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#230046")
}

func (t *TintLavandula) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#4f4a7f")
}

func (t *TintLavandula) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#58777f")
}

func (t *TintLavandula) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#337e6f")
}

func (t *TintLavandula) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#5a3f7f")
}

func (t *TintLavandula) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#7d1625")
}

func (t *TintLavandula) White() lipgloss.TerminalColor {
	return lipgloss.Color("#736e7d")
}

func (t *TintLavandula) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#7f6f49")
}