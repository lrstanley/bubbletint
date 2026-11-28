// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ghostty

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/lrstanley/bubbletint/sources/v2"
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

	got, err := first(Load(filepath.Join(cfg, "ghostty", "config"), true))
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
	orig := installThemeDirs
	installThemeDirs = func() []string { return nil }
	t.Cleanup(func() { installThemeDirs = orig })

	// No theme file anywhere, so it falls back to bubbletint's tint of that name.
	testutil.Write(t, filepath.Join(cfg, "ghostty", "config"), "theme = light:Dracula+,dark:Dracula+\nbackground = #000001\n")

	got, err := first(Load(filepath.Join(cfg, "ghostty", "config"), true))
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
	got, err = first(Load(filepath.Join(cfg, "ghostty", "config"), true))
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
	if _, err := Load(filepath.Join(cfg, "ghostty", "config"), true); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestLoad_RequiredIncludeMissing(t *testing.T) {
	cfg := t.TempDir()
	testutil.Write(t, filepath.Join(cfg, "config"), "config-file = nope.conf\nbackground = #000000\n")
	if _, err := Load(filepath.Join(cfg, "config"), true); err == nil {
		t.Fatal("expected an error for a missing, non-optional include")
	}
}

// first is the active tint Load returned.
func first(tints []*tint.Tint, err error) (*tint.Tint, error) {
	if err != nil {
		return nil, err
	}
	return tints[0], nil
}

func TestPickTheme(t *testing.T) {
	t.Parallel()

	for in, want := range map[string]themes{
		"nord":                           {dark: "nord", light: "nord"},
		"light:Rose Pine Dawn,dark:Nord": {dark: "Nord", light: "Rose Pine Dawn"},
		"dark:Nord, light:Rose Pine":     {dark: "Nord", light: "Rose Pine"},
		"light:Solo":                     {dark: "Solo", light: "Solo"},
	} {
		if got := pickTheme(in); got != want {
			t.Errorf("pickTheme(%q) = %+v, want %+v", in, got, want)
		}
	}
}

// A theme for each of light and dark mode: the active tint is the mode's,
// and both are returned, each with the config's own colors over it.
func TestLoad_Pair(t *testing.T) {
	cfg := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", cfg)
	t.Setenv("GHOSTTY_RESOURCES_DIR", "")
	orig := installThemeDirs
	installThemeDirs = func() []string { return nil }
	t.Cleanup(func() { installThemeDirs = orig })

	path := filepath.Join(cfg, "ghostty", "config")
	testutil.Write(t, path, "theme = light:Tokyo Night Day,dark:Tokyo Night\ncursor-color = #abcdef\n")

	for _, dark := range []bool{true, false} {
		tints, err := Load(path, dark)
		if err != nil {
			t.Fatal(err)
		}
		if len(tints) != 3 || tints[0].ID != "current" || tints[1].ID != sources.DarkID || tints[2].ID != sources.LightID {
			t.Fatalf("dark %v: got %d tints", dark, len(tints))
		}
		want, name := tint.TintTokyoNightDay, "Tokyo Night Day"
		if dark {
			want, name = tint.TintTokyoNight, "Tokyo Night"
		}
		if got := tints[0]; got.Bg.Hex() != want.Bg.Hex() || got.DisplayName != name {
			t.Errorf("dark %v: current is %s %s, want %s", dark, got.DisplayName, got.Bg.Hex(), name)
		}
		if tints[1].Bg.Hex() != tint.TintTokyoNight.Bg.Hex() || tints[2].Bg.Hex() != tint.TintTokyoNightDay.Bg.Hex() {
			t.Errorf("dark %v: pair is %s and %s", dark, tints[1].Bg.Hex(), tints[2].Bg.Hex())
		}
		for _, tn := range tints {
			if tn.Cursor == nil || tn.Cursor.Hex() != "#abcdef" {
				t.Errorf("dark %v: %s lost the config's cursor color", dark, tn.ID)
			}
		}
	}

	// SetDark switches the active tint, and wakes a watcher to reload.
	src := New(sources.WithPath(path), sources.WithPollInterval(time.Hour))
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	changes, err := src.Watch(ctx)
	if err != nil {
		t.Fatal(err)
	}
	select { // A watch starts with a reload.
	case <-changes:
	case <-time.After(5 * time.Second):
		t.Fatal("the watch didn't start with a reload")
	}
	src.SetDark(true) // unchanged: no reload
	src.SetDark(false)
	select {
	case <-changes:
	case <-time.After(5 * time.Second):
		t.Fatal("SetDark didn't wake the watcher")
	}
	tints, err := src.Load(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if tints[0].Bg.Hex() != tint.TintTokyoNightDay.Bg.Hex() || src.Dark() {
		t.Errorf("after SetDark(false), current is %s", tints[0].DisplayName)
	}
	select {
	case <-changes:
		t.Error("an unchanged SetDark woke the watcher")
	default:
	}
}
