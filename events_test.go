// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"slices"
	"sync"
	"testing"
)

// recorder collects events from a registry.
type recorder struct {
	mu     sync.Mutex
	events []Event
}

func (r *recorder) record(ev Event) {
	r.mu.Lock()
	r.events = append(r.events, ev)
	r.mu.Unlock()
}

// drain returns, and clears, the recorded events as "Kind:ID" strings (where ID
// is that of Event.Tint, or "<nil>").
func (r *recorder) drain() []string {
	r.mu.Lock()
	defer r.mu.Unlock()

	out := make([]string, 0, len(r.events))
	for _, ev := range r.events {
		id := "<nil>"
		if ev.Tint != nil {
			id = ev.Tint.ID
		}
		out = append(out, ev.Kind.String()+":"+id)
	}
	r.events = nil
	return out
}

func newTestTint(id string) *Tint {
	return &Tint{ID: id, DisplayName: id, Bg: FromHex("#000000"), Fg: FromHex("#ffffff")}
}

func newRecordedRegistry(t *testing.T, tints ...*Tint) (*Registry, *recorder) {
	t.Helper()

	reg := NewRegistry(nil, tints...)
	rec := &recorder{}
	cancel := reg.Subscribe(rec.record)
	t.Cleanup(cancel)
	return reg, rec
}

func assertEvents(t *testing.T, rec *recorder, want ...string) {
	t.Helper()

	got := rec.drain()
	if want == nil {
		want = []string{}
	}
	if !slices.Equal(got, want) {
		t.Fatalf("unexpected events:\n got: %v\nwant: %v", got, want)
	}
}

func TestRegistry_Events_Register(t *testing.T) {
	t.Parallel()

	a, b := newTestTint("a"), newTestTint("b")
	reg, rec := newRecordedRegistry(t, a)

	reg.Register(b)
	assertEvents(t, rec, "TintAdded:b")

	reg.Register(b) // Same pointer, nothing changed.
	assertEvents(t, rec)

	b2 := newTestTint("b")
	reg.Register(b2)
	assertEvents(t, rec, "TintUpdated:b")

	// Replacing the current tint also reports the current tint changing.
	a2 := newTestTint("a")
	reg.Register(a2)
	assertEvents(t, rec, "TintUpdated:a", "CurrentChanged:a")
	if reg.Current() != a2 {
		t.Fatalf("expected current tint to be the replacement")
	}
}

func TestRegistry_Events_SetTint(t *testing.T) {
	t.Parallel()

	a, b, c := newTestTint("a"), newTestTint("b"), newTestTint("c")
	reg, rec := newRecordedRegistry(t, a, b)

	reg.SetTint(b)
	assertEvents(t, rec, "CurrentChanged:b")

	reg.SetTint(b) // No change.
	assertEvents(t, rec)

	reg.SetTint(c) // Not yet registered.
	assertEvents(t, rec, "TintAdded:c", "CurrentChanged:c")

	if reg.SetTintID("missing") {
		t.Fatal("expected SetTintID to fail for an unknown ID")
	}
	assertEvents(t, rec)

	reg.SetTintID("a")
	assertEvents(t, rec, "CurrentChanged:a")

	reg.NextTint()
	assertEvents(t, rec, "CurrentChanged:b")
}

func TestRegistry_Events_Previous(t *testing.T) {
	t.Parallel()

	a, b := newTestTint("a"), newTestTint("b")
	reg := NewRegistry(nil, a, b)

	var got Event
	reg.Subscribe(func(ev Event) { got = ev })

	reg.SetTintID("b")
	if got.Kind != CurrentChanged || got.Tint != b || got.Previous != a || got.Source != "" {
		t.Fatalf("unexpected event: %+v", got)
	}
}

func TestRegistry_Unregister_CurrentFallsBack(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		current    string
		remove     []string
		wantID     string
		wantEvents []string
	}{
		{"next neighbour", "b", []string{"b"}, "c", []string{"TintRemoved:b", "CurrentChanged:c"}},
		{"last falls back to previous", "c", []string{"c"}, "b", []string{"TintRemoved:c", "CurrentChanged:b"}},
		{"first", "a", []string{"a"}, "b", []string{"TintRemoved:a", "CurrentChanged:b"}},
		{"several including current", "b", []string{"b", "c"}, "a", []string{"TintRemoved:b", "TintRemoved:c", "CurrentChanged:a"}},
		{"not current", "a", []string{"c"}, "a", []string{"TintRemoved:c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			reg, rec := newRecordedRegistry(t, newTestTint("a"), newTestTint("b"), newTestTint("c"))
			reg.SetTintID(tt.current)
			rec.drain()

			reg.UnregisterIDs(tt.remove...)

			if got := reg.Current(); got == nil || got.ID != tt.wantID {
				t.Fatalf("expected current %q, got %v", tt.wantID, got)
			}
			assertEvents(t, rec, tt.wantEvents...)
		})
	}
}

func TestRegistry_Unregister_DoesNotModifyCallerSlice(t *testing.T) {
	t.Parallel()

	reg := NewRegistry(nil, newTestTint("a"), newTestTint("b"), newTestTint("c"))
	ids := []string{"a", "z"}
	reg.UnregisterIDs(ids...)

	if !slices.Equal(ids, []string{"a", "z"}) {
		t.Fatalf("caller's slice was modified: %v", ids)
	}
}

func TestRegistry_UnregisterAll(t *testing.T) {
	t.Parallel()

	reg, rec := newRecordedRegistry(t, newTestTint("a"), newTestTint("b"))
	reg.UnregisterAll()

	got := rec.drain()
	want := []string{"TintRemoved:a", "TintRemoved:b", "CurrentChanged:<nil>"}
	slices.Sort(got[:2])
	if !slices.Equal(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected Current() to panic with an empty registry")
			}
		}()
		reg.Current()
	}()

	// Registering again recovers.
	reg.Register(newTestTint("z"))
	if reg.Current().ID != "z" {
		t.Fatalf("expected current to recover to z, got %v", reg.Current())
	}
}

func TestRegistry_Current_Recovers(t *testing.T) {
	t.Parallel()

	// A registry that has tints, but no current tint selected.
	reg := &Registry{tints: map[string]*Tint{"b": newTestTint("b"), "a": newTestTint("a")}}
	if got := reg.Current(); got == nil || got.ID != "a" {
		t.Fatalf("expected current to recover to the first tint, got %v", got)
	}
	if reg.Current().ID != "a" {
		t.Fatal("expected the recovered tint to be remembered")
	}
}

func TestRegistry_Subscribe(t *testing.T) {
	t.Parallel()

	reg := NewRegistry(nil, newTestTint("a"))

	var calls, calls2 int
	cancel := reg.Subscribe(func(Event) { calls++ })
	reg.Subscribe(func(Event) { calls2++ })

	reg.Register(newTestTint("b"))
	cancel()
	cancel() // Safe to call twice.
	reg.Register(newTestTint("c"))

	if calls != 1 || calls2 != 2 {
		t.Fatalf("expected 1 and 2 calls, got %d and %d", calls, calls2)
	}

	defer func() {
		if recover() == nil {
			t.Fatal("expected Subscribe(nil) to panic")
		}
	}()
	reg.Subscribe(nil)
}

func TestRegistry_Subscribe_Reentrant(t *testing.T) {
	t.Parallel()

	reg := NewRegistry(nil, newTestTint("a"))
	reg.Subscribe(func(ev Event) {
		if ev.Kind == TintAdded && ev.Tint.ID == "b" {
			reg.Register(newTestTint("derived")) // Would deadlock if emitted under the lock.
			_ = reg.Current()
		}
	})

	reg.Register(newTestTint("b"))
	if _, ok := reg.GetTint("derived"); !ok {
		t.Fatal("expected the subscriber's registration to have happened")
	}
}

func TestRegistry_Subscribe_Concurrent(t *testing.T) {
	t.Parallel()

	reg := NewRegistry(nil, newTestTint("a"))

	var mu sync.Mutex
	var count int
	reg.Subscribe(func(Event) {
		mu.Lock()
		count++
		mu.Unlock()
	})

	var wg sync.WaitGroup
	for i := range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t := newTestTint(string(rune('b' + i)))
			reg.Register(t)
			reg.SetTint(t)
		}()
	}
	wg.Wait()

	if count < 20 {
		t.Fatalf("expected at least 20 events, got %d", count)
	}
}
