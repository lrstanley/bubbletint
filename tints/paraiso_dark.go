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

// TintParaisoDark (Paraiso Dark) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Paraiso+Dark
type TintParaisoDark struct{}

// DisplayName returns the display name of the tint.
func (t *TintParaisoDark) DisplayName() string {
	return "Paraiso Dark"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintParaisoDark) ID() string {
	return "paraiso_dark"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintParaisoDark) About() string {
	return `Tint: Paraiso Dark`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintParaisoDark) Fg() lipgloss.Color {
	return lipgloss.Color("#a39e9b")
}

// Bg returns the recommended default background color for this tint.
func (t *TintParaisoDark) Bg() lipgloss.Color {
	return lipgloss.Color("#2f1e2e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintParaisoDark) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintParaisoDark) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintParaisoDark) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#776e71")
}

func (t *TintParaisoDark) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#06b6ef")
}

func (t *TintParaisoDark) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#5bc4bf")
}

func (t *TintParaisoDark) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#48b685")
}

func (t *TintParaisoDark) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#815ba4")
}

func (t *TintParaisoDark) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ef6155")
}

func (t *TintParaisoDark) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#e7e9db")
}

func (t *TintParaisoDark) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fec418")
}

func (t *TintParaisoDark) Black() lipgloss.Color {
	return lipgloss.Color("#2f1e2e")
}

func (t *TintParaisoDark) Blue() lipgloss.Color {
	return lipgloss.Color("#06b6ef")
}

func (t *TintParaisoDark) Cyan() lipgloss.Color {
	return lipgloss.Color("#5bc4bf")
}

func (t *TintParaisoDark) Green() lipgloss.Color {
	return lipgloss.Color("#48b685")
}

func (t *TintParaisoDark) Purple() lipgloss.Color {
	return lipgloss.Color("#815ba4")
}

func (t *TintParaisoDark) Red() lipgloss.Color {
	return lipgloss.Color("#ef6155")
}

func (t *TintParaisoDark) White() lipgloss.Color {
	return lipgloss.Color("#a39e9b")
}

func (t *TintParaisoDark) Yellow() lipgloss.Color {
	return lipgloss.Color("#fec418")
}
