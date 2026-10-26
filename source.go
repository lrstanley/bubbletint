// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var (
	// ErrNotFound is returned (wrapped) by a [Source] when whatever it reads
	// from isn't present, e.g. the terminal it supports isn't installed, or has
	// no config file. Check for it with [errors.Is].
	ErrNotFound = errors.New("tint: source not found")

	// ErrAlreadyAttached is returned by [Registry.Attach] when a source with the
	// same name is already attached to the registry.
	ErrAlreadyAttached = errors.New("tint: source already attached")
)

// Source provides tints from somewhere outside of this package, such as a
// terminal's config file. Sources are opt-in: they live in their own packages
// and are added to a registry with [Registry.Attach].
type Source interface {
	// Name is a short, stable, lowercase identifier for the source (e.g.
	// "kitty"). It must be non-empty, and must not contain a colon. The ID of
	// every tint the source provides is prefixed with it ("kitty:current"),
	// so tints from a source can never overlap with your own, or with another
	// source's.
	Name() string

	// Load returns the tints that are currently available from the source. The
	// first tint returned is the source's active tint (e.g. the theme the
	// terminal is currently configured to use), which is what the registry
	// follows. The IDs of the returned tints should not include the source's
	// name, as [Registry.Attach] adds it.
	//
	// Load should return an error wrapping [ErrNotFound] when the source isn't
	// available.
	Load(ctx context.Context) ([]*Tint, error)
}

// Watcher can optionally be implemented by a [Source] that is able to tell when
// what it provides has changed. If it is, [Registry.Attach] keeps the registry
// up to date.
type Watcher interface {
	// Watch returns a channel that receives a value each time the source's
	// contents may have changed. The channel must be closed when ctx is done.
	// Sends should not block the source: coalescing is fine, as every received
	// value results in a full [Source.Load].
	Watch(ctx context.Context) (<-chan struct{}, error)
}

type attachOptions struct {
	follow  bool
	onError func(error)
}

// AttachOption configures [Registry.Attach].
type AttachOption func(*attachOptions)

// WithFollow controls whether the registry follows the source, i.e. makes the
// source's active tint the current tint when the source is attached, and again
// whenever the source's active tint changes. The default is true. When false,
// the source's tints are only registered, and it's up to you to select one.
func WithFollow(follow bool) AttachOption {
	return func(o *attachOptions) { o.follow = follow }
}

// WithErrorHandler sets a function that's called with errors that occur while
// reloading a source in the background (the initial load's error is returned
// by [Registry.Attach] instead). When a reload fails, the previously loaded
// tints are kept.
func WithErrorHandler(fn func(error)) AttachOption {
	return func(o *attachOptions) { o.onError = fn }
}

type attachment struct {
	reg    *Registry
	source Source
	prefix string
	opts   attachOptions

	cancel context.CancelFunc

	mu     sync.Mutex // Serializes reloads and detaching.
	closed bool
	active string // ID of the active tint, after prefixing.
}

// Attach loads the provided source and registers its tints in the registry, with
// the ID of each tint prefixed by the source's name and a colon (see
// [Source.Name]). Unless disabled with [WithFollow], the source's active tint
// also becomes the current tint.
//
// If the source implements [Watcher], the registry is kept in sync with it in
// the background until ctx is done, or the returned detach function is called.
// Tints that disappear from the source are unregistered, tints that changed are
// replaced (emitting [TintUpdated]), and, when following, the current tint
// follows the source's active tint.
//
// The initial load happens before Attach returns. If it fails, nothing is
// attached, and the error is returned, wrapped, so it can still be checked for
// [ErrNotFound] with [errors.Is].
//
// detach stops watching and unregisters everything the source provided. It is
// safe to call more than once. Cancelling ctx only stops watching.
func (r *Registry) Attach(ctx context.Context, source Source, opts ...AttachOption) (detach func(), err error) {
	name := source.Name()
	if name == "" || strings.Contains(name, ":") {
		return nil, fmt.Errorf("tint: invalid source name %q: must be non-empty and not contain a colon", name)
	}

	a := &attachment{
		reg:    r,
		source: source,
		prefix: name + ":",
		opts:   attachOptions{follow: true},
	}
	for _, opt := range opts {
		opt(&a.opts)
	}

	r.mu.Lock()
	if _, ok := r.attached[name]; ok {
		r.mu.Unlock()
		return nil, fmt.Errorf("%w: %s", ErrAlreadyAttached, name)
	}
	if r.attached == nil {
		r.attached = make(map[string]*attachment)
	}
	r.attached[name] = a // Reserve the name while loading.
	r.mu.Unlock()

	if err = a.reload(ctx, true); err != nil {
		r.mu.Lock()
		delete(r.attached, name)
		r.mu.Unlock()
		return nil, err
	}

	var changes <-chan struct{}
	if w, ok := source.(Watcher); ok {
		var watchCtx context.Context
		watchCtx, a.cancel = context.WithCancel(ctx)
		changes, err = w.Watch(watchCtx)
		if err != nil {
			a.cancel()
			a.unregister()
			return nil, fmt.Errorf("tint: watching source %q: %w", name, err)
		}
		go a.watch(watchCtx, changes)
	} else {
		a.cancel = func() {}
	}

	return sync.OnceFunc(func() {
		a.cancel()
		a.unregister()
	}), nil
}

func (a *attachment) watch(ctx context.Context, changes <-chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			return
		case _, ok := <-changes:
			if !ok {
				return
			}
			if err := a.reload(ctx, false); err != nil && a.opts.onError != nil && ctx.Err() == nil {
				a.opts.onError(err)
			}
		}
	}
}

// unregister removes everything the source registered, and stops it from being
// reloaded.
func (a *attachment) unregister() {
	a.mu.Lock()
	a.closed = true
	a.mu.Unlock()

	name := a.source.Name()

	a.reg.mu.Lock()
	delete(a.reg.attached, name)
	events := a.reg.unregisterLocked(a.reg.idsFromLocked(name)...)
	a.reg.mu.Unlock()

	a.reg.subs.emit(events)
}

// reload loads the source and reconciles the registry with the result.
func (a *attachment) reload(ctx context.Context, initial bool) error {
	loaded, err := a.source.Load(ctx)
	if err != nil {
		return fmt.Errorf("tint: loading source %q: %w", a.source.Name(), err)
	}

	a.mu.Lock()
	if a.closed {
		a.mu.Unlock()
		return nil
	}

	name := a.source.Name()
	next := make(map[string]*Tint, len(loaded))
	order := make([]string, 0, len(loaded))

	for _, t := range loaded {
		if t == nil || t.ID == "" {
			continue
		}

		// Copy, so the source's tint is never modified.
		prefixed := *t
		prefixed.ID = a.prefix + t.ID

		if _, dup := next[prefixed.ID]; dup {
			continue
		}
		next[prefixed.ID] = &prefixed
		order = append(order, prefixed.ID)
	}

	activeID := ""
	if len(order) > 0 {
		activeID = order[0]
	}

	a.reg.mu.Lock()

	var events []Event
	activeChanged := initial || activeID != a.active
	for _, id := range order {
		previous, exists := a.reg.tints[id]
		if exists && reflect.DeepEqual(previous, next[id]) {
			continue
		}
		if id == activeID && exists {
			activeChanged = true
		}
		events = append(events, a.reg.registerLocked(name, next[id])...)
	}

	// Follow before removing stale tints, so that if the current tint is one of
	// them, the registry doesn't need to fall back to a tint we're about to
	// replace anyway.
	if a.opts.follow && activeChanged && activeID != "" && a.reg.currentTint != activeID {
		events = append(events, Event{
			Kind:     CurrentChanged,
			Tint:     a.reg.tints[activeID],
			Previous: a.reg.tints[a.reg.currentTint],
			Source:   name,
		})
		a.reg.currentTint = activeID
	}

	// Whatever the source registered before, and no longer provides, is stale.
	var stale []string
	for _, id := range a.reg.idsFromLocked(name) {
		if _, ok := next[id]; !ok {
			stale = append(stale, id)
		}
	}
	events = append(events, a.reg.unregisterLocked(stale...)...)
	a.reg.mu.Unlock()

	a.active = activeID
	a.mu.Unlock()

	// Emit without holding any locks, so subscribers can call back into the
	// registry, including detaching this source.
	a.reg.subs.emit(events)
	return nil
}
