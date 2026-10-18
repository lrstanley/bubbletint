// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const schemesRepoURL = "https://github.com/mbadolato/iTerm2-Color-Schemes.git"

// schemesTagPattern matches the release tags published by iTerm2-Color-Schemes.
var schemesTagPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$`)

func schemesCacheDir() string {
	return filepath.Join(os.TempDir(), "bubbletint", "iTerm2-Color-Schemes")
}

// loadSchemes clones (or reuses) the pinned iTerm2-Color-Schemes tag, reads
// vhs/*.json, and overlays extra colors from matching yaml/*.yml files.
func loadSchemes(ctx context.Context, tag string) ([]Tint, error) {
	root, err := schemesRepo(ctx, tag)
	if err != nil {
		return nil, err
	}
	tints, err := loadTints(filepath.Join(root, "vhs"))
	if err != nil {
		return nil, err
	}
	if yamlErr := applyYAML(tints, filepath.Join(root, "yaml")); yamlErr != nil {
		return nil, yamlErr
	}
	return tints, nil
}

func schemesRepo(ctx context.Context, tag string) (string, error) {
	if !schemesTagPattern.MatchString(tag) {
		return "", fmt.Errorf("invalid schemes tag %q", tag)
	}

	dir := schemesCacheDir()
	if isGitRepo(dir) {
		match, err := headMatchesTag(ctx, dir, tag)
		if err != nil {
			return "", err
		}
		if match {
			logger.InfoContext(ctx, "cached schemes repo already at tag", "dir", dir, "tag", tag)
		} else {
			logger.InfoContext(ctx, "updating cached schemes repo", "dir", dir, "tag", tag)
			if fetchErr := fetchTag(ctx, dir, tag); fetchErr != nil {
				return "", fetchErr
			}
		}
	} else {
		logger.InfoContext(ctx, "cloning schemes repo", "dir", dir, "tag", tag)
		if err := cloneTag(ctx, dir, tag); err != nil {
			return "", err
		}
	}

	if err := ensureSparse(ctx, dir); err != nil {
		return "", err
	}
	return dir, nil
}

func cloneTag(ctx context.Context, dir, tag string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("remove stale schemes cache: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(dir), 0o750); err != nil {
		return fmt.Errorf("create schemes cache parent: %w", err)
	}

	if _, err := runGit(ctx, "", "clone", "--filter=blob:none", "--sparse", "--depth", "1", "--branch", tag, schemesRepoURL, dir); err != nil {
		if rerr := os.RemoveAll(dir); rerr != nil {
			err = errors.Join(err, fmt.Errorf("remove partial clone: %w", rerr))
		}
		return err
	}
	return ensureLocalTag(ctx, dir, tag)
}

func fetchTag(ctx context.Context, dir, tag string) error {
	ref := "refs/tags/" + tag
	if _, err := runGit(ctx, dir, "fetch", "--depth", "1", "origin", ref+":"+ref); err != nil {
		return err
	}
	_, err := runGit(ctx, dir, "checkout", "--force", tag)
	return err
}

func ensureSparse(ctx context.Context, dir string) error {
	_, err := runGit(ctx, dir, "sparse-checkout", "set", "vhs", "yaml")
	return err
}

// ensureLocalTag records the requested tag locally when a shallow clone checked
// out the commit without creating refs/tags/<tag>. Later runs compare HEAD to
// that ref and skip fetching when it already matches.
func ensureLocalTag(ctx context.Context, dir, tag string) error {
	if _, err := runGit(ctx, dir, "rev-parse", "--verify", "--quiet", "refs/tags/"+tag+"^{commit}"); err == nil {
		return nil
	}
	if _, err := runGit(ctx, dir, "tag", "--force", tag, "HEAD"); err != nil {
		return fmt.Errorf("pin schemes tag: %w", err)
	}
	return nil
}

func headMatchesTag(ctx context.Context, dir, tag string) (bool, error) {
	head, err := gitOutput(ctx, dir, "rev-parse", "HEAD")
	if err != nil {
		return false, fmt.Errorf("read cached schemes HEAD: %w", err)
	}

	tagged, ok := verifyRev(ctx, dir, "refs/tags/"+tag+"^{commit}")
	if !ok {
		return false, nil
	}
	return head == tagged, nil
}

// verifyRev reports the commit for rev. ok is false when rev does not exist.
func verifyRev(ctx context.Context, dir, rev string) (string, bool) {
	out, err := runGit(ctx, dir, "rev-parse", "--verify", "--quiet", rev)
	if err != nil {
		return "", false
	}
	return strings.TrimSpace(string(out)), true
}

func isGitRepo(dir string) bool {
	info, err := os.Stat(filepath.Join(dir, ".git"))
	return err == nil && info.IsDir()
}

func loadTints(dir string) ([]Tint, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read schemes dir: %w", err)
	}

	tints := make([]Tint, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		tint, ok, readErr := readTint(filepath.Join(dir, entry.Name()))
		if readErr != nil {
			return nil, readErr
		}
		if ok {
			tints = append(tints, tint)
		}
	}
	if len(tints) == 0 {
		return nil, fmt.Errorf("no schemes found in %s", dir)
	}
	return tints, nil
}

func readTint(path string) (Tint, bool, error) {
	b, err := os.ReadFile(path) //nolint:gosec // Path is a file inside the cloned schemes checkout.
	if err != nil {
		return Tint{}, false, fmt.Errorf("read scheme: %w", err)
	}

	var tint Tint
	if err = json.Unmarshal(b, &tint); err != nil {
		return Tint{}, false, fmt.Errorf("parse scheme %s: %w", filepath.Base(path), err)
	}
	if err = val.Struct(tint); err != nil {
		logger.Error("invalid tint", "file", filepath.Base(path), "error", err) //nolint:all
		return Tint{}, false, nil
	}
	return tint, true, nil
}

func runGit(ctx context.Context, dir string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "git", args...) //nolint:gosec // Args are fixed git subcommands plus a validated tag.
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Env = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")

	out, err := cmd.CombinedOutput()
	if err != nil {
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			return nil, fmt.Errorf("git %s: %w", strings.Join(args, " "), err)
		}
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, msg)
	}
	return out, nil
}

func gitOutput(ctx context.Context, dir string, args ...string) (string, error) {
	out, err := runGit(ctx, dir, args...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
