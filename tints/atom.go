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

// TintAtom (Atom) is a collection of lipgloss styles.
type TintAtom struct{}

// DisplayName returns the display name of the tint.
func (t *TintAtom) DisplayName() string {
	return "Atom"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintAtom) ID() string {
	return "atom"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintAtom) About() string {
	return `Tint: Atom`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintAtom) Fg() lipgloss.Color {
	return lipgloss.Color("#c5c8c6")
}

// Bg returns the recommended default background color for this tint.
func (t *TintAtom) Bg() lipgloss.Color {
	return lipgloss.Color("#161719")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintAtom) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintAtom) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintAtom) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintAtom) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#96cbfe")
}

func (t *TintAtom) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#85befd")
}

func (t *TintAtom) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#94fa36")
}

func (t *TintAtom) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#b9b6fc")
}

func (t *TintAtom) BrightRed() lipgloss.Color {
	return lipgloss.Color("#fd5ff1")
}

func (t *TintAtom) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#e0e0e0")
}

func (t *TintAtom) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#f5ffa8")
}

func (t *TintAtom) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintAtom) Blue() lipgloss.Color {
	return lipgloss.Color("#85befd")
}

func (t *TintAtom) Cyan() lipgloss.Color {
	return lipgloss.Color("#85befd")
}

func (t *TintAtom) Green() lipgloss.Color {
	return lipgloss.Color("#87c38a")
}

func (t *TintAtom) Purple() lipgloss.Color {
	return lipgloss.Color("#b9b6fc")
}

func (t *TintAtom) Red() lipgloss.Color {
	return lipgloss.Color("#fd5ff1")
}

func (t *TintAtom) White() lipgloss.Color {
	return lipgloss.Color("#e0e0e0")
}

func (t *TintAtom) Yellow() lipgloss.Color {
	return lipgloss.Color("#ffd7b1")
}
