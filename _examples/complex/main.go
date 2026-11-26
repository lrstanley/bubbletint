// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lrstanley/bubbletint/roles/v2"
	tint "github.com/lrstanley/bubbletint/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

// This is a modified version of this example, supporting full screen, dynamic
// resizing, and clickable models (tabs, lists, dialogs, etc). Every color comes
// from semantic roles derived from the current tint, so dialogs, buttons and
// text stay readable on any tint, dark or light, and at any color depth.
// 	https://github.com/charmbracelet/lipgloss/blob/master/example

var (
	// profile is the terminal's colors, as Bubble Tea detects them.
	profile = colorprofile.TrueColor
	// f is the current tint's roles, fitted to profile. Components style from
	// it when they get a ThemeChangedMsg.
	f *roles.Fitted
)

func fit() *roles.Fitted {
	return roles.Derive(tint.Current()).Fit(roles.Target{Profile: profile})
}

type ThemeChangedMsg struct{}

type model struct {
	height int
	width  int

	// Styles.
	themeTitleStyle lipgloss.Style
	themeHintStyle  lipgloss.Style

	// Child components.
	tabs    *tabs
	dialog  *dialog
	list1   *list
	list2   *list
	history *history
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) isInitialized() bool {
	return m.height != 0 && m.width != 0
}

func (m *model) setStyles() {
	m.themeTitleStyle = f.Text(roles.Accent).Bold(true).Padding(0, 1)
	m.themeHintStyle = f.Text(roles.Subtle).Padding(0, 1)
}

// restyle fits the current tint's roles, and tells every component.
func (m *model) restyle() tea.Cmd {
	f = fit()
	m.setStyles()
	return func() tea.Msg { return ThemeChangedMsg{} }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+e" {
			zone.SetEnabled(!zone.Enabled())
			return m, nil
		}

		if msg.String() == "left" {
			tint.PreviousTint()
			cmd := m.restyle()
			return m, cmd
		}

		if msg.String() == "right" {
			tint.NextTint()
			cmd := m.restyle()
			return m, cmd
		}

		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.ColorProfileMsg:
		profile = msg.Profile
		cmd := m.restyle()
		return m, cmd
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, m.propagate(msg) //nolint:gocritic
}

func (m model) propagate(msg tea.Msg) tea.Cmd {
	// Propagate to all children.
	cmds := []tea.Cmd{
		m.tabs.Update(msg),
		m.dialog.Update(msg),
		m.list1.Update(msg),
		m.list2.Update(msg),
	}

	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		msg.Height -= m.tabs.GetHeight() +
			max(m.list1.GetHeight(), m.list2.GetHeight(), m.dialog.GetHeight()) +
			3 // +1 for bottom margin on tabs, +1 for top margin on history + 1 for tint changer.

		cmds = append(cmds, m.history.Update(msg))
		return tea.Batch(cmds...)
	}
	return tea.Batch(append(cmds, m.history.Update(msg))...)
}

func (m model) View() tea.View {
	var view tea.View
	view.AltScreen = true
	view.MouseMode = tea.MouseModeCellMotion
	view.BackgroundColor = f.Color(roles.Bg)
	view.ForegroundColor = f.Color(roles.Fg)

	if !m.isInitialized() {
		return view
	}

	s := lipgloss.NewStyle().MaxHeight(m.height).MaxWidth(m.width)

	view.SetContent(zone.Scan(s.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.themeTitleStyle.Render("theme "+tint.Current().ID),
				lipgloss.PlaceHorizontal(
					m.width-len(tint.Current().ID)-5-5, // 5 = theme, 5 = padding
					lipgloss.Right,
					m.themeHintStyle.Render("[←left / right→] to change theme"),
					lipgloss.WithWhitespaceChars(" "),
				),
			),
			lipgloss.NewStyle().MarginBottom(1).Render(m.tabs.View()),
			lipgloss.PlaceHorizontal(
				m.width, lipgloss.Center,
				lipgloss.JoinHorizontal(
					lipgloss.Top,
					m.list1.View(), m.list2.View(), m.dialog.View(),
				),
				lipgloss.WithWhitespaceChars(" "),
			),
			lipgloss.NewStyle().MarginTop(1).Render(m.history.View()),
		),
	)))

	return view
}

func main() {
	// Initialize a global zone manager, so we don't have to pass around the manager
	// throughout components.
	zone.NewGlobal()

	// Initialize the default tint we want.
	tint.NewDefaultRegistry()
	tint.SetTint(tint.TintDraculaPlus)
	f = fit()

	m := &model{
		tabs:   newTabs("Lip Gloss", "Blush", "Eye Shadow", "Mascara", "Foundation"),
		dialog: newDialog("Are you sure you want to eat marmalade?"),
		list1: newList(
			"Citrus Fruits to Try",
			listItem{name: "Grapefruit", done: true},
			listItem{name: "Yuzu", done: false},
			listItem{name: "Citron", done: false},
			listItem{name: "Kumquat", done: true},
			listItem{name: "Pomelo", done: false},
		),
		list2: newList(
			"Actual Lip Gloss Vendors",
			listItem{name: "Glossier", done: true},
			listItem{name: "Claire's Boutique", done: true},
			listItem{name: "Nyx", done: false},
			listItem{name: "Mac", done: false},
			listItem{name: "Milk", done: false},
		),
		history: newHistory(
			"The Romans learned from the Greeks that quinces slowly cooked with honey would “set” when cool. The Apicius gives a recipe for preserving whole quinces, stems and leaves attached, in a bath of honey diluted with defrutum: Roman marmalade. Preserves of quince and lemon appear (along with rose, apple, plum and pear) in the Book of ceremonies of the Byzantine Emperor Constantine VII Porphyrogennetos.",
			"Medieval quince preserves, which went by the French name cotignac, produced in a clear version and a fruit pulp version, began to lose their medieval seasoning of spices in the 16th century. In the 17th century, La Varenne provided recipes for both thick and clear cotignac.",
			"In 1524, Henry VIII, King of England, received a “box of marmalade” from Mr. Hull of Exeter. This was probably marmelada, a solid quince paste from Portugal, still made and sold in southern Europe today. It became a favourite treat of Anne Boleyn and her ladies in waiting.",
		),
	}
	m.setStyles()

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err) //nolint:forbidigo
		os.Exit(1)
	}
}
