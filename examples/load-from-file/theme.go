// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss/v2"
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

type styles struct {
	tint   *tint.Tint
	app    lipgloss.Style
	header lipgloss.Style
	info   lipgloss.Style
	help   lipgloss.Style
}

func createStyles(theme *tint.Tint) styles {
	return styles{
		tint: theme,
		app: lipgloss.NewStyle().
			Background(theme.Bg).
			Foreground(theme.Fg).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(theme.BrightGreen),
		header: lipgloss.NewStyle().
			Background(theme.Bg).
			Foreground(theme.Blue).
			Bold(true).
			MarginBottom(1),
		info: lipgloss.NewStyle().
			Background(theme.Bg).
			Foreground(theme.Green),
		help: lipgloss.NewStyle().
			Background(theme.Bg).
			Foreground(theme.Yellow).
			Italic(true),
	}
}
