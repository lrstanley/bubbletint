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

// TintMidnightInMojave (midnight-in-mojave) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#MidnightInMojave
type TintMidnightInMojave struct{}

// DisplayName returns the display name of the tint.
func (t *TintMidnightInMojave) DisplayName() string {
	return "midnight-in-mojave"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintMidnightInMojave) ID() string {
	return "midnight_in_mojave"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintMidnightInMojave) About() string {
	return `Tint: midnight-in-mojave`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintMidnightInMojave) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

// Bg returns the recommended default background color for this tint.
func (t *TintMidnightInMojave) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1e1e1e")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintMidnightInMojave) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintMidnightInMojave) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintMidnightInMojave) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#1e1e1e")
}

func (t *TintMidnightInMojave) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#0a84ff")
}

func (t *TintMidnightInMojave) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5ac8fa")
}

func (t *TintMidnightInMojave) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#32d74b")
}

func (t *TintMidnightInMojave) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#bf5af2")
}

func (t *TintMidnightInMojave) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff453a")
}

func (t *TintMidnightInMojave) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintMidnightInMojave) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd60a")
}

func (t *TintMidnightInMojave) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#1e1e1e")
}

func (t *TintMidnightInMojave) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0a84ff")
}

func (t *TintMidnightInMojave) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#5ac8fa")
}

func (t *TintMidnightInMojave) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#32d74b")
}

func (t *TintMidnightInMojave) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#bf5af2")
}

func (t *TintMidnightInMojave) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff453a")
}

func (t *TintMidnightInMojave) White() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintMidnightInMojave) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffd60a")
}
