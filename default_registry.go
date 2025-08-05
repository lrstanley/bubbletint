// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

// Register registers one or more tints, using the default registry. It does not
// change to the provided tint.
func Register(tint ...*Tint) {
	DefaultRegistry.Register(tint...)
}

// Unregister unregisters one or more provided tints, using the default registry.
func Unregister(tint ...*Tint) {
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
func Tints() []*Tint {
	return DefaultRegistry.Tints()
}

// TintIDs returns a list of all registered tint IDs, sorted alphabetically by their
// ID, using the default registry.
func TintIDs() (ids []string) { //nolint:revive
	return DefaultRegistry.TintIDs()
}

// GetTint returns the tint with the provided ID, or nil if it doesn't exist, using
// the default registry.
func GetTint(id string) (tint *Tint, ok bool) {
	return DefaultRegistry.GetTint(id)
}

// SetTint sets the current tint to the provided tint, and adds it to the list of
// registered tints if it isn't already, using the default registry. Panics if the
// provided tint is nil.
func SetTint(tint *Tint) {
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

// Current returns the current tint, using the default registry. Panics if no tint
// has been set yet, and the registry has no tints registered that it can fall back
// to.
func Current() (tint *Tint) {
	return DefaultRegistry.Current()
}
