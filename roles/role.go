// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"slices"
	"strconv"
	"sync"

	tint "github.com/lrstanley/bubbletint/v2"
)

// Role is a semantic color: what a color is for, such as [Muted] text or a
// [Raised] dialog, rather than which color it is. The core roles are
// constants; an app adds its own with [Define]. Any role can be put in a
// [State], such as Overlay.Hover() for a menu's row under the pointer.
type Role uint16

// The layers a screen is built from. Bg to Overlay are worked out from the
// background, so they're the same wherever they're drawn; Sunken and Control
// are worked out from what they're drawn on.
const (
	// Bg is the background.
	Bg Role = iota
	// Surface is a bar or side panel: one step off the background.
	Surface
	// Raised is a dialog or toast.
	Raised
	// Overlay is a menu or dropdown, which can float over a dialog too.
	Overlay
	// Sunken is a well, such as an input or a code block, sunk into whatever
	// it's drawn on: an input on a dialog sinks into the dialog.
	Sunken

	// Control is a neutral button or chip: a step off whatever it's drawn
	// on, as far as its text still reads. It's worked out from the surface
	// given to [Palette.On], so a button on a dialog stands off the dialog.
	Control

	// Fg is body text. It, and the text and lines below, are worked out from
	// what they're drawn on: on a tone's fill, body text is whatever reads on
	// the fill.
	Fg
	// Muted is secondary text. It still reads as body text.
	Muted
	// Subtle is hints, placeholders and disabled text.
	Subtle
	// Border is an outline or divider.
	Border
	// BorderStrong is an emphasized outline, or a scrollbar thumb.
	BorderStrong

	// Neutral is a tone without a hue: a gray for labels and pills that say
	// nothing about status.
	Neutral
	// Accent is the tint's own color: focus, keys, the current thing.
	Accent
	// Accent2 is a second color, for variety without meaning.
	Accent2
	// Success, Warning, Danger and Info are status.
	Success
	Warning
	Danger
	Info
	// Red to Magenta are the tint's hues, for things colored by the user,
	// such as a profile set to "blue". See [Hue].
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan

	// numCore counts the roles above. The soft variant of each tone follows,
	// then the roles defined by apps.
	numCore
)

const (
	firstText  = Fg
	firstTone  = Neutral
	numTones   = numCore - firstTone
	numBuiltin = numCore + numTones

	// A role's state is kept in its top bits; the rest is the role at rest.
	stateShift        = 12
	restMask     Role = 1<<stateShift - 1
	selectedBit  Role = 1 << stateShift
	hoverBit     Role = 2 << stateShift
	pressedBit   Role = 4 << stateShift
	stateMask         = selectedBit | hoverBit | pressedBit
	maxRoles          = int(restMask) + 1
	firstDefined      = numBuiltin
)

// State is what's happening to a role: under the pointer, pressed, or
// selected. A selected role can be hovered or pressed too.
type State uint8

// States. [StateHover] and [StatePressed] don't combine; [StateSelected]
// combines with either.
const (
	StateSelected State = 1 << iota
	StateHover
	StatePressed
)

var stateNames = [...]string{"selected", "hover", "pressed"}

// String returns the states' names, joined by ":", or "rest".
func (s State) String() string {
	if s == 0 {
		return "rest"
	}
	var out string
	for i, n := range stateNames {
		if s&(1<<i) != 0 {
			if out != "" {
				out += ":"
			}
			out += n
		}
	}
	return out
}

// Kind is how a role is drawn, which decides how it is guarded and fitted.
type Kind uint8

const (
	// KindSurface is a background: something drawn on.
	KindSurface Kind = iota
	// KindText is a foreground: text, marks and lines, which must stand out
	// from the surface they are drawn on.
	KindText
	// KindTone is a hue's base color, used as text or marks, or as a solid
	// fill with the text that reads on it.
	KindTone
	// KindSoft is a hue faded into the surface it's drawn on, as a fill. The
	// tone's base color reads on it. See [Role.Soft].
	KindSoft
)

// Hover returns r under the pointer: a step from r toward the text on the
// page, kept apart from r with 256 colors, and marked with an attribute
// where no step can be drawn. It replaces [Role.Pressed].
func (r Role) Hover() Role { return r&^pressedBit | hoverBit }

// Pressed returns r while it's pressed: a step further than [Role.Hover]. It
// replaces Hover.
func (r Role) Pressed() Role { return r&^hoverBit | pressedBit }

// Selected returns r while it's selected, such as a list's current row: r
// tinted toward the accent, as far as text still reads on it.
func (r Role) Selected() Role { return r | selectedBit }

// In returns r in the states s, added to those it's in.
func (r Role) In(s State) Role {
	if s&StateSelected != 0 {
		r = r.Selected()
	}
	switch {
	case s&StatePressed != 0:
		r = r.Pressed()
	case s&StateHover != 0:
		r = r.Hover()
	}
	return r
}

// State returns the states r is in.
func (r Role) State() State {
	var s State
	if r&selectedBit != 0 {
		s |= StateSelected
	}
	if r&hoverBit != 0 {
		s |= StateHover
	}
	if r&pressedBit != 0 {
		s |= StatePressed
	}
	return s
}

// Rest returns r without its states.
func (r Role) Rest() Role { return r & restMask }

// IsTone reports whether r is a tone at rest or in a state: a hue with a soft
// variant.
func (r Role) IsTone() bool {
	rest := r.Rest()
	return rest >= firstTone && rest < numCore
}

// Soft returns the tone's color faded into the surface it's drawn on, as a
// fill, in r's states. It returns r for anything but a tone.
func (r Role) Soft() Role {
	if !r.IsTone() {
		return r
	}
	return r + numTones
}

// Base returns the tone a soft role belongs to, in r's states. Any other role
// is returned as is.
func (r Role) Base() Role {
	if rest := r.Rest(); rest >= numCore && rest < numBuiltin {
		return r - numTones
	}
	return r
}

// Kind returns how r is drawn. A role in a state is drawn as it is at rest.
func (r Role) Kind() Kind {
	switch rest := r.Rest(); {
	case rest < firstText:
		return KindSurface
	case rest < firstTone:
		return KindText
	case rest < numCore:
		return KindTone
	case rest < numBuiltin:
		return KindSoft
	}
	if s, ok := specOf(r); ok {
		return s.Kind
	}
	return KindText
}

// Hue returns the tone for one of the tint's chromatic slots: [Red] for
// [tint.SlotRed] or [tint.SlotBrightRed], and so on. Black, white and their
// bright variants have no hue, and return [Neutral].
func Hue(s tint.Slot) Role {
	switch s.Normal() { //nolint:exhaustive // black and white have no tone
	case tint.SlotRed:
		return Red
	case tint.SlotGreen:
		return Green
	case tint.SlotYellow:
		return Yellow
	case tint.SlotBlue:
		return Blue
	case tint.SlotMagenta:
		return Magenta
	case tint.SlotCyan:
		return Cyan
	default:
		return Neutral
	}
}

var coreNames = [numCore]string{
	"bg", "surface", "raised", "overlay", "sunken", "control",
	"fg", "muted", "subtle", "border", "border-strong",
	"neutral", "accent", "accent2", "success", "warning", "danger", "info",
	"red", "green", "yellow", "blue", "magenta", "cyan",
}

// String returns the role's name, such as "danger-soft", with its states
// after a colon, such as "overlay:hover" or "bg:selected:hover".
func (r Role) String() string {
	name := r.restName()
	if s := r.State(); s != 0 {
		name += ":" + s.String()
	}
	return name
}

func (r Role) restName() string {
	rest := r.Rest()
	switch {
	case rest < numCore:
		return coreNames[rest]
	case rest < numBuiltin:
		return coreNames[rest.Base()] + "-soft"
	}
	if s, ok := specOf(rest); ok {
		return s.Name
	}
	return "role(" + strconv.Itoa(int(rest)) + ")"
}

// minFor is the contrast a role keeps against the surface it's drawn on,
// when the guards are on. 0 is no floor.
func minFor(r Role) float64 {
	switch r.Rest() { //nolint:exhaustive // tones and defined roles follow
	case Fg:
		return minText + fgMargin
	case Muted:
		return minText
	case Subtle:
		return minSubtle
	case Border:
		return minBorder
	}
	if r.Kind() == KindTone {
		return minColor
	}
	if s, ok := specOf(r); ok && s.Kind != KindSurface {
		return s.Min
	}
	return 0
}

// Spec describes a role an app adds with [Define].
type Spec struct {
	// Name is the role's name, as [Role.String] returns it.
	Name string
	// Kind is how the role is drawn: [KindSurface] for a background, or
	// [KindText] for text, marks and lines. Any other kind is taken as
	// [KindText].
	Kind Kind
	// Derive works out the role's color from the palette's other roles. It
	// may use the core roles, and roles defined before this one. It runs for
	// every surface drawn on, with [Palette.On]'s palette: [Palette.Ground]
	// is the surface, and [Palette.GroundColor] its color, so a role can be
	// worked out from whatever it's on.
	Derive func(p *Palette) tint.Color
	// Min is the contrast the role keeps against the surface it's drawn on,
	// when the guards are on: 4.5 for body text, 3 for marks. 0 is none. It
	// applies to [KindText] roles.
	Min float64
	// Distinct are roles this one shouldn't share a color with when it's
	// fitted to 256 colors. A surface is always kept apart from the surface
	// it's drawn on. It gives way to Min: a readable color beats a distinct
	// one. Roles in a state are ignored here.
	Distinct []Role
	// Fallback is how the role is drawn with 16 colors or none.
	Fallback Fallback
}

// Fallback is how a role is drawn where its own color can't be: with only
// the terminal's 16 colors, or none.
type Fallback struct {
	// Slots are the terminal colors to try, in order, with 16 colors. The
	// first that keeps the role's Min against what it's drawn on is used.
	Slots []tint.Slot
	// Accent tries the accent's slots after Slots.
	Accent bool
	// Attr is added where no slot fits, and with no color at all.
	Attr Attr
}

var registry struct {
	mu    sync.RWMutex
	specs []Spec
}

// Define adds a role for an app, derived alongside the core roles and fitted
// like them. Call it once per role, from a package-level var, before deriving
// a palette:
//
//	// Shadow is the shadow under a dragged pane.
//	var Shadow = roles.Define(roles.Spec{
//		Name: "shadow",
//		Kind: roles.KindSurface,
//		Derive: func(p *roles.Palette) tint.Color {
//			return p.Color(roles.Bg).Mix(tint.Color{A: 0xff}, 0.5)
//		},
//		Fallback: roles.Fallback{Slots: []tint.Slot{tint.SlotBlack}},
//	})
//
// It panics if Derive is nil, or if there's no room for another role.
func Define(s Spec) Role {
	if s.Derive == nil {
		panic("roles: Define with a nil Derive")
	}
	if s.Kind != KindSurface {
		s.Kind = KindText
	}
	s.Distinct = slices.Clone(s.Distinct)
	s.Fallback.Slots = slices.Clone(s.Fallback.Slots)
	registry.mu.Lock()
	defer registry.mu.Unlock()
	if int(firstDefined)+len(registry.specs) >= maxRoles {
		panic("roles: too many roles defined")
	}
	registry.specs = append(registry.specs, s)
	return firstDefined + Role(len(registry.specs)-1) //nolint:gosec // bounded by maxRoles
}

func specOf(r Role) (Spec, bool) {
	rest := r.Rest()
	if rest < firstDefined {
		return Spec{}, false
	}
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	if i := int(rest - firstDefined); i < len(registry.specs) {
		return registry.specs[i], true
	}
	return Spec{}, false
}

// defined returns every role defined so far.
func defined() []Spec {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	return slices.Clip(registry.specs)
}
