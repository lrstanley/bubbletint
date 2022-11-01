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

// TintCiapre (Ciapre) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Ciapre
type TintCiapre struct{}

// DisplayName returns the display name of the tint.
func (t *TintCiapre) DisplayName() string {
	return "Ciapre"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintCiapre) ID() string {
	return "ciapre"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintCiapre) About() string {
	return `Tint: Ciapre`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintCiapre) Fg() lipgloss.Color {
	return lipgloss.Color("#aea47a")
}

// Bg returns the recommended default background color for this tint.
func (t *TintCiapre) Bg() lipgloss.Color {
	return lipgloss.Color("#191c27")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintCiapre) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintCiapre) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintCiapre) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#555555")
}

func (t *TintCiapre) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#3097c6")
}

func (t *TintCiapre) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#f3dbb2")
}

func (t *TintCiapre) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#a6a75d")
}

func (t *TintCiapre) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#d33061")
}

func (t *TintCiapre) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ac3835")
}

func (t *TintCiapre) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#f4f4f4")
}

func (t *TintCiapre) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#dcdf7c")
}

func (t *TintCiapre) Black() lipgloss.Color {
	return lipgloss.Color("#181818")
}

func (t *TintCiapre) Blue() lipgloss.Color {
	return lipgloss.Color("#576d8c")
}

func (t *TintCiapre) Cyan() lipgloss.Color {
	return lipgloss.Color("#5c4f4b")
}

func (t *TintCiapre) Green() lipgloss.Color {
	return lipgloss.Color("#48513b")
}

func (t *TintCiapre) Purple() lipgloss.Color {
	return lipgloss.Color("#724d7c")
}

func (t *TintCiapre) Red() lipgloss.Color {
	return lipgloss.Color("#810009")
}

func (t *TintCiapre) White() lipgloss.Color {
	return lipgloss.Color("#aea47f")
}

func (t *TintCiapre) Yellow() lipgloss.Color {
	return lipgloss.Color("#cc8b3f")
}
