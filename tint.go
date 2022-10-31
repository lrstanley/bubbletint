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
	Fg() lipgloss.Color

	// Bg returns the recommended default background color for this tint.
	Bg() lipgloss.Color

	// SelectionBg returns the recommended background color for selected text.
	SelectionBg() lipgloss.Color

	// Cursor returns the recommended color for the cursor.
	Cursor() lipgloss.Color

	BrightBlack() lipgloss.Color
	BrightBlue() lipgloss.Color
	BrightCyan() lipgloss.Color
	BrightGreen() lipgloss.Color
	BrightPurple() lipgloss.Color
	BrightRed() lipgloss.Color
	BrightWhite() lipgloss.Color
	BrightYellow() lipgloss.Color

	Black() lipgloss.Color
	Blue() lipgloss.Color
	Cyan() lipgloss.Color
	Green() lipgloss.Color
	Purple() lipgloss.Color
	Red() lipgloss.Color
	White() lipgloss.Color
	Yellow() lipgloss.Color
}

var currentTint = atomic.Pointer[string]{}

func SetTint(tint Tint) {
	id := tint.ID()
	currentTint.Store(&id)
}

func SetTintName(id string) (ok bool) {
	_, ok = GetTint(id)
	if !ok {
		return ok
	}

	currentTint.Store(&id)
	return true
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
func Fg() lipgloss.Color {
	return GetCurrentTint().Fg()
}

// Bg returns the recommended default background color for this tint.
func Bg() lipgloss.Color {
	return GetCurrentTint().Bg()
}

// SelectionBg returns the recommended background color for selected text.
func SelectionBg() lipgloss.Color {
	return GetCurrentTint().SelectionBg()
}

// Cursor returns the recommended color for the cursor.
func Cursor() lipgloss.Color {
	return GetCurrentTint().Cursor()
}

func BrightBlack() lipgloss.Color {
	return GetCurrentTint().BrightBlack()
}

func BrightBlue() lipgloss.Color {
	return GetCurrentTint().BrightBlue()
}

func BrightCyan() lipgloss.Color {
	return GetCurrentTint().BrightCyan()
}

func BrightGreen() lipgloss.Color {
	return GetCurrentTint().BrightGreen()
}

func BrightPurple() lipgloss.Color {
	return GetCurrentTint().BrightPurple()
}

func BrightRed() lipgloss.Color {
	return GetCurrentTint().BrightRed()
}

func BrightWhite() lipgloss.Color {
	return GetCurrentTint().BrightWhite()
}

func BrightYellow() lipgloss.Color {
	return GetCurrentTint().BrightYellow()
}

func Black() lipgloss.Color {
	return GetCurrentTint().Black()
}

func Blue() lipgloss.Color {
	return GetCurrentTint().Blue()
}

func Cyan() lipgloss.Color {
	return GetCurrentTint().Cyan()
}

func Green() lipgloss.Color {
	return GetCurrentTint().Green()
}

func Purple() lipgloss.Color {
	return GetCurrentTint().Purple()
}

func Red() lipgloss.Color {
	return GetCurrentTint().Red()
}

func White() lipgloss.Color {
	return GetCurrentTint().White()
}

func Yellow() lipgloss.Color {
	return GetCurrentTint().Yellow()
}
