// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/charmbracelet/bubbles/v2/viewport"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/lrstanley/bubbletint/chromatint/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

const viewportContent = `func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.viewport.SetWidth(msg.Width)
		m.viewport.SetHeight(msg.Height - 1)
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "left":
			m.theme.PreviousTint()
			m.setStyles()
			m.setContent()
			return m, tea.SetBackgroundColor(m.theme.Current().Bg)
		case "right":
			m.theme.NextTint()
			m.setStyles()
			m.setContent()
			return m, tea.SetBackgroundColor(m.theme.Current().Bg)
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	default:
		return m, nil
	}
}`

type model struct { //nolint:recvcheck
	// Core state.
	height int
	width  int
	theme  *tint.Registry

	// Child components.
	viewport viewport.Model

	// Styles.
	helpStyle lipgloss.Style
}

func newModel() *model {
	m := &model{
		viewport: viewport.New(),
		theme: tint.NewRegistry(
			tint.TintGrape,
			tint.DefaultDarkTints()...,
		),
	}

	m.setStyles()
	m.setContent()
	return m
}

func (m *model) setStyles() {
	m.viewport.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(m.theme.Current().BrightPurple).
		PaddingRight(2)

	m.helpStyle = lipgloss.NewStyle().
		Foreground(m.theme.Current().Fg).
		Faint(true)
}

func (m *model) setContent() {
	lexer := lexers.Get("go")

	style := chroma.MustNewStyle("app", chromatint.StyleEntry(m.theme.Current(), false))

	formatter := formatters.TTY256 // Adjust this based off what the terminal supports.

	iterator, err := lexer.Tokenise(nil, viewportContent)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	err = formatter.Format(&buf, style, iterator)
	if err != nil {
		return
	}
	m.viewport.SetContent(buf.String())
}

func (m model) Init() tea.Cmd {
	return tea.SetBackgroundColor(m.theme.Current().Bg)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.viewport.SetWidth(msg.Width)
		m.viewport.SetHeight(msg.Height - 1) // -1=help.
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "left":
			m.theme.PreviousTint()
			m.setStyles()
			m.setContent()
			return m, tea.SetBackgroundColor(m.theme.Current().Bg)
		case "right":
			m.theme.NextTint()
			m.setStyles()
			m.setContent()
			return m, tea.SetBackgroundColor(m.theme.Current().Bg)
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	default:
		return m, nil
	}
}

func (m model) View() string {
	return m.viewport.View() + m.helpView()
}

func (m model) helpView() string {
	return m.helpStyle.Render("\n ←/→: Change Tint (current: " + m.theme.Current().DisplayName + ") • ↑/↓: Navigate • q: Quit\n")
}

func main() {
	_, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Println("Bummer, there's been an error:", err) //nolint:forbidigo
		os.Exit(1)
	}
}
