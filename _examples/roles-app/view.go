// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lrstanley/bubbletint/examples/roles-app/theme"
)

const (
	sideWidth = 28
	menuWidth = 34
)

// View draws the frame as layers: the page, the side panel's rows, and the
// dialog and menu over them. The layers are kept for the mouse, which finds
// what it's on by their IDs.
func (m *model) View() tea.View {
	var v tea.View
	v.AltScreen = true
	v.MouseMode = tea.MouseModeAllMotion
	if m.width == 0 || m.th == nil {
		return v
	}
	th, w, h := m.th, m.width, m.height

	layers := []*lipgloss.Layer{
		lipgloss.NewLayer(th.Screen().Width(w).Height(h).Render("")),
		lipgloss.NewLayer(m.bar(th, w)).Z(1),
		lipgloss.NewLayer(m.status(th, w)).Y(h - 1).Z(1),
		lipgloss.NewLayer(th.Panel().Width(sideWidth).Height(h-2).Padding(1, 1).
			Render(th.Ink(theme.OnPanel).Hint().Render("SERVICES"))).Y(1).Z(1),
	}
	focused := !m.dialog && !m.menu.open
	for i := range services {
		id := fmt.Sprintf(idRow, i)
		// Rows only respond to the pointer while the list has focus.
		row := m.list.row(th, i, sideWidth, focused, focused && m.ui.hover == id)
		layers = append(layers, lipgloss.NewLayer(row).ID(id).Y(3+i).Z(2))
	}

	// A block of lines is drawn through the page's style, which pads the short
	// lines with the page's background: a layer is as wide as its widest
	// line, and would draw blank cells after the others.
	detail := th.Screen().Render(m.detail(th))
	x := sideWidth + 3
	layers = append(layers,
		lipgloss.NewLayer(detail).X(x).Y(2).Z(1),
		lipgloss.NewLayer(th.PageButton(theme.Danger, m.ui.button(idRollBack)).Render("Roll back")).
			ID(idRollBack).X(x).Y(3+lipgloss.Height(detail)).Z(1),
	)
	if m.dialog {
		layers = append(layers, m.dialogLayers(th, w, h)...)
	}
	if m.menu.open {
		layers = append(layers, m.menuLayers(th, w)...)
	}

	m.ui.comp = lipgloss.NewCompositor(layers...)
	v.SetContent(m.ui.comp.Render())
	return v
}

// bar is the top bar: the app's name, and the tint and colors in use.
func (m *model) bar(th *theme.Theme, w int) string {
	ink := th.Ink(theme.OnBar)
	left := th.Title("deploys", theme.OnBar) + ink.Muted().Render(" / production")
	right := ink.Muted().Render(fmt.Sprintf("%s · %s", th.Tint.DisplayName, profileLabel(th)))
	gap := ink.Body().Render(strings.Repeat(" ", max(1, w-lipgloss.Width(left)-lipgloss.Width(right)-2)))
	return th.Bar().Width(w).Padding(0, 1).Render(left + gap + right)
}

// status is the status bar: key hints.
func (m *model) status(th *theme.Theme, w int) string {
	ink := th.Ink(theme.OnBar)
	right := ""
	if th.Mono() {
		right = ink.Hint().Render("status marked by weight")
	}
	left := hints(ink)
	gap := ink.Body().Render(strings.Repeat(" ", max(1, w-lipgloss.Width(left)-lipgloss.Width(right)-2)))
	return th.Bar().Width(w).Padding(0, 1).Render(left + gap + right)
}

// detail is the selected service's details, on the page.
func (m *model) detail(th *theme.Theme) string {
	s, ink := services[m.list.selected], th.Ink(theme.OnPage)
	gap := ink.Body().Render(" ")
	var tones []string
	for _, tn := range []theme.Tone{theme.Success, theme.Warning, theme.Danger, theme.Info, theme.Accent, theme.Neutral} {
		tones = append(tones, th.Pill(tn).Render(toneName(tn)))
	}
	return strings.Join([]string{
		ink.Heading().Render(s.name) + gap + th.Pill(s.status).Render(s.label),
		ink.Muted().Render(s.version + " · deployed " + s.age + " ago"),
		"",
		th.Tag(theme.Info).Render("production") + gap + th.Tag(theme.Neutral).Render("us-east-1") + gap +
			th.Tag(s.status).Render(s.label),
		"",
		ink.Hint().Render("every tone, as a pill:"),
		strings.Join(tones, gap),
	}, "\n")
}

// dialogLayers are the roll back dialog, its shadow, and its buttons, in the
// middle of the screen.
func (m *model) dialogLayers(th *theme.Theme, w, h int) []*lipgloss.Layer {
	const inner = 44
	s, ink := services[m.list.selected], th.Ink(theme.OnDialog)
	cancel := th.Button(theme.Neutral, m.ui.button(idCancel)).Render("Cancel")
	confirm := th.Button(theme.Danger, m.ui.button(idConfirm)).Render("Roll back")
	field := th.Field().Width(inner).Render(th.Ink(theme.OnField).Hint().Render("type " + s.name + " to confirm"))
	lines := []string{
		ink.Heading().Render("Roll back " + s.name + "?"),
		ink.Muted().Render("Redeploys v2.13.2 to production."),
		"",
		field,
		"",
		"", // the buttons' row, which they're drawn over
	}
	st := th.Dialog()
	box := st.Width(inner + st.GetHorizontalFrameSize()).Render(strings.Join(lines, "\n"))
	bw, bh := lipgloss.Width(box), lipgloss.Height(box)
	x, y := max(0, (w-bw)/2), max(1, (h-bh)/2)

	// The buttons sit on the dialog's last row of content, on the right.
	by := y + st.GetBorderTopSize() + st.GetPaddingTop() + len(lines) - 1
	bx := x + st.GetBorderLeftSize() + st.GetPaddingLeft() + inner - lipgloss.Width(cancel) - 2 - lipgloss.Width(confirm)
	return []*lipgloss.Layer{
		lipgloss.NewLayer(th.Shadow().Width(bw).Height(bh).Render("")).X(x + 2).Y(y + 1).Z(5),
		lipgloss.NewLayer(box).ID(idDialog).X(x).Y(y).Z(6),
		lipgloss.NewLayer(cancel).ID(idCancel).X(bx).Y(by).Z(7),
		lipgloss.NewLayer(confirm).ID(idConfirm).X(bx + lipgloss.Width(cancel) + 2).Y(by).Z(7),
	}
}

// menuLayers are the settings menu, under the top bar on the right.
func (m *model) menuLayers(th *theme.Theme, w int) []*lipgloss.Layer {
	x := max(0, w-menuWidth-1)
	rows := len(settingRows)
	layers := []*lipgloss.Layer{
		lipgloss.NewLayer(th.Menu().Width(menuWidth).Height(rows + 2).Render("")).ID(idMenu).X(x).Y(1).Z(8),
	}
	for i, s := range settingRows {
		id := fmt.Sprintf(idSetting, i)
		hovered := m.ui.hover == id || (m.ui.hover == "" && i == m.menu.selected)
		row := th.MenuRow(hovered).Width(menuWidth).Padding(0, 1).
			Render(th.MenuInk(hovered).Body().Render(s.label(m)))
		layers = append(layers, lipgloss.NewLayer(row).ID(id).X(x).Y(2+i).Z(9))
	}
	return layers
}

func toneName(tn theme.Tone) string {
	return [...]string{
		theme.Neutral: "neutral", theme.Accent: "accent", theme.Success: "success",
		theme.Warning: "warning", theme.Danger: "danger", theme.Info: "info",
	}[tn]
}
