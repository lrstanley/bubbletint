// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	tint "github.com/lrstanley/bubbletint/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

type history struct {
	// Core state.
	id     string
	height int
	width  int
	active string
	items  []string

	// Styles.
	historyStyle       lipgloss.Style
	activeHistoryStyle lipgloss.Style
}

func newHistory(items ...string) *history {
	m := &history{
		id:    zone.NewPrefix(),
		items: items,
	}
	m.setStyles()
	return m
}

func (m *history) setStyles() {
	m.historyStyle = lipgloss.NewStyle().
		Align(lipgloss.Left).
		Background(adaptBright(tint.Current().Bg, 0.15)).
		Foreground(adaptBright(tint.Current().Fg, 0.15)).
		Margin(0, 1).
		Padding(1, 2)

	m.activeHistoryStyle = m.historyStyle.
		Background(adaptBright(tint.Current().Bg, 0.35)).
		Foreground(adaptBright(tint.Current().Fg, 0.25))
}

func (m *history) Init() tea.Cmd {
	return nil
}

func (m *history) Update(msg tea.Msg) tea.Cmd { //nolint:unparam
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case ThemeChangedMsg:
		m.setStyles()
	case tea.MouseReleaseMsg:
		if msg.Button != tea.MouseLeft {
			return nil
		}
		for _, item := range m.items {
			// Check each item to see if it's in bounds.
			if zone.Get(m.id + item).InBounds(msg) {
				m.active = item
				break
			}
		}
	}
	return nil
}

func (m *history) View() string {
	history := m.historyStyle.
		Width((m.width / len(m.items)) - 2).
		Height(m.height).
		MaxHeight(m.height)

	active := m.activeHistoryStyle.
		Width((m.width / len(m.items)) - 2).
		Height(m.height).
		MaxHeight(m.height)

	out := []string{}

	for _, item := range m.items {
		if item == m.active {
			// Customize the active item.
			out = append(out, zone.Mark(
				m.id+item,
				active.Render(item),
			))
		} else {
			// Make sure to mark all zones.
			out = append(out, zone.Mark(m.id+item, history.Render(item)))
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, out...)
}
