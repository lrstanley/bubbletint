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

// TintKonsolas (Konsolas) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Konsolas
type TintKonsolas struct{}

// DisplayName returns the display name of the tint.
func (t *TintKonsolas) DisplayName() string {
	return "Konsolas"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintKonsolas) ID() string {
	return "konsolas"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintKonsolas) About() string {
	return `Tint: Konsolas`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintKonsolas) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c8c1c1")
}

// Bg returns the recommended default background color for this tint.
func (t *TintKonsolas) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#060606")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintKonsolas) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintKonsolas) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintKonsolas) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#7b716e")
}

func (t *TintKonsolas) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#4b4bff")
}

func (t *TintKonsolas) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#69ffff")
}

func (t *TintKonsolas) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#5fff5f")
}

func (t *TintKonsolas) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff54ff")
}

func (t *TintKonsolas) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff4141")
}

func (t *TintKonsolas) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintKonsolas) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff55")
}

func (t *TintKonsolas) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintKonsolas) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2323a5")
}

func (t *TintKonsolas) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#42b0c8")
}

func (t *TintKonsolas) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#18b218")
}

func (t *TintKonsolas) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ad1edc")
}

func (t *TintKonsolas) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#aa1717")
}

func (t *TintKonsolas) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c8c1c1")
}

func (t *TintKonsolas) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ebae1f")
}
