// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	variantDark  = "dark"
	variantLight = "light"
)

// yamlExtra is the subset of a yaml/*.yml theme that the vhs JSON export does
// not include.
type yamlExtra struct {
	Name          string `yaml:"name"`
	Author        string `yaml:"author"`
	Variant       string `yaml:"variant"`
	SelectionText string `yaml:"selection_text"`
	Bold          string `yaml:"bold"`
	CursorText    string `yaml:"cursor_text"`
}

// applyYAML overlays author-provided variant and extra colors onto tints loaded
// from vhs JSON. Themes without a YAML file are left unchanged.
func applyYAML(tints []Tint, dir string) error {
	extras, err := loadYAMLExtras(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for i := range tints {
		extra, ok := extras[tints[i].Name]
		if !ok {
			continue
		}
		tints[i].Author = extra.Author
		tints[i].Variant = extra.Variant
		tints[i].SelectionText = extra.SelectionText
		tints[i].Bold = extra.Bold
		tints[i].CursorText = extra.CursorText
	}
	return nil
}

func loadYAMLExtras(dir string) (map[string]yamlExtra, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read schemes yaml dir: %w", err)
	}

	extras := make(map[string]yamlExtra, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".yml") {
			continue
		}

		extra, readErr := readYAMLExtra(filepath.Join(dir, entry.Name()))
		if readErr != nil {
			return nil, readErr
		}
		stem := strings.TrimSuffix(entry.Name(), ".yml")
		if extra.Name == "" {
			extra.Name = stem
		}
		extras[extra.Name] = extra
		if _, exists := extras[stem]; !exists {
			extras[stem] = extra
		}
	}
	return extras, nil
}

func readYAMLExtra(path string) (yamlExtra, error) {
	b, err := os.ReadFile(path) //nolint:gosec // Path is a file inside the cloned schemes checkout.
	if err != nil {
		return yamlExtra{}, fmt.Errorf("read scheme: %w", err)
	}

	var extra yamlExtra
	if err = yaml.Unmarshal(b, &extra); err != nil {
		return yamlExtra{}, fmt.Errorf("parse scheme %s: %w", filepath.Base(path), err)
	}
	extra.Name = strings.TrimSpace(extra.Name)
	extra.Author = strings.TrimSpace(extra.Author)
	extra.Variant = themeVariant(extra.Variant)
	extra.SelectionText = optionalHex(filepath.Base(path), "selection_text", extra.SelectionText)
	extra.Bold = optionalHex(filepath.Base(path), "bold", extra.Bold)
	extra.CursorText = optionalHex(filepath.Base(path), "cursor_text", extra.CursorText)
	return extra, nil
}

func themeVariant(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case variantDark:
		return variantDark
	case variantLight:
		return variantLight
	default:
		return ""
	}
}

func optionalHex(file, field, value string) string {
	value = normalizeHex(value)
	if value == "" {
		return ""
	}
	if _, err := colorful.Hex(value); err != nil {
		logger.Error("invalid theme color", "file", file, "field", field, "error", err) //nolint:all
		return ""
	}
	return value
}

// normalizeHex accepts the "#rrggbb" form used by most themes and the bare
// "rrggbb" form used by a few YAML files.
func normalizeHex(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || strings.HasPrefix(value, "#") {
		return value
	}
	if len(value) != 3 && len(value) != 6 {
		return value
	}
	for _, r := range value {
		switch {
		case r >= '0' && r <= '9', r >= 'a' && r <= 'f', r >= 'A' && r <= 'F':
		default:
			return value
		}
	}
	return "#" + value
}
