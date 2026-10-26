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
	Load func(path string) (tints []*tint.Tint, files []string, err error)
}

// FileSource is a [tint.Source] and [tint.Watcher] that reads tints from a config
// file on disk, and watches it, and anything it includes, for changes.
type FileSource struct {
	spec Spec
	cfg  Config

	mu    sync.Mutex
	files []string // Files read by the last load.
}

// Single adapts a function that reads one tint, as most terminals' configs hold
// a single theme, to a [Spec.Load].
func Single(load func(path string) (*tint.Tint, []string, error)) func(path string) ([]*tint.Tint, []string, error) {
	return func(path string) ([]*tint.Tint, []string, error) {
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
	return s
}

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

	tints, files, err := s.spec.Load(path)

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
	return poll(ctx, s.cfg.PollInterval, s.watched), nil
}
