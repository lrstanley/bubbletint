// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ghostty

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/lrstanley/bubbletint/sources/v2/internal/testutil"

	tint "github.com/lrstanley/bubbletint/v2"
)

func TestLoad_ThemeFile(t *testing.T) {
	cfg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", cfg)
	t.Setenv("GHOSTTY_RESOURCES_DIR", "")

	testutil.Write(t, filepath.Join(cfg, "ghostty", "themes", "mytheme"), `
background = #101010
foreground = #eeeeee
palette = 1=#ff0000
palette = 2=#00ff00
selection-background = #333333
`)
	testutil.Write(t, filepath.Join(cfg, "ghostty", "config"), `
# comment
font-size = 12
theme = mytheme
background = "#202020"
palette = 2=#00aa00
config-file = extra.conf
config-file = ?missing.conf
`)
	testutil.Write(t, filepath.Join(cfg, "ghostty", "extra.conf"), "cursor-color = #abcdef\n")

	got, err := Load(filepath.Join(cfg, "ghostty", "config"))
	if err != nil {
		t.Fatal(err)
	}

	for name, c := range map[string][2]string{
		"bg (config overrides theme)": {got.Bg.Hex(), "#202020"},
		"fg (from theme)":             {got.Fg.Hex(), "#eeeeee"},
		"red (from theme)":            {got.Red.Hex(), "#ff0000"},
		"green (config overrides)":    {got.Green.Hex(), "#00aa00"},
		"selection (from theme)":      {got.SelectionBg.Hex(), "#333333"},
		"cursor (from config-file)":   {got.Cursor.Hex(), "#abcdef"},
	} {
		if c[0] != c[1] {
			t.Errorf("%s: got %s, want %s", name, c[0], c[1])
		}
	}
	if got.DisplayName != "mytheme" || got.ID != "current" {
		t.Errorf("unexpected names: %q %q", got.ID, got.DisplayName)
	}
}

func TestLoad_BuiltinThemeFallback(t *testing.T) {
	cfg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", cfg)
	t.Setenv("GHOSTTY_RESOURCES_DIR", "")

	// No theme file anywhere, so it falls back to bubbletint's tint of that name.
	testutil.Write(t, filepath.Join(cfg, "ghostty", "config"), "theme = light:Dracula+,dark:Dracula+\nbackground = #000001\n")

	got, err := Load(filepath.Join(cfg, "ghostty", "config"))
	if err != nil {
		t.Fatal(err)
	}

	if got.Bg.Hex() != "#000001" {
		t.Errorf("expected config to override the built-in theme, got %s", got.Bg.Hex())
	}
	if want := tint.TintDraculaPlus.Fg.Hex(); got.Fg.Hex() != want {
		t.Errorf("expected fg from the built-in theme %s, got %s", want, got.Fg.Hex())
	}
	if got.Red.Hex() != tint.TintDraculaPlus.Red.Hex() {
		t.Errorf("expected ANSI colors from the built-in theme")
	}

	// And with nothing to override it, the built-in tint is used as-is (copied).
	testutil.Write(t, filepath.Join(cfg, "ghostty", "config"), "theme = Dracula+\n")
	got, err = Load(filepath.Join(cfg, "ghostty", "config"))
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != "current" || got.Bg.Hex() != tint.TintDraculaPlus.Bg.Hex() || tint.TintDraculaPlus.ID == "current" {
		t.Errorf("unexpected tint %q %s (built-in ID %q)", got.ID, got.Bg.Hex(), tint.TintDraculaPlus.ID)
	}
}

func TestLoad_NoColors(t *testing.T) {
	cfg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", cfg)

	testutil.Write(t, filepath.Join(cfg, "ghostty", "config"), "font-size = 12\ntheme = does-not-exist\n")
	if _, err := Load(filepath.Join(cfg, "ghostty", "config")); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestLoad_RequiredIncludeMissing(t *testing.T) {
	cfg := t.TempDir()
	testutil.Write(t, filepath.Join(cfg, "config"), "config-file = nope.conf\nbackground = #000000\n")
	if _, err := Load(filepath.Join(cfg, "config")); err == nil {
		t.Fatal("expected an error for a missing, non-optional include")
	}
}

func TestPickTheme(t *testing.T) {
	t.Parallel()

	for in, want := range map[string]string{
		"nord":                           "nord",
		"light:Rose Pine Dawn,dark:Nord": "Nord",
		"light:Solo":                     "Solo",
	} {
		if got := pickTheme(in); got != want {
			t.Errorf("pickTheme(%q) = %q, want %q", in, got, want)
		}
	}
}
