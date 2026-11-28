// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package sources provides [tint.Source] implementations that read the color
// scheme of terminals from their config files, so that an application can match
// the terminal it's running in. They are opt-in: nothing here is used unless
// you import it and attach it to a registry.
//
// Each supported terminal has its own package (kitty, ghostty, alacritty,
// foot, windowsterminal, xresources and itermcolors), and the terminal
// package picks the right one for the terminal the process is running in:
//
//	src, err := terminal.Detect()
//	if err == nil {
//		detach, err := tint.Attach(ctx, src) // Follows the terminal's theme.
//	}
//
// This package contains the building blocks shared between them, which can also
// be used to write sources for other config formats.
package sources

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	tint "github.com/lrstanley/bubbletint/v2"
)

// DefaultPollInterval is how often a [FileSource] checks its files for changes.
const DefaultPollInterval = time.Second

// Config is the configuration of a [FileSource].
type Config struct {
	// Path overrides where the config is read from. When empty, the source's
	// default locations are searched.
	Path string

	// PollInterval is how often files are checked for changes while watching.
	PollInterval time.Duration

	// Light is whether the terminal is in light mode, for a config that names
	// a theme for each of light and dark mode. The zero value is dark.
	Light bool
}

// Option configures a source.
type Option func(*Config)

// WithPath reads the config from path, instead of searching the terminal's
// default locations.
func WithPath(path string) Option {
	return func(c *Config) { c.Path = path }
}

// WithPollInterval sets how often the config files are checked for changes while
// the source is being watched (see [tint.Registry.Attach]). The default is
// [DefaultPollInterval]. Files are polled, rather than using filesystem
// notifications, so that watching works the same on every platform and keeps
// working with editors that replace files, and with symlinked dotfiles.
func WithPollInterval(d time.Duration) Option {
	return func(c *Config) { c.PollInterval = d }
}

// WithDark sets whether the terminal is in dark mode, the default, or light
// mode, for a config that names a theme for each, such as Ghostty's
// "theme = light:X,dark:Y". See [FileSource.SetDark] to change it later.
func WithDark(dark bool) Option {
	return func(c *Config) { c.Light = !dark }
}

// Spec describes how to read a config for a [FileSource].
type Spec struct {
	// Name is the [tint.Source] name, e.g. "kitty".
	Name string

	// Candidates returns the config paths to look for, in order of preference.
	// The first one that exists is used. Ignored if [WithPath] is used.
	Candidates func() []string

	// Load parses the config at path. It returns the tints it found (the first
	// being the active one), and every file it read, so that all of them can be
	// watched. Load should return the files it read even if it fails.
	//
	// dark is whether the terminal is in dark mode. Where the config names a
	// theme for each of light and dark mode, the active tint is the one for
	// dark's mode, and both are returned too, as [DarkID] and [LightID].
	Load func(path string, dark bool) (tints []*tint.Tint, files []string, err error)
}

// Variants returns the tints for a config that names a theme for each of
// light and dark mode: current, the one for the mode the terminal is in, with
// the ID "current", then dark and light, as [DarkID] and [LightID]. A nil
// dark or light, one that can't be found, is left out. Each is a copy.
func Variants(current, dark, light *tint.Tint) []*tint.Tint {
	out := []*tint.Tint{withID(current, "current")}
	if dark != nil {
		out = append(out, withID(dark, DarkID))
	}
	if light != nil {
		out = append(out, withID(light, LightID))
	}
	return out
}

func withID(t *tint.Tint, id string) *tint.Tint {
	cp := *t
	cp.ID = id
	return &cp
}

// The IDs of the tints a source returns, besides its active one, where its
// config names a theme for each of light and dark mode, before the source's
// name is added. An app that picks between them itself, rather than with
// [FileSource.SetDark], attaches with [tint.WithFollow] off.
const (
	// DarkID is the tint for dark mode.
	DarkID = "current_dark"
	// LightID is the tint for light mode.
	LightID = "current_light"
)

// FileSource is a [tint.Source] and [tint.Watcher] that reads tints from a config
// file on disk, and watches it, and anything it includes, for changes.
type FileSource struct {
	spec Spec
	cfg  Config

	mu    sync.Mutex
	files []string // Files read by the last load.

	light    atomic.Bool
	watchers map[chan struct{}]struct{} // Each tells a watcher to reload now.
}

// Single adapts a function that reads one tint, as most terminals' configs hold
// a single theme, to a [Spec.Load].
func Single(load func(path string) (*tint.Tint, []string, error)) func(path string, dark bool) ([]*tint.Tint, []string, error) {
	return func(path string, _ bool) ([]*tint.Tint, []string, error) {
		t, files, err := load(path)
		if err != nil {
			return nil, files, err
		}
		return []*tint.Tint{t}, files, nil
	}
}

// New returns a source for the provided spec.
func New(spec Spec, opts ...Option) *FileSource {
	s := &FileSource{spec: spec, cfg: Config{PollInterval: DefaultPollInterval}}
	for _, opt := range opts {
		opt(&s.cfg)
	}
	if s.cfg.PollInterval <= 0 {
		s.cfg.PollInterval = DefaultPollInterval
	}
	s.light.Store(s.cfg.Light)
	return s
}

// SetDark sets whether the terminal is in dark mode or light mode, for a
// config that names a theme for each. Terminals that do, such as Ghostty and
// Windows Terminal, follow the system's appearance, which the source can't
// see; an app can, from the terminal's background (Bubble Tea delivers it as
// a tea.BackgroundColorMsg, with IsDark). If the source is being watched, a
// change reloads it, so every registry following it follows the other theme.
func (s *FileSource) SetDark(dark bool) {
	if s.light.Swap(!dark) == !dark {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for nudge := range s.watchers {
		select {
		case nudge <- struct{}{}:
		default: // A reload is already pending.
		}
	}
}

// Dark reports whether the source loads the theme for dark mode, where its
// config names one for each. See [FileSource.SetDark].
func (s *FileSource) Dark() bool { return !s.light.Load() }

// Name implements [tint.Source].
func (s *FileSource) Name() string { return s.spec.Name }

// candidates returns the paths that could provide the config.
func (s *FileSource) candidates() []string {
	if s.cfg.Path != "" {
		return []string{ExpandPath(s.cfg.Path)}
	}
	if s.spec.Candidates == nil {
		return nil
	}
	return s.spec.Candidates()
}

// Find returns the path of the config that would be loaded, or an error wrapping
// [tint.ErrNotFound] if there isn't one.
func (s *FileSource) Find() (string, error) {
	candidates := s.candidates()
	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	if len(candidates) == 0 {
		return "", fmt.Errorf("%w: %s has no config locations on this platform", tint.ErrNotFound, s.spec.Name)
	}
	return "", fmt.Errorf("%w: no %s config in any of: %s", tint.ErrNotFound, s.spec.Name, strings.Join(candidates, ", "))
}

// Load implements [tint.Source].
func (s *FileSource) Load(ctx context.Context) ([]*tint.Tint, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	path, err := s.Find()
	if err != nil {
		return nil, err
	}

	tints, files, err := s.spec.Load(path, s.Dark())

	s.mu.Lock()
	s.files = append([]string{path}, files...)
	s.mu.Unlock()

	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	return tints, nil
}

// watched returns every file that should be checked for changes: the files that
// were read, and everywhere the config might appear.
func (s *FileSource) watched() []string {
	s.mu.Lock()
	files := append([]string(nil), s.files...)
	s.mu.Unlock()
	return append(files, s.candidates()...)
}

// Watch implements [tint.Watcher].
func (s *FileSource) Watch(ctx context.Context) (<-chan struct{}, error) {
	// A watcher starts with a reload queued: the mode may have changed since
	// the load it follows, before anything was watching to be told, as from
	// a subscriber to that load's events. A registry ignores a reload that
	// changes nothing.
	nudge := make(chan struct{}, 1)
	nudge <- struct{}{}
	s.mu.Lock()
	if s.watchers == nil {
		s.watchers = make(map[chan struct{}]struct{})
	}
	s.watchers[nudge] = struct{}{}
	s.mu.Unlock()
	context.AfterFunc(ctx, func() {
		s.mu.Lock()
		delete(s.watchers, nudge)
		s.mu.Unlock()
	})
	return poll(ctx, s.cfg.PollInterval, s.watched, nudge), nil
}
