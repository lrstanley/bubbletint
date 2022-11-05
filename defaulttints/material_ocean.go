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

// TintMaterialOcean (MaterialOcean) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=MaterialOcean
type TintMaterialOcean struct{}

// DisplayName returns the display name of the tint.
func (t *TintMaterialOcean) DisplayName() string {
	return "MaterialOcean"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintMaterialOcean) ID() string {
	return "material_ocean"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintMaterialOcean) About() string {
	return `Tint: MaterialOcean`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintMaterialOcean) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#8f93a2")
}

// Bg returns the recommended default background color for this tint.
func (t *TintMaterialOcean) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#0f111a")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintMaterialOcean) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintMaterialOcean) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintMaterialOcean) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#546e7a")
}

func (t *TintMaterialOcean) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#82aaff")
}

func (t *TintMaterialOcean) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#89ddff")
}

func (t *TintMaterialOcean) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#c3e88d")
}

func (t *TintMaterialOcean) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#c792ea")
}

func (t *TintMaterialOcean) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5370")
}

func (t *TintMaterialOcean) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintMaterialOcean) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcb6b")
}

func (t *TintMaterialOcean) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#546e7a")
}

func (t *TintMaterialOcean) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#82aaff")
}

func (t *TintMaterialOcean) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#89ddff")
}

func (t *TintMaterialOcean) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#c3e88d")
}

func (t *TintMaterialOcean) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#c792ea")
}

func (t *TintMaterialOcean) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5370")
}

func (t *TintMaterialOcean) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintMaterialOcean) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcb6b")
}