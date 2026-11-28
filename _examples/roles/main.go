// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// This example styles a small deploy dashboard from semantic roles instead of
// a tint's raw colors: a top bar and sidebar on surfaces, a list with a
// selected row, a dialog with an input and buttons, and status pills. Every
// color is derived from the current tint, so it holds up on dark, light and
// single-hue tints alike, and is fitted to the terminal's color depth.
//
// Keys: up/down select a service, left/right switch tint, d cycles the color
// depth (to see how 256, 16 and no colors are handled), q quits.
package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lrstanley/bubbletint/roles/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

type service struct {
	name, version, age string
	status             roles.Role // a tone: Success, Warning, Danger or Info
	label              string
}

var services = []service{
	{name: "api", version: "v2.14.0", age: "3m", status: roles.Success, label: "live"},
	{name: "web", version: "v2.14.0", age: "3m", status: roles.Danger, label: "failed"},
	{name: "worker", version: "v2.13.2", age: "1h", status: roles.Warning, label: "degraded"},
	{name: "scheduler", version: "v2.13.2", age: "1h", status: roles.Success, label: "live"},
	{name: "search", version: "v2.12.0", age: "2d", status: roles.Info, label: "rolling"},
}

type model struct {
	width, height int
	selected      int
	detected      colorprofile.Profile // the terminal's colors
	profile       colorprofile.Profile // the colors drawn with: detected, or fewer with d

	f *roles.Fitted // the roles, fitted to profile
}

// restyle derives the current tint's roles and fits them to the terminal. It
// runs when the tint or profile changes, not on every render.
func (m *model) restyle() {
	m.f = roles.Derive(tint.Current()).Fit(roles.Target{Profile: m.profile})
}

func (m *model) Init() tea.Cmd { return nil }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.ColorProfileMsg:
		m.detected = msg.Profile
		m.profile = m.detected
		m.restyle()
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "k":
			m.selected = (m.selected + len(services) - 1) % len(services)
		case "down", "j":
			m.selected = (m.selected + 1) % len(services)
		case "left", "h":
			tint.PreviousTint()
			m.restyle()
		case "right", "l":
			tint.NextTint()
			m.restyle()
		case "d":
			// Fewer colors each time, then back to the terminal's.
			if m.profile == colorprofile.Unknown {
				m.profile = colorprofile.TrueColor
			}
			m.profile--
			if m.profile < colorprofile.ASCII {
				m.profile = m.detected
			}
			m.restyle()
		}
	}
	return m, nil
}

func (m *model) View() tea.View {
	var v tea.View
	v.AltScreen = true
	if m.f == nil || m.width == 0 {
		return v
	}
	f := m.f
	v.BackgroundColor = f.Color(roles.Bg)
	page := f.Text(roles.Fg)

	// Top bar, on a surface. Text on it comes from f.On(roles.Surface), which
	// keeps it readable there. Each piece is rendered on its own and joined
	// onto the bar, so it's a Span, which draws on the surface's color too;
	// Text would leave a hole in the bar after each piece.
	t := tint.Current()
	surface := f.On(roles.Surface)
	left := surface.Span(roles.Accent).Bold(true).Render("deploys") + surface.Span(roles.Muted).Render(" / production")
	right := surface.Span(roles.Muted).Render(fmt.Sprintf("%s · %s", t.DisplayName, profileLabel(m.profile)))
	gap := max(1, m.width-lipgloss.Width(left)-lipgloss.Width(right)-2)
	bar := f.Fill(roles.Surface).Width(m.width).Padding(0, 1).
		Render(left + strings.Repeat(" ", gap) + right)

	// Sidebar, on the same surface, with a border between it and the page.
	side := []string{surface.Span(roles.Subtle).Render("ENVIRONMENTS"), ""}
	for i, env := range []string{"production", "staging", "preview"} {
		if i == 0 {
			side = append(side, surface.Span(roles.Accent).Bold(true).Render("● "+env))
			continue
		}
		side = append(side, surface.Span(roles.Muted).Render("  "+env))
	}
	// Everything below the bar. Height includes the padding.
	body := max(len(services)+12, m.height-1)
	sidebar := f.Fill(roles.Surface).Width(18).Height(body).Padding(1, 1).
		BorderStyle(lipgloss.NormalBorder()).BorderRight(true).
		BorderForeground(f.Color(roles.Border)).BorderBackground(f.Color(roles.Surface)).
		Render(strings.Join(side, "\n"))

	// The list. The selected row is the page's background, selected, and its
	// text comes from f.On of it, so hints and status stay readable on the
	// highlight.
	width := max(40, m.width-sidebarWidth-4)
	rows := []string{f.Text(roles.Subtle).Render(fmt.Sprintf("  %-11s %-9s %-10s %s", "SERVICE", "VERSION", "STATUS", "AGE"))}
	selected := roles.Bg.Selected()
	for i, s := range services {
		rf, row := f, lipgloss.NewStyle()
		if i == m.selected {
			rf, row = f.On(selected), f.Fill(selected)
		}
		cell := func(r roles.Role, text string) string {
			style := rf.Span(r)
			if i == m.selected {
				style = style.Bold(true)
			}
			return style.Render(text)
		}
		cells := cell(roles.Fg, fmt.Sprintf("  %-11s ", s.name)) +
			cell(roles.Muted, fmt.Sprintf("%-9s ", s.version)) +
			cell(s.status, fmt.Sprintf("● %-8s ", s.label)) +
			cell(roles.Subtle, s.age)
		rows = append(rows, row.Width(width).Render(cells))
	}

	// A dialog on a raised surface, with an input sunk into it, a neutral
	// button that steps off the dialog, and the focused one, pressed.
	sel := services[m.selected]
	raised, onRaised := f.Fill(roles.Raised), f.On(roles.Raised)
	input := onRaised.Fill(roles.Sunken).Width(34).Padding(0, 1).Render(
		onRaised.On(roles.Sunken).Span(roles.Subtle).Render("type " + sel.name + " to confirm"))
	buttons := onRaised.Fill(roles.Control).Padding(0, 2).Render("Cancel") + raised.Render("  ") +
		onRaised.Fill(roles.Danger.Pressed()).Padding(0, 2).Render("Roll back")
	dialog := raised.Padding(1, 2).Width(40).Render(strings.Join([]string{
		onRaised.Span(roles.Fg).Bold(true).Render("Roll back " + sel.name + "?"),
		onRaised.Span(roles.Muted).Render("Redeploys v2.13.2 to production."),
		"",
		input,
		"",
		lipgloss.PlaceHorizontal(36, lipgloss.Right, buttons, lipgloss.WithWhitespaceStyle(raised)),
	}, "\n"))

	// Status pills: solid fills, and soft ones.
	var pills []string
	for _, tone := range []roles.Role{roles.Success, roles.Warning, roles.Danger, roles.Info, roles.Accent, roles.Neutral} {
		pills = append(pills, f.Fill(tone).Padding(0, 1).Render(tone.String()))
	}
	pills = append(pills, " ")
	for _, tone := range []roles.Role{roles.Success, roles.Warning, roles.Danger, roles.Info} {
		pills = append(pills, f.Fill(tone.Soft()).Padding(0, 1).Render(tone.String()))
	}

	help := f.Text(roles.Subtle).Render("↑/↓ select  ←/→ tint  d color depth  q quit")
	main := lipgloss.NewStyle().Padding(1, 2).Render(strings.Join([]string{
		strings.Join(rows, "\n"),
		"",
		dialog,
		"",
		strings.Join(pills, ""),
		"",
		help,
	}, "\n"))

	v.SetContent(page.Render(bar + "\n" + lipgloss.JoinHorizontal(lipgloss.Top, sidebar, main)))
	return v
}

const sidebarWidth = 21

func profileLabel(p colorprofile.Profile) string {
	switch p { //nolint:exhaustive // the rest draw no color
	case colorprofile.TrueColor, colorprofile.Unknown:
		return "true color"
	case colorprofile.ANSI256:
		return "256 colors"
	case colorprofile.ANSI:
		return "16 colors"
	}
	return "no color"
}

func main() {
	// The default registry has every built-in tint, starting at Tokyo Night.
	tint.NewDefaultRegistry()
	tint.SetTint(tint.TintTokyoNight)
	m := &model{}
	m.restyle()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
