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

// TintSnazzy (Snazzy) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Snazzy
type TintSnazzy struct{}

// DisplayName returns the display name of the tint.
func (t *TintSnazzy) DisplayName() string {
	return "Snazzy"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSnazzy) ID() string {
	return "snazzy"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSnazzy) About() string {
	return `Tint: Snazzy`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSnazzy) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ebece6")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSnazzy) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1e1f29")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSnazzy) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSnazzy) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSnazzy) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintSnazzy) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#49baff")
}

func (t *TintSnazzy) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8be9fe")
}

func (t *TintSnazzy) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#50fb7c")
}

func (t *TintSnazzy) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#fc4cb4")
}

func (t *TintSnazzy) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fc4346")
}

func (t *TintSnazzy) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ededec")
}

func (t *TintSnazzy) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f0fb8c")
}

func (t *TintSnazzy) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintSnazzy) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#49baff")
}

func (t *TintSnazzy) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8be9fe")
}

func (t *TintSnazzy) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#50fb7c")
}

func (t *TintSnazzy) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#fc4cb4")
}

func (t *TintSnazzy) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#fc4346")
}

func (t *TintSnazzy) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ededec")
}

func (t *TintSnazzy) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f0fb8c")
}
