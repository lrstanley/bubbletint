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

// TintTokyoNight (TokyoNight) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=TokyoNight
type TintTokyoNight struct{}

// DisplayName returns the display name of the tint.
func (t *TintTokyoNight) DisplayName() string {
	return "TokyoNight"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintTokyoNight) ID() string {
	return "tokyo_night"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintTokyoNight) About() string {
	return `Tint: TokyoNight`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintTokyoNight) Fg() lipgloss.Color {
	return lipgloss.Color("#787c99")
}

// Bg returns the recommended default background color for this tint.
func (t *TintTokyoNight) Bg() lipgloss.Color {
	return lipgloss.Color("#16161e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintTokyoNight) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintTokyoNight) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintTokyoNight) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#363b54")
}

func (t *TintTokyoNight) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#7aa2f7")
}

func (t *TintTokyoNight) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#7dcfff")
}

func (t *TintTokyoNight) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#41a6b5")
}

func (t *TintTokyoNight) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#bb9af7")
}

func (t *TintTokyoNight) BrightRed() lipgloss.Color {
	return lipgloss.Color("#f7768e")
}

func (t *TintTokyoNight) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#acb0d0")
}

func (t *TintTokyoNight) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#e0af68")
}

func (t *TintTokyoNight) Black() lipgloss.Color {
	return lipgloss.Color("#363b54")
}

func (t *TintTokyoNight) Blue() lipgloss.Color {
	return lipgloss.Color("#7aa2f7")
}

func (t *TintTokyoNight) Cyan() lipgloss.Color {
	return lipgloss.Color("#7dcfff")
}

func (t *TintTokyoNight) Green() lipgloss.Color {
	return lipgloss.Color("#41a6b5")
}

func (t *TintTokyoNight) Purple() lipgloss.Color {
	return lipgloss.Color("#bb9af7")
}

func (t *TintTokyoNight) Red() lipgloss.Color {
	return lipgloss.Color("#f7768e")
}

func (t *TintTokyoNight) White() lipgloss.Color {
	return lipgloss.Color("#787c99")
}

func (t *TintTokyoNight) Yellow() lipgloss.Color {
	return lipgloss.Color("#e0af68")
}
