// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintRosePine (Rose Pine) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Rose+Pine
type TintRosePine struct{}

// DisplayName returns the display name of the tint.
func (t *TintRosePine) DisplayName() string {
	return "Rose Pine"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintRosePine) ID() string {
	return "rose_pine"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintRosePine) About() string {
	return `Tint: Rose Pine`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintRosePine) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#e0def4")
}

// Bg returns the recommended default background color for this tint.
func (t *TintRosePine) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#191724")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintRosePine) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintRosePine) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintRosePine) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#706e86")
}

func (t *TintRosePine) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#31748f")
}

func (t *TintRosePine) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ebbcba")
}

func (t *TintRosePine) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#9ccfd8")
}

func (t *TintRosePine) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c4a7e7")
}

func (t *TintRosePine) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#eb6f92")
}

func (t *TintRosePine) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#e0def4")
}

func (t *TintRosePine) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f6c177")
}

func (t *TintRosePine) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#706e86")
}

func (t *TintRosePine) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#31748f")
}

func (t *TintRosePine) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#ebbcba")
}

func (t *TintRosePine) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#9ccfd8")
}

func (t *TintRosePine) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c4a7e7")
}

func (t *TintRosePine) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#eb6f92")
}

func (t *TintRosePine) White() lipgloss.TerminalColor {
	return lipgloss.Color("#e0def4")
}

func (t *TintRosePine) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f6c177")
}