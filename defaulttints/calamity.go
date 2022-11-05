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

// TintCalamity (Calamity) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Calamity
type TintCalamity struct{}

// DisplayName returns the display name of the tint.
func (t *TintCalamity) DisplayName() string {
	return "Calamity"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintCalamity) ID() string {
	return "calamity"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintCalamity) About() string {
	return `Tint: Calamity`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintCalamity) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#d5ced9")
}

// Bg returns the recommended default background color for this tint.
func (t *TintCalamity) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#2f2833")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintCalamity) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintCalamity) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintCalamity) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7e6c88")
}

func (t *TintCalamity) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#3b79c7")
}

func (t *TintCalamity) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#74d3de")
}

func (t *TintCalamity) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#a5f69c")
}

func (t *TintCalamity) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#f92672")
}

func (t *TintCalamity) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fc644d")
}

func (t *TintCalamity) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintCalamity) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e9d7a5")
}

func (t *TintCalamity) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#2f2833")
}

func (t *TintCalamity) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#3b79c7")
}

func (t *TintCalamity) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#74d3de")
}

func (t *TintCalamity) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#a5f69c")
}

func (t *TintCalamity) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#f92672")
}

func (t *TintCalamity) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#fc644d")
}

func (t *TintCalamity) White() lipgloss.TerminalColor {
	return lipgloss.Color("#d5ced9")
}

func (t *TintCalamity) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e9d7a5")
}