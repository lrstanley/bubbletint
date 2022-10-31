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

// TintIRBlack (IR_Black) is a collection of lipgloss styles.
type TintIRBlack struct{}

// DisplayName returns the display name of the tint.
func (t *TintIRBlack) DisplayName() string {
	return "IR_Black"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintIRBlack) ID() string {
	return "ir_black"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintIRBlack) About() string {
	return `Tint: IR_Black`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintIRBlack) Fg() lipgloss.Color {
	return lipgloss.Color("#f1f1f1")
}

// Bg returns the recommended default background color for this tint.
func (t *TintIRBlack) Bg() lipgloss.Color {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintIRBlack) SelectionBg() lipgloss.Color {
	return lipgloss.Color("")
}

// Cursor returns the recommended color for the cursor.
func (t *TintIRBlack) Cursor() lipgloss.Color {
	return lipgloss.Color("")
}

func (t *TintIRBlack) BrightBlack() lipgloss.Color {
	return lipgloss.Color("#7b7b7b")
}

func (t *TintIRBlack) BrightBlue() lipgloss.Color {
	return lipgloss.Color("#b5dcff")
}

func (t *TintIRBlack) BrightCyan() lipgloss.Color {
	return lipgloss.Color("#e0e0fe")
}

func (t *TintIRBlack) BrightGreen() lipgloss.Color {
	return lipgloss.Color("#cfffab")
}

func (t *TintIRBlack) BrightPurple() lipgloss.Color {
	return lipgloss.Color("#fb9cfe")
}

func (t *TintIRBlack) BrightRed() lipgloss.Color {
	return lipgloss.Color("#fcb6b0")
}

func (t *TintIRBlack) BrightWhite() lipgloss.Color {
	return lipgloss.Color("#ffffff")
}

func (t *TintIRBlack) BrightYellow() lipgloss.Color {
	return lipgloss.Color("#ffffcc")
}

func (t *TintIRBlack) Black() lipgloss.Color {
	return lipgloss.Color("#4f4f4f")
}

func (t *TintIRBlack) Blue() lipgloss.Color {
	return lipgloss.Color("#96cafe")
}

func (t *TintIRBlack) Cyan() lipgloss.Color {
	return lipgloss.Color("#c6c5fe")
}

func (t *TintIRBlack) Green() lipgloss.Color {
	return lipgloss.Color("#a8ff60")
}

func (t *TintIRBlack) Purple() lipgloss.Color {
	return lipgloss.Color("#fa73fd")
}

func (t *TintIRBlack) Red() lipgloss.Color {
	return lipgloss.Color("#fa6c60")
}

func (t *TintIRBlack) White() lipgloss.Color {
	return lipgloss.Color("#efedef")
}

func (t *TintIRBlack) Yellow() lipgloss.Color {
	return lipgloss.Color("#fffeb7")
}
