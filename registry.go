// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"slices"
	"sort"
	"sync"
)

// NewRegistry returns a new registry, with the provided tints registered.
// When tints is empty, the built-in catalog is registered. When defaultTint
// is nil, the first registered tint becomes current.
func NewRegistry(defaultTint *Tint, tints ...*Tint) *Registry {
	r := &Registry{
		tints: make(map[string]*Tint),
	}

	if len(tints) == 0 {
		tints = DefaultTints()
	}

	r.Register(tints...)

	if defaultTint != nil {
		r.SetTint(defaultTint)
	} else {
		r.SetTint(tints[0])
	}

	return r
}

// Registry holds a set of tints, and tracks which one is current. Changes to it
// can be observed with [Registry.Subscribe]. A Registry is safe for concurrent
// use.
type Registry struct {
	mu          sync.RWMutex
	currentTint string
	tints       map[string]*Tint
	origins     map[string]string // Tint ID -> name of the [Source] that registered it.
	attached    map[string]*attachment
	subs        subscribers
}

// Register registers one or more tints within the registry. It does not change
// to the provided tint. Registering a tint with an ID that is already registered
// replaces the previous tint.
func (r *Registry) Register(tint ...*Tint) {
	var events []Event

	r.mu.Lock()
	for _, t := range tint {
		events = append(events, r.registerLocked("", t)...)
	}
	r.mu.Unlock()

	r.subs.emit(events)
}

// registerLocked registers t, returning the events that resulted. r.mu must be
// held for writing.
func (r *Registry) registerLocked(source string, t *Tint) (events []Event) {
	previous, exists := r.tints[t.ID]
	if exists && previous == t {
		return nil
	}

	r.tints[t.ID] = t

	if source == "" {
		delete(r.origins, t.ID)
	} else {
		if r.origins == nil {
			r.origins = make(map[string]string)
		}
		r.origins[t.ID] = source
	}

	if !exists {
		return []Event{{Kind: TintAdded, Tint: t, Source: source}}
	}

	events = append(events, Event{Kind: TintUpdated, Tint: t, Previous: previous, Source: source})
	if r.currentTint == t.ID {
		events = append(events, Event{Kind: CurrentChanged, Tint: t, Previous: previous, Source: source})
	}
	return events
}

// idsFromLocked returns the sorted IDs of the tints registered by source. r.mu
// must be held.
func (r *Registry) idsFromLocked(source string) []string {
	var ids []string
	for id, origin := range r.origins {
		if origin == source {
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	return ids
}

// Unregister unregisters one or more provided tints from the registry. If the
// current tint is unregistered, the registry falls back to the next tint (by
// sorted ID), or the previous one if there is no next tint, so [Registry.Current]
// never returns nil unless the registry is empty.
func (r *Registry) Unregister(tint ...*Tint) {
	ids := make([]string, 0, len(tint))
	for _, t := range tint {
		ids = append(ids, t.ID)
	}
	r.UnregisterIDs(ids...)
}

// UnregisterIDs unregisters the tints with the provided ID(s) from the registry.
// See [Registry.Unregister] for how the current tint is handled.
func (r *Registry) UnregisterIDs(id ...string) {
	r.mu.Lock()
	events := r.unregisterLocked(id...)
	r.mu.Unlock()

	r.subs.emit(events)
}

// UnregisterAll unregisters all tints from the registry. As there is nothing to
// fall back to, [Registry.Current] panics until a tint is registered.
func (r *Registry) UnregisterAll() {
	r.mu.Lock()
	ids := make([]string, 0, len(r.tints))
	for id := range r.tints {
		ids = append(ids, id)
	}
	events := r.unregisterLocked(ids...)
	r.mu.Unlock()

	r.subs.emit(events)
}

// unregisterLocked removes the provided IDs, returning the events that resulted.
// r.mu must be held for writing.
func (r *Registry) unregisterLocked(ids ...string) (events []Event) {
	previous := r.tints[r.currentTint]
	removedCurrent := false

	for _, id := range ids {
		t, ok := r.tints[id]
		if !ok {
			continue
		}

		delete(r.tints, id)
		events = append(events, Event{Kind: TintRemoved, Tint: t, Source: r.origins[id]})
		delete(r.origins, id)

		if id == r.currentTint {
			removedCurrent = true
		}
	}

	if !removedCurrent {
		return events
	}

	// The current tint is gone. Fall back to its sorted neighbour: the next tint,
	// or the last one if the removed tint sorted after all of the remaining ones.
	removedID := r.currentTint
	r.currentTint = ""

	remaining := make([]string, 0, len(r.tints))
	for id := range r.tints {
		remaining = append(remaining, id)
	}
	sort.Strings(remaining)

	var next *Tint
	if len(remaining) > 0 {
		i, _ := slices.BinarySearch(remaining, removedID)
		if i >= len(remaining) {
			i = len(remaining) - 1
		}
		next = r.tints[remaining[i]]
		r.currentTint = next.ID
	}

	return append(events, Event{Kind: CurrentChanged, Tint: next, Previous: previous})
}

// Tints returns a list of all registered tints, sorted alphabetically by their ID.
func (r *Registry) Tints() []*Tint {
	return r.list(func(*Tint) bool { return true })
}

// TintsByHue returns registered tints whose Hue equals hue, sorted by ID.
// An empty hue matches tints that did not set one.
func (r *Registry) TintsByHue(hue Hue) []*Tint {
	return r.list(func(t *Tint) bool { return t.Hue == hue })
}

// TintsByTone returns registered tints whose Tone equals tone, sorted by ID.
// An empty tone matches tints with no tone.
func (r *Registry) TintsByTone(tone Tone) []*Tint {
	return r.list(func(t *Tint) bool { return t.Tone == tone })
}

// TintsByHueAndTone returns registered tints matching both hue and tone, sorted
// by ID. An empty hue or tone matches only tints that did not set that field.
func (r *Registry) TintsByHueAndTone(hue Hue, tone Tone) []*Tint {
	return r.list(func(t *Tint) bool { return t.Hue == hue && t.Tone == tone })
}

// list returns registered tints for which match reports true, sorted by ID.
func (r *Registry) list(match func(*Tint) bool) []*Tint {
	v := []*Tint{}

	r.mu.RLock()
	for _, tint := range r.tints {
		if match(tint) {
			v = append(v, tint)
		}
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

	r.mu.Lock()
	previousID := r.currentTint
	previous := r.tints[previousID]

	events := r.registerLocked("", tint)
	if previousID != tint.ID {
		r.currentTint = tint.ID
		events = append(events, Event{Kind: CurrentChanged, Tint: tint, Previous: previous})
	}
	r.mu.Unlock()

	r.subs.emit(events)
}

// SetTintID sets the current tint to the provided tint ID.
func (r *Registry) SetTintID(id string) (ok bool) {
	r.mu.Lock()
	tint, ok := r.tints[id]
	if !ok {
		r.mu.Unlock()
		return false
	}

	var events []Event
	if r.currentTint != id {
		events = []Event{{Kind: CurrentChanged, Tint: tint, Previous: r.tints[r.currentTint]}}
		r.currentTint = id
	}
	r.mu.Unlock()

	r.subs.emit(events)
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

// Current returns the current tint. If no tint has been set yet, it falls back
// to (and sets) the first registered tint. Panics if no tints are registered.
func (r *Registry) Current() (tint *Tint) {
	r.mu.RLock()
	tint = r.tints[r.currentTint]
	r.mu.RUnlock()

	if tint != nil {
		return tint
	}

	tints := r.Tints()
	if len(tints) == 0 {
		panic("no tint set, and no tints registered to fall back to") //nolint:panic
	}
	r.SetTint(tints[0])
	return tints[0]
}

func (r *Registry) getCurrentTintID() (id string) {
	r.mu.RLock()
	id = r.currentTint
	r.mu.RUnlock()
	return id
}
