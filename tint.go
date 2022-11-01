// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bubbletint

import (
	"sync/atomic"

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

var currentTint = atomic.Pointer[string]{}

// SetTint sets the current tint to the provided tint, and adds it to the list of
// registered tints if it isn't already.
func SetTint(tint Tint) {
	Register(tint) // Register if not already done.
	id := tint.ID()
	currentTint.Store(&id)
}

// SetTintID sets the current tint to the provided tint ID.
func SetTintID(id string) (ok bool) {
	_, ok = GetTint(id)
	if !ok {
		return ok
	}

	currentTint.Store(&id)
	return true
}

// PreviousTint sets the current tint to the previous tint in the list of
// registered tints. If the current tint is the first tint in the list, it will
// pin to the first tint in the list.
//
// PreviousTint uses a sorted list of tint IDs.
func PreviousTint() {
	id := ID()
	tints := TintIDs()

	if len(tints) == 0 {
		return
	}

	if id == "" {
		SetTintID(tints[0])
		return
	}

	currentI := 0
	for i, t := range tints {
		if t == id {
			currentI = i
			break
		}
	}

	prevI := currentI - 1
	if prevI < 1 {
		prevI = 0
	}

	SetTintID(tints[prevI])
}

// NextTint sets the current tint to the next tint in the list of registered
// tints. If the current tint is the last tint in the list, it will pin to the
// last tint in the list.
//
// NextTint uses a sorted list of tint IDs.
func NextTint() {
	id := ID()
	tints := TintIDs()

	if len(tints) == 0 {
		return
	}

	if id == "" {
		SetTintID(tints[0])
		return
	}

	currentI := 0
	for i, t := range tints {
		if t == id {
			currentI = i
			break
		}
	}

	nextI := currentI + 1
	if nextI >= len(tints) {
		nextI = len(tints) - 1
	}

	SetTintID(tints[nextI])
}

func GetCurrentTint() (tint Tint) {
	id := ID()
	if id == "" {
		panic("no tint set")
	}

	tint, _ = GetTint(id)
	return tint
}

// DisplayName returns the display name of the tint.
func DisplayName() string {
	return GetCurrentTint().DisplayName()
}

// ID returns the name of the tint (normalized, snakecase style).
func ID() string {
	return *currentTint.Load()
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func About() string {
	return GetCurrentTint().About()
}

// Fg returns the recommended default foreground color for this tint.
func Fg() lipgloss.TerminalColor {
	return GetCurrentTint().Fg()
}

// Bg returns the recommended default background color for this tint.
func Bg() lipgloss.TerminalColor {
	return GetCurrentTint().Bg()
}

// SelectionBg returns the recommended background color for selected text.
func SelectionBg() lipgloss.TerminalColor {
	return GetCurrentTint().SelectionBg()
}

// Cursor returns the recommended color for the cursor.
func Cursor() lipgloss.TerminalColor {
	return GetCurrentTint().Cursor()
}

func BrightBlack() lipgloss.TerminalColor {
	return GetCurrentTint().BrightBlack()
}

func BrightBlue() lipgloss.TerminalColor {
	return GetCurrentTint().BrightBlue()
}

func BrightCyan() lipgloss.TerminalColor {
	return GetCurrentTint().BrightCyan()
}

func BrightGreen() lipgloss.TerminalColor {
	return GetCurrentTint().BrightGreen()
}

func BrightPurple() lipgloss.TerminalColor {
	return GetCurrentTint().BrightPurple()
}

func BrightRed() lipgloss.TerminalColor {
	return GetCurrentTint().BrightRed()
}

func BrightWhite() lipgloss.TerminalColor {
	return GetCurrentTint().BrightWhite()
}

func BrightYellow() lipgloss.TerminalColor {
	return GetCurrentTint().BrightYellow()
}

func Black() lipgloss.TerminalColor {
	return GetCurrentTint().Black()
}

func Blue() lipgloss.TerminalColor {
	return GetCurrentTint().Blue()
}

func Cyan() lipgloss.TerminalColor {
	return GetCurrentTint().Cyan()
}

func Green() lipgloss.TerminalColor {
	return GetCurrentTint().Green()
}

func Purple() lipgloss.TerminalColor {
	return GetCurrentTint().Purple()
}

func Red() lipgloss.TerminalColor {
	return GetCurrentTint().Red()
}

func White() lipgloss.TerminalColor {
	return GetCurrentTint().White()
}

func Yellow() lipgloss.TerminalColor {
	return GetCurrentTint().Yellow()
}
