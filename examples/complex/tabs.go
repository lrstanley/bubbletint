// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	tint "github.com/lrstanley/bubbletint/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		Padding(0, 1)
)

type tabs struct {
	// Core state.
	id     string
	width  int
	active string
	items  []string

	// Styles.
	inactiveStyle lipgloss.Style
	activeStyle   lipgloss.Style
	tabGapStyle   lipgloss.Style
}

func newTabs(items ...string) *tabs {
	m := &tabs{
		id:     zone.NewPrefix(),
		items:  items,
		active: items[0],
	}
	m.setStyles()
	return m
}

func (m *tabs) setStyles() {
	m.inactiveStyle = lipgloss.NewStyle().
		Border(tabBorder, true).
		Padding(0, 1).
		BorderForeground(tint.Current().BrightPurple).
		Foreground(tint.Current().Fg)

	m.activeStyle = m.inactiveStyle.
		Border(activeTabBorder, true).
		Bold(true).
		Foreground(adaptBright(tint.Current().BrightPurple, 0.25))

	m.tabGapStyle = tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false).
		BorderForeground(tint.Current().BrightPurple)
}

func (m *tabs) Init() tea.Cmd {
	return nil
}

func (m *tabs) GetHeight() int {
	return lipgloss.Height(m.View())
}

func (m *tabs) Update(msg tea.Msg) tea.Cmd { //nolint:unparam
	switch msg := msg.(type) {
	case ThemeChangedMsg:
		m.setStyles()
	case tea.WindowSizeMsg:
		m.width = msg.Width
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

func (m *tabs) View() string {
	out := []string{}

	for _, item := range m.items {
		// Make sure to mark each tab when rendering.
		if item == m.active {
			out = append(out, zone.Mark(m.id+item, m.activeStyle.Render(item)))
		} else {
			out = append(out, zone.Mark(m.id+item, m.inactiveStyle.Render(item)))
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, out...)
	return lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		row,
		m.tabGapStyle.Render(strings.Repeat(" ", max(0, m.width-lipgloss.Width(row)-2))),
	)
}
