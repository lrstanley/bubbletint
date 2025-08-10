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

type listItem struct {
	name string
	done bool
}

type list struct {
	// Core state.
	id    string
	title string
	items []listItem

	// Styles.
	listStyle       lipgloss.Style
	listHeaderStyle lipgloss.Style
	listItemStyle   lipgloss.Style
	checkMarkStyle  lipgloss.Style
	listDoneStyle   lipgloss.Style
}

func newList(title string, items ...listItem) *list {
	m := &list{
		id:    zone.NewPrefix(),
		title: title,
		items: items,
	}
	m.setStyles()
	return m
}

func (m *list) setStyles() {
	m.listStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(tint.Current().Fg).
		Foreground(tint.Current().Fg).
		MarginRight(2)

	m.listHeaderStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(tint.Current().Fg).
		Foreground(tint.Current().Fg).
		MarginRight(2)

	m.listItemStyle = lipgloss.NewStyle().
		Foreground(tint.Current().Fg).
		PaddingLeft(2)

	m.checkMarkStyle = lipgloss.NewStyle().
		Foreground(tint.Current().BrightGreen).
		SetString("✓").
		PaddingRight(1)

	m.listDoneStyle = lipgloss.NewStyle().
		Strikethrough(true).
		Foreground(tint.Current().BrightGreen)
}

func (m *list) Init() tea.Cmd {
	return nil
}

func (m *list) GetHeight() int {
	return lipgloss.Height(m.View())
}

func (m *list) Update(msg tea.Msg) tea.Cmd { //nolint:unparam
	switch msg := msg.(type) {
	case ThemeChangedMsg:
		m.setStyles()
	case tea.MouseReleaseMsg:
		if msg.Button != tea.MouseLeft {
			return nil
		}
		for i, item := range m.items {
			// Check each item to see if it's in bounds.
			if zone.Get(m.id + item.name).InBounds(msg) {
				m.items[i].done = !m.items[i].done
				break
			}
		}
	}
	return nil
}

func (m *list) View() string {
	out := []string{m.listHeaderStyle.Render(m.title)}

	for _, item := range m.items {
		if item.done {
			out = append(out, zone.Mark(
				m.id+item.name,
				m.checkMarkStyle.String()+
					m.listDoneStyle.Render(item.name),
			))
			continue
		}

		out = append(out, zone.Mark(m.id+item.name, m.listItemStyle.Render(item.name)))
	}

	return m.listStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, out...),
	)
}
