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

// TintScarletProtocol (Scarlet Protocol) is a collection of lipgloss styles.
//
// Reference: https://windowsterminalthemes.dev/?theme=Scarlet+Protocol
type TintScarletProtocol struct{}

// DisplayName returns the display name of the tint.
func (t *TintScarletProtocol) DisplayName() string {
	return "Scarlet Protocol"
}

// ID returns the name of the tint (normalized, snakecase style).
func (t *TintScarletProtocol) ID() string {
	return "scarlet_protocol"
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (t *TintScarletProtocol) About() string {
	return `Tint: Scarlet Protocol`
}

// Fg returns the recommended default foreground color for this tint.
func (t *TintScarletProtocol) Fg() lipgloss.TerminalColor {
	return lipgloss.Color("#e41951")
}

// Bg returns the recommended default background color for this tint.
func (t *TintScarletProtocol) Bg() lipgloss.TerminalColor {
	return lipgloss.Color("#1c153d")
}

// SelectionBg returns the recommended background color for selected text.
func (t *TintScarletProtocol) SelectionBg() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

// Cursor returns the recommended color for the cursor.
func (t *TintScarletProtocol) Cursor() lipgloss.TerminalColor {
	return lipgloss.NoColor{}
}

func (t *TintScarletProtocol) BrightBlack() lipgloss.TerminalColor {
	return lipgloss.Color("#686868")
}

func (t *TintScarletProtocol) BrightBlue() lipgloss.TerminalColor {
	return lipgloss.Color("#6871ff")
}

func (t *TintScarletProtocol) BrightCyan() lipgloss.TerminalColor {
	return lipgloss.Color("#60fdff")
}

func (t *TintScarletProtocol) BrightGreen() lipgloss.TerminalColor {
	return lipgloss.Color("#5ffa68")
}

func (t *TintScarletProtocol) BrightPurple() lipgloss.TerminalColor {
	return lipgloss.Color("#bd35ec")
}

func (t *TintScarletProtocol) BrightRed() lipgloss.TerminalColor {
	return lipgloss.Color("#ff6e67")
}

func (t *TintScarletProtocol) BrightWhite() lipgloss.TerminalColor {
	return lipgloss.Color("#ffffff")
}

func (t *TintScarletProtocol) BrightYellow() lipgloss.TerminalColor {
	return lipgloss.Color("#fffc67")
}

func (t *TintScarletProtocol) Black() lipgloss.TerminalColor {
	return lipgloss.Color("#101116")
}

func (t *TintScarletProtocol) Blue() lipgloss.TerminalColor {
	return lipgloss.Color("#0271b6")
}

func (t *TintScarletProtocol) Cyan() lipgloss.TerminalColor {
	return lipgloss.Color("#00c5c7")
}

func (t *TintScarletProtocol) Green() lipgloss.TerminalColor {
	return lipgloss.Color("#00dc84")
}

func (t *TintScarletProtocol) Purple() lipgloss.TerminalColor {
	return lipgloss.Color("#ca30c7")
}

func (t *TintScarletProtocol) Red() lipgloss.TerminalColor {
	return lipgloss.Color("#ff0051")
}

func (t *TintScarletProtocol) White() lipgloss.TerminalColor {
	return lipgloss.Color("#c7c7c7")
}

func (t *TintScarletProtocol) Yellow() lipgloss.TerminalColor {
	return lipgloss.Color("#faf945")
}