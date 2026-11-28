// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"log"
	"time"

	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lrstanley/bubbletint/roles/v2"
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

	m := model{
		styles:    createStyles(theme, colorprofile.TrueColor),
		textarea:  ta,
		startTime: time.Now(),
	}
	styleTextarea(&m.textarea, m.styles.f)
	return m
}

// styleTextarea styles a textarea from the same roles as the rest.
func styleTextarea(ta *textarea.Model, f *roles.Fitted) {
	state := textarea.StyleState{
		Base:             f.Fill(roles.Bg),
		Text:             f.Text(roles.Fg),
		LineNumber:       f.Text(roles.Subtle),
		CursorLineNumber: f.On(roles.Bg.Hover()).Span(roles.Muted),
		CursorLine:       f.Fill(roles.Bg.Hover()),
		EndOfBuffer:      f.Text(roles.Subtle),
		Placeholder:      f.Text(roles.Subtle),
		Prompt:           f.Text(roles.Accent),
		Selection:        f.Fill(roles.Bg.Selected()),
	}
	s := ta.Styles()
	s.Focused, s.Blurred = state, state
	s.Cursor.Color = f.Color(roles.Accent)
	ta.SetStyles(s)
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.ColorProfileMsg:
		m.styles = createStyles(m.styles.tint, msg.Profile)
		styleTextarea(&m.textarea, m.styles.f)
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

func (m model) View() tea.View {
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
	ta := m.styles.f.Fill(roles.Bg).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.styles.f.Color(roles.Accent)).
		BorderBackground(m.styles.f.Color(roles.Bg)).
		Render(m.textarea.View())

	elapsed := time.Since(m.startTime).Round(time.Second)
	info := m.styles.info.Width(ew).Render(
		fmt.Sprintf("Elapsed: %s | Last key: %s", elapsed, m.lastKey),
	)

	help := m.styles.help.Width(ew).Render(
		"Press 'q' or Ctrl+C to quit • Type in the textarea above",
	)

	content := lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		palette,
		ta,
		info,
		help,
	)

	var view tea.View
	view.AltScreen = true
	view.BackgroundColor = m.styles.f.Color(roles.Bg)
	view.SetContent(m.styles.app.Render(content))
	return view
}

// renderColorPalette shows the tint's hues as solid fills, each with the text
// that reads on it.
func (m model) renderColorPalette(w int) string {
	f := m.styles.f
	gap := f.Fill(roles.Bg)

	var blocks []string
	for _, tone := range []roles.Role{roles.Red, roles.Green, roles.Blue, roles.Yellow, roles.Magenta, roles.Cyan} {
		blocks = append(blocks, f.Fill(tone).
			Padding(0, 2).
			Margin(0, 1).
			MarginBackground(f.Color(roles.Bg)).
			Render(tone.String()))
	}

	return gap.Width(w).Render(lipgloss.JoinHorizontal(lipgloss.Top, blocks...))
}

func main() {
	// Load a custom theme from file.
	theme, err := loadTheme("custom-theme.json")
	if err != nil {
		log.Fatalf("Failed to load theme: %v", err)
	}

	p := tea.NewProgram(initialModel(theme))
	_, err = p.Run()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
