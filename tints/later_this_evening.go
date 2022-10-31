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

// TintLaterThisEvening (Later This Evening) is a collection of lipgloss styles.
type TintLaterThisEvening struct{}

// DisplayName returns the display name of the tint.
func (t *TintLaterThisEvening) DisplayName() string {
	return "Later This Evening"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintLaterThisEvening) ID() string {
	return "later_this_evening"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintLaterThisEvening) About() string {
	return `Tint: Later This Evening`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintLaterThisEvening) Fg() lipgloss.Color {
	return lipgloss.Color("#959595")
}

// Bg returns the recommended default background color for this tint.
func (t *TintLaterThisEvening) Bg() lipgloss.Color {
	return lipgloss.Color("#222222")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintLaterThisEvening) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintLaterThisEvening) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintLaterThisEvening) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#454747")
}

func (t *TintLaterThisEvening) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#6699d6")
}

func (t *TintLaterThisEvening) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#5fc0ae")
}

func (t *TintLaterThisEvening) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#aabb39")
}

func (t *TintLaterThisEvening) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#ab53d6")
}

func (t *TintLaterThisEvening) BrightRed() lipgloss.Color {
	return lipgloss.Color("#d3232f")
}

func (t *TintLaterThisEvening) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#c1c2c2")
}

func (t *TintLaterThisEvening) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#e5be39")
}

func (t *TintLaterThisEvening) Black() lipgloss.Color {
	return lipgloss.Color("#2b2b2b")
}

func (t *TintLaterThisEvening) Blue() lipgloss.Color {
	return lipgloss.Color("#a0bad6")
}

func (t *TintLaterThisEvening) Cyan() lipgloss.Color {
	return lipgloss.Color("#91bfb7")
}

func (t *TintLaterThisEvening) Green() lipgloss.Color {
	return lipgloss.Color("#afba67")
}

func (t *TintLaterThisEvening) Purple() lipgloss.Color {
	return lipgloss.Color("#c092d6")
}

func (t *TintLaterThisEvening) Red() lipgloss.Color {
	return lipgloss.Color("#d45a60")
}

func (t *TintLaterThisEvening) White() lipgloss.Color {
	return lipgloss.Color("#3c3d3d")
}

func (t *TintLaterThisEvening) Yellow() lipgloss.Color {
	return lipgloss.Color("#e5d289")
}
