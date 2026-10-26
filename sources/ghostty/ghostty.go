// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package ghostty is a [tint.Source] for the colors of the [Ghostty] terminal.
//
// It reads Ghostty's config (including "config-file" includes), resolves the
// "theme" it names, and provides a single tint, "current". Colors set directly
// in the config take precedence over those from the theme, as they do in
// Ghostty. Themes are looked up in Ghostty's themes directories (the user's
// config, then the install's resources), and fall back to bubbletint's built-in
// tints by name, as Ghostty's themes come from the same collection. When a
// theme is given for both light and dark mode, the dark one is used.
//
// Locations searched, first match wins:
//
//   - $XDG_CONFIG_HOME/ghostty, ~/.config/ghostty
//   - ~/Library/Application Support/com.mitchellh.ghostty (macOS)
//   - %APPDATA%\ghostty (Windows)
//
// Both "config.ghostty" and the older "config" file names are supported.
//
// [Ghostty]: https://ghostty.org
package ghostty

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Name is the name of the source.
const Name = "ghostty"

// New returns a source that reads Ghostty's config.
func New(opts ...sources.Option) *sources.FileSource {
	return sources.New(sources.Spec{Name: Name, Candidates: Candidates, Load: sources.Single(load)}, opts...)
}

func configDirs() []string {
	var dirs []string
	for _, dir := range sources.ConfigDirs() {
		dirs = append(dirs, filepath.Join(dir, "ghostty"))
	}
	if home := sources.HomeDir(); home != "" && runtime.GOOS == "darwin" {
		dirs = append(dirs, filepath.Join(home, "Library", "Application Support", "com.mitchellh.ghostty"))
	}
	return dirs
}

// Candidates returns the paths Ghostty's config is searched for, in order.
func Candidates() []string {
	var paths []string
	for _, dir := range configDirs() {
		paths = append(paths, filepath.Join(dir, "config.ghostty"), filepath.Join(dir, "config"))
	}
	return paths
}

// themeDirs returns the directories themes are searched for in.
func themeDirs() []string {
	dirs := sources.JoinAll(configDirs(), "themes")
	if dir := os.Getenv("GHOSTTY_RESOURCES_DIR"); dir != "" {
		dirs = append(dirs, filepath.Join(sources.ExpandPath(dir), "themes"))
	}

	switch runtime.GOOS {
	case "darwin":
		dirs = append(dirs, "/Applications/Ghostty.app/Contents/Resources/ghostty/themes")
	case "windows":
	default:
		dirs = append(dirs, "/usr/share/ghostty/themes", "/usr/local/share/ghostty/themes")
		if home := sources.HomeDir(); home != "" {
			dirs = append(dirs, filepath.Join(home, ".local", "share", "ghostty", "themes"))
		}
	}
	return dirs
}

// Load reads the Ghostty config at path, and returns its colors as a tint with
// the ID "current".
func Load(path string) (*tint.Tint, error) {
	t, _, err := load(path)
	return t, err
}

type entry struct{ key, value string }

type loader struct{ walker sources.Walker }

func load(path string) (*tint.Tint, []string, error) {
	l := &loader{}
	entries, err := l.read(path, true)
	if err != nil {
		return nil, l.walker.Files, err
	}

	var palette sources.Palette
	name := "Ghostty"

	// The theme is the base, and anything the config sets itself replaces it.
	if theme := pickTheme(lastValue(entries, "theme")); theme != "" {
		name = theme

		var themeFile string
		palette, themeFile = loadTheme(theme)
		if themeFile != "" {
			l.walker.Files = append(l.walker.Files, themeFile)
		}
	}
	apply(&palette, entries)

	if palette.Empty() {
		return nil, l.walker.Files, sources.ErrNoColors(Name, path)
	}
	return palette.Tint("current", name), l.walker.Files, nil
}

// read returns the settings in the file, followed by those of its "config-file"
// includes, as Ghostty loads them after the file itself. Later settings take
// precedence.
func (l *loader) read(path string, required bool) ([]entry, error) {
	var entries []entry
	var includes []string

	err := l.walker.Scan(path, required, func(line string) error {
		if line[0] == '#' {
			return nil
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			return nil
		}
		e := entry{strings.TrimSpace(key), unquote(strings.TrimSpace(value))}

		if e.key == "config-file" {
			includes = append(includes, e.value)
		} else {
			entries = append(entries, e)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	for _, include := range includes {
		optional := strings.HasPrefix(include, "?")
		include = strings.TrimPrefix(include, "?")
		if include == "" {
			continue
		}

		included, includeErr := l.read(sources.ResolvePath(filepath.Dir(path), include), !optional)
		if includeErr != nil {
			return nil, includeErr
		}
		entries = append(entries, included...)
	}
	return entries, nil
}

func lastValue(entries []entry, key string) (value string) {
	for _, e := range entries {
		if e.key == key {
			value = e.value
		}
	}
	return value
}

// pickTheme returns the theme to use from the "theme" setting, which is either a
// name or path, or a "light:X,dark:Y" pair.
func pickTheme(value string) string {
	if !strings.Contains(value, "light:") && !strings.Contains(value, "dark:") {
		return value
	}

	var light, dark string
	for _, part := range strings.Split(value, ",") {
		k, v, _ := strings.Cut(strings.TrimSpace(part), ":")
		switch k {
		case "light":
			light = v
		case "dark":
			dark = v
		}
	}
	if dark != "" {
		return dark
	}
	return light
}

// loadTheme returns the colors of the theme, found by path, in the theme
// directories (in which case file is where it was found, so it can be watched),
// or finally among bubbletint's built-in tints, as Ghostty's themes come from
// the same collection. The palette is empty if the theme can't be found.
func loadTheme(theme string) (p sources.Palette, file string) {
	candidates := sources.JoinAll(themeDirs(), theme)
	if expanded := sources.ExpandPath(theme); filepath.IsAbs(expanded) {
		candidates = []string{expanded}
	}

	for _, path := range candidates {
		if !sources.IsFile(path) {
			continue
		}
		if entries, err := (&loader{}).read(path, true); err == nil {
			apply(&p, entries)
			return p, path
		}
	}

	// Prefer an exact name, as IDs lose information ("Dracula+" and "Dracula").
	builtin := tint.DefaultTints()
	for _, t := range builtin {
		if strings.EqualFold(t.DisplayName, theme) {
			return sources.PaletteOf(t), ""
		}
	}
	for _, t := range builtin {
		if t.ID == sources.ID(theme) {
			return sources.PaletteOf(t), ""
		}
	}
	return p, ""
}

// apply sets the colors in entries on the palette.
func apply(p *sources.Palette, entries []entry) {
	for _, e := range entries {
		if e.key == "palette" {
			idx, color, ok := strings.Cut(e.value, "=")
			i, err := strconv.Atoi(strings.TrimSpace(idx))
			c, cok := sources.ParseColor(color)
			if ok && err == nil && cok {
				p.SetANSI(i, c)
			}
			continue
		}

		c, ok := sources.ParseColor(e.value)
		if !ok {
			continue
		}

		switch e.key {
		case "foreground":
			p.Foreground = c
		case "background":
			p.Background = c
		case "cursor-color":
			p.Cursor = c
		case "cursor-text":
			p.CursorText = c
		case "selection-foreground":
			p.SelectionForeground = c
		case "selection-background":
			p.SelectionBackground = c
		}
	}
}

func unquote(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		if u, err := strconv.Unquote(s); err == nil {
			return u
		}
		return s[1 : len(s)-1]
	}
	return s
}
