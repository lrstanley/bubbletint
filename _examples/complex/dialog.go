// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lrstanley/bubbletint/roles/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

type dialog struct {
	// Core state.
	id       string
	active   string
	question string

	// Styles.
	dialogBoxStyle    lipgloss.Style
	buttonStyle       lipgloss.Style
	activeButtonStyle lipgloss.Style
}

func newDialog(question string) *dialog {
	m := &dialog{
		id:       zone.NewPrefix(),
		question: question,
		active:   "confirm",
	}
	m.setStyles()
	return m
}

func (m *dialog) setStyles() {
	m.dialogBoxStyle = f.Text(roles.Fg).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(f.Color(roles.Accent)).
		Padding(1, 0)

	// A neutral button, a step off the page.
	m.buttonStyle = f.Fill(roles.Control).
		Padding(0, 3).
		MarginTop(1).
		MarginRight(2)

	// The accent as a solid fill, with the text that reads on it.
	m.activeButtonStyle = f.Fill(roles.Accent).
		Padding(0, 3).
		MarginTop(1).
		MarginRight(2).
		Underline(true)
}

func (m *dialog) Init() tea.Cmd {
	return nil
}

func (m *dialog) GetHeight() int {
	return lipgloss.Height(m.View())
}

func (m *dialog) Update(msg tea.Msg) tea.Cmd { //nolint:unparam
	switch msg := msg.(type) {
	case ThemeChangedMsg:
		m.setStyles()
	case tea.MouseReleaseMsg:
		if msg.Button != tea.MouseLeft {
			return nil
		}
		if zone.Get(m.id + "confirm").InBounds(msg) {
			m.active = "confirm"
		} else if zone.Get(m.id + "cancel").InBounds(msg) {
			m.active = "cancel"
		}
	}
	return nil
}

func (m *dialog) View() string {
	var okButton, cancelButton string

	if m.active == "confirm" {
		okButton = m.activeButtonStyle.Render("Yes")
		cancelButton = m.buttonStyle.Render("Maybe")
	} else {
		okButton = m.buttonStyle.Render("Yes")
		cancelButton = m.activeButtonStyle.Render("Maybe")
	}

	question := lipgloss.NewStyle().Width(27).Align(lipgloss.Center).Render(m.question)
	buttons := lipgloss.JoinHorizontal(
		lipgloss.Top,
		zone.Mark(m.id+"confirm", okButton),
		zone.Mark(m.id+"cancel", cancelButton),
	)
	return m.dialogBoxStyle.Render(lipgloss.JoinVertical(lipgloss.Center, question, buttons))
}
