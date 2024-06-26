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

// TintOceanicMaterial (OceanicMaterial) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#OceanicMaterial
type TintOceanicMaterial struct{}

// DisplayName returns the display name of the tint.
func (t *TintOceanicMaterial) DisplayName() string {
	return "OceanicMaterial"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintOceanicMaterial) ID() string {
	return "oceanic_material"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintOceanicMaterial) About() string {
	return `Tint: OceanicMaterial`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintOceanicMaterial) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#c2c8d7")
}

// Bg returns the recommended default background color for this tint.
func (t *TintOceanicMaterial) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1c262b")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintOceanicMaterial) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintOceanicMaterial) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintOceanicMaterial) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#777777")
}

func (t *TintOceanicMaterial) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#54a4f3")
}

func (t *TintOceanicMaterial) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#42c7da")
}

func (t *TintOceanicMaterial) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#70be71")
}

func (t *TintOceanicMaterial) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#aa4dbc")
}

func (t *TintOceanicMaterial) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#dc5c60")
}

func (t *TintOceanicMaterial) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintOceanicMaterial) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fff163")
}

func (t *TintOceanicMaterial) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintOceanicMaterial) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#1e80f0")
}

func (t *TintOceanicMaterial) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#16afca")
}

func (t *TintOceanicMaterial) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#40a33f")
}

func (t *TintOceanicMaterial) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#8800a0")
}

func (t *TintOceanicMaterial) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ee2b2a")
}

func (t *TintOceanicMaterial) White() lipgloss.TerminalColor {
	return lipgloss.Color("#a4a4a4")
}

func (t *TintOceanicMaterial) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffea2e")
}
