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

// TintGlacier (Glacier) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Glacier
type TintGlacier struct{}

// DisplayName returns the display name of the tint.
func (t *TintGlacier) DisplayName() string {
	return "Glacier"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintGlacier) ID() string {
	return "glacier"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintGlacier) About() string {
	return `Tint: Glacier`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintGlacier) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintGlacier) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0c1115")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintGlacier) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintGlacier) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintGlacier) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#404a55")
}

func (t *TintGlacier) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#2a8bc1")
}

func (t *TintGlacier) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#a0b6d3")
}

func (t *TintGlacier) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#49e998")
}

func (t *TintGlacier) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ea4727")
}

func (t *TintGlacier) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#bd0f2f")
}

func (t *TintGlacier) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintGlacier) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fddf6e")
}

func (t *TintGlacier) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#2e343c")
}

func (t *TintGlacier) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#1f5872")
}

func (t *TintGlacier) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#778397")
}

func (t *TintGlacier) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#35a770")
}

func (t *TintGlacier) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#bd2523")
}

func (t *TintGlacier) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#bd0f2f")
}

func (t *TintGlacier) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintGlacier) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fb9435")
}
