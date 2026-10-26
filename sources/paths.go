// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package sources

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// HomeDir returns the user's home directory, or an empty string if it can't be
// determined.
func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home
}

// ExpandPath expands a leading "~" to the user's home directory, and
// environment variables ($VAR and ${VAR}) in path.
func ExpandPath(path string) string {
	path = os.ExpandEnv(path)
	if path == "~" || strings.HasPrefix(path, "~/") || strings.HasPrefix(path, `~\`) {
		if home := HomeDir(); home != "" {
			path = filepath.Join(home, path[1:])
		}
	}
	return path
}

// ResolvePath expands path (see [ExpandPath]), and, if it is still relative,
// resolves it relative to dir.
func ResolvePath(dir, path string) string {
	path = ExpandPath(path)
	if path != "" && !filepath.IsAbs(path) {
		path = filepath.Join(dir, path)
	}
	return filepath.Clean(path)
}

// IsFile reports whether path exists and is a regular file (after following
// symlinks).
func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
}

// IsDir reports whether path exists and is a directory.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// ConfigDirs returns the directories in which applications keep their per-user
// config on the current platform, in order of preference, without duplicates.
// On every platform that is $XDG_CONFIG_HOME (or ~/.config), as many of the
// cross-platform terminals use it even on macOS and Windows. In addition:
//
//   - macOS: ~/Library/Application Support and ~/Library/Preferences
//   - Windows: %APPDATA% and %LOCALAPPDATA%
func ConfigDirs() []string {
	var dirs []string
	add := func(dir string) {
		if dir == "" {
			return
		}
		dir = filepath.Clean(dir)
		for _, d := range dirs {
			if d == dir {
				return
			}
		}
		dirs = append(dirs, dir)
	}

	home := HomeDir()

	if xdg := os.Getenv("XDG_CONFIG_HOME"); filepath.IsAbs(xdg) {
		add(xdg)
	}
	if home != "" {
		add(filepath.Join(home, ".config"))
	}

	switch runtime.GOOS {
	case "darwin":
		if home != "" {
			add(filepath.Join(home, "Library", "Application Support"))
			add(filepath.Join(home, "Library", "Preferences"))
		}
	case "windows":
		add(os.Getenv("APPDATA"))
		add(os.Getenv("LOCALAPPDATA"))
	}

	return dirs
}

// SystemConfigDirs returns the directories in which system-wide config is kept
// on the current platform ($XDG_CONFIG_DIRS, or /etc/xdg, on Unix-likes, and
// %ProgramData% on Windows).
func SystemConfigDirs() []string {
	if runtime.GOOS == "windows" {
		if dir := os.Getenv("ProgramData"); dir != "" {
			return []string{dir}
		}
		return nil
	}

	var dirs []string
	for _, dir := range filepath.SplitList(os.Getenv("XDG_CONFIG_DIRS")) {
		if filepath.IsAbs(dir) {
			dirs = append(dirs, dir)
		}
	}
	if len(dirs) == 0 {
		dirs = []string{"/etc/xdg"}
	}
	return dirs
}

// JoinAll joins each of dirs with elem.
func JoinAll(dirs []string, elem ...string) []string {
	out := make([]string, 0, len(dirs))
	for _, dir := range dirs {
		out = append(out, filepath.Join(append([]string{dir}, elem...)...))
	}
	return out
}
