// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes
//
// Additional credit to:
//  * monokai (https://monokai.pro)

package tints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintMonokaiPro (Monokai Pro) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Monokai+Pro
//
// Credit to:
//   - monokai (https://monokai.pro)
type TintMonokaiPro struct{}

// DisplayName returns the display name of the tint.
func (t *TintMonokaiPro) DisplayName() string {
	return "Monokai Pro"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintMonokaiPro) ID() string {
	return "monokai_pro"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintMonokaiPro) About() string {
	return `Tint: Monokai Pro
Tint credits:
  * monokai (https://monokai.pro)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintMonokaiPro) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#fcfcfa")
}

// Bg returns the recommended default background color for this tint.
func (t *TintMonokaiPro) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#403e41")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintMonokaiPro) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.Color("#fcfcfa")
}

// Cursor returns the recommended color for the cursor.
func (t *TintMonokaiPro) Cursor() lipgloss.TerminalColor {
	return lipgloss.Color("#fcfcfa")
}

func (t *TintMonokaiPro) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#727072")
}

func (t *TintMonokaiPro) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#fc9867")
}

func (t *TintMonokaiPro) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#78dce8")
}

func (t *TintMonokaiPro) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#a9dc76")
}

func (t *TintMonokaiPro) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ab9df2")
}

func (t *TintMonokaiPro) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6188")
}

func (t *TintMonokaiPro) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#fcfcfa")
}

func (t *TintMonokaiPro) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd866")
}

func (t *TintMonokaiPro) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#403e41")
}

func (t *TintMonokaiPro) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#fc9867")
}

func (t *TintMonokaiPro) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#78dce8")
}

func (t *TintMonokaiPro) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#a9dc76")
}

func (t *TintMonokaiPro) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ab9df2")
}

func (t *TintMonokaiPro) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6188")
}

func (t *TintMonokaiPro) White() lipgloss.TerminalColor {
	return lipgloss.Color("#fcfcfa")
}

func (t *TintMonokaiPro) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd866")
}
