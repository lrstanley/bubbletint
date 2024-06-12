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

// TintTinaciousDesignLight (Tinacious Design (Light)) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#TinaciousDesignLight
type TintTinaciousDesignLight struct{}

// DisplayName returns the display name of the tint.
func (t *TintTinaciousDesignLight) DisplayName() string {
	return "Tinacious Design (Light)"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintTinaciousDesignLight) ID() string {
	return "tinacious_design_light"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintTinaciousDesignLight) About() string {
	return `Tint: Tinacious Design (Light)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintTinaciousDesignLight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#1d1d26")
}

// Bg returns the recommended default background color for this tint.
func (t *TintTinaciousDesignLight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#f8f8ff")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintTinaciousDesignLight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintTinaciousDesignLight) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintTinaciousDesignLight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#636667")
}

func (t *TintTinaciousDesignLight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#00cbff")
}

func (t *TintTinaciousDesignLight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00d5d4")
}

func (t *TintTinaciousDesignLight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#00d364")
}

func (t *TintTinaciousDesignLight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#d783ff")
}

func (t *TintTinaciousDesignLight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff2f92")
}

func (t *TintTinaciousDesignLight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#d5d6f3")
}

func (t *TintTinaciousDesignLight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd479")
}

func (t *TintTinaciousDesignLight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1d1d26")
}

func (t *TintTinaciousDesignLight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#00cbff")
}

func (t *TintTinaciousDesignLight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00ceca")
}

func (t *TintTinaciousDesignLight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#00d364")
}

func (t *TintTinaciousDesignLight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#cc66ff")
}

func (t *TintTinaciousDesignLight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff3399")
}

func (t *TintTinaciousDesignLight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#cbcbf0")
}

func (t *TintTinaciousDesignLight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffcc66")
}
