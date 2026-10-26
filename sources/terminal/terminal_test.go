// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package terminal

import (
	"errors"
	"testing"

	tint "github.com/lrstanley/bubbletint/v2"
)

func TestDetect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		env  map[string]string
		host string
		want string
	}{
		{"kitty window", map[string]string{"KITTY_WINDOW_ID": "1", "TERM": "xterm-256color"}, "", "kitty"},
		{"kitty term", map[string]string{"TERM": "xterm-kitty"}, "", "kitty"},
		{"ghostty resources", map[string]string{"GHOSTTY_RESOURCES_DIR": "/x", "TERM": "xterm-256color"}, "", "ghostty"},
		{"ghostty program", map[string]string{"TERM_PROGRAM": "ghostty"}, "", "ghostty"},
		{"alacritty socket", map[string]string{"ALACRITTY_SOCKET": "/x"}, "", "alacritty"},
		{"alacritty term", map[string]string{"TERM": "alacritty"}, "", "alacritty"},
		{"windows terminal", map[string]string{"WT_SESSION": "abc", "TERM": "xterm-256color"}, "", "windowsterminal"},
		{"foot", map[string]string{"TERM": "foot-direct"}, "", "foot"},
		{"windows terminal host", map[string]string{}, "windowsterminal.exe", "windowsterminal"},
		{"xterm", map[string]string{"XTERM_VERSION": "XTerm(379)", "TERM": "xterm-256color"}, "", "xresources"},
		{"urxvt", map[string]string{"TERM": "rxvt-unicode-256color"}, "", "xresources"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			src, err := detect(func(k string) string { return tt.env[k] }, tt.host, nil)
			if err != nil {
				t.Fatal(err)
			}
			if src.Name() != tt.want {
				t.Errorf("got %q, want %q", src.Name(), tt.want)
			}
		})
	}

	// Plenty of terminals just say xterm-256color. That tells us nothing.
	_, err := detect(func(k string) string {
		return map[string]string{"TERM": "xterm-256color", "TERM_PROGRAM": "Apple_Terminal"}[k]
	}, "explorer.exe", nil)
	if !errors.Is(err, ErrUnknown) || !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrUnknown wrapping ErrNotFound, got %v", err)
	}
}

func TestAll(t *testing.T) {
	t.Parallel()

	seen := map[string]bool{}
	for _, src := range All() {
		if seen[src.Name()] {
			t.Errorf("duplicate source name %q", src.Name())
		}
		seen[src.Name()] = true
	}
	if len(seen) < 6 {
		t.Errorf("expected every terminal to be included, got %v", seen)
	}
}
