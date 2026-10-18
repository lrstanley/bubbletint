// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/glamour/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lrstanley/bubbletint/glamourtint/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

const sampleMD = `# Bubbletint

Map a tint onto a [Glamour](https://github.com/charmbracelet/glamour) stylesheet.

## Headings

H1 is a banner. Later levels keep the heading color and use prefixes.

### Lists

- themes from iTerm2 color schemes
- inline ` + "`code`" + ` sits on a chip
- [x] fenced blocks follow the tint
- [ ] a task that is still open

> Quotes use the tint foreground.

---

` + "```go\nfunc main() {\n\tfmt.Println(\"hello\")\n}\n```\n"

type model struct { //nolint:recvcheck
	height    int
	width     int
	theme     *tint.Registry
	viewport  viewport.Model
	helpStyle lipgloss.Style
	formatter string
}

func newModel() *model {
	m := &model{
		viewport:  viewport.New(),
		formatter: "terminal256",
		theme: tint.NewRegistry(
			tint.TintDraculaPlus,
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

func (m *model) contentWidth() int {
	width := m.width
	if width == 0 {
		return 80
	}
	frame := m.viewport.Style.GetHorizontalFrameSize()
	if width > frame {
		return width - frame
	}
	return 1
}

func (m *model) setContent() {
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(glamourtint.StyleConfig(m.theme.Current())),
		glamour.WithWordWrap(m.contentWidth()),
		glamour.WithChromaFormatter(m.formatter),
	)
	if err != nil {
		m.viewport.SetContent(err.Error())
		return
	}

	out, err := renderer.Render(sampleMD)
	if err != nil {
		m.viewport.SetContent(err.Error())
		return
	}
	m.viewport.SetContent(out)
}

func formatterFor(profile colorprofile.Profile) string {
	switch profile {
	case colorprofile.TrueColor:
		return "terminal16m"
	case colorprofile.ANSI256, colorprofile.ANSI:
		return "terminal256"
	case colorprofile.Unknown, colorprofile.NoTTY, colorprofile.ASCII:
		return "terminal"
	default:
		return "terminal"
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.viewport.SetWidth(msg.Width)
		m.viewport.SetHeight(msg.Height - 1)
		m.setContent()
		return m, nil
	case tea.ColorProfileMsg:
		next := formatterFor(msg.Profile)
		if next == m.formatter {
			return m, nil
		}
		m.formatter = next
		m.setContent()
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "left":
			m.theme.PreviousTint()
			m.setStyles()
			m.setContent()
			return m, nil
		case "right":
			m.theme.NextTint()
			m.setStyles()
			m.setContent()
			return m, nil
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	default:
		return m, nil
	}
}

func (m model) View() tea.View {
	var view tea.View
	view.AltScreen = true
	view.BackgroundColor = m.theme.Current().Bg
	view.ForegroundColor = m.theme.Current().Fg
	view.SetContent(m.viewport.View() + m.helpView())
	return view
}

func (m model) helpView() string {
	return m.helpStyle.Render("\n ←/→: Change Tint (current: " + m.theme.Current().DisplayName + ") • ↑/↓: Navigate • q: Quit\n")
}

func main() {
	_, err := tea.NewProgram(newModel()).Run()
	if err != nil {
		fmt.Println("Bummer, there's been an error:", err) //nolint:forbidigo
		os.Exit(1)
	}
}
