// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package terminal picks the [tint.Source] that matches the terminal the process
// is running in, so an application can follow the user's terminal theme without
// knowing which terminal that is.
//
//	src, err := terminal.Detect()
//	if err != nil {
//		return err // Unknown terminal. Keep the app's own theme.
//	}
//	detach, err := tint.Attach(ctx, src)
//
// Terminals are recognized by the environment variables they set for the
// programs they run (and $TERM), which works the same way on Linux, macOS,
// Windows and BSDs. On Windows, Windows Terminal is also recognized by being the
// host of the console, as programs it didn't start itself (when it's the default
// terminal application) don't have its environment variables. Note that neither
// tells which config file is in use, so the default locations are searched; use
// the terminal's own package with sources.WithPath for others.
package terminal

import (
	"fmt"
	"os"
	"strings"

	"github.com/lrstanley/bubbletint/sources/v2"
	"github.com/lrstanley/bubbletint/sources/v2/alacritty"
	"github.com/lrstanley/bubbletint/sources/v2/foot"
	"github.com/lrstanley/bubbletint/sources/v2/ghostty"
	"github.com/lrstanley/bubbletint/sources/v2/kitty"
	"github.com/lrstanley/bubbletint/sources/v2/windowsterminal"
	"github.com/lrstanley/bubbletint/sources/v2/xresources"
	tint "github.com/lrstanley/bubbletint/v2"
)

// ErrUnknown is returned by [Detect] when the terminal isn't one that has a
// source. It wraps [tint.ErrNotFound].
var ErrUnknown = fmt.Errorf("%w: unrecognized terminal", tint.ErrNotFound)

// Detect returns the source for the terminal the process is running in.
func Detect(opts ...sources.Option) (tint.Source, error) {
	host, hostErr := hostProcess()

	src, err := detect(os.Getenv, host, opts)
	switch {
	case err == nil:
	case hostErr != nil:
		err = fmt.Errorf("%w (console host unknown: %w)", err, hostErr)
	case host != "":
		err = fmt.Errorf("%w (console host: %s)", err, host)
	}
	return src, err
}

// All returns a source for each of the supported terminals, whether or not it's
// installed. Attaching one of them for a terminal that has no config fails with
// an error wrapping [tint.ErrNotFound].
func All(opts ...sources.Option) []tint.Source {
	return []tint.Source{
		kitty.New(opts...),
		ghostty.New(opts...),
		alacritty.New(opts...),
		foot.New(opts...),
		windowsterminal.New(opts...),
		xresources.New(opts...),
	}
}

// detect picks a source from the environment, and failing that, from the name of
// the process hosting the console, if known (see hostProcess).
func detect(env func(string) string, host string, opts []sources.Option) (tint.Source, error) {
	term := env("TERM")
	program := strings.ToLower(env("TERM_PROGRAM"))

	switch {
	case env("KITTY_WINDOW_ID") != "" || term == "xterm-kitty":
		return kitty.New(opts...), nil
	case env("GHOSTTY_RESOURCES_DIR") != "" || env("GHOSTTY_BIN_DIR") != "" ||
		program == "ghostty" || term == "xterm-ghostty":
		return ghostty.New(opts...), nil
	case env("ALACRITTY_WINDOW_ID") != "" || env("ALACRITTY_SOCKET") != "" ||
		env("ALACRITTY_LOG") != "" || term == "alacritty" || term == "alacritty-direct":
		return alacritty.New(opts...), nil
	case env("WT_SESSION") != "":
		return windowsterminal.New(opts...), nil
	case strings.HasPrefix(term, "foot"):
		return foot.New(opts...), nil
	case env("XTERM_VERSION") != "" || strings.HasPrefix(term, "rxvt") || term == "st" || strings.HasPrefix(term, "st-"):
		// xterm, urxvt and st are all configured through X resources.
		return xresources.New(opts...), nil
	case host == "windowsterminal.exe":
		return windowsterminal.New(opts...), nil
	}
	return nil, ErrUnknown
}
