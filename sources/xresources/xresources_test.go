// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package xresources

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
	testutil.Write(t, filepath.Join(dir, "colors"), "*.color2: #00ff00\n")
	testutil.Write(t, filepath.Join(dir, ".Xresources"), `
! comment
#define bg #1d1f21
#include "colors"
Xft.dpi: 96
*.foreground:   #c5c8c6
*.background:   bg
URxvt*color1: #cc6666
xterm*color0:#000000
*.cursorColor: rgb:ff/ff/00
*color9: #ff0000
*.colorBD: #ffffff
`)

	got, err := Load(filepath.Join(dir, ".Xresources"))
	if err != nil {
		t.Fatal(err)
	}

	for name, c := range map[string][2]string{
		"fg":     {got.Fg.Hex(), "#c5c8c6"},
		"bg":     {got.Bg.Hex(), "#1d1f21"},
		"black":  {got.Black.Hex(), "#000000"},
		"red":    {got.Red.Hex(), "#cc6666"},
		"green":  {got.Green.Hex(), "#00ff00"},
		"bright": {got.BrightRed.Hex(), "#ff0000"},
		"cursor": {got.Cursor.Hex(), "#ffff00"},
		"bold":   {got.Bold.Hex(), "#ffffff"},
	} {
		if c[0] != c[1] {
			t.Errorf("%s: got %s, want %s", name, c[0], c[1])
		}
	}
}

func TestLoad_NoColors(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), ".Xresources")
	testutil.Write(t, path, "Xft.dpi: 96\n")
	if _, err := Load(path); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
