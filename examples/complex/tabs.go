// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
	zone "github.com/lrstanley/bubblezone"
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
)

type tabs struct {
	id     string
	height int
	width  int

	active string
	items  []string
}

func (m tabs) Init() tea.Cmd {
	return nil
}

func (m tabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.MouseMsg:
		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
			return m, nil
		}

		for _, item := range m.items {
			// Check each item to see if it's in bounds.
			if zone.Get(m.id + item).InBounds(msg) {
				m.active = item
				break
			}
		}

		return m, nil
	}
	return m, nil
}

func (m tabs) View() string {
	tab := lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(tint.BrightPurple()).
		Padding(0, 1)

	activeTab := tab.Border(activeTabBorder, true)

	tabGap := tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	out := []string{}

	for _, item := range m.items {
		// Make sure to mark each tab when rendering.
		if item == m.active {
			out = append(out, zone.Mark(m.id+item, activeTab.Render(item)))
		} else {
			out = append(out, zone.Mark(m.id+item, tab.Render(item)))
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, out...)
	gap := tabGap.Render(strings.Repeat(" ", max(0, m.width-lipgloss.Width(row)-2)))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
	return row
}
