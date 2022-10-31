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

// TintFahrenheit (Fahrenheit) is a collection of lipgloss styles.
type TintFahrenheit struct{}

// DisplayName returns the display name of the tint.
func (t *TintFahrenheit) DisplayName() string {
	return "Fahrenheit"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintFahrenheit) ID() string {
	return "fahrenheit"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintFahrenheit) About() string {
	return `Tint: Fahrenheit`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintFahrenheit) Fg() lipgloss.Color {
	return lipgloss.Color("#ffffce")
}

// Bg returns the recommended default background color for this tint.
func (t *TintFahrenheit) Bg() lipgloss.Color {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintFahrenheit) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintFahrenheit) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintFahrenheit) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintFahrenheit) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#cb4a05")
}

func (t *TintFahrenheit) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#fed04d")
}

func (t *TintFahrenheit) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#cc734d")
}

func (t *TintFahrenheit) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#4e739f")
}

func (t *TintFahrenheit) BrightRed() lipgloss.Color {
	return lipgloss.Color("#fecea0")
}

func (t *TintFahrenheit) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintFahrenheit) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fd9f4d")
}

func (t *TintFahrenheit) Black() lipgloss.Color {
	return lipgloss.Color("#1d1d1d")
}

func (t *TintFahrenheit) Blue() lipgloss.Color {
	return lipgloss.Color("#720102")
}

func (t *TintFahrenheit) Cyan() lipgloss.Color {
	return lipgloss.Color("#979797")
}

func (t *TintFahrenheit) Green() lipgloss.Color {
	return lipgloss.Color("#9e744d")
}

func (t *TintFahrenheit) Purple() lipgloss.Color {
	return lipgloss.Color("#734c4d")
}

func (t *TintFahrenheit) Red() lipgloss.Color {
	return lipgloss.Color("#cda074")
}

func (t *TintFahrenheit) White() lipgloss.Color {
	return lipgloss.Color("#ffffce")
}

func (t *TintFahrenheit) Yellow() lipgloss.Color {
	return lipgloss.Color("#fecf75")
}
