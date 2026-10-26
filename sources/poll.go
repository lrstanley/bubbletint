// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package sources

import (
	"context"
	"maps"
	"os"
	"time"
)

type fileState struct {
	exists  bool
	modTime time.Time
	size    int64
}

type snapshot map[string]fileState

func takeSnapshot(paths []string) snapshot {
	s := make(snapshot, len(paths))
	for _, path := range paths {
		if _, ok := s[path]; ok {
			continue
		}
		// Stat follows symlinks, which is what we want for symlinked dotfiles.
		info, err := os.Stat(path)
		if err != nil {
			s[path] = fileState{}
			continue
		}
		s[path] = fileState{exists: true, modTime: info.ModTime(), size: info.Size()}
	}
	return s
}

func (s snapshot) equal(other snapshot) bool { return maps.Equal(s, other) }

// poll sends on the returned channel when any of the files returned by paths
// change. paths is re-evaluated each tick, as the set of files can change as a
// config is edited. A change is only reported once the files stop changing, so a
// config that is mid-write isn't read half-way. The channel is closed when ctx
// is done.
func poll(ctx context.Context, interval time.Duration, paths func() []string) <-chan struct{} {
	out := make(chan struct{}, 1)

	// Take the baseline before returning, so changes made right after the caller
	// gets the channel aren't missed.
	stable := takeSnapshot(paths())

	go func() {
		defer close(out)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		var pending snapshot

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}

			current := takeSnapshot(paths())

			switch {
			case pending != nil && current.equal(pending):
				// The change has settled.
				stable, pending = current, nil
				select {
				case out <- struct{}{}:
				default: // Already has a pending notification.
				}
			case !current.equal(stable):
				pending = current
			default:
				pending = nil
			}
		}
	}()

	return out
}
