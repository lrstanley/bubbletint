// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package sources_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/lrstanley/bubbletint/sources/v2/internal/testutil"

	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

func TestParseColor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   string
		want string
		ok   bool
	}{
		{"#ff8000", "#ff8000", true},
		{"#F80", "#ff8800", true},
		{"ff8000", "#ff8000", true},
		{"0xff8000", "#ff8000", true},
		{`"#ff8000"`, "#ff8000", true},
		{"#ff8000cc", "#ff8000", true},
		{"rgb:ff/80/00", "#ff8000", true},
		{"rgb:f/8/0", "#ff8800", true},
		{"rgb:ffff/8080/0000", "#ff8000", true},
		{"red", "", false},
		{"none", "", false},
		{"", "", false},
		{"#12345", "", false},
		{"#gggggg", "", false},
		{"rgb:ff/80", "", false},
	}
	for _, tt := range tests {
		got, ok := sources.ParseColor(tt.in)
		if ok != tt.ok || (ok && got.Hex() != tt.want) {
			t.Errorf("ParseColor(%q) = %v, %v; want %s, %v", tt.in, got, ok, tt.want, tt.ok)
		}
	}
}

func TestID(t *testing.T) {
	t.Parallel()

	for in, want := range map[string]string{
		"One Half Dark":     "one_half_dark",
		"  Dracula+ ":       "dracula",
		"tokyo-night_storm": "tokyo_night_storm",
		"3024 Day":          "3024_day",
	} {
		if got := sources.ID(in); got != want {
			t.Errorf("ID(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestPalette_Tint(t *testing.T) {
	t.Parallel()

	var p sources.Palette
	if !p.Empty() {
		t.Fatal("expected palette to be empty")
	}

	p.Background, _ = sources.ParseColor("#ffffff")
	p.SetANSI(1, &tint.Color{R: 1, G: 2, B: 3, A: 255})
	p.SetANSI(99, &tint.Color{}) // Ignored.

	got := p.Tint("x", "X")
	if got.Dark {
		t.Error("expected a light background to be a light tint")
	}
	if got.Fg.Hex() != "#000000" {
		t.Errorf("expected missing fg to contrast with bg, got %s", got.Fg.Hex())
	}
	if got.Red.Hex() != "#010203" || got.Blue.Hex() != "#0000ee" || got.BrightWhite.Hex() != "#ffffff" {
		t.Errorf("unexpected ANSI colors: red=%s blue=%s bright_white=%s", got.Red.Hex(), got.Blue.Hex(), got.BrightWhite.Hex())
	}
	if got.Cursor != nil {
		t.Error("expected unset optional colors to stay nil")
	}
}

func testSpec(name string) sources.Spec {
	return sources.Spec{
		Name: name,
		Load: func(path string) ([]*tint.Tint, []string, error) {
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, nil, err
			}
			c, ok := sources.ParseColor(string(data))
			if !ok {
				return nil, nil, errors.New("bad color")
			}
			p := sources.Palette{Background: c}
			return []*tint.Tint{p.Tint("current", "Test")}, nil, nil
		},
	}
}

func TestFileSource_NotFound(t *testing.T) {
	t.Parallel()

	src := sources.New(testSpec("test"), sources.WithPath(filepath.Join(t.TempDir(), "missing")))
	if _, err := src.Load(context.Background()); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	// No locations at all (e.g. a terminal that doesn't exist on this platform).
	if _, err := sources.New(testSpec("test")).Load(context.Background()); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestFileSource_Candidates(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	first, second := filepath.Join(dir, "a"), filepath.Join(dir, "b")
	testutil.Write(t, second, "#112233")

	spec := testSpec("test")
	spec.Candidates = func() []string { return []string{first, second} }
	src := sources.New(spec)

	got, err := src.Find()
	if err != nil || got != second {
		t.Fatalf("expected %s, got %q (%v)", second, got, err)
	}

	testutil.Write(t, first, "#445566") // Higher priority wins.
	if got, _ = src.Find(); got != first {
		t.Fatalf("expected %s, got %q", first, got)
	}
}

// TestFileSource_Attach exercises the whole flow: a config file on disk, attached
// to a registry, being edited, and removed.
func TestFileSource_Attach(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "theme")
	testutil.Write(t, path, "#112233")

	reg := tint.NewRegistry(nil, &tint.Tint{ID: "user", Bg: &tint.Color{A: 255}, Fg: &tint.Color{A: 255}})

	var mu sync.Mutex
	var events []tint.Event
	reg.Subscribe(func(ev tint.Event) {
		mu.Lock()
		events = append(events, ev)
		mu.Unlock()
	})

	src := sources.New(testSpec("test"), sources.WithPath(path), sources.WithPollInterval(5*time.Millisecond))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var errCount int
	detach, err := reg.Attach(ctx, src, tint.WithErrorHandler(func(error) {
		mu.Lock()
		errCount++
		mu.Unlock()
	}))
	if err != nil {
		t.Fatal(err)
	}
	defer detach()

	if got := reg.Current(); got.ID != "test:current" || got.Bg.Hex() != "#112233" {
		t.Fatalf("unexpected current tint %s %s", got.ID, got.Bg.Hex())
	}

	waitFor(t, func() bool { return reg.Current().Bg.Hex() == "#112233" })

	// Edit the file, making sure the modification time moves on.
	testutil.Write(t, path, "#445566")
	future := time.Now().Add(time.Minute)
	if err = os.Chtimes(path, future, future); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return reg.Current().Bg.Hex() == "#445566" })

	// A broken edit keeps the last good tint, and reports the error.
	testutil.Write(t, path, "garbage")
	later := future.Add(time.Minute)
	if err = os.Chtimes(path, later, later); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { mu.Lock(); defer mu.Unlock(); return errCount > 0 })
	if reg.Current().Bg.Hex() != "#445566" {
		t.Fatalf("expected the last good tint to be kept, got %s", reg.Current().Bg.Hex())
	}

	mu.Lock()
	defer mu.Unlock()
	var updated bool
	for _, ev := range events {
		if ev.Kind == tint.TintUpdated && ev.Source == "test" && ev.Tint.ID == "test:current" {
			updated = true
		}
	}
	if !updated {
		t.Fatalf("expected a TintUpdated event from the source, got %v", events)
	}
}

func waitFor(t *testing.T, cond func() bool) {
	t.Helper()

	deadline := time.Now().Add(5 * time.Second)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatal("timed out waiting for condition")
		}
		time.Sleep(5 * time.Millisecond)
	}
}
