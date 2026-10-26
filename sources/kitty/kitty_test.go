// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package kitty

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/lrstanley/bubbletint/sources/v2/internal/testutil"

	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	testutil.Write(t, filepath.Join(dir, "kitty.conf"), `
# A comment
font_size 12
include current-theme.conf
globinclude extra/*.conf
include missing.conf
include kitty.conf
color1	#aa0000
cursor none
`)
	testutil.Write(t, filepath.Join(dir, "current-theme.conf"), `# vim:ft=kitty
## name: Dracula
## author: someone
foreground #f8f8f2
background #282a36
selection_foreground #ffffff
selection_background #44475a
cursor #f8f8f2
cursor_text_color background
color0 #21222c
color1 #ff5555
color15 #ffffff
color16 #123456
`)
	testutil.Write(t, filepath.Join(dir, "extra", "more.conf"), "color2 #50fa7b\n")

	got, err := Load(filepath.Join(dir, "kitty.conf"))
	if err != nil {
		t.Fatal(err)
	}

	for name, c := range map[string][2]string{
		"fg":        {got.Fg.Hex(), "#f8f8f2"},
		"bg":        {got.Bg.Hex(), "#282a36"},
		"selection": {got.SelectionBg.Hex(), "#44475a"},
		"cursor":    {got.Cursor.Hex(), "#f8f8f2"},
		"black":     {got.Black.Hex(), "#21222c"},
		"red":       {got.Red.Hex(), "#aa0000"}, // Set after the include, so it wins.
		"green":     {got.Green.Hex(), "#50fa7b"},
		"white":     {got.BrightWhite.Hex(), "#ffffff"},
	} {
		if c[0] != c[1] {
			t.Errorf("%s: got %s, want %s", name, c[0], c[1])
		}
	}

	if got.ID != "current" || got.DisplayName != "Dracula" || !got.Dark {
		t.Errorf("unexpected metadata: %q %q dark=%v", got.ID, got.DisplayName, got.Dark)
	}
	if got.CursorText != nil {
		t.Error("expected non-color cursor_text_color to be ignored")
	}
}

func TestSource(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	testutil.Write(t, filepath.Join(dir, "kitty.conf"), "include theme.conf\n")
	testutil.Write(t, filepath.Join(dir, "theme.conf"), "background #000000\n")

	src := New(sources.WithPath(filepath.Join(dir, "kitty.conf")))
	tints, err := src.Load(context.Background())
	if err != nil || len(tints) != 1 || tints[0].DisplayName != "Kitty" {
		t.Fatalf("unexpected result: %v, %v", tints, err)
	}
}

func TestLoad_NoColors(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "kitty.conf")
	testutil.Write(t, path, "font_size 12\n")

	_, err := Load(path)
	if !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestCandidates(t *testing.T) {
	t.Setenv("KITTY_CONFIG_DIRECTORY", "/custom/kitty")
	t.Setenv("XDG_CONFIG_HOME", "/xdg")

	got := Candidates()
	if len(got) < 3 || got[0] != filepath.FromSlash("/custom/kitty/kitty.conf") || got[1] != filepath.FromSlash("/xdg/kitty/kitty.conf") {
		t.Fatalf("unexpected candidates: %v", got)
	}
}
