// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const exampleScheme = `{
  "name": "Example",
  "black": "#000000",
  "red": "#ff0000",
  "green": "#00ff00",
  "yellow": "#ffff00",
  "blue": "#0000ff",
  "purple": "#800080",
  "cyan": "#00ffff",
  "white": "#ffffff",
  "brightBlack": "#808080",
  "brightRed": "#ff8080",
  "brightGreen": "#80ff80",
  "brightYellow": "#ffff80",
  "brightBlue": "#8080ff",
  "brightPurple": "#ff80ff",
  "brightCyan": "#80ffff",
  "brightWhite": "#ffffff",
  "background": "#000000",
  "foreground": "#ffffff",
  "cursor": "#ffffff",
  "selection": "#808080"
}`

func TestLoadTints(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "Example.json"), exampleScheme)
	writeFile(t, filepath.Join(dir, "Broken.json"), `{"name":"Broken"}`)
	writeFile(t, filepath.Join(dir, "notes.txt"), "ignore me")

	tints, err := loadTints(dir)
	if err != nil {
		t.Fatalf("loadTints() error = %v", err)
	}
	if len(tints) != 1 {
		t.Fatalf("loadTints() len = %d, want 1", len(tints))
	}
	if tints[0].Name != "Example" {
		t.Errorf("Name = %q, want Example", tints[0].Name)
	}
	if tints[0].Purple != "#800080" {
		t.Errorf("Purple = %q, want #800080", tints[0].Purple)
	}
	if tints[0].BrightPurple != "#ff80ff" {
		t.Errorf("BrightPurple = %q, want #ff80ff", tints[0].BrightPurple)
	}
	if tints[0].IsDark() != true {
		t.Error("IsDark() = false, want true from black background")
	}
}

func TestApplyYAML(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "Example.yml"), `
name: "Example"
variant: "light"
selection_text: "eeeeee"
bold: "#ff00ff"
cursor_text: "#111111"
author: "Ada (https://example.com)"
`)

	tints := []Tint{{Name: "Example", Bg: "#000000"}}
	if err := applyYAML(tints, dir); err != nil {
		t.Fatalf("applyYAML() error = %v", err)
	}
	got := tints[0]
	if got.SelectionText != "#eeeeee" || got.Bold != "#ff00ff" || got.CursorText != "#111111" {
		t.Errorf("extra colors = selection %q bold %q cursor_text %q", got.SelectionText, got.Bold, got.CursorText)
	}
	if got.Author != "Ada (https://example.com)" {
		t.Errorf("Author = %q", got.Author)
	}
	if got.IsDark() {
		t.Error("IsDark() = true, want false from variant light")
	}
}

func TestNormalizeHex(t *testing.T) {
	t.Parallel()

	if got := normalizeHex("062329"); got != "#062329" {
		t.Errorf("normalizeHex() = %q, want #062329", got)
	}
	if got := normalizeHex("#ffffff"); got != "#ffffff" {
		t.Errorf("normalizeHex() = %q, want #ffffff", got)
	}
}

func TestLoadTintsParseError(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "Bad.json"), "not-json")

	_, err := loadTints(dir)
	if err == nil {
		t.Fatal("loadTints() error = nil, want parse error")
	}
	if !strings.Contains(err.Error(), "Bad.json") {
		t.Errorf("loadTints() error = %v, want file name", err)
	}
}

func TestVhsDirRejectsTag(t *testing.T) {
	t.Parallel()

	_, err := schemesRepo(t.Context(), "--upload-pack=evil")
	if err == nil {
		t.Fatal("schemesRepo() error = nil, want invalid tag")
	}
}

func TestHeadMatchesTag(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "README"), "x")
	git(t, dir, "init")
	git(t, dir, "add", "README")
	git(t, dir, "commit", "-m", "init")
	git(t, dir, "tag", "release-test")

	match, err := headMatchesTag(t.Context(), dir, "release-test")
	if err != nil {
		t.Fatalf("headMatchesTag() error = %v", err)
	}
	if !match {
		t.Fatal("headMatchesTag() = false, want true")
	}

	match, err = headMatchesTag(t.Context(), dir, "release-missing")
	if err != nil {
		t.Fatalf("headMatchesTag() missing tag error = %v", err)
	}
	if match {
		t.Fatal("headMatchesTag() missing tag = true, want false")
	}

	writeFile(t, filepath.Join(dir, "README"), "y")
	git(t, dir, "add", "README")
	git(t, dir, "commit", "-m", "second")

	match, err = headMatchesTag(t.Context(), dir, "release-test")
	if err != nil {
		t.Fatalf("headMatchesTag() moved HEAD error = %v", err)
	}
	if match {
		t.Fatal("headMatchesTag() moved HEAD = true, want false")
	}
}

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func git(t *testing.T, dir string, args ...string) {
	t.Helper()

	cmd := exec.Command("git", append([]string{
		"-c", "commit.gpgsign=false",
		"-c", "tag.gpgsign=false",
		"-c", "user.name=tintgen",
		"-c", "user.email=tintgen@example.com",
	}, args...)...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		"GIT_AUTHOR_NAME=tintgen",
		"GIT_AUTHOR_EMAIL=tintgen@example.com",
		"GIT_COMMITTER_NAME=tintgen",
		"GIT_COMMITTER_EMAIL=tintgen@example.com",
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %s: %v: %s", strings.Join(args, " "), err, out)
	}
}
