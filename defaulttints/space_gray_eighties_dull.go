// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
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

// TintSpaceGrayEightiesDull (SpaceGray Eighties Dull) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#SpaceGrayEightiesDull
type TintSpaceGrayEightiesDull struct{}

// DisplayName returns the display name of the tint.
func (t *TintSpaceGrayEightiesDull) DisplayName() string {
	return "SpaceGray Eighties Dull"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSpaceGrayEightiesDull) ID() string {
	return "space_gray_eighties_dull"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSpaceGrayEightiesDull) About() string {
	return `Tint: SpaceGray Eighties Dull`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSpaceGrayEightiesDull) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c9c6bc")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSpaceGrayEightiesDull) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#222222")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSpaceGrayEightiesDull) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintSpaceGrayEightiesDull) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintSpaceGrayEightiesDull) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintSpaceGrayEightiesDull) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#5486c0")
}

func (t *TintSpaceGrayEightiesDull) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#58c2c1")
}

func (t *TintSpaceGrayEightiesDull) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#89e986")
}

func (t *TintSpaceGrayEightiesDull) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#bf83c1")
}

func (t *TintSpaceGrayEightiesDull) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ec5f67")
}

func (t *TintSpaceGrayEightiesDull) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintSpaceGrayEightiesDull) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fec254")
}

func (t *TintSpaceGrayEightiesDull) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#15171c")
}

func (t *TintSpaceGrayEightiesDull) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#7c8fa5")
}

func (t *TintSpaceGrayEightiesDull) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#80cdcb")
}

func (t *TintSpaceGrayEightiesDull) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#92b477")
}

func (t *TintSpaceGrayEightiesDull) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#a5789e")
}

func (t *TintSpaceGrayEightiesDull) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#b24a56")
}

func (t *TintSpaceGrayEightiesDull) White() lipgloss.TerminalColor {
	return lipgloss.Color("#b3b8c3")
}

func (t *TintSpaceGrayEightiesDull) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#c6735a")
}
