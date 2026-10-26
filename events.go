// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"slices"
	"sync"
)

// EventKind describes what happened to a [Registry].
type EventKind uint8

const (
	// TintAdded is emitted when a tint is registered under an ID that wasn't
	// previously registered.
	TintAdded EventKind = iota + 1

	// TintUpdated is emitted when a tint is registered under an ID that already
	// had a different tint registered. [Event.Previous] holds the old tint.
	TintUpdated

	// TintRemoved is emitted when a tint is unregistered.
	TintRemoved

	// CurrentChanged is emitted when the tint returned by [Registry.Current]
	// changes. This includes the current tint being replaced in-place by a
	// [TintUpdated] (same ID, new tint), or the current tint being unregistered
	// and the registry falling back to another tint. [Event.Tint] is nil only
	// when the last tint was removed. This is the event to listen for if all
	// you need to know is when to rebuild your styles.
	CurrentChanged
)

// String returns the name of the event kind.
func (k EventKind) String() string {
	switch k {
	case TintAdded:
		return "TintAdded"
	case TintUpdated:
		return "TintUpdated"
	case TintRemoved:
		return "TintRemoved"
	case CurrentChanged:
		return "CurrentChanged"
	default:
		return "Unknown"
	}
}

// Event describes a single change to a [Registry].
type Event struct {
	// Kind is what happened.
	Kind EventKind

	// Tint is the tint the event is about: the added/updated tint, the removed
	// tint, or the new current tint. It is nil only for a [CurrentChanged] that
	// left the registry without any tints.
	Tint *Tint

	// Previous is the tint that was replaced, for [TintUpdated] and
	// [CurrentChanged]. It is nil when there was no previous tint.
	Previous *Tint

	// Source is the name of the [Source] that caused the event, or empty if
	// the change was made directly (e.g. via [Registry.Register]).
	Source string
}

type subscription struct {
	id uint64
	fn func(Event)
}

type subscribers struct {
	mu   sync.RWMutex
	next uint64
	subs []subscription
}

func (s *subscribers) add(fn func(Event)) (cancel func()) {
	s.mu.Lock()
	s.next++
	id := s.next
	s.subs = append(s.subs, subscription{id: id, fn: fn})
	s.mu.Unlock()

	return func() {
		s.mu.Lock()
		s.subs = slices.DeleteFunc(s.subs, func(sub subscription) bool { return sub.id == id })
		s.mu.Unlock()
	}
}

// emit delivers events to every subscriber, in subscription order. It must not
// be called with the registry lock held.
func (s *subscribers) emit(events []Event) {
	if len(events) == 0 {
		return
	}

	s.mu.RLock()
	subs := slices.Clone(s.subs)
	s.mu.RUnlock()

	for _, ev := range events {
		for _, sub := range subs {
			sub.fn(ev)
		}
	}
}

// Subscribe registers fn to be called for every change to the registry (tints
// being added, updated, removed, or the current tint changing), and returns a
// function that removes the subscription. The returned function is safe to call
// more than once.
//
// fn is called synchronously, on the goroutine that made the change, after the
// registry's lock has been released. This means fn is free to call back into
// the registry, but it should return quickly: a slow fn delays whoever made the
// change. Changes made concurrently from different goroutines may reach fn
// concurrently, and without any ordering guarantee between them, so fn must be
// safe for concurrent use. To receive events from within Bubble Tea, have fn
// forward them with Bubble Tea's Program.Send.
//
// Panics if fn is nil.
func (r *Registry) Subscribe(fn func(Event)) (cancel func()) {
	if fn == nil {
		panic("tint: subscriber is nil") //nolint:panic
	}
	return sync.OnceFunc(r.subs.add(fn))
}
