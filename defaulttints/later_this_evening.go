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

// TintLaterThisEvening (Later This Evening) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#LaterThisEvening
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
func (t *TintLaterThisEvening) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#959595")
}

// Bg returns the recommended default background color for this tint.
func (t *TintLaterThisEvening) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#222222")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintLaterThisEvening) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintLaterThisEvening) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintLaterThisEvening) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#454747")
}

func (t *TintLaterThisEvening) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6699d6")
}

func (t *TintLaterThisEvening) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5fc0ae")
}

func (t *TintLaterThisEvening) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#aabb39")
}

func (t *TintLaterThisEvening) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ab53d6")
}

func (t *TintLaterThisEvening) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#d3232f")
}

func (t *TintLaterThisEvening) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#c1c2c2")
}

func (t *TintLaterThisEvening) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5be39")
}

func (t *TintLaterThisEvening) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#2b2b2b")
}

func (t *TintLaterThisEvening) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#a0bad6")
}

func (t *TintLaterThisEvening) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#91bfb7")
}

func (t *TintLaterThisEvening) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#afba67")
}

func (t *TintLaterThisEvening) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c092d6")
}

func (t *TintLaterThisEvening) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#d45a60")
}

func (t *TintLaterThisEvening) White() lipgloss.TerminalColor {
	return lipgloss.Color("#3c3d3d")
}

func (t *TintLaterThisEvening) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#e5d289")
}
