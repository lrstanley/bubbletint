// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package sources_test

import (
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/lrstanley/bubbletint/sources/v2"
	"github.com/lrstanley/bubbletint/sources/v2/internal/testutil"
)

func TestWalker_Scan(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	a, b := filepath.Join(dir, "a"), filepath.Join(dir, "b")
	testutil.Write(t, a, "one\n\n  two  \ninclude b\nfour\ninclude missing\n")
	testutil.Write(t, b, "three\ninclude a\n") // A cycle.

	var w sources.Walker
	var lines []string

	var handle func(line string) error
	handle = func(line string) error {
		if inc, ok := strings.CutPrefix(line, "include "); ok {
			return w.Scan(filepath.Join(dir, inc), false, handle)
		}
		lines = append(lines, line)
		return nil
	}

	if err := w.Scan(a, true, handle); err != nil {
		t.Fatal(err)
	}

	// Includes are followed in place, blank lines are dropped, and the cycle ends.
	if want := []string{"one", "two", "three", "four"}; !slices.Equal(lines, want) {
		t.Errorf("got lines %v, want %v", lines, want)
	}

	// Every file is recorded, so it can be watched, even the one that doesn't exist.
	if want := []string{a, b, a, filepath.Join(dir, "missing")}; !slices.Equal(w.Files, want) {
		t.Errorf("got files %v, want %v", w.Files, want)
	}

	// A required file that's missing is an error.
	if err := w.Scan(filepath.Join(dir, "missing"), true, handle); err == nil {
		t.Error("expected an error for a missing, required file")
	}
}
