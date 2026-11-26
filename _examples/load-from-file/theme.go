// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lrstanley/bubbletint/roles/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// loadTheme loads a theme from a JSON file.
func loadTheme(filename string) (*tint.Tint, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read theme file: %w", err)
	}

	theme := &tint.Tint{}
	err = json.Unmarshal(data, &theme)
	if err != nil {
		return nil, fmt.Errorf("failed to parse theme file: %w", err)
	}

	return theme, nil
}

// styles are the example's styles, from semantic roles derived from the
// loaded tint and fitted to the terminal's color depth.
type styles struct {
	tint   *tint.Tint
	f      *roles.Fitted
	app    lipgloss.Style
	header lipgloss.Style
	info   lipgloss.Style
	help   lipgloss.Style
}

func createStyles(theme *tint.Tint, profile colorprofile.Profile) styles {
	f := roles.Derive(theme).Fit(roles.Target{Profile: profile})
	return styles{
		tint: theme,
		f:    f,
		app: f.Fill(roles.Bg).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(f.Color(roles.BorderStrong)).
			BorderBackground(f.Color(roles.Bg)),
		header: f.Text(roles.Accent).Bold(true).MarginBottom(1),
		info:   f.Text(roles.Muted),
		help:   f.Text(roles.Subtle).Italic(true),
	}
}
