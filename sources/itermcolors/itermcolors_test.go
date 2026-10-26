// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package itermcolors

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	tint "github.com/lrstanley/bubbletint/v2"
)

func colorDict(r, g, b float64) string {
	return fmt.Sprintf(`<dict>
		<key>Alpha Component</key><real>1</real>
		<key>Blue Component</key><real>%v</real>
		<key>Color Space</key><string>sRGB</string>
		<key>Green Component</key><real>%v</real>
		<key>Red Component</key><real>%v</real>
	</dict>`, b, g, r)
}

func preset(bg, red string) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Ansi 1 Color</key>` + red + `
	<key>Background Color</key>` + bg + `
	<key>Foreground Color</key>` + colorDict(1, 1, 1) + `
	<key>Cursor Color</key>` + colorDict(0, 1, 0) + `
	<key>Some Flag</key><true/>
	<key>List</key><array><string>a</string><integer>1</integer></array>
</dict>
</plist>`
}

func TestParse(t *testing.T) {
	t.Parallel()

	got, err := Parse(strings.NewReader(preset(colorDict(0.1, 0.2, 0.4), colorDict(1, 0, 0))), "x", "X")
	if err != nil {
		t.Fatal(err)
	}

	for name, c := range map[string][2]string{
		"bg":     {got.Bg.Hex(), "#1a3366"},
		"fg":     {got.Fg.Hex(), "#ffffff"},
		"red":    {got.Red.Hex(), "#ff0000"},
		"cursor": {got.Cursor.Hex(), "#00ff00"},
	} {
		if c[0] != c[1] {
			t.Errorf("%s: got %s, want %s", name, c[0], c[1])
		}
	}
	if got.ID != "x" || got.DisplayName != "X" {
		t.Errorf("unexpected names %q %q", got.ID, got.DisplayName)
	}

	for _, bad := range []string{"", "not xml", "<plist><array/></plist>", "<plist><dict></dict></plist>"} {
		if _, err = Parse(strings.NewReader(bad), "x", "X"); err == nil {
			t.Errorf("expected an error parsing %q", bad)
		}
	}
}

func TestSource(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	write := func(name, content string) string {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return path
	}

	file := write("My Theme.itermcolors", preset(colorDict(0, 0, 0), colorDict(1, 0, 0)))
	write("b.itermcolors", preset(colorDict(1, 1, 1), colorDict(1, 0, 0)))
	write("broken.itermcolors", "garbage")
	write("readme.txt", "ignored")

	tints, err := New(file).Load(context.Background())
	if err != nil || len(tints) != 1 || tints[0].ID != "current" || tints[0].DisplayName != "My Theme" {
		t.Fatalf("unexpected result for a file: %v, %v", tints, err)
	}

	tints, err = New(dir).Load(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(tints) != 2 || tints[0].ID != "b" || tints[1].ID != "my_theme" {
		t.Fatalf("unexpected result for a directory: %v", tints)
	}

	if _, err = New(t.TempDir()).Load(context.Background()); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound for an empty directory, got %v", err)
	}
	if _, err = New(filepath.Join(dir, "missing")).Load(context.Background()); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound for a missing path, got %v", err)
	}
}
