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

// TintThayerBright (Thayer Bright) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#ThayerBright
type TintThayerBright struct{}

// DisplayName returns the display name of the tint.
func (t *TintThayerBright) DisplayName() string {
	return "Thayer Bright"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintThayerBright) ID() string {
	return "thayer_bright"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintThayerBright) About() string {
	return `Tint: Thayer Bright`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintThayerBright) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#f8f8f8")
}

// Bg returns the recommended default background color for this tint.
func (t *TintThayerBright) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1b1d1e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintThayerBright) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintThayerBright) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintThayerBright) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#505354")
}

func (t *TintThayerBright) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#3f78ff")
}

func (t *TintThayerBright) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#23cfd5")
}

func (t *TintThayerBright) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#b6e354")
}

func (t *TintThayerBright) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#9e6ffe")
}

func (t *TintThayerBright) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5995")
}

func (t *TintThayerBright) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#f8f8f2")
}

func (t *TintThayerBright) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#feed6c")
}

func (t *TintThayerBright) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1b1d1e")
}

func (t *TintThayerBright) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#2757d6")
}

func (t *TintThayerBright) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#38c8b5")
}

func (t *TintThayerBright) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#4df840")
}

func (t *TintThayerBright) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#8c54fe")
}

func (t *TintThayerBright) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#f92672")
}

func (t *TintThayerBright) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ccccc6")
}

func (t *TintThayerBright) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f4fd22")
}
