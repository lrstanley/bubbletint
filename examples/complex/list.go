// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
	zone "github.com/lrstanley/bubblezone"
)

type listItem struct {
	name string
	done bool
}

type list struct {
	id     string
	height int
	width  int

	title string
	items []listItem
}

func (m list) Init() tea.Cmd {
	return nil
}

func (m list) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.MouseMsg:
		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
			return m, nil
		}

		for i, item := range m.items {
			// Check each item to see if it's in bounds.
			if zone.Get(m.id + item.name).InBounds(msg) {
				m.items[i].done = !m.items[i].done
				break
			}
		}

		return m, nil
	}
	return m, nil
}

func (m list) View() string {
	listStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(tint.Fg()).
		MarginRight(2)

	listHeader := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(tint.Fg()).
		MarginRight(2).
		Render

	listItemStyle := lipgloss.NewStyle().PaddingLeft(2).Render

	checkMark := lipgloss.NewStyle().SetString("✓").
		Foreground(tint.BrightGreen()).
		PaddingRight(1).
		String()

	listDoneStyle := func(s string) string {
		return checkMark + lipgloss.NewStyle().
			Strikethrough(true).
			Foreground(tint.BrightBlack()).
			Render(s)
	}

	out := []string{listHeader(m.title)}

	for _, item := range m.items {
		if item.done {
			out = append(out, zone.Mark(m.id+item.name, listDoneStyle(item.name)))
			continue
		}

		out = append(out, zone.Mark(m.id+item.name, listItemStyle(item.name)))
	}

	return listStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, out...),
	)
}
