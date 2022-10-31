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

// TintLovelace (lovelace) is a collection of lipgloss styles.
type TintLovelace struct{}

// DisplayName returns the display name of the tint.
func (t *TintLovelace) DisplayName() string {
	return "lovelace"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintLovelace) ID() string {
	return "lovelace"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintLovelace) About() string {
	return `Tint: lovelace`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintLovelace) Fg() lipgloss.Color {
	return lipgloss.Color("#fdfdfd")
}

// Bg returns the recommended default background color for this tint.
func (t *TintLovelace) Bg() lipgloss.Color {
	return lipgloss.Color("#1d1f28")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintLovelace) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintLovelace) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintLovelace) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#414458")
}

func (t *TintLovelace) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#556fff")
}

func (t *TintLovelace) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#3fdcee")
}

func (t *TintLovelace) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#18e3c8")
}

func (t *TintLovelace) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#b043d1")
}

func (t *TintLovelace) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ff4971")
}

func (t *TintLovelace) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#bebec1")
}

func (t *TintLovelace) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#ff8037")
}

func (t *TintLovelace) Black() lipgloss.Color {
	return lipgloss.Color("#282a36")
}

func (t *TintLovelace) Blue() lipgloss.Color {
	return lipgloss.Color("#8897f4")
}

func (t *TintLovelace) Cyan() lipgloss.Color {
	return lipgloss.Color("#79e6f3")
}

func (t *TintLovelace) Green() lipgloss.Color {
	return lipgloss.Color("#5adecd")
}

func (t *TintLovelace) Purple() lipgloss.Color {
	return lipgloss.Color("#c574dd")
}

func (t *TintLovelace) Red() lipgloss.Color {
	return lipgloss.Color("#f37f97")
}

func (t *TintLovelace) White() lipgloss.Color {
	return lipgloss.Color("#fdfdfd")
}

func (t *TintLovelace) Yellow() lipgloss.Color {
	return lipgloss.Color("#f2a272")
}
