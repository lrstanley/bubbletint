// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package foot

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
	testutil.Write(t, filepath.Join(dir, "themes", "gruvbox"), `
[colors]
background=282828
foreground=ebdbb2
regular0=282828
regular1=cc241d
bright0=928374
bright7=ebdbb2
`)
	testutil.Write(t, filepath.Join(dir, "foot.ini"), `
# comment
include=themes/gruvbox
font=monospace:size=12

[cursor]
color=111111 dcdccc

[colors-dark]
background=000002
regular1=ff0000
`)

	got, err := Load(filepath.Join(dir, "foot.ini"))
	if err != nil {
		t.Fatal(err)
	}

	// [colors-dark] takes priority over [colors], which isn't merged with it.
	if got.Bg.Hex() != "#000002" || got.Red.Hex() != "#ff0000" {
		t.Errorf("expected [colors-dark] to be used, got bg=%s red=%s", got.Bg.Hex(), got.Red.Hex())
	}
	if got.Cursor.Hex() != "#dcdccc" || got.CursorText.Hex() != "#111111" {
		t.Errorf("unexpected cursor %s / %s", got.Cursor.Hex(), got.CursorText.Hex())
	}

	// With only an included [colors] section, that's used.
	testutil.Write(t, filepath.Join(dir, "plain.ini"), "include=themes/gruvbox\n")
	got, err = Load(filepath.Join(dir, "plain.ini"))
	if err != nil {
		t.Fatal(err)
	}
	if got.Bg.Hex() != "#282828" || got.Red.Hex() != "#cc241d" || got.BrightBlack.Hex() != "#928374" || got.BrightWhite.Hex() != "#ebdbb2" {
		t.Errorf("unexpected colors from include: %s %s %s %s",
			got.Bg.Hex(), got.Red.Hex(), got.BrightBlack.Hex(), got.BrightWhite.Hex())
	}
}

func TestLoad_NoColors(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "foot.ini")
	testutil.Write(t, path, "font=monospace\n")
	if _, err := Load(path); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
