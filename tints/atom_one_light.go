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

// TintAtomOneLight (AtomOneLight) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=AtomOneLight
type TintAtomOneLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintAtomOneLight) DisplayName() string {
	return "AtomOneLight"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintAtomOneLight) ID() string {
	return "atom_one_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintAtomOneLight) About() string {
	return `Tint: AtomOneLight`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintAtomOneLight) Fg() lipgloss.Color {
	return lipgloss.Color("#2a2c33")
}

// Bg returns the recommended default background color for this tint.
func (t *TintAtomOneLight) Bg() lipgloss.Color {
	return lipgloss.Color("#f9f9f9")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintAtomOneLight) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintAtomOneLight) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintAtomOneLight) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintAtomOneLight) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#2f5af3")
}

func (t *TintAtomOneLight) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#3f953a")
}

func (t *TintAtomOneLight) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#3f953a")
}

func (t *TintAtomOneLight) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#a00095")
}

func (t *TintAtomOneLight) BrightRed() lipgloss.Color {
	return lipgloss.Color("#de3e35")
}

func (t *TintAtomOneLight) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintAtomOneLight) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#d2b67c")
}

func (t *TintAtomOneLight) Black() lipgloss.Color {
	return lipgloss.Color("#000000")
}

func (t *TintAtomOneLight) Blue() lipgloss.Color {
	return lipgloss.Color("#2f5af3")
}

func (t *TintAtomOneLight) Cyan() lipgloss.Color {
	return lipgloss.Color("#3f953a")
}

func (t *TintAtomOneLight) Green() lipgloss.Color {
	return lipgloss.Color("#3f953a")
}

func (t *TintAtomOneLight) Purple() lipgloss.Color {
	return lipgloss.Color("#950095")
}

func (t *TintAtomOneLight) Red() lipgloss.Color {
	return lipgloss.Color("#de3e35")
}

func (t *TintAtomOneLight) White() lipgloss.Color {
	return lipgloss.Color("#bbbbbb")
}

func (t *TintAtomOneLight) Yellow() lipgloss.Color {
	return lipgloss.Color("#d2b67c")
}
