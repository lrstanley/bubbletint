// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/v2/textarea"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

type model struct {
	height    int
	width     int
	styles    styles
	textarea  textarea.Model
	lastKey   string
	startTime time.Time
}

func initialModel(theme *tint.Tint) model {
	ta := textarea.New()
	ta.Placeholder = "Type something here..."
	ta.Focus()

	return model{
		styles:    createStyles(theme),
		textarea:  ta,
		startTime: time.Now(),
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			m.lastKey = msg.String()
		}
	}

	var taCmd tea.Cmd
	m.textarea, taCmd = m.textarea.Update(msg)

	return m, taCmd
}

func (m model) View() string {
	app := m.styles.app.
		Height(m.height).
		Width(m.width)

	// Effective width of app.
	ew := app.GetWidth() - app.GetHorizontalPadding() - app.GetHorizontalBorderSize()

	header := m.styles.header.Width(ew).Render(
		fmt.Sprintf("🎨 Theme Demo: %s (%s)", m.styles.tint.DisplayName, m.styles.tint.ID),
	)

	palette := m.renderColorPalette(ew)

	m.textarea.SetWidth(ew - 2) // -2=border
	m.textarea.Styles.Focused.Base.Background(m.styles.tint.Bg)
	ta := m.textarea.View()
	ta = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.tint.Blue).
		Render(ta)

	elapsed := time.Since(m.startTime).Round(time.Second)
	info := m.styles.info.Width(ew).Render(
		fmt.Sprintf("⏱️  Elapsed: %s | ⌨️  Last key: %s", elapsed, m.lastKey),
	)

	help := m.styles.help.Width(ew).Render(
		"💡  Press 'q' or Ctrl+C to quit • Type in the textarea above",
	)

	content := lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		palette,
		ta,
		info,
		help,
	)

	return m.styles.app.Render(content)
}

func (m model) renderColorPalette(w int) string {
	colors := []struct {
		name  string
		color *tint.Color
	}{
		{"Red", m.styles.tint.Red},
		{"Green", m.styles.tint.Green},
		{"Blue", m.styles.tint.Blue},
		{"Yellow", m.styles.tint.Yellow},
		{"Purple", m.styles.tint.Purple},
		{"Cyan", m.styles.tint.Cyan},
	}

	var colorBlocks []string
	for _, c := range colors {
		if c.color != nil {
			block := lipgloss.NewStyle().
				Background(c.color).
				Foreground(m.styles.tint.Bg).
				Padding(0, 2).
				Margin(0, 1).
				MarginBackground(m.styles.tint.Bg).
				Render(c.name)
			colorBlocks = append(colorBlocks, block)
		}
	}

	return lipgloss.NewStyle().
		Background(m.styles.tint.Bg).
		Width(w).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, colorBlocks...))
}

func main() {
	// Load a custom theme from file.
	theme, err := loadTheme("custom-theme.json")
	if err != nil {
		log.Fatalf("Failed to load theme: %v", err)
	}

	p := tea.NewProgram(initialModel(theme), tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
