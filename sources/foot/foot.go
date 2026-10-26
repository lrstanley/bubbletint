// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package foot is a [tint.Source] for the colors of the [foot] terminal.
//
// It reads foot.ini (including files named by "include", which is where foot
// themes normally go) and provides a single tint, "current". The [colors-dark]
// section is used when present (foot's default), and otherwise [colors], then
// [colors-light]. Locations searched, first match wins:
//
//   - $XDG_CONFIG_HOME/foot/foot.ini, ~/.config/foot/foot.ini
//   - $XDG_CONFIG_DIRS/foot/foot.ini (/etc/xdg/foot/foot.ini)
//
// [foot]: https://codeberg.org/dnkl/foot
package foot

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Name is the name of the source.
const Name = "foot"

// New returns a source that reads foot's config.
func New(opts ...sources.Option) *sources.FileSource {
	return sources.New(sources.Spec{Name: Name, Candidates: Candidates, Load: sources.Single(load)}, opts...)
}

// Candidates returns the paths foot's config is searched for, in order.
func Candidates() []string {
	dirs := append(sources.ConfigDirs(), sources.SystemConfigDirs()...)
	return sources.JoinAll(dirs, "foot", "foot.ini")
}

// Load reads the foot config at path, and returns its colors as a tint with the
// ID "current".
func Load(path string) (*tint.Tint, error) {
	t, _, err := load(path)
	return t, err
}

// sections maps section names to their keys, with later values replacing
// earlier ones.
type sections map[string]map[string]string

type loader struct {
	walker   sources.Walker
	sections sections
}

func load(path string) (*tint.Tint, []string, error) {
	l := &loader{sections: make(sections)}
	if err := l.scan(path, true); err != nil {
		return nil, l.walker.Files, err
	}

	colors := l.sections["colors-dark"]
	if colors == nil {
		colors = l.sections["colors"]
	}
	if colors == nil {
		colors = l.sections["colors-light"]
	}

	var p sources.Palette
	get := func(key string) *tint.Color {
		c, _ := sources.ParseColor(colors[key])
		return c
	}

	p.Background, p.Foreground = get("background"), get("foreground")
	p.SelectionBackground, p.SelectionForeground = get("selection-background"), get("selection-foreground")
	for i := range 8 {
		n := strconv.Itoa(i)
		p.SetANSI(i, get("regular"+n))
		p.SetANSI(i+8, get("bright"+n))
	}

	// "color" in [cursor] is "<text color> <cursor color>".
	if fields := strings.Fields(l.sections["cursor"]["color"]); len(fields) == 2 {
		p.CursorText, _ = sources.ParseColor(fields[0])
		p.Cursor, _ = sources.ParseColor(fields[1])
	}

	if p.Empty() {
		return nil, l.walker.Files, sources.ErrNoColors(Name, path)
	}
	return p.Tint("current", "Foot"), l.walker.Files, nil
}

// scan reads the file into l.sections, following includes where they appear.
func (l *loader) scan(path string, required bool) error {
	section := "main"

	return l.walker.Scan(path, required, func(line string) error {
		switch line[0] {
		case '#', ';':
			return nil
		case '[':
			if end := strings.IndexByte(line, ']'); end > 0 {
				section = line[1:end]
			}
			return nil
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			return nil
		}
		key, value = strings.TrimSpace(key), strings.TrimSpace(value)

		if section == "main" && key == "include" {
			return l.scan(sources.ResolvePath(filepath.Dir(path), value), false)
		}

		if l.sections[section] == nil {
			l.sections[section] = make(map[string]string)
		}
		l.sections[section][key] = value
		return nil
	})
}
