// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"

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
	help      string
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

func (m *model) resize() {
	if m.width <= 0 || m.height <= 0 {
		return
	}

	m.help = m.helpStyle.Width(m.width).Render("←/→: " + m.theme.Current().DisplayName + " • ↑/↓: scroll • q: quit")
	m.viewport.SetWidth(m.width)
	m.viewport.SetHeight(max(1, m.height-max(1, lipgloss.Height(m.help))))
	m.setContent()
}

func (m *model) contentWidth() int {
	width := m.viewport.Width()
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
	m.viewport.SetContent(strings.ReplaceAll(out, "\t", "    "))
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
		m.resize()
		return m, nil
	case tea.ColorProfileMsg:
		next := formatterFor(msg.Profile)
		if next == m.formatter {
			return m, nil
		}
		m.formatter = next
		m.resize()
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "left":
			m.theme.PreviousTint()
			m.setStyles()
			m.resize()
			return m, nil
		case "right":
			m.theme.NextTint()
			m.setStyles()
			m.resize()
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
	view.SetContent(m.viewport.View() + "\n" + m.help)
	return view
}

func main() {
	_, err := tea.NewProgram(newModel()).Run()
	if err != nil {
		fmt.Println("Bummer, there's been an error:", err) //nolint:forbidigo
		os.Exit(1)
	}
}
