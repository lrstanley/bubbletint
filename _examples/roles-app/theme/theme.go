// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package theme is the one place the app's colors come from. Components ask
// it for what they draw, such as a panel, a selected row or a danger button,
// and never pick a color, or a role, themselves. Only this package imports
// bubbletint's roles.
//
// A [Theme] is immutable. The app keeps the current one, and builds a new one
// when the tint, a setting, or the terminal changes; a component never keeps
// a style past the frame it drew, so the next frame is drawn in the new one.
package theme

import (
	"strings"
	"sync"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lrstanley/bubbletint/roles/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Tone is what something is about, which picks its color.
type Tone uint8

// Tones.
const (
	Neutral Tone = iota
	Accent
	Success
	Warning
	Danger
	Info
)

func (tn Tone) role() roles.Role {
	return [...]roles.Role{
		Neutral: roles.Neutral, Accent: roles.Accent, Success: roles.Success,
		Warning: roles.Warning, Danger: roles.Danger, Info: roles.Info,
	}[tn]
}

// Settings are the user's choices, as read from the config file. Each field
// parses itself, so a bad value is an error when the file is read.
type Settings struct {
	// Tints are the tint IDs for dark and light terminals, or one ID for both.
	Tints tint.Pair `json:"tints"`
	// Accent is "auto", one of the tint's colors, such as "purple", or "#rrggbb".
	Accent roles.AccentChoice `json:"accent"`
	// Colors caps how many colors are drawn: "auto", "256", "16" or "none".
	Colors roles.ColorLimit `json:"colors"`
	// Flat draws panels, dialogs and menus on the background, with no shade.
	Flat bool `json:"flat"`
}

// DefaultSettings are the settings without a config file.
func DefaultSettings() Settings {
	return Settings{Tints: tint.Pair{Dark: "tokyo_night", Light: "tokyo_night_day"}}
}

// Env is what the terminal says about itself.
type Env struct {
	// Profile is the terminal's color support. The zero value is unknown,
	// taken as true color.
	Profile colorprofile.Profile
	// Light is set when the terminal's background is light. The zero value
	// is dark, until the terminal says.
	Light bool
}

// shadow is a role the app adds to bubbletint's: the shadow a dialog casts on
// the page, the background darkened. It's derived and fitted like the core
// roles, and with 16 colors it's black.
var shadow = roles.Define(roles.Spec{
	Name: "shadow",
	Kind: roles.KindSurface,
	Derive: func(p *roles.Palette) tint.Color {
		return p.Color(roles.Bg).Mix(tint.Color{A: 0xff}, p.ByDark(0.45, 0.25))
	},
	Fallback: roles.Fallback{Slots: []tint.Slot{tint.SlotBlack}},
})

// Theme is everything the app draws with.
type Theme struct {
	// Tint is the tint the theme was made from.
	Tint *tint.Tint
	// Settings are the choices it was made with.
	Settings Settings

	f      *roles.Fitted
	styles func() *styles // built the first time they're asked for
}

// New makes the theme for tint t, the user's settings, and the terminal.
func New(t *tint.Tint, s Settings, env Env) *Theme {
	opts := []roles.Option{s.Accent.Option()}
	if s.Flat {
		opts = append(opts, roles.WithFlatSurfaces())
	}
	f := roles.Derive(t, opts...).Fit(roles.Target{Profile: s.Colors.Cap(env.Profile)})
	return newTheme(t, s, f)
}

func newTheme(t *tint.Tint, s Settings, f *roles.Fitted) *Theme {
	th := &Theme{Tint: t, Settings: s, f: f}
	th.styles = sync.OnceValue(th.build)
	return th
}

// Lerp returns a theme partway from a to b, for crossfading between tints: p
// runs from 0 (a) to 1 (b). Below 256 colors there's nothing between, so it's
// b.
func Lerp(a, b *Theme, p float64) *Theme {
	if a == nil || b == nil {
		return b
	}
	switch f := roles.Lerp(a.f, b.f, p); f {
	case a.f:
		return a
	case b.f:
		return b
	default:
		return newTheme(b.Tint, b.Settings, f)
	}
}

// Profile is how many colors the theme draws.
func (t *Theme) Profile() colorprofile.Profile { return t.f.Profile() }

// Mono reports whether status can't be told apart by hue, so danger and
// warning are marked with weight instead.
func (t *Theme) Mono() bool { return t.f.Mono() }

// styles are the styles that don't depend on what's drawn, built once per
// theme. A crossfade makes a theme a frame, so these are built a frame at a
// time then, and only those a frame asks for.
type styles struct {
	screen, bar, panel, menu, shadow, field lipgloss.Style
	dialog                                  lipgloss.Style
}

func (t *Theme) build() *styles {
	f := t.f
	dialog := f.Fill(roles.Raised).Padding(1, 2)
	// Where the dialog's shade wouldn't show against the page, as with flat
	// surfaces, 16 colors or none, it's outlined instead.
	if !f.Distinct(roles.Raised, roles.Bg) {
		dialog = dialog.Padding(0, 1).Border(lipgloss.RoundedBorder()).
			BorderForeground(f.Color(roles.Accent)).BorderBackground(f.Color(roles.Raised))
	}
	return &styles{
		screen: f.Fill(roles.Bg),
		bar:    f.Fill(roles.Surface),
		panel:  f.Fill(roles.Surface),
		menu:   f.Fill(roles.Overlay),
		shadow: f.Fill(shadow),
		field:  f.On(roles.Raised).Fill(roles.Sunken).Padding(0, 1),
		dialog: dialog,
	}
}

// Screen is the whole screen.
func (t *Theme) Screen() lipgloss.Style { return t.styles().screen }

// Bar is the top and status bars.
func (t *Theme) Bar() lipgloss.Style { return t.styles().bar }

// Panel is the side panel.
func (t *Theme) Panel() lipgloss.Style { return t.styles().panel }

// Dialog is a box over the page: its shade, or an outline where shades don't
// show.
func (t *Theme) Dialog() lipgloss.Style { return t.styles().dialog }

// Shadow is the shadow a dialog casts.
func (t *Theme) Shadow() lipgloss.Style { return t.styles().shadow }

// Field is a text input in a dialog: a well, sunk into the dialog.
func (t *Theme) Field() lipgloss.Style { return t.styles().field }

// Menu is a dropdown menu.
func (t *Theme) Menu() lipgloss.Style { return t.styles().menu }

// rowRole is a side panel's row: the panel, selected, and stepped under the
// pointer. While the list doesn't have focus, its selection is quieter.
func rowRole(selected, focused, hovered bool) roles.Role {
	r := roles.Surface
	switch {
	case selected && focused:
		r = r.Selected()
	case selected:
		r = r.Hover()
	}
	if hovered {
		r = r.Hover()
	}
	return r
}

// Row is a row of the side panel's list.
func (t *Theme) Row(selected, focused, hovered bool) lipgloss.Style {
	return t.f.Fill(rowRole(selected, focused, hovered))
}

// RowInk is the ink for text on a row of the side panel's list: kept
// readable on the row's highlight, not only on the panel.
func (t *Theme) RowInk(selected, focused, hovered bool) Ink {
	return Ink{t.f.On(rowRole(selected, focused, hovered))}
}

// menuRole is a menu's row: the menu, stepped under the pointer.
func menuRole(hovered bool) roles.Role {
	if hovered {
		return roles.Overlay.Hover()
	}
	return roles.Overlay
}

// MenuRow is a row of a menu.
func (t *Theme) MenuRow(hovered bool) lipgloss.Style { return t.f.Fill(menuRole(hovered)) }

// MenuInk is the ink for text on a menu's row.
func (t *Theme) MenuInk(hovered bool) Ink { return Ink{t.f.On(menuRole(hovered))} }

// ButtonState is how a button is being touched.
type ButtonState uint8

// Button states.
const (
	ButtonRest ButtonState = iota
	ButtonHover
	ButtonPressed
)

// Button is a dialog's button: in its tone, or for the neutral way out, a
// control a step off the dialog.
func (t *Theme) Button(tn Tone, st ButtonState) lipgloss.Style {
	r := tn.role()
	if tn == Neutral {
		r = roles.Control
	}
	return t.f.On(roles.Raised).Fill(st.of(r)).Bold(true).Padding(0, 2)
}

// PageButton is a button on the page, in its tone.
func (t *Theme) PageButton(tn Tone, st ButtonState) lipgloss.Style {
	return t.f.Fill(st.of(tn.role())).Bold(true).Padding(0, 2)
}

// of is r in the state: a step toward the text under the pointer, and a
// step further while pressed.
func (st ButtonState) of(r roles.Role) roles.Role {
	switch st {
	case ButtonHover:
		return r.Hover()
	case ButtonPressed:
		return r.Pressed()
	case ButtonRest:
	}
	return r
}

// Pill is a solid label in a tone, such as a service's status.
func (t *Theme) Pill(tn Tone) lipgloss.Style { return t.f.Fill(tn.role()).Bold(true).Padding(0, 1) }

// Tag is a quieter label: the tone faded into the page.
func (t *Theme) Tag(tn Tone) lipgloss.Style { return t.f.Fill(tn.role().Soft()).Padding(0, 1) }

// Where is somewhere text is drawn.
type Where uint8

// Where text is drawn.
const (
	OnPage Where = iota
	OnBar
	OnPanel
	OnDialog
	OnField
)

// Ink is the text styles for somewhere text is drawn.
func (t *Theme) Ink(w Where) Ink {
	switch w {
	case OnBar, OnPanel:
		return Ink{t.f.On(roles.Surface)}
	case OnDialog:
		return Ink{t.f.On(roles.Raised)}
	case OnField:
		return Ink{t.f.On(roles.Raised).On(roles.Sunken)}
	case OnPage:
	}
	return Ink{t.f}
}

// Title draws s from the accent to the second accent, a character at a time,
// such as the app's name. Below 256 colors it's the accent.
func (t *Theme) Title(s string, w Where) string {
	ink := t.Ink(w)
	r := []rune(s)
	if t.Profile() < colorprofile.ANSI256 || len(r) < 2 {
		return ink.Key().Render(s)
	}
	p := t.f.Palette()
	var b strings.Builder
	for i, ch := range r {
		c := p.Color(roles.Accent).Mix(p.Color(roles.Accent2), float64(i)/float64(len(r)-1))
		b.WriteString(ink.Key().Foreground(t.f.Fit(c, roles.Accent)).Render(string(ch)))
	}
	return b.String()
}

// Ink is the text styles that read on one surface. Each is drawn on the
// surface's color, as text rendered on its own and joined onto a fill must
// be, or it leaves a hole in the fill.
type Ink struct{ f *roles.Fitted }

// Body is body text.
func (i Ink) Body() lipgloss.Style { return i.f.Span(roles.Fg) }

// Heading is a title.
func (i Ink) Heading() lipgloss.Style { return i.f.Span(roles.Fg).Bold(true) }

// Muted is secondary text.
func (i Ink) Muted() lipgloss.Style { return i.f.Span(roles.Muted) }

// Hint is a hint or a placeholder.
func (i Ink) Hint() lipgloss.Style { return i.f.Span(roles.Subtle) }

// Key is a key in a hint, or something current.
func (i Ink) Key() lipgloss.Style { return i.f.Span(roles.Accent).Bold(true) }

// Tone is text in a tone, such as a status.
func (i Ink) Tone(tn Tone) lipgloss.Style { return i.f.Span(tn.role()) }
