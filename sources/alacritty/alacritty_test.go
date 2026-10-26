// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package alacritty

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/lrstanley/bubbletint/sources/v2/internal/testutil"

	tint "github.com/lrstanley/bubbletint/v2"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	testutil.Write(t, filepath.Join(dir, "themes", "nord.toml"), `
[colors.primary]
background = "0x2e3440"
foreground = "#d8dee9"
[colors.cursor]
text = "#2e3440"
cursor = "#d8dee9"
[colors.normal]
black   = "#3b4252"
red     = "#bf616a"
magenta = "#b48ead"
[colors.bright]
black = "#4c566a"
white = "#eceff4"
`)
	testutil.Write(t, filepath.Join(dir, "other.toml"), "[colors.normal]\nred = \"#111111\"\n")
	testutil.Write(t, filepath.Join(dir, "alacritty.toml"), `
[general]
import = ["themes/nord.toml", "other.toml", "missing.toml", "alacritty.toml"]

[font]
size = 12

[colors.normal]
black = "#000001"
`)

	got, err := Load(filepath.Join(dir, "alacritty.toml"))
	if err != nil {
		t.Fatal(err)
	}

	for name, c := range map[string][2]string{
		"bg (0x form)":            {got.Bg.Hex(), "#2e3440"},
		"fg":                      {got.Fg.Hex(), "#d8dee9"},
		"cursor":                  {got.Cursor.Hex(), "#d8dee9"},
		"cursor text":             {got.CursorText.Hex(), "#2e3440"},
		"black (main file wins)":  {got.Black.Hex(), "#000001"},
		"red (later import wins)": {got.Red.Hex(), "#111111"},
		"purple (magenta)":        {got.Purple.Hex(), "#b48ead"},
		"bright black":            {got.BrightBlack.Hex(), "#4c566a"},
		"bright white":            {got.BrightWhite.Hex(), "#eceff4"},
	} {
		if c[0] != c[1] {
			t.Errorf("%s: got %s, want %s", name, c[0], c[1])
		}
	}
	if got.DisplayName != "nord" {
		t.Errorf("expected name from the theme import, got %q", got.DisplayName)
	}
}

func TestLoad_LegacyImport(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	testutil.Write(t, filepath.Join(dir, "theme.toml"), "[colors.primary]\nbackground = \"#ffffff\"\n")
	testutil.Write(t, filepath.Join(dir, "alacritty.toml"), "import = [\"theme.toml\"]\n")

	got, err := Load(filepath.Join(dir, "alacritty.toml"))
	if err != nil {
		t.Fatal(err)
	}
	if got.Bg.Hex() != "#ffffff" || got.Dark {
		t.Errorf("unexpected tint: %s dark=%v", got.Bg.Hex(), got.Dark)
	}
}

func TestLoad_Errors(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()

	testutil.Write(t, filepath.Join(dir, "nocolors.toml"), "[font]\nsize = 12\n")
	if _, err := Load(filepath.Join(dir, "nocolors.toml")); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	testutil.Write(t, filepath.Join(dir, "broken.toml"), "[colors\n")
	if _, err := Load(filepath.Join(dir, "broken.toml")); err == nil {
		t.Fatal("expected an error for invalid TOML")
	}

	// A broken import is skipped, like Alacritty does, rather than failing.
	testutil.Write(t, filepath.Join(dir, "broken-import.toml"), "[colors\n")
	testutil.Write(t, filepath.Join(dir, "main.toml"), "import = [\"broken-import.toml\"]\n[colors.primary]\nbackground = \"#000000\"\n")
	if _, err := Load(filepath.Join(dir, "main.toml")); err != nil {
		t.Fatalf("expected a broken import to be skipped, got %v", err)
	}
}
