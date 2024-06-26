// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// All tints can be previewed here:
//  * https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes
//  * https://github.com/mbadolato/iTerm2-Color-Schemes

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintNightLionV1 (NightLion v1) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#NightLionV1
type TintNightLionV1 struct{}

// DisplayName returns the display name of the tint.
func (t *TintNightLionV1) DisplayName() string {
	return "NightLion v1"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintNightLionV1) ID() string {
	return "night_lion_v_1"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintNightLionV1) About() string {
	return `Tint: NightLion v1`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintNightLionV1) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#bbbbbb")
}

// Bg returns the recommended default background color for this tint.
func (t *TintNightLionV1) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintNightLionV1) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintNightLionV1) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintNightLionV1) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintNightLionV1) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#5555ff")
}

func (t *TintNightLionV1) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#55ffff")
}

func (t *TintNightLionV1) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#55ff55")
}

func (t *TintNightLionV1) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff55ff")
}

func (t *TintNightLionV1) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5555")
}

func (t *TintNightLionV1) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintNightLionV1) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff55")
}

func (t *TintNightLionV1) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#4c4c4c")
}

func (t *TintNightLionV1) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#276bd8")
}

func (t *TintNightLionV1) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00dadf")
}

func (t *TintNightLionV1) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#5fde8f")
}

func (t *TintNightLionV1) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#bb00bb")
}

func (t *TintNightLionV1) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#bb0000")
}

func (t *TintNightLionV1) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bbbbbb")
}

func (t *TintNightLionV1) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f3f167")
}
