// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package kitty is a [tint.Source] for the colors of the [kitty] terminal.
//
// It reads kitty.conf, including everything it includes (which is where
// "kitten themes" puts the chosen theme, in current-theme.conf), and provides a
// single tint, "current". Locations searched, first match wins:
//
//   - $KITTY_CONFIG_DIRECTORY
//   - $XDG_CONFIG_HOME/kitty, ~/.config/kitty
//   - ~/Library/Preferences/kitty (macOS)
//   - $XDG_CONFIG_DIRS/kitty (/etc/xdg/kitty)
//
// [kitty]: https://sw.kovidgoyal.net/kitty/
package kitty

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Name is the name of the source.
const Name = "kitty"

// New returns a source that reads kitty's config.
func New(opts ...sources.Option) *sources.FileSource {
	return sources.New(sources.Spec{Name: Name, Candidates: Candidates, Load: sources.Single(load)}, opts...)
}

// Candidates returns the paths kitty.conf is searched for, in order.
func Candidates() []string {
	var dirs []string
	if dir := os.Getenv("KITTY_CONFIG_DIRECTORY"); dir != "" {
		dirs = append(dirs, sources.ExpandPath(dir))
	}
	dirs = append(dirs, sources.JoinAll(sources.ConfigDirs(), "kitty")...)
	dirs = append(dirs, sources.JoinAll(sources.SystemConfigDirs(), "kitty")...)
	return sources.JoinAll(dirs, "kitty.conf")
}

// Load reads the kitty config at path, and returns its colors as a tint with the
// ID "current".
func Load(path string) (*tint.Tint, error) {
	t, _, err := load(path)
	return t, err
}

type parser struct {
	walker  sources.Walker
	dir     string // Includes are relative to the main config's directory.
	palette sources.Palette
	name    string
}

func load(path string) (*tint.Tint, []string, error) {
	p := &parser{dir: filepath.Dir(path)}
	if err := p.walker.Scan(path, true, p.parse); err != nil {
		return nil, p.walker.Files, err
	}
	if p.palette.Empty() {
		return nil, p.walker.Files, sources.ErrNoColors(Name, path)
	}

	name := p.name
	if name == "" {
		name = "Kitty"
	}
	return p.palette.Tint("current", name), p.walker.Files, nil
}

// parse handles a line of the config, or of anything it includes.
func (p *parser) parse(line string) error {
	if line[0] == '#' {
		// Themes carry their name in a "## name: Foo" comment.
		if rest, ok := strings.CutPrefix(line, "## name:"); ok {
			p.name = strings.TrimSpace(rest)
		}
		return nil
	}

	key, value := line, ""
	if i := strings.IndexAny(line, " \t"); i >= 0 {
		key, value = line[:i], strings.TrimSpace(line[i+1:])
	}

	switch key {
	case "include":
		if value != "" {
			return p.walker.Scan(sources.ResolvePath(p.dir, value), false, p.parse)
		}
	case "globinclude":
		matches, _ := filepath.Glob(sources.ResolvePath(p.dir, value))
		for _, m := range matches {
			if err := p.walker.Scan(m, false, p.parse); err != nil {
				return err
			}
		}
	default:
		p.set(key, value)
	}
	return nil
}

func (p *parser) set(key, value string) {
	c, ok := sources.ParseColor(value)
	if !ok {
		return
	}

	switch key {
	case "foreground":
		p.palette.Foreground = c
	case "background":
		p.palette.Background = c
	case "selection_foreground":
		p.palette.SelectionForeground = c
	case "selection_background":
		p.palette.SelectionBackground = c
	case "cursor":
		p.palette.Cursor = c
	case "cursor_text_color":
		p.palette.CursorText = c
	default:
		if n, found := strings.CutPrefix(key, "color"); found {
			if i, err := strconv.Atoi(n); err == nil {
				p.palette.SetANSI(i, c)
			}
		}
	}
}
