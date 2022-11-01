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

// TintSlate (Slate) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Slate
type TintSlate struct{}

// DisplayName returns the display name of the tint.
func (t *TintSlate) DisplayName() string {
	return "Slate"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSlate) ID() string {
	return "slate"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSlate) About() string {
	return `Tint: Slate`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSlate) Fg() lipgloss.Color {
	return lipgloss.Color("#35b1d2")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSlate) Bg() lipgloss.Color {
	return lipgloss.Color("#222222")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSlate) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintSlate) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintSlate) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintSlate) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#7ab0d2")
}

func (t *TintSlate) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#8cdfe0")
}

func (t *TintSlate) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#beffa8")
}

func (t *TintSlate) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#c5a7d9")
}

func (t *TintSlate) BrightRed() lipgloss.Color {
	return lipgloss.Color("#ffcdd9")
}

func (t *TintSlate) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#e0e0e0")
}

func (t *TintSlate) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#d0ccca")
}

func (t *TintSlate) Black() lipgloss.Color {
	return lipgloss.Color("#222222")
}

func (t *TintSlate) Blue() lipgloss.Color {
	return lipgloss.Color("#264b49")
}

func (t *TintSlate) Cyan() lipgloss.Color {
	return lipgloss.Color("#15ab9c")
}

func (t *TintSlate) Green() lipgloss.Color {
	return lipgloss.Color("#81d778")
}

func (t *TintSlate) Purple() lipgloss.Color {
	return lipgloss.Color("#a481d3")
}

func (t *TintSlate) Red() lipgloss.Color {
	return lipgloss.Color("#e2a8bf")
}

func (t *TintSlate) White() lipgloss.Color {
	return lipgloss.Color("#02c5e0")
}

func (t *TintSlate) Yellow() lipgloss.Color {
	return lipgloss.Color("#c4c9c0")
}
