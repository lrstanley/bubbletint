// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package sources

import (
	"bufio"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Walker reads line-based config files that can include one another, keeping
// track of which files were read so that all of them can be watched.
type Walker struct {
	// Files is every file that was scanned, in order, including includes that
	// turned out not to exist.
	Files []string

	visited map[string]bool
}

// Scan calls fn with each non-empty line of the file at path, with surrounding
// whitespace removed. fn may call Scan to follow an include, and returning an error
// from it stops the scan.
//
// A file that was already scanned by this Walker is skipped, which ends include
// cycles. A file that doesn't exist is an error only if required, as most
// terminals ignore includes they can't find.
func (w *Walker) Scan(path string, required bool, fn func(line string) error) error {
	w.Files = append(w.Files, path)

	abs, err := filepath.Abs(path)
	if err != nil {
		abs = path
	}
	if w.visited[abs] {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		if !required && errors.Is(err, fs.ErrNotExist) {
			return nil
		}
		return err
	}
	defer f.Close()

	if w.visited == nil {
		w.visited = make(map[string]bool)
	}
	w.visited[abs] = true

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if fnErr := fn(line); fnErr != nil {
			return fnErr
		}
	}
	return scanner.Err()
}
