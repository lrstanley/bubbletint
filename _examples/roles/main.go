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

// control is a role the app adds for itself: a dialog's neutral button. It's
// derived alongside the core roles, keeps text readable, never shares a color
// with the dialog with 256 colors, and is reversed with 16 or none.
var control = roles.Define(roles.Spec{
	Name: "control",
	Kind: roles.KindSurface,
	Derive: func(p *roles.Palette) tint.Color {
		fg := p.Color(roles.Fg)
		return p.Color(roles.Raised).MixReadable(fg, fg, p.ByDark(0.22, 0.17), 4.5)
	},
	Distinct: []roles.Role{roles.Raised},
	Fallback: roles.Fallback{Attr: roles.Reverse},
})

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
	detected      roles.Depth // the terminal's depth
	depth         roles.Depth // the depth drawn at: detected, or fewer colors with d

	f   *roles.Fitted // the roles, fitted to depth
	row *roles.Fitted // the same, for text on the selected row
}

// restyle derives the current tint's roles and fits them to the terminal. It
// runs when the tint or depth changes, not on every render.
func (m *model) restyle() {
	m.f = roles.Derive(tint.Current()).Fit(roles.Target{Depth: m.depth})
	m.row = m.f.On(roles.Selection)
}

func (m *model) Init() tea.Cmd { return nil }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.ColorProfileMsg:
		m.detected = depthOf(msg.Profile)
		m.depth = m.detected
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
			m.depth++
			if m.depth > roles.NoColor {
				m.depth = m.detected
			}
			m.restyle()
		}
	}
	return m, nil
}

// depthOf maps Bubble Tea's detected color profile to a depth.
func depthOf(p colorprofile.Profile) roles.Depth {
	switch p {
	case colorprofile.TrueColor, colorprofile.Unknown:
		return roles.TrueColor
	case colorprofile.ANSI256:
		return roles.Colors256
	case colorprofile.ANSI:
		return roles.Colors16
	case colorprofile.ASCII, colorprofile.NoTTY:
	}
	return roles.NoColor
}

// fg is text in a role's color, plus any attribute that stands in for the
// color where the terminal can't draw it.
func fg(f *roles.Fitted, r roles.Role) lipgloss.Style {
	return attrs(lipgloss.NewStyle().Foreground(f.Color(r)), f.Attr(r))
}

// on is a surface filled with a role's color.
func on(f *roles.Fitted, r roles.Role) lipgloss.Style {
	return attrs(lipgloss.NewStyle().Background(f.Color(r)), f.Attr(r))
}

// fill is a solid fill of a tone, with the tone's own text color on it.
func fill(f *roles.Fitted, tone roles.Role) lipgloss.Style {
	return attrs(lipgloss.NewStyle().Background(f.Color(tone)).Foreground(f.Color(tone.On())), f.Attr(tone.On()))
}

// soft is a tone faded into the background, with the tone as text on it.
func soft(f *roles.Fitted, tone roles.Role) lipgloss.Style {
	return fg(f, tone).Background(f.Color(tone.Soft()))
}

// attrs adds the attributes in a. Only those set are added, so a style that
// inherits from this one still picks up the rest.
func attrs(s lipgloss.Style, a roles.Attr) lipgloss.Style {
	if a&roles.Bold != 0 {
		s = s.Bold(true)
	}
	if a&roles.Faint != 0 {
		s = s.Faint(true)
	}
	if a&roles.Italic != 0 {
		s = s.Italic(true)
	}
	if a&roles.Underline != 0 {
		s = s.Underline(true)
	}
	if a&roles.Reverse != 0 {
		s = s.Reverse(true)
	}
	if a&roles.Strikethrough != 0 {
		s = s.Strikethrough(true)
	}
	return s
}

func (m *model) View() tea.View {
	var v tea.View
	v.AltScreen = true
	if m.f == nil || m.width == 0 {
		return v
	}
	f := m.f
	v.BackgroundColor = f.Color(roles.Bg)
	page := lipgloss.NewStyle().Foreground(f.Color(roles.Fg))

	// Top bar, on a surface.
	t := tint.Current()
	left := fg(f, roles.Accent).Bold(true).Render("deploys") + fg(f, roles.Muted).Render(" / production")
	right := fg(f, roles.Muted).Render(fmt.Sprintf("%s · %s", t.DisplayName, depthLabel(m.depth)))
	gap := max(1, m.width-lipgloss.Width(left)-lipgloss.Width(right)-2)
	bar := on(f, roles.Surface).Width(m.width).Padding(0, 1).
		Render(left + strings.Repeat(" ", gap) + right)

	// Sidebar, on the same surface, with a border between it and the page.
	side := []string{fg(f, roles.Subtle).Render("ENVIRONMENTS"), ""}
	for i, env := range []string{"production", "staging", "preview"} {
		if i == 0 {
			side = append(side, fg(f, roles.Accent).Bold(true).Render("● "+env))
			continue
		}
		side = append(side, fg(f, roles.Muted).Render("  "+env))
	}
	// Everything below the bar. Height includes the padding.
	body := max(len(services)+12, m.height-1)
	sidebar := on(f, roles.Surface).Width(18).Height(body).Padding(1, 1).
		BorderStyle(lipgloss.NormalBorder()).BorderRight(true).
		BorderForeground(f.Color(roles.Border)).BorderBackground(f.Color(roles.Surface)).
		Render(strings.Join(side, "\n"))

	// The list. The selected row's text comes from m.row, fitted for drawing
	// on the selection, so hints and status stay readable on the highlight.
	width := max(40, m.width-sidebarWidth-4)
	rows := []string{fg(f, roles.Subtle).Render(fmt.Sprintf("  %-11s %-9s %-10s %s", "SERVICE", "VERSION", "STATUS", "AGE"))}
	for i, s := range services {
		rf := f
		row := lipgloss.NewStyle()
		if i == m.selected {
			rf = m.row
			row = on(f, roles.Selection).Bold(true)
		}
		cells := fg(rf, roles.Fg).Inherit(row).Render(fmt.Sprintf("  %-11s ", s.name)) +
			fg(rf, roles.Muted).Inherit(row).Render(fmt.Sprintf("%-9s ", s.version)) +
			fg(rf, s.status).Inherit(row).Render(fmt.Sprintf("● %-8s ", s.label)) +
			fg(rf, roles.Subtle).Inherit(row).Render(s.age)
		rows = append(rows, row.Width(width).Render(cells))
	}

	// A dialog on a raised surface, with an input sunk into it.
	sel := services[m.selected]
	raised := on(f, roles.Raised)
	input := on(f, roles.Sunken).Width(34).Padding(0, 1).Render(
		fg(f, roles.Subtle).Background(f.Color(roles.Sunken)).Render("type " + sel.name + " to confirm"))
	buttons := on(f, control).Foreground(f.Color(roles.Fg)).Padding(0, 2).Render("Cancel") + raised.Render("  ") +
		fill(f, roles.Danger).Padding(0, 2).Render("Roll back")
	dialog := raised.Padding(1, 2).Width(40).Render(strings.Join([]string{
		raised.Inherit(page).Bold(true).Render("Roll back " + sel.name + "?"),
		raised.Inherit(fg(f, roles.Muted)).Render("Redeploys v2.13.2 to production."),
		"",
		input,
		"",
		lipgloss.PlaceHorizontal(36, lipgloss.Right, buttons, lipgloss.WithWhitespaceStyle(raised)),
	}, "\n"))

	// Status pills: solid fills, and soft ones.
	var pills []string
	for _, tone := range []roles.Role{roles.Success, roles.Warning, roles.Danger, roles.Info, roles.Accent} {
		pills = append(pills, fill(f, tone).Padding(0, 1).Render(tone.String()))
	}
	pills = append(pills, " ")
	for _, tone := range []roles.Role{roles.Success, roles.Warning, roles.Danger, roles.Info} {
		pills = append(pills, soft(f, tone).Padding(0, 1).Render(tone.String()))
	}

	help := fg(f, roles.Subtle).Render("↑/↓ select  ←/→ tint  d color depth  q quit")
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

func depthLabel(d roles.Depth) string {
	switch d {
	case roles.TrueColor:
		return "true color"
	case roles.Colors256:
		return "256 colors"
	case roles.Colors16:
		return "16 colors"
	case roles.NoColor:
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
