// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bubbletint

import (
	"github.com/charmbracelet/lipgloss"
)

// Tint is an interface that represents each tint in this package.
type Tint interface {
	// DisplayName returns the display name of the tint.
	DisplayName() string

	// ID returns the name of the tint (normalized, snakecase style).
	ID() string

	// About returns information about the tint (and if we have credit for who
	// assisted with/created it).
	About() string

	// Fg returns the recommended default foreground color for this tint.
	Fg() lipgloss.TerminalColor

	// Bg returns the recommended default background color for this tint.
	Bg() lipgloss.TerminalColor

	// SelectionBg returns the recommended background color for selected text.
	SelectionBg() lipgloss.TerminalColor

	// Cursor returns the recommended color for the cursor.
	Cursor() lipgloss.TerminalColor

	BrightBlack() lipgloss.TerminalColor
	BrightBlue() lipgloss.TerminalColor
	BrightCyan() lipgloss.TerminalColor
	BrightGreen() lipgloss.TerminalColor
	BrightPurple() lipgloss.TerminalColor
	BrightRed() lipgloss.TerminalColor
	BrightWhite() lipgloss.TerminalColor
	BrightYellow() lipgloss.TerminalColor

	Black() lipgloss.TerminalColor
	Blue() lipgloss.TerminalColor
	Cyan() lipgloss.TerminalColor
	Green() lipgloss.TerminalColor
	Purple() lipgloss.TerminalColor
	Red() lipgloss.TerminalColor
	White() lipgloss.TerminalColor
	Yellow() lipgloss.TerminalColor
}
