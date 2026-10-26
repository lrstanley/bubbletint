// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package xresources is a [tint.Source] for the colors of terminals configured
// through X resources: xterm, urxvt, st (patched), and most other terminals for
// the X Window System.
//
// It reads the user's Xresources file (including #include'd files and simple
// #define'd values) and provides a single tint, "current", made of the
// "foreground", "background", "cursorColor", "colorBD" and "color0" through
// "color15" resources. As the resources of different applications can't be told
// apart without knowing which one is in use, the last value set for each wins,
// no matter the application it's for. Locations searched, first match wins:
//
//   - ~/.Xresources, ~/.Xdefaults
//   - $XDG_CONFIG_HOME/X11/Xresources
package xresources

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Name is the name of the source.
const Name = "xresources"

// New returns a source that reads the user's X resources.
func New(opts ...sources.Option) *sources.FileSource {
	return sources.New(sources.Spec{Name: Name, Candidates: Candidates, Load: sources.Single(load)}, opts...)
}

// Candidates returns the paths X resources are searched for, in order.
func Candidates() []string {
	var paths []string
	if home := sources.HomeDir(); home != "" {
		paths = append(paths, filepath.Join(home, ".Xresources"), filepath.Join(home, ".Xdefaults"))
	}
	return append(paths, sources.JoinAll(sources.ConfigDirs(), "X11", "Xresources")...)
}

// Load reads the X resources at path, and returns their colors as a tint with the
// ID "current".
func Load(path string) (*tint.Tint, error) {
	t, _, err := load(path)
	return t, err
}

type parser struct {
	walker  sources.Walker
	palette sources.Palette
	defines map[string]string
}

func load(path string) (*tint.Tint, []string, error) {
	p := &parser{defines: make(map[string]string)}
	if err := p.scan(path, true); err != nil {
		return nil, p.walker.Files, err
	}
	if p.palette.Empty() {
		return nil, p.walker.Files, sources.ErrNoColors(Name, path)
	}
	return p.palette.Tint("current", "X resources"), p.walker.Files, nil
}

func (p *parser) scan(path string, required bool) error {
	return p.walker.Scan(path, required, func(line string) error {
		switch line[0] {
		case '!':
			return nil
		case '#':
			directive, rest, _ := strings.Cut(strings.TrimSpace(line[1:]), " ")
			rest = strings.TrimSpace(rest)

			switch directive {
			case "define":
				if name, value, ok := strings.Cut(rest, " "); ok {
					p.defines[name] = strings.TrimSpace(value)
				}
			case "include":
				return p.scan(sources.ResolvePath(filepath.Dir(path), strings.Trim(rest, `"<>`)), false)
			}
			return nil
		}

		if name, value, ok := strings.Cut(line, ":"); ok {
			p.set(name, strings.TrimSpace(value))
		}
		return nil
	})
}

// set applies a resource. name is the full resource specifier, such as
// "URxvt*color1", of which only the last component matters.
func (p *parser) set(name, value string) {
	if v, ok := p.defines[value]; ok {
		value = v
	}

	if i := strings.LastIndexAny(name, ".*"); i >= 0 {
		name = name[i+1:]
	}

	c, ok := sources.ParseColor(value)
	if !ok {
		return
	}

	switch strings.ToLower(name) {
	case "foreground":
		p.palette.Foreground = c
	case "background":
		p.palette.Background = c
	case "cursorcolor":
		p.palette.Cursor = c
	case "colorbd":
		p.palette.Bold = c
	case "highlightcolor":
		p.palette.SelectionBackground = c
	case "highlighttextcolor":
		p.palette.SelectionForeground = c
	default:
		if n, found := strings.CutPrefix(strings.ToLower(name), "color"); found {
			if i, err := strconv.Atoi(n); err == nil {
				p.palette.SetANSI(i, c)
			}
		}
	}
}
