// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package alacritty is a [tint.Source] for the colors of the [Alacritty]
// terminal.
//
// It reads alacritty.toml, merging in everything listed in "general.import" (or
// the older top-level "import", which is where theme files normally go), and
// provides a single tint, "current". The tint's display name is that of the last
// imported file that sets the primary colors, which is the theme if you use one. The legacy YAML config
// is not supported. Locations searched, first match wins:
//
//   - $XDG_CONFIG_HOME/alacritty/alacritty.toml, $XDG_CONFIG_HOME/alacritty.toml
//   - ~/.config/alacritty/alacritty.toml, ~/.alacritty.toml
//   - %APPDATA%\alacritty\alacritty.toml (Windows)
//   - $XDG_CONFIG_DIRS/alacritty/alacritty.toml (/etc/alacritty/alacritty.toml)
//
// [Alacritty]: https://alacritty.org
package alacritty

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Name is the name of the source.
const Name = "alacritty"

// maxImportDepth is how deeply imports are followed, matching Alacritty.
const maxImportDepth = 5

// New returns a source that reads Alacritty's config.
func New(opts ...sources.Option) *sources.FileSource {
	return sources.New(sources.Spec{Name: Name, Candidates: Candidates, Load: sources.Single(load)}, opts...)
}

// Candidates returns the paths Alacritty's config is searched for, in order.
func Candidates() []string {
	var paths []string
	for _, dir := range sources.ConfigDirs() {
		paths = append(paths,
			filepath.Join(dir, "alacritty", "alacritty.toml"),
			filepath.Join(dir, "alacritty.toml"),
		)
	}
	if home := sources.HomeDir(); home != "" {
		paths = append(paths, filepath.Join(home, ".alacritty.toml"))
	}
	for _, dir := range sources.SystemConfigDirs() {
		paths = append(paths, filepath.Join(dir, "alacritty", "alacritty.toml"))
	}
	return append(paths, "/etc/alacritty/alacritty.toml")
}

// Load reads the Alacritty config at path, and returns its colors as a tint with
// the ID "current".
func Load(path string) (*tint.Tint, error) {
	t, _, err := load(path)
	return t, err
}

type loader struct {
	files []string
	name  string // Name of the last import that set the primary colors.
}

func load(path string) (*tint.Tint, []string, error) {
	l := &loader{}

	cfg, err := l.read(path, 0, make(map[string]bool))
	if err != nil {
		return nil, l.files, err
	}

	colors, _ := cfg["colors"].(map[string]any)

	var p sources.Palette
	section := func(name string) map[string]any {
		m, _ := colors[name].(map[string]any)
		return m
	}
	color := func(m map[string]any, key string) *tint.Color {
		s, _ := m[key].(string)
		c, _ := sources.ParseColor(s)
		return c
	}

	primary := section("primary")
	p.Background, p.Foreground = color(primary, "background"), color(primary, "foreground")
	p.Bold = color(primary, "bright_foreground")

	cursor := section("cursor")
	p.Cursor, p.CursorText = color(cursor, "cursor"), color(cursor, "text")

	selection := section("selection")
	p.SelectionBackground, p.SelectionForeground = color(selection, "background"), color(selection, "text")

	// Alacritty calls magenta what bubbletint calls purple; sources.Palette uses
	// the standard ANSI order.
	names := [...]string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	normal, bright := section("normal"), section("bright")
	for i, n := range names {
		p.SetANSI(i, color(normal, n))
		p.SetANSI(i+8, color(bright, n))
	}

	if p.Empty() {
		return nil, l.files, sources.ErrNoColors(Name, path)
	}

	name := "Alacritty"
	if l.name != "" {
		name = l.name
	}
	return p.Tint("current", name), l.files, nil
}

// read reads the config at path, with its imports merged in (the file's own
// settings taking precedence over its imports, and later imports over earlier
// ones).
func (l *loader) read(path string, depth int, seen map[string]bool) (map[string]any, error) {
	l.files = append(l.files, path)

	if abs, err := filepath.Abs(path); err == nil {
		path = abs
	}
	if seen[path] {
		return map[string]any{}, nil
	}
	seen[path] = true
	defer delete(seen, path) // Only guard against cycles, not diamond imports.

	var cfg map[string]any
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}
	if cfg == nil {
		cfg = map[string]any{}
	}

	merged := map[string]any{}
	if depth < maxImportDepth {
		for _, imp := range imports(cfg) {
			impPath := sources.ResolvePath(filepath.Dir(path), imp)

			sub, err := l.read(impPath, depth+1, seen)
			if err != nil {
				// Alacritty only warns about imports it can't read.
				var perr toml.ParseError
				if errors.Is(err, fs.ErrNotExist) || errors.As(err, &perr) {
					continue
				}
				return nil, fmt.Errorf("import %s: %w", impPath, err)
			}
			if colors, _ := sub["colors"].(map[string]any); colors["primary"] != nil {
				l.name = strings.TrimSuffix(filepath.Base(impPath), filepath.Ext(impPath))
			}
			deepMerge(merged, sub)
		}
	}

	deepMerge(merged, cfg)
	return merged, nil
}

// imports returns the paths listed in the config's imports, both the current
// ("general.import") and deprecated (top-level "import") location.
func imports(cfg map[string]any) []string {
	var out []string
	collect := func(v any) {
		list, _ := v.([]any)
		for _, item := range list {
			if s, ok := item.(string); ok && s != "" {
				out = append(out, s)
			}
		}
	}

	collect(cfg["import"])
	if general, ok := cfg["general"].(map[string]any); ok {
		collect(general["import"])
	}
	return out
}

// deepMerge merges src into dst, recursing into tables. Values in src win.
func deepMerge(dst, src map[string]any) {
	for k, v := range src {
		if srcMap, ok := v.(map[string]any); ok {
			if dstMap, isMap := dst[k].(map[string]any); isMap {
				deepMerge(dstMap, srcMap)
				continue
			}
			cp := map[string]any{}
			deepMerge(cp, srcMap)
			dst[k] = cp
			continue
		}
		dst[k] = v
	}
}
