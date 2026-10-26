// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// fakeSource is a [Source] whose contents are controlled by the test.
type fakeSource struct {
	name string

	mu    sync.Mutex
	tints []*Tint
	err   error
	loads int
}

func (s *fakeSource) Name() string { return s.name }

func (s *fakeSource) Load(context.Context) ([]*Tint, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.loads++
	return s.tints, s.err
}

func (s *fakeSource) set(err error, tints ...*Tint) {
	s.mu.Lock()
	s.tints, s.err = tints, err
	s.mu.Unlock()
}

// watchedSource is a fakeSource that is also a [Watcher].
type watchedSource struct {
	fakeSource
	changes chan struct{}
}

func newWatchedSource(name string, tints ...*Tint) *watchedSource {
	return &watchedSource{
		fakeSource: fakeSource{name: name, tints: tints},
		changes:    make(chan struct{}, 8),
	}
}

func (s *watchedSource) Watch(ctx context.Context) (<-chan struct{}, error) {
	out := make(chan struct{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.changes:
				select {
				case out <- struct{}{}:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out, nil
}

// notify signals the change, and waits until the attachment processed it.
func (s *watchedSource) notify(t *testing.T) {
	t.Helper()

	s.mu.Lock()
	before := s.loads
	s.mu.Unlock()

	s.changes <- struct{}{}

	waitFor(t, func() bool {
		s.mu.Lock()
		defer s.mu.Unlock()
		return s.loads > before
	})
	time.Sleep(10 * time.Millisecond) // Reconciliation happens right after Load.
}

func waitFor(t *testing.T, cond func() bool) {
	t.Helper()

	deadline := time.Now().Add(2 * time.Second)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatal("timed out waiting for condition")
		}
		time.Sleep(time.Millisecond)
	}
}

func TestRegistry_Attach(t *testing.T) {
	t.Parallel()

	reg, rec := newRecordedRegistry(t, newTestTint("user"))
	src := &fakeSource{name: "fake", tints: []*Tint{newTestTint("current"), newTestTint("other")}}

	detach, err := reg.Attach(testContext(t), src)
	if err != nil {
		t.Fatal(err)
	}

	assertEvents(t, rec,
		"TintAdded:fake:current", "TintAdded:fake:other", "CurrentChanged:fake:current")

	if got := reg.Current().ID; got != "fake:current" {
		t.Fatalf("expected to follow the source, current is %q", got)
	}
	if src.tints[0].ID != "current" {
		t.Fatal("source's tints must not be modified")
	}

	reg.mu.RLock()
	origin := reg.origins["fake:current"]
	reg.mu.RUnlock()
	if origin != "fake" {
		t.Fatalf("expected origin to be tracked, got %q", origin)
	}

	// Attaching the same source twice isn't allowed.
	if _, err = reg.Attach(testContext(t), src); !errors.Is(err, ErrAlreadyAttached) {
		t.Fatalf("expected ErrAlreadyAttached, got %v", err)
	}

	detach()
	detach() // Safe to call twice.

	assertEvents(t, rec, "TintRemoved:fake:current", "TintRemoved:fake:other", "CurrentChanged:user")
	if got := reg.Current().ID; got != "user" {
		t.Fatalf("expected to fall back to user's tint, current is %q", got)
	}

	// And it can be attached again after detaching.
	if _, err = reg.Attach(testContext(t), src); err != nil {
		t.Fatal(err)
	}
}

func TestRegistry_Attach_NoFollow(t *testing.T) {
	t.Parallel()

	reg, rec := newRecordedRegistry(t, newTestTint("user"))
	src := &fakeSource{name: "fake", tints: []*Tint{newTestTint("current")}}

	if _, err := reg.Attach(testContext(t), src, WithFollow(false)); err != nil {
		t.Fatal(err)
	}

	assertEvents(t, rec, "TintAdded:fake:current")
	if got := reg.Current().ID; got != "user" {
		t.Fatalf("expected current to be unchanged, got %q", got)
	}
}

func TestRegistry_Attach_Errors(t *testing.T) {
	t.Parallel()

	reg := NewRegistry(nil, newTestTint("user"))

	for _, name := range []string{"", "a:b"} {
		if _, err := reg.Attach(testContext(t), &fakeSource{name: name}); err == nil {
			t.Errorf("expected an error for source name %q", name)
		}
	}

	src := &fakeSource{name: "fake", err: fmt.Errorf("no config: %w", ErrNotFound)}
	if _, err := reg.Attach(testContext(t), src); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	// A failed attach must not reserve the name.
	src.set(nil, newTestTint("current"))
	if _, err := reg.Attach(testContext(t), src); err != nil {
		t.Fatalf("expected attach to succeed after a failed one, got %v", err)
	}
}

func TestRegistry_Attach_Watch(t *testing.T) {
	t.Parallel()

	reg, rec := newRecordedRegistry(t, newTestTint("user"))
	src := newWatchedSource("fake", newTestTint("current"), newTestTint("other"))

	var errs []error
	var errMu sync.Mutex
	_, err := reg.Attach(testContext(t), src, WithErrorHandler(func(err error) {
		errMu.Lock()
		errs = append(errs, err)
		errMu.Unlock()
	}))
	if err != nil {
		t.Fatal(err)
	}
	rec.drain()

	// Nothing changed: no events.
	src.notify(t)
	assertEvents(t, rec)

	// The active tint's colors changed: replaced, and the current tint is told.
	updated := newTestTint("current")
	updated.Bg = FromHex("#112233")
	src.set(nil, updated, newTestTint("other"))
	src.notify(t)
	assertEvents(t, rec, "TintUpdated:fake:current", "CurrentChanged:fake:current")
	if reg.Current().Bg.Hex() != "#112233" {
		t.Fatalf("expected current to reflect the update, got %s", reg.Current().Bg.Hex())
	}

	// The user picks something else; unchanged reloads don't steal it back.
	reg.SetTintID("user")
	rec.drain()
	src.notify(t)
	assertEvents(t, rec)

	// A different active tint, a new one, and a removed one.
	src.set(nil, newTestTint("third"), updated)
	src.notify(t)
	assertEvents(t, rec,
		"TintAdded:fake:third", "CurrentChanged:fake:third", "TintRemoved:fake:other")

	// A failed reload keeps what's there, and reports the error.
	src.set(fmt.Errorf("broken: %w", ErrNotFound))
	src.notify(t)
	assertEvents(t, rec)
	errMu.Lock()
	n := len(errs)
	errMu.Unlock()
	if n != 1 || !errors.Is(errs[0], ErrNotFound) {
		t.Fatalf("expected one ErrNotFound reported to the handler, got %v", errs)
	}
}

func TestRegistry_Attach_DetachFromSubscriber(t *testing.T) {
	t.Parallel()

	reg := NewRegistry(nil, newTestTint("user"))
	src := newWatchedSource("fake", newTestTint("current"))

	var detach func()
	var err error
	detach, err = reg.Attach(testContext(t), src)
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	var once sync.Once
	reg.Subscribe(func(ev Event) {
		if ev.Kind == TintUpdated {
			detach() // Runs on the watch goroutine; must not deadlock.
			once.Do(func() { close(done) })
		}
	})

	updated := newTestTint("current")
	updated.Bg = FromHex("#112233")
	src.set(nil, updated)
	src.changes <- struct{}{}

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("deadlocked detaching from a subscriber")
	}

	waitFor(t, func() bool { _, ok := reg.GetTint("fake:current"); return !ok })
}

func TestAttach_DefaultRegistry(t *testing.T) {
	old := DefaultRegistry
	t.Cleanup(func() { DefaultRegistry = old })
	DefaultRegistry = NewRegistry(nil, newTestTint("user"))

	var got []Event
	cancel := Subscribe(func(ev Event) { got = append(got, ev) })
	defer cancel()

	detach, err := Attach(testContext(t), &fakeSource{name: "fake", tints: []*Tint{newTestTint("current")}})
	if err != nil {
		t.Fatal(err)
	}
	defer detach()

	if len(got) != 2 || Current().ID != "fake:current" {
		t.Fatalf("unexpected state: events=%v current=%v", got, Current().ID)
	}
}

// testContext returns a context that's cancelled when the test finishes
// (testing.T.Context requires a newer Go version than this module supports).
func testContext(t *testing.T) context.Context {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}

func TestRegistry_Attach_UserTakesOverID(t *testing.T) {
	t.Parallel()

	reg := NewRegistry(nil, newTestTint("user"))
	detach, err := reg.Attach(testContext(t), &fakeSource{name: "fake", tints: []*Tint{newTestTint("current")}})
	if err != nil {
		t.Fatal(err)
	}

	// Registering over an attached tint makes it yours: it's no longer the
	// source's to remove.
	mine := newTestTint("fake:current")
	reg.Register(mine)
	detach()

	if got, ok := reg.GetTint("fake:current"); !ok || got != mine {
		t.Fatal("expected a tint the user registered to survive detaching")
	}
}
