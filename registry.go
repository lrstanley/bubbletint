// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"sort"
	"sync"

	"github.com/charmbracelet/lipgloss"
)

// NewRegistry returns a new Registry, with the provided tints registered.
func NewRegistry(defaultTint Tint, tints ...Tint) *Registry {
	r := &Registry{
		tints: make(map[string]Tint),
	}

	r.Register(tints...)

	if defaultTint != nil {
		r.SetTint(defaultTint)
	}

	return r
}

type Registry struct {
	mu          sync.RWMutex
	currentTint string
	tints       map[string]Tint
}

// Register registers one or more tints within the registry. It does not change
// to the provided tint.
func (r *Registry) Register(tint ...Tint) {
	r.mu.Lock()
	for _, t := range tint {
		r.tints[t.ID()] = t
	}
	r.mu.Unlock()
}

// Unregister unregisters one or more provided tints from the registry.
func (r *Registry) Unregister(tint ...Tint) {
	r.mu.Lock()
	for _, t := range tint {
		delete(r.tints, t.ID())
	}
	r.mu.Unlock()
}

// UnregisterIDs unregisters the tints with the provided ID(s) from the registry.
func (r *Registry) UnregisterIDs(id ...string) {
	r.mu.Lock()
	for _, t := range id {
		delete(r.tints, t)
	}
	r.mu.Unlock()
}

// UnregisterAll unregisters all tints from the registry.
func (r *Registry) UnregisterAll() {
	r.mu.Lock()
	for id := range r.tints {
		delete(r.tints, id)
	}
	r.mu.Unlock()
}

// Tints returns a list of all registered tints, sorted alphabetically by their ID.
func (r *Registry) Tints() []Tint {
	v := []Tint{}

	r.mu.RLock()
	for _, tint := range r.tints {
		v = append(v, tint)
	}
	r.mu.RUnlock()

	sort.Slice(v, func(i, j int) bool {
		return v[i].ID() < v[j].ID()
	})

	return v
}

// TintIDs returns a list of all registered tint IDs, sorted alphabetically by their
// ID.
func (r *Registry) TintIDs() (ids []string) {
	r.mu.RLock()
	for _, tint := range r.tints {
		ids = append(ids, tint.ID())
	}
	r.mu.RUnlock()

	sort.Strings(ids)
	return ids
}

// GetTint returns the tint with the provided ID, or nil if it doesn't exist.
func (r *Registry) GetTint(id string) (tint Tint, ok bool) {
	r.mu.RLock()
	tint, ok = r.tints[id]
	r.mu.RUnlock()
	return tint, ok
}

// SetTint sets the current tint to the provided tint, and adds it to the list of
// registered tints if it isn't already.
func (r *Registry) SetTint(tint Tint) {
	r.Register(tint) // Register if not already done.
	id := tint.ID()

	r.mu.Lock()
	r.currentTint = id
	r.mu.Unlock()
}

// SetTintID sets the current tint to the provided tint ID.
func (r *Registry) SetTintID(id string) (ok bool) {
	_, ok = r.GetTint(id)
	if !ok {
		return ok
	}

	r.mu.Lock()
	r.currentTint = id
	r.mu.Unlock()

	return true
}

// PreviousTint sets the current tint to the previous tint in the list of
// registered tints. If the current tint is the first tint in the list, it will
// pin to the first tint in the list (i.e. won't rollback to the end).
//
// PreviousTint uses a sorted list of tint IDs.
func (r *Registry) PreviousTint() {
	id := r.ID()

	tints := r.TintIDs()

	if len(tints) == 0 {
		return
	}

	if id == "" {
		r.SetTintID(tints[0])
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

	r.SetTintID(tints[prevI])
}

// NextTint sets the current tint to the next tint in the list of registered
// tints. If the current tint is the last tint in the list, it will pin to the
// last tint in the list (i.e. won't rollback to the start).
//
// NextTint uses a sorted list of tint IDs.
func (r *Registry) NextTint() {
	id := r.ID()

	tints := r.TintIDs()

	if len(tints) == 0 {
		return
	}

	if id == "" {
		r.SetTintID(tints[0])
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

	r.SetTintID(tints[nextI])
}

// GetCurrenTint returns the current tint.
func (r *Registry) GetCurrentTint() (tint Tint) {
	id := r.ID()

	if id == "" {
		panic("no tint set")
	}

	tint, _ = r.GetTint(id)
	return tint
}

// DisplayName returns the display name of the tint.
func (r *Registry) DisplayName() string {
	return r.GetCurrentTint().DisplayName()
}

// ID returns the name of the current tint (normalized, snakecase style).
func (r *Registry) ID() (id string) {
	r.mu.RLock()
	id = r.currentTint
	r.mu.RUnlock()

	return id
}

// About returns information about the tint (and if we have credit for who
// assisted with/created it).
func (r *Registry) About() string {
	return r.GetCurrentTint().About()
}

// Fg returns the recommended default foreground color for this tint.
func (r *Registry) Fg() lipgloss.TerminalColor {
	return r.GetCurrentTint().Fg()
}

// Bg returns the recommended default background color for this tint.
func (r *Registry) Bg() lipgloss.TerminalColor {
	return r.GetCurrentTint().Bg()
}

// SelectionBg returns the recommended background color for selected text, using
func (r *Registry) SelectionBg() lipgloss.TerminalColor {
	return r.GetCurrentTint().SelectionBg()
}

// Cursor returns the recommended color for the cursor.
func (r *Registry) Cursor() lipgloss.TerminalColor {
	return r.GetCurrentTint().Cursor()
}

// BrightBlack returns the recommended color for bright black.
func (r *Registry) BrightBlack() lipgloss.TerminalColor {
	return r.GetCurrentTint().BrightBlack()
}

// BrightBlue returns the recommended color for bright blue.
func (r *Registry) BrightBlue() lipgloss.TerminalColor {
	return r.GetCurrentTint().BrightBlue()
}

// BrightCyan returns the recommended color for bright cyan.
func (r *Registry) BrightCyan() lipgloss.TerminalColor {
	return r.GetCurrentTint().BrightCyan()
}

// BrightGreen returns the recommended color for bright green.
func (r *Registry) BrightGreen() lipgloss.TerminalColor {
	return r.GetCurrentTint().BrightGreen()
}

// BrightPurple returns the recommended color for bright purple.
func (r *Registry) BrightPurple() lipgloss.TerminalColor {
	return r.GetCurrentTint().BrightPurple()
}

// BrightRed returns the recommended color for bright red.
func (r *Registry) BrightRed() lipgloss.TerminalColor {
	return r.GetCurrentTint().BrightRed()
}

// BrightWhite returns the recommended color for bright white.
func (r *Registry) BrightWhite() lipgloss.TerminalColor {
	return r.GetCurrentTint().BrightWhite()
}

// BrightYellow returns the recommended color for bright yellow.
func (r *Registry) BrightYellow() lipgloss.TerminalColor {
	return r.GetCurrentTint().BrightYellow()
}

// Black returns the recommended color for black.
func (r *Registry) Black() lipgloss.TerminalColor {
	return r.GetCurrentTint().Black()
}

// Blue returns the recommended color for blue.
func (r *Registry) Blue() lipgloss.TerminalColor {
	return r.GetCurrentTint().Blue()
}

// Cyan returns the recommended color for cyan.
func (r *Registry) Cyan() lipgloss.TerminalColor {
	return r.GetCurrentTint().Cyan()
}

// Green returns the recommended color for green.
func (r *Registry) Green() lipgloss.TerminalColor {
	return r.GetCurrentTint().Green()
}

// Purple returns the recommended color for purple.
func (r *Registry) Purple() lipgloss.TerminalColor {
	return r.GetCurrentTint().Purple()
}

// Red returns the recommended color for red.
func (r *Registry) Red() lipgloss.TerminalColor {
	return r.GetCurrentTint().Red()
}

// White returns the recommended color for white.
func (r *Registry) White() lipgloss.TerminalColor {
	return r.GetCurrentTint().White()
}

// Yellow returns the recommended color for yellow.
func (r *Registry) Yellow() lipgloss.TerminalColor {
	return r.GetCurrentTint().Yellow()
}
