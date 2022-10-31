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

// TintAyu (ayu) is a collection of lipgloss styles.
type TintAyu struct{}

// DisplayName returns the display name of the tint.
func (t *TintAyu) DisplayName() string {
	return "ayu"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintAyu) ID() string {
	return "ayu"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintAyu) About() string {
	return `Tint: ayu`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintAyu) Fg() lipgloss.Color {
	return lipgloss.Color("#e6e1cf")
}

// Bg returns the recommended default background color for this tint.
func (t *TintAyu) Bg() lipgloss.Color {
	return lipgloss.Color("#0f1419")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintAyu) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintAyu) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintAyu) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#323232")
}

func (t *TintAyu) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#68d5ff")
}

func (t *TintAyu) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#c7fffd")
}

func (t *TintAyu) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#eafe84")
}

func (t *TintAyu) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#ffa3aa")
}

func (t *TintAyu) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ff6565")
}

func (t *TintAyu) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintAyu) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fff779")
}

func (t *TintAyu) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintAyu) Blue() lipgloss.Color {
	return lipgloss.Color("#36a3d9")
}

func (t *TintAyu) Cyan() lipgloss.Color {
	return lipgloss.Color("#95e6cb")
}

func (t *TintAyu) Green() lipgloss.Color {
	return lipgloss.Color("#b8cc52")
}

func (t *TintAyu) Purple() lipgloss.Color {
	return lipgloss.Color("#f07178")
}

func (t *TintAyu) Red() lipgloss.Color {
	return lipgloss.Color("#ff3333")
}

func (t *TintAyu) White() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintAyu) Yellow() lipgloss.Color {
	return lipgloss.Color("#e7c547")
}
