// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.
//
// Generated using tints/themes from:
//  * https://github.com/atomcorp/themes

package tints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintIdea (idea) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=idea
type TintIdea struct{}

// DisplayName returns the display name of the tint.
func (t *TintIdea) DisplayName() string {
	return "idea"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintIdea) ID() string {
	return "idea"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintIdea) About() string {
	return `Tint: idea`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintIdea) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#adadad")
}

// Bg returns the recommended default background color for this tint.
func (t *TintIdea) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#202020")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintIdea) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintIdea) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintIdea) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintIdea) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6c9ced")
}

func (t *TintIdea) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#248887")
}

func (t *TintIdea) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#98b61c")
}

func (t *TintIdea) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#fc7eff")
}

func (t *TintIdea) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#fc7072")
}

func (t *TintIdea) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#181818")
}

func (t *TintIdea) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ffff0b")
}

func (t *TintIdea) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#adadad")
}

func (t *TintIdea) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#437ee7")
}

func (t *TintIdea) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#248887")
}

func (t *TintIdea) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#98b61c")
}

func (t *TintIdea) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#9d74b0")
}

func (t *TintIdea) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#fc5256")
}

func (t *TintIdea) White() lipgloss.TerminalColor {
	return lipgloss.Color("#181818")
}

func (t *TintIdea) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#ccb444")
}
