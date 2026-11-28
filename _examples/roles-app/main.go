// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// This example shows how an app is structured around bubbletint's roles: a
// theme package that's the only place colors come from, components that ask
// it for what they draw, and an app that holds one current theme and builds
// a new one when anything changes.
//
//   - The tint follows the terminal's background: dark or light, from the
//     settings' pair of tints. ←/→ picks another tint, which crossfades in.
//   - The settings (accent, color limit, flat surfaces) come from an optional
//     JSON config file, and change live from the settings menu (s).
//   - The color profile follows the terminal, capped by the settings.
//   - Rows and buttons respond to the mouse: hover, press and click.
//
// Run it with a config file of your own, such as:
//
//	{"tints": {"dark": "gruvbox_dark", "light": "gruvbox_light"}, "accent": "purple", "colors": "256"}
//
// Keys: ↑/↓ select, enter rolls back, ←/→ tint, s settings, q quits.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lrstanley/bubbletint/examples/roles-app/theme"
	"github.com/lrstanley/bubbletint/roles/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// tints are the tints ←/→ cycle through.
var tints = []string{
	"tokyo_night", "tokyo_night_day", "catppuccin_mocha", "catppuccin_latte",
	"gruvbox_dark", "nord", "amber_crt_retro", "rose_pine_dawn",
}

// accents are the accents the settings menu cycles through.
var accents = []string{"auto", "blue", "purple", "green", "#ff8800"}

// fadeFor is how long a change of tint crossfades for.
const fadeFor = 260 * time.Millisecond

// fadeMsg is a frame of the crossfade with the number fade, so a frame of
// one that was replaced is dropped rather than running a second loop.
type fadeMsg struct{ fade int }

func fadeTick(fade int) tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg { return fadeMsg{fade} })
}

type model struct {
	width, height int

	settings theme.Settings
	env      theme.Env
	picked   string // a tint picked with ←/→, or "" to follow the settings

	// th is the current theme, which everything draws with. During a
	// crossfade it's a frame between from and to.
	th       *theme.Theme
	from, to *theme.Theme
	fadeAt   time.Time
	fades    int // the crossfade running, counting from 1

	list   list
	menu   menu
	dialog bool

	ui ui // what the last frame drew where, for the mouse
}

// tint is the tint to draw with: the one picked, or the settings' tint for
// the terminal's background.
func (m *model) tint() *tint.Tint {
	id := m.picked
	if id == "" {
		id = m.settings.Tints.For(!m.env.Light)
	}
	if t := tint.DefaultTintsByID(id); t != nil {
		return t
	}
	return tint.TintTokyoNight
}

// retheme builds the theme from the tint, the settings and the terminal. A
// new tint crossfades in; anything else changes at once.
func (m *model) retheme(fade bool) tea.Cmd {
	next := theme.New(m.tint(), m.settings, m.env)
	if !fade || m.th == nil {
		m.th, m.from, m.to = next, nil, nil
		return nil
	}
	m.from, m.to, m.fadeAt = m.th, next, time.Now()
	m.fades++
	return fadeTick(m.fades)
}

// Init asks the terminal for its background, to pick the dark or light tint.
// The app draws its own background over the terminal's, but never changes
// the terminal's own, so the answer is always the terminal's.
func (m *model) Init() tea.Cmd { return tea.RequestBackgroundColor }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.ColorProfileMsg:
		m.env.Profile = msg.Profile
		cmd = m.retheme(false)
	case tea.BackgroundColorMsg:
		m.env.Light = !msg.IsDark()
		cmd = m.retheme(m.picked == "")
	case fadeMsg:
		if msg.fade == m.fades {
			cmd = m.fade()
		}
	case tea.KeyPressMsg:
		cmd = m.key(msg.String())
	case tea.MouseMotionMsg:
		m.ui.hover = m.ui.hit(msg.X, msg.Y)
	case tea.MouseClickMsg:
		m.ui.pressed = m.ui.hit(msg.X, msg.Y)
	case tea.MouseReleaseMsg:
		id, pressed := m.ui.hit(msg.X, msg.Y), m.ui.pressed
		m.ui.pressed = ""
		if id != "" && id == pressed {
			cmd = m.click(id)
		}
	}
	return m, cmd
}

// fade moves a crossfade on a frame, to the next theme once it's done.
func (m *model) fade() tea.Cmd {
	if m.to == nil {
		return nil
	}
	p := float64(time.Since(m.fadeAt)) / float64(fadeFor)
	if p >= 1 {
		m.th, m.from, m.to = m.to, nil, nil
		return nil
	}
	m.th = theme.Lerp(m.from, m.to, p)
	return fadeTick(m.fades)
}

func (m *model) key(k string) tea.Cmd {
	switch k {
	case "q", "ctrl+c":
		return tea.Quit
	case "esc":
		m.dialog, m.menu.open = false, false
	case "s":
		m.menu.open = !m.menu.open
	case "left", "right":
		step := 1
		if k == "left" {
			step = len(tints) - 1
		}
		m.picked = tints[(max(slices.Index(tints, m.tint().ID), 0)+step)%len(tints)]
		return m.retheme(true)
	case "up", "k", "down", "j":
		step := 1
		if k == "up" || k == "k" {
			step = -1
		}
		switch {
		case m.menu.open:
			m.menu.move(step)
		case !m.dialog:
			m.list.move(step)
		}
	case "enter":
		switch {
		case m.menu.open:
			return m.apply(m.menu.selected)
		default:
			m.dialog = !m.dialog
		}
	}
	return nil
}

// click runs what was clicked, by the ID the view gave it.
func (m *model) click(id string) tea.Cmd {
	var i int
	switch {
	case id == idRollBack:
		m.dialog = true
	case id == idCancel || id == idConfirm:
		m.dialog = false
	case scan(id, idRow, &i):
		m.list.selected = i
	case scan(id, idSetting, &i):
		m.menu.selected = i
		return m.apply(i)
	}
	return nil
}

// apply changes setting i to its next value, and rethemes.
func (m *model) apply(i int) tea.Cmd {
	switch settingRows[i] {
	case settingAccent:
		next := accents[(slices.Index(accents, m.settings.Accent.String())+1)%len(accents)]
		m.settings.Accent, _ = roles.ParseAccent(next)
	case settingColors:
		limits := roles.ColorLimits()
		m.settings.Colors = limits[(slices.Index(limits, m.settings.Colors)+1)%len(limits)]
	case settingFlat:
		m.settings.Flat = !m.settings.Flat
	case settingTint:
		m.picked = ""
		return m.retheme(true)
	}
	return m.retheme(false)
}

func scan(s, format string, v *int) bool {
	_, err := fmt.Sscanf(s, format, v)
	return err == nil
}

func main() {
	config := flag.String("config", "", "a JSON settings file")
	flag.Parse()

	m := &model{settings: theme.DefaultSettings()}
	if *config != "" {
		data, err := os.ReadFile(*config)
		if err == nil {
			err = json.Unmarshal(data, &m.settings)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading settings:", err)
			os.Exit(1)
		}
	}
	m.retheme(false)

	// Motion that stays on what the pointer was on changes nothing, so it's
	// dropped before it reaches Update, and draws no frame.
	moved := func(_ tea.Model, msg tea.Msg) tea.Msg {
		if mm, ok := msg.(tea.MouseMotionMsg); ok && m.ui.hit(mm.X, mm.Y) == m.ui.hover {
			return nil
		}
		return msg
	}
	if _, err := tea.NewProgram(m, tea.WithFilter(moved)).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
