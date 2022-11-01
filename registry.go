// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bubbletint

import (
	"sort"
	"sync"
)

var registryLock = sync.RWMutex{}

// Register registers a tint within the builtin registry. It does not change to the
// provided tint.
func Register(tint Tint) {
	registryLock.Lock()
	defaultRegistry[tint.ID()] = tint
	registryLock.Unlock()
}

// Unregister unregisters the provided tint from the builtin registry.
func Unregister(tint Tint) {
	registryLock.Lock()
	delete(defaultRegistry, tint.ID())
	registryLock.Unlock()
}

// UnregisterName unregisters the tint with the provided name from the builtin
// registry.
func UnregisterName(id string) {
	registryLock.Lock()
	delete(defaultRegistry, id)
	registryLock.Unlock()
}

// UnregisterAll unregisters all tints from the builtin registry.
func UnregisterAll() {
	registryLock.Lock()
	defaultRegistry = map[string]Tint{}
	registryLock.Unlock()
}

// Tints returns a list of all registered tints, sorted alphabetically by their ID.
func Tints() []Tint {
	v := []Tint{}

	registryLock.RLock()
	for _, tint := range defaultRegistry {
		v = append(v, tint)
	}
	registryLock.RUnlock()

	sort.Slice(v, func(i, j int) bool {
		return v[i].ID() < v[j].ID()
	})

	return v
}

// TintIDs returns a list of all registered tint IDs, sorted alphabetically by their
// ID.
func TintIDs() (ids []string) {
	registryLock.RLock()
	for _, tint := range defaultRegistry {
		ids = append(ids, tint.ID())
	}
	registryLock.RUnlock()

	sort.Strings(ids)
	return ids
}

// GetTint returns the tint with the provided ID, or nil if it doesn't exist.
func GetTint(id string) (tint Tint, ok bool) {
	registryLock.RLock()
	tint, ok = defaultRegistry[id]
	registryLock.RUnlock()
	return tint, ok
}
