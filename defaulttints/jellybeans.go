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

// TintJellybeans (Jellybeans) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Jellybeans
type TintJellybeans struct{}

// DisplayName returns the display name of the tint.
func (t *TintJellybeans) DisplayName() string {
	return "Jellybeans"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintJellybeans) ID() string {
	return "jellybeans"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintJellybeans) About() string {
	return `Tint: Jellybeans`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintJellybeans) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#dedede")
}

// Bg returns the recommended default background color for this tint.
func (t *TintJellybeans) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#121212")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintJellybeans) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintJellybeans) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintJellybeans) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#bdbdbd")
}

func (t *TintJellybeans) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#b1d8f6")
}

func (t *TintJellybeans) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#1ab2a8")
}

func (t *TintJellybeans) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#bddeab")
}

func (t *TintJellybeans) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#fbdaff")
}

func (t *TintJellybeans) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ffa1a1")
}

func (t *TintJellybeans) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintJellybeans) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffdca0")
}

func (t *TintJellybeans) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#929292")
}

func (t *TintJellybeans) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#97bedc")
}

func (t *TintJellybeans) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00988e")
}

func (t *TintJellybeans) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#94b979")
}

func (t *TintJellybeans) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#e1c0fa")
}

func (t *TintJellybeans) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#e27373")
}

func (t *TintJellybeans) White() lipgloss.TerminalColor {
	return lipgloss.Color("#dedede")
}

func (t *TintJellybeans) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffba7b")
}
