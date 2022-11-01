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
//  * Michael Andreuzza (https://github.com/michael-andreuzza)
//  * Chaphasilor (https://github.com/Chaphasilor)

package tints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintSerendipityMidnight (Serendipity Midnight) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Serendipity+Midnight
//
// Credit to:
//   - Michael Andreuzza (https://github.com/michael-andreuzza)
//   - Chaphasilor (https://github.com/Chaphasilor)
type TintSerendipityMidnight struct{}

// DisplayName returns the display name of the tint.
func (t *TintSerendipityMidnight) DisplayName() string {
	return "Serendipity Midnight"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintSerendipityMidnight) ID() string {
	return "serendipity_midnight"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintSerendipityMidnight) About() string {
	return `Tint: Serendipity Midnight
Tint credits:
  * Michael Andreuzza (https://github.com/michael-andreuzza)
  * Chaphasilor (https://github.com/Chaphasilor)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintSerendipityMidnight) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#DEE0EF")
}

// Bg returns the recommended default background color for this tint.
func (t *TintSerendipityMidnight) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1C1E2D")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintSerendipityMidnight) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.Color("#2C2E3D")
}

// Cursor returns the recommended color for the cursor.
func (t *TintSerendipityMidnight) Cursor() lipgloss.TerminalColor {
	return lipgloss.Color("#6B6D7C")
}

func (t *TintSerendipityMidnight) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#8D8F9E")
}

func (t *TintSerendipityMidnight) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#5BA2D0")
}

func (t *TintSerendipityMidnight) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#94B8FF")
}

func (t *TintSerendipityMidnight) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#9CCFD8")
}

func (t *TintSerendipityMidnight) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#A78BFA")
}

func (t *TintSerendipityMidnight) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#EE8679")
}

func (t *TintSerendipityMidnight) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#DEE0EF")
}

func (t *TintSerendipityMidnight) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#F8D2C9")
}

func (t *TintSerendipityMidnight) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#232534")
}

func (t *TintSerendipityMidnight) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#5BA2D0")
}

func (t *TintSerendipityMidnight) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#94B8FF")
}

func (t *TintSerendipityMidnight) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#9CCFD8")
}

func (t *TintSerendipityMidnight) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#A78BFA")
}

func (t *TintSerendipityMidnight) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#EE8679")
}

func (t *TintSerendipityMidnight) White() lipgloss.TerminalColor {
	return lipgloss.Color("#DEE0EF")
}

func (t *TintSerendipityMidnight) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#F8D2C9")
}
