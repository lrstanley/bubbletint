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

// TintDeep (deep) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=deep
type TintDeep struct{}

// DisplayName returns the display name of the tint.
func (t *TintDeep) DisplayName() string {
	return "deep"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDeep) ID() string {
	return "deep"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDeep) About() string {
	return `Tint: deep`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDeep) Fg() lipgloss.Color {
	return lipgloss.Color("#cdcdcd")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDeep) Bg() lipgloss.Color {
	return lipgloss.Color("#090909")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDeep) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintDeep) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintDeep) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#535353")
}

func (t *TintDeep) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#9fa9ff")
}

func (t *TintDeep) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#8df9ff")
}

func (t *TintDeep) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#22ff18")
}

func (t *TintDeep) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#e09aff")
}

func (t *TintDeep) BrightRed() lipgloss.Color {
	return lipgloss.Color("#fb0007")
}

func (t *TintDeep) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintDeep) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fedc2b")
}

func (t *TintDeep) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintDeep) Blue() lipgloss.Color {
	return lipgloss.Color("#5665ff")
}

func (t *TintDeep) Cyan() lipgloss.Color {
	return lipgloss.Color("#50d2da")
}

func (t *TintDeep) Green() lipgloss.Color {
	return lipgloss.Color("#1cd915")
}

func (t *TintDeep) Purple() lipgloss.Color {
	return lipgloss.Color("#b052da")
}

func (t *TintDeep) Red() lipgloss.Color {
	return lipgloss.Color("#d70005")
}

func (t *TintDeep) White() lipgloss.Color {
	return lipgloss.Color("#e0e0e0")
}

func (t *TintDeep) Yellow() lipgloss.Color {
	return lipgloss.Color("#d9bd26")
}
