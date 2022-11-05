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

// TintDracula (Dracula) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Dracula
type TintDracula struct{}

// DisplayName returns the display name of the tint.
func (t *TintDracula) DisplayName() string {
	return "Dracula"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintDracula) ID() string {
	return "dracula"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintDracula) About() string {
	return `Tint: Dracula`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintDracula) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#f8f8f2")
}

// Bg returns the recommended default background color for this tint.
func (t *TintDracula) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1e1f29")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintDracula) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintDracula) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintDracula) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintDracula) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#bd93f9")
}

func (t *TintDracula) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8be9fd")
}

func (t *TintDracula) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#50fa7b")
}

func (t *TintDracula) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff79c6")
}

func (t *TintDracula) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5555")
}

func (t *TintDracula) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintDracula) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f1fa8c")
}

func (t *TintDracula) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintDracula) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#bd93f9")
}

func (t *TintDracula) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#8be9fd")
}

func (t *TintDracula) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#50fa7b")
}

func (t *TintDracula) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ff79c6")
}

func (t *TintDracula) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff5555")
}

func (t *TintDracula) White() lipgloss.TerminalColor {
	return lipgloss.Color("#bbbbbb")
}

func (t *TintDracula) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#f1fa8c")
}