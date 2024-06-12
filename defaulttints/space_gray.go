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

// TintSpaceGray (SpaceGray) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#SpaceGray
type TintSpaceGray struct{}

// DisplayName returns the display name of the tint.
func (t *TintSpaceGray) DisplayName() string {
	return "SpaceGray"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSpaceGray) ID() string {
	return "space_gray"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSpaceGray) About() string {
	return `Tint: SpaceGray`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSpaceGray) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#b3b8c3")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSpaceGray) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#20242d")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSpaceGray) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSpaceGray) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSpaceGray) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintSpaceGray) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#7d8fa4")
}

func (t *TintSpaceGray) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#85a7a5")
}

func (t *TintSpaceGray) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#87b379")
}

func (t *TintSpaceGray) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#a47996")
}

func (t *TintSpaceGray) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#b04b57")
}

func (t *TintSpaceGray) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintSpaceGray) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5c179")
}

func (t *TintSpaceGray) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintSpaceGray) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#7d8fa4")
}

func (t *TintSpaceGray) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#85a7a5")
}

func (t *TintSpaceGray) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#87b379")
}

func (t *TintSpaceGray) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a47996")
}

func (t *TintSpaceGray) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#b04b57")
}

func (t *TintSpaceGray) White() lipgloss.TerminalColor {
	return lipgloss.Color("#b3b8c3")
}

func (t *TintSpaceGray) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5c179")
}
