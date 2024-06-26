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

// TintBlueMatrix (Blue Matrix) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#BlueMatrix
type TintBlueMatrix struct{}

// DisplayName returns the display name of the tint.
func (t *TintBlueMatrix) DisplayName() string {
	return "Blue Matrix"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBlueMatrix) ID() string {
	return "blue_matrix"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBlueMatrix) About() string {
	return `Tint: Blue Matrix`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBlueMatrix) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#00a2ff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBlueMatrix) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#101116")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBlueMatrix) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintBlueMatrix) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintBlueMatrix) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#686868")
}

func (t *TintBlueMatrix) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6871ff")
}

func (t *TintBlueMatrix) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#60fdff")
}

func (t *TintBlueMatrix) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#5ffa68")
}

func (t *TintBlueMatrix) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d682ec")
}

func (t *TintBlueMatrix) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6e67")
}

func (t *TintBlueMatrix) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintBlueMatrix) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fffc67")
}

func (t *TintBlueMatrix) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#101116")
}

func (t *TintBlueMatrix) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#00b0ff")
}

func (t *TintBlueMatrix) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#76c1ff")
}

func (t *TintBlueMatrix) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#00ff9c")
}

func (t *TintBlueMatrix) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#d57bff")
}

func (t *TintBlueMatrix) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5680")
}

func (t *TintBlueMatrix) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c7c7c7")
}

func (t *TintBlueMatrix) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fffc58")
}
