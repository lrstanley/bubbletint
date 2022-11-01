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

// TintParasioDark (Parasio Dark) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Parasio+Dark
type TintParasioDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintParasioDark) DisplayName() string {
	return "Parasio Dark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintParasioDark) ID() string {
	return "parasio_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintParasioDark) About() string {
	return `Tint: Parasio Dark`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintParasioDark) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#a39e9b")
}

// Bg returns the recommended default background color for this tint.
func (t *TintParasioDark) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#2f1e2e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintParasioDark) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintParasioDark) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintParasioDark) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#776e71")
}

func (t *TintParasioDark) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#06b6ef")
}

func (t *TintParasioDark) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5bc4bf")
}

func (t *TintParasioDark) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#48b685")
}

func (t *TintParasioDark) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#815ba4")
}

func (t *TintParasioDark) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ef6155")
}

func (t *TintParasioDark) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e7e9db")
}

func (t *TintParasioDark) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fec418")
}

func (t *TintParasioDark) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#2f1e2e")
}

func (t *TintParasioDark) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#06b6ef")
}

func (t *TintParasioDark) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5bc4bf")
}

func (t *TintParasioDark) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#48b685")
}

func (t *TintParasioDark) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#815ba4")
}

func (t *TintParasioDark) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ef6155")
}

func (t *TintParasioDark) White() lipgloss.TerminalColor {
	return lipgloss.Color("#a39e9b")
}

func (t *TintParasioDark) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fec418")
}
