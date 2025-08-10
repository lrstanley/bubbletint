// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"sort"
	"sync"
)

// NewRegistry returns a new registry, with the provided tints registered.
func NewRegistry(defaultTint *Tint, tints ...*Tint) *Registry {
	r := &Registry{
		tints: make(map[string]*Tint),
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
	tints       map[string]*Tint
}

// Register registers one or more tints within the registry. It does not change
// to the provided tint.
func (r *Registry) Register(tint ...*Tint) {
	r.mu.Lock()
	for _, t := range tint {
		r.tints[t.ID] = t
	}
	r.mu.Unlock()
}

// Unregister unregisters one or more provided tints from the registry.
func (r *Registry) Unregister(tint ...*Tint) {
	r.mu.Lock()
	for _, t := range tint {
		delete(r.tints, t.ID)
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
func (r *Registry) Tints() []*Tint {
	v := []*Tint{}

	r.mu.RLock()
	for _, tint := range r.tints {
		v = append(v, tint)
	}
	r.mu.RUnlock()

	sort.Slice(v, func(i, j int) bool {
		return v[i].ID < v[j].ID
	})

	return v
}

// TintIDs returns a list of all registered tint IDs, sorted alphabetically by their
// ID.
func (r *Registry) TintIDs() (ids []string) {
	r.mu.RLock()
	for _, tint := range r.tints {
		ids = append(ids, tint.ID)
	}
	r.mu.RUnlock()
	sort.Strings(ids)
	return ids
}

// GetTint returns the tint with the provided ID, or nil if it doesn't exist.
func (r *Registry) GetTint(id string) (tint *Tint, ok bool) {
	r.mu.RLock()
	tint, ok = r.tints[id]
	r.mu.RUnlock()
	return tint, ok
}

// SetTint sets the current tint to the provided tint, and adds it to the list of
// registered tints if it isn't already. Panics if the provided tint is nil.
func (r *Registry) SetTint(tint *Tint) {
	if tint == nil {
		panic("tint is nil") //nolint:panic
	}

	r.Register(tint) // Register if not already done.
	id := tint.ID

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
	current := r.getCurrentTintID()

	tints := r.TintIDs()

	if len(tints) == 0 {
		return
	}

	if current == "" {
		r.SetTintID(tints[0])
		return
	}

	currentI := 0
	for i, t := range tints {
		if t == current {
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
	current := r.getCurrentTintID()

	tints := r.TintIDs()

	if len(tints) == 0 {
		return
	}

	if current == "" {
		r.SetTintID(tints[0])
		return
	}

	currentI := 0
	for i, t := range tints {
		if t == current {
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

// Current returns the current tint. Panics if no tint has been set yet, and the
// registry has no tints registered that it can fall back to.
func (r *Registry) Current() (tint *Tint) {
	current := r.getCurrentTintID()
	if current == "" {
		// Attempt to recover by returning the first tint.
		tints := r.Tints()
		if len(tints) == 0 {
			panic("no tint set, and no tints registered to fall back to")
		}
		r.SetTint(tint)
		return tints[0]
	}
	tint, _ = r.GetTint(current)
	return tint
}

func (r *Registry) getCurrentTintID() (id string) {
	r.mu.RLock()
	id = r.currentTint
	r.mu.RUnlock()
	return id
}
