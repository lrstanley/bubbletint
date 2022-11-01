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

// TintHurtado (Hurtado) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Hurtado
type TintHurtado struct{}

// DisplayName returns the display name of the tint.
func (t *TintHurtado) DisplayName() string {
	return "Hurtado"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintHurtado) ID() string {
	return "hurtado"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintHurtado) About() string {
	return `Tint: Hurtado`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintHurtado) Fg() lipgloss.Color {
	return lipgloss.Color("#dbdbdb")
}

// Bg returns the recommended default background color for this tint.
func (t *TintHurtado) Bg() lipgloss.Color {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintHurtado) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintHurtado) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintHurtado) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#262626")
}

func (t *TintHurtado) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#89beff")
}

func (t *TintHurtado) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#86eafe")
}

func (t *TintHurtado) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#a5df55")
}

func (t *TintHurtado) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#c001c1")
}

func (t *TintHurtado) BrightRed() lipgloss.Color {
	return lipgloss.Color("#d51d00")
}

func (t *TintHurtado) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#dbdbdb")
}

func (t *TintHurtado) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fbe84a")
}

func (t *TintHurtado) Black() lipgloss.Color {
	return lipgloss.Color("#575757")
}

func (t *TintHurtado) Blue() lipgloss.Color {
	return lipgloss.Color("#496487")
}

func (t *TintHurtado) Cyan() lipgloss.Color {
	return lipgloss.Color("#86e9fe")
}

func (t *TintHurtado) Green() lipgloss.Color {
	return lipgloss.Color("#a5e055")
}

func (t *TintHurtado) Purple() lipgloss.Color {
	return lipgloss.Color("#fd5ff1")
}

func (t *TintHurtado) Red() lipgloss.Color {
	return lipgloss.Color("#ff1b00")
}

func (t *TintHurtado) White() lipgloss.Color {
	return lipgloss.Color("#cbcccb")
}

func (t *TintHurtado) Yellow() lipgloss.Color {
	return lipgloss.Color("#fbe74a")
}
