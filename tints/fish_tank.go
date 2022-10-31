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

// TintFishTank (FishTank) is a collection of lipgloss styles.
type TintFishTank struct{}

// DisplayName returns the display name of the tint.
func (t *TintFishTank) DisplayName() string {
	return "FishTank"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintFishTank) ID() string {
	return "fish_tank"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintFishTank) About() string {
	return `Tint: FishTank`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintFishTank) Fg() lipgloss.Color {
	return lipgloss.Color("#ecf0fe")
}

// Bg returns the recommended default background color for this tint.
func (t *TintFishTank) Bg() lipgloss.Color {
	return lipgloss.Color("#232537")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintFishTank) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintFishTank) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintFishTank) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#6c5b30")
}

func (t *TintFishTank) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#b2befa")
}

func (t *TintFishTank) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#a5bd86")
}

func (t *TintFishTank) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#dbffa9")
}

func (t *TintFishTank) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#fda5cd")
}

func (t *TintFishTank) BrightRed() lipgloss.Color {
	return lipgloss.Color("#da4b8a")
}

func (t *TintFishTank) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#f6ffec")
}

func (t *TintFishTank) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#fee6a9")
}

func (t *TintFishTank) Black() lipgloss.Color {
	return lipgloss.Color("#03073c")
}

func (t *TintFishTank) Blue() lipgloss.Color {
	return lipgloss.Color("#525fb8")
}

func (t *TintFishTank) Cyan() lipgloss.Color {
	return lipgloss.Color("#968763")
}

func (t *TintFishTank) Green() lipgloss.Color {
	return lipgloss.Color("#acf157")
}

func (t *TintFishTank) Purple() lipgloss.Color {
	return lipgloss.Color("#986f82")
}

func (t *TintFishTank) Red() lipgloss.Color {
	return lipgloss.Color("#c6004a")
}

func (t *TintFishTank) White() lipgloss.Color {
	return lipgloss.Color("#ecf0fc")
}

func (t *TintFishTank) Yellow() lipgloss.Color {
	return lipgloss.Color("#fecd5e")
}
