// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bubbletint

import (
	"sort"
	"sync"
)

var registryLock = sync.RWMutex{}

func Register(tint Tint) {
	registryLock.Lock()
	defaultRegistry[tint.ID()] = tint
	registryLock.Unlock()
}

func Unregister(tint Tint) {
	registryLock.Lock()
	delete(defaultRegistry, tint.ID())
	registryLock.Unlock()
}

func UnregisterName(id string) {
	registryLock.Lock()
	delete(defaultRegistry, id)
	registryLock.Unlock()
}

func UnregisterAll() {
	registryLock.Lock()
	defaultRegistry = map[string]Tint{}
	registryLock.Unlock()
}

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

func TintIDs() (ids []string) {
	registryLock.RLock()
	for _, tint := range defaultRegistry {
		ids = append(ids, tint.ID())
	}
	registryLock.RUnlock()

	sort.Strings(ids)
	return ids
}

func GetTint(id string) (tint Tint, ok bool) {
	registryLock.RLock()
	tint, ok = defaultRegistry[id]
	registryLock.RUnlock()
	return tint, ok
}
