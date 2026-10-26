// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lrstanley/bubbletint/sources/v2/terminal"
	tint "github.com/lrstanley/bubbletint/v2"
)

// tintEventMsg is sent to the program for every change to the registry.
type tintEventMsg tint.Event

type model struct {
	registry *tint.Registry
	status   string // Result of attaching the terminal's config.
	last     string // Last registry event.
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "left":
			m.registry.PreviousTint()
		case "right":
			m.registry.NextTint()
		}
	case statusMsg:
		m.status = string(msg)
	case tintEventMsg:
		id := "<none>"
		if msg.Tint != nil {
			id = msg.Tint.ID
		}
		m.last = fmt.Sprintf("%s %s", msg.Kind, id)
		if msg.Source != "" {
			m.last += " (from " + msg.Source + ")"
		}

		// Nothing else to do: styles are derived from the current tint on every
		// render, so receiving the message is what triggers the redraw. An app
		// that caches styles would rebuild them here when msg.Kind is
		// tint.CurrentChanged.
	}
	return m, nil
}

func (m model) View() tea.View {
	t := m.registry.Current()

	base := lipgloss.NewStyle().Background(t.Bg).Foreground(t.Fg)
	title := base.Foreground(t.Blue).Bold(true).Render(fmt.Sprintf("Current tint: %s (%s)", t.DisplayName, t.ID))

	var swatches []string
	for _, c := range []*tint.Color{t.Red, t.Green, t.Yellow, t.Blue, t.Purple, t.Cyan} {
		swatches = append(swatches, lipgloss.NewStyle().Background(c).Render("   "))
	}

	content := strings.Join([]string{
		title,
		"",
		strings.Join(swatches, " "),
		"",
		base.Foreground(t.Green).Render(m.status),
		base.Foreground(t.Yellow).Render("Last event: " + m.last),
		"",
		base.Foreground(t.BrightBlack).Render("Edit your terminal's theme and watch this update. ←/→ switch tint, q quits."),
	}, "\n")

	var view tea.View
	view.AltScreen = true
	view.BackgroundColor = t.Bg
	view.SetContent(base.Padding(1, 2).Render(content))
	return view
}

func main() {
	// The registry has all built-in tints. The terminal's own theme is added to
	// it, and followed, below.
	registry := tint.NewRegistry(tint.TintDraculaPlus)

	m := model{registry: registry, status: "Detecting terminal..."}
	p := tea.NewProgram(m)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Attaching happens off the main goroutine, as sending to the program only
	// works once it's running. The subscription turns registry events into
	// messages.
	go func() {
		registry.Subscribe(func(ev tint.Event) { p.Send(tintEventMsg(ev)) })

		src, err := terminal.Detect()
		if err != nil {
			p.Send(statusMsg(fmt.Sprintf(
				"%v, using the built-in theme. (TERM=%q TERM_PROGRAM=%q WT_SESSION=%q)",
				err, os.Getenv("TERM"), os.Getenv("TERM_PROGRAM"), os.Getenv("WT_SESSION"),
			)))
			return
		}

		// Attach follows the terminal's theme by default: its tint becomes the
		// current one, and stays up to date as the terminal's config changes. Use
		// tint.WithFollow(false) to only register it, as "<terminal>:current".
		if _, err = registry.Attach(ctx, src); err != nil {
			status := "Couldn't read the terminal's theme: " + err.Error()
			if errors.Is(err, tint.ErrNotFound) {
				status = "No config found for the terminal, using the built-in theme."
			}
			p.Send(statusMsg(status))
			return
		}
		p.Send(statusMsg("Following the theme of " + src.Name() + "."))
	}()

	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type statusMsg string
