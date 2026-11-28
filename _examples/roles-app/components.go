// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lrstanley/bubbletint/examples/roles-app/theme"
)

// The IDs the view gives what can be clicked, and what blocks clicks from
// reaching what's under it.
const (
	idRow      = "row:%d"
	idSetting  = "setting:%d"
	idRollBack = "rollback"
	idCancel   = "cancel"
	idConfirm  = "confirm"
	idDialog   = "dialog"
	idMenu     = "menu"
)

// ui is what the last frame drew where, and what the pointer is on.
type ui struct {
	comp           *lipgloss.Compositor
	hover, pressed string // IDs
}

// hit is the ID of what's at x, y, or "".
func (u *ui) hit(x, y int) string {
	if u.comp == nil {
		return ""
	}
	return u.comp.Hit(x, y).ID()
}

// button is the state of the button with the ID id.
func (u *ui) button(id string) theme.ButtonState {
	switch {
	case u.hover == id && u.pressed == id:
		return theme.ButtonPressed
	case u.hover == id:
		return theme.ButtonHover
	}
	return theme.ButtonRest
}

type service struct {
	name, version, age, label string
	status                    theme.Tone
}

var services = []service{
	{name: "api", version: "v2.14.0", age: "3m", status: theme.Success, label: "live"},
	{name: "web", version: "v2.14.0", age: "3m", status: theme.Danger, label: "failed"},
	{name: "worker", version: "v2.13.2", age: "1h", status: theme.Warning, label: "degraded"},
	{name: "scheduler", version: "v2.13.2", age: "1h", status: theme.Success, label: "live"},
	{name: "search", version: "v2.12.0", age: "2d", status: theme.Info, label: "rolling"},
}

// list is the side panel's list of services.
type list struct{ selected int }

func (l *list) move(step int) {
	l.selected = (l.selected + step + len(services)) % len(services)
}

// row draws service i, width cells wide.
func (l *list) row(th *theme.Theme, i, width int, focused, hovered bool) string {
	s, selected := services[i], i == l.selected
	ink := th.RowInk(selected, focused, hovered)
	name := ink.Body()
	if selected {
		name = ink.Heading()
	}
	text := ink.Tone(s.status).Render(" ● ") + name.Render(fmt.Sprintf("%-10s", s.name)) +
		ink.Muted().Render(fmt.Sprintf("%-9s", s.version)) + ink.Hint().Render(s.age)
	return th.Row(selected, focused, hovered).Width(width).Render(text)
}

// setting is a row of the settings menu.
type setting uint8

const (
	settingTint setting = iota
	settingAccent
	settingColors
	settingFlat
)

var settingRows = []setting{settingTint, settingAccent, settingColors, settingFlat}

// label is the setting's row, as the model has it now.
func (s setting) label(m *model) string {
	switch s {
	case settingTint:
		if m.picked == "" {
			return "tint     follows the terminal"
		}
		return "tint     " + m.picked + " (reset)"
	case settingAccent:
		return "accent   " + m.settings.Accent.String()
	case settingColors:
		return "colors   " + m.settings.Colors.String()
	case settingFlat:
		if m.settings.Flat {
			return "surfaces flat"
		}
		return "surfaces layered"
	}
	return ""
}

// menu is the settings menu.
type menu struct {
	open     bool
	selected int
}

func (n *menu) move(step int) {
	n.selected = (n.selected + step + len(settingRows)) % len(settingRows)
}

// hints are the status bar's key hints.
func hints(ink theme.Ink) string {
	var b strings.Builder
	for _, h := range [][2]string{{"↑↓", "select"}, {"enter", "roll back"}, {"←→", "tint"}, {"s", "settings"}, {"q", "quit"}} {
		b.WriteString(ink.Key().Render(h[0]) + ink.Hint().Render(" "+h[1]+"   "))
	}
	return b.String()
}

func profileLabel(th *theme.Theme) string {
	switch th.Profile() { //nolint:exhaustive // the rest draw no color
	case colorprofile.TrueColor:
		return "true color"
	case colorprofile.ANSI256:
		return "256 colors"
	case colorprofile.ANSI:
		return "16 colors"
	}
	return "no color"
}
