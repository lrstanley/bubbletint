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
//
// Additional credit to:
//  * jos3s (https://github.com/jos3s)

package defaulttints

import (
	"github.com/charmbracelet/lipgloss"
)

// TintBubbles (Bubbles) is a collection of lipgloss styles.
//
// Reference: https://github.com/lrstanley/bubbletint/blob/master/DEFAULT_TINTS.md#Bubbles
//
// Credit to:
//   - jos3s (https://github.com/jos3s)
type TintBubbles struct{}

// DisplayName returns the display name of the tint.
func (t *TintBubbles) DisplayName() string {
	return "Bubbles"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintBubbles) ID() string {
	return "bubbles"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintBubbles) About() string {
	return `Tint: Bubbles
Tint credits:
  * jos3s (https://github.com/jos3s)`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintBubbles) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#FFFFFF")
}

// Bg returns the recommended default background color for this tint.
func (t *TintBubbles) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#161D38")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintBubbles) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.Color("#B5D5FF")
}

// Cursor returns the recommended color for the cursor.
func (t *TintBubbles) Cursor() lipgloss.TerminalColor {
	return lipgloss.Color("#BBBBBB")
}

func (t *TintBubbles) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#555555")
}

func (t *TintBubbles) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#589DF6")
}

func (t *TintBubbles) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#3979BC")
}

func (t *TintBubbles) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#35BB9A")
}

func (t *TintBubbles) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#E75699")
}

func (t *TintBubbles) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#FA8C8F")
}

func (t *TintBubbles) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#FFFFFF")
}

func (t *TintBubbles) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#FFFF55")
}

func (t *TintBubbles) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#000000")
}

func (t *TintBubbles) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#589DF6")
}

func (t *TintBubbles) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#1F9EE7")
}

func (t *TintBubbles) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#21B089")
}

func (t *TintBubbles) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#944D95")
}

func (t *TintBubbles) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#F9555F")
}

func (t *TintBubbles) White() lipgloss.TerminalColor {
	return lipgloss.Color("#BBBBBB")
}

func (t *TintBubbles) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#FEF02A")
}
