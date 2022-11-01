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

// TintProLight (Pro Light) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Pro+Light
type TintProLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintProLight) DisplayName() string {
	return "Pro Light"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintProLight) ID() string {
	return "pro_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintProLight) About() string {
	return `Tint: Pro Light`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintProLight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#191919")
}

// Bg returns the recommended default background color for this tint.
func (t *TintProLight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintProLight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintProLight) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintProLight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#9f9f9f")
}

func (t *TintProLight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#0082ff")
}

func (t *TintProLight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#61f7f8")
}

func (t *TintProLight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#61ef57")
}

func (t *TintProLight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff7eff")
}

func (t *TintProLight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6640")
}

func (t *TintProLight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f2f2f2")
}

func (t *TintProLight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f2f156")
}

func (t *TintProLight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintProLight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#3b75ff")
}

func (t *TintProLight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#4ed2de")
}

func (t *TintProLight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#50d148")
}

func (t *TintProLight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ed66e8")
}

func (t *TintProLight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#e5492b")
}

func (t *TintProLight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#dcdcdc")
}

func (t *TintProLight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#c6c440")
}
