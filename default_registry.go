// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bubbletint

import (
	"github.com/charmbracelet/lipgloss"
)

// Register registers one or more tints, using the default registry. It does not
// change to the provided tint.
func Register(tint ...Tint) {
	DefaultRegistry.Register(tint...)
}

// Unregister unregisters one or more provided tints, using the default registry.
func Unregister(tint ...Tint) {
	DefaultRegistry.Unregister(tint...)
}

// UnregisterIDs unregisters the tints with the provided ID(s), using the default
// registry.
func UnregisterIDs(id ...string) {
	DefaultRegistry.UnregisterIDs(id...)
}

// UnregisterAll unregisters all tints, using the default registry.
func UnregisterAll() {
	DefaultRegistry.UnregisterAll()
}

// Tints returns a list of all registered tints, sorted alphabetically by their ID,
// using the default registry.
func Tints() []Tint {
	return DefaultRegistry.Tints()
}

// TintIDs returns a list of all registered tint IDs, sorted alphabetically by their
// ID, using the default registry.
func TintIDs() (ids []string) {
	return DefaultRegistry.TintIDs()
}

// GetTint returns the tint with the provided ID, or nil if it doesn't exist, using
// the default registry.
func GetTint(id string) (tint Tint, ok bool) {
	return DefaultRegistry.GetTint(id)
}

// SetTint sets the current tint to the provided tint, and adds it to the list of
// registered tints if it isn't already, using the default registry.
func SetTint(tint Tint) {
	DefaultRegistry.SetTint(tint)
}

// SetTintID sets the current tint to the provided tint ID, using the default
// registry.
func SetTintID(id string) (ok bool) {
	return DefaultRegistry.SetTintID(id)
}

// PreviousTint sets the current tint to the previous tint in the list of
// registered tints, using the default registry. If the current tint is the first
// tint in the list, it will pin to the first tint in the list (i.e. won't rollback
// to the end).
//
// PreviousTint uses a sorted list of tint IDs.
func PreviousTint() {
	DefaultRegistry.PreviousTint()
}

// NextTint sets the current tint to the next tint in the list of registered
// tints, using the default registry. If the current tint is the last tint in the
// list, it will pin to the last tint in the list (i.e. won't rollback to the start).
//
// NextTint uses a sorted list of tint IDs.
func NextTint() {
	DefaultRegistry.NextTint()
}

// GetCurrenTint returns the current tint, using the default registry.
func GetCurrentTint() (tint Tint) {
	return DefaultRegistry.GetCurrentTint()
}

// DisplayName returns the display name of the tint, using the default registry.
func DisplayName() string {
	return DefaultRegistry.GetCurrentTint().DisplayName()
}

// ID returns the name of the tint (normalized, snakecase style), using the default
// registry.
func ID() string {
	return DefaultRegistry.ID()
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it), using the default registry.
func About() string {
	return DefaultRegistry.GetCurrentTint().About()
}

// Fg returns the recommended default foreground color for this tint, using the
// default registry.
func Fg() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Fg()
}

// Bg returns the recommended default background color for this tint, using the
// default registry.
func Bg() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Bg()
}

// SelectionBg returns the recommended background color for selected text, using
// the default registry.
func SelectionBg() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().SelectionBg()
}

// Cursor returns the recommended color for the cursor, using the default registry.
func Cursor() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Cursor()
}

// BrightBlack returns the recommended color for bright black, using the default
// registry.
func BrightBlack() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().BrightBlack()
}

// BrightBlue returns the recommended color for bright blue, using the default
// registry.
func BrightBlue() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().BrightBlue()
}

// BrightCyan returns the recommended color for bright cyan, using the default
// registry.
func BrightCyan() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().BrightCyan()
}

// BrightGreen returns the recommended color for bright green, using the default
// registry.
func BrightGreen() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().BrightGreen()
}

// BrightPurple returns the recommended color for bright purple, using the default
// registry.
func BrightPurple() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().BrightPurple()
}

// BrightRed returns the recommended color for bright red, using the default
// registry.
func BrightRed() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().BrightRed()
}

// BrightWhite returns the recommended color for bright white, using the default
// registry.
func BrightWhite() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().BrightWhite()
}

// BrightYellow returns the recommended color for bright yellow, using the default
// registry.
func BrightYellow() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().BrightYellow()
}

// Black returns the recommended color for black, using the default registry.
func Black() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Black()
}

// Blue returns the recommended color for blue, using the default registry.
func Blue() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Blue()
}

// Cyan returns the recommended color for cyan, using the default registry.
func Cyan() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Cyan()
}

// Green returns the recommended color for green, using the default registry.
func Green() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Green()
}

// Purple returns the recommended color for purple, using the default registry.
func Purple() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Purple()
}

// Red returns the recommended color for red, using the default registry.
func Red() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Red()
}

// White returns the recommended color for white, using the default registry.
func White() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().White()
}

// Yellow returns the recommended color for yellow, using the default registry.
func Yellow() lipgloss.TerminalColor {
	return DefaultRegistry.GetCurrentTint().Yellow()
}
