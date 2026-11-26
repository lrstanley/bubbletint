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
// constants; an app adds its own with [Define].
type Role uint16

// Surfaces, text and lines. Every [Palette] has a color for each.
const (
	// Bg is the background.
	Bg Role = iota
	// Fg is body text.
	Fg
	// Surface is a bar or side panel: one step off the background.
	Surface
	// Raised is a dialog or toast.
	Raised
	// Overlay is a menu or dropdown, which can float over a dialog too.
	Overlay
	// Sunken is a well, such as an input or a code block.
	Sunken
	// Hover is a row under the pointer.
	Hover
	// Selection is the selected row's highlight. Draw text on it with
	// [Palette.On].
	Selection
	// Muted is secondary text. It still reads as body text.
	Muted
	// Subtle is hints, placeholders and disabled text.
	Subtle
	// Border is an outline or divider.
	Border
	// BorderStrong is an emphasized outline, or a scrollbar thumb.
	BorderStrong

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

	// numCore counts the roles above. The soft and on variants of each tone
	// follow, then the roles defined by apps.
	numCore
)

const (
	firstTone  = Accent
	numTones   = numCore - firstTone
	numBuiltin = numCore + 2*numTones
)

// Kind is how a role is drawn, which decides how it is guarded and fitted.
type Kind uint8

const (
	// KindSurface is a background: something drawn on.
	KindSurface Kind = iota
	// KindText is a foreground: text, marks and lines, which must stand out
	// from the surface they are drawn on.
	KindText
	// KindTone is a hue's base color, used as text or marks, or as a solid
	// fill with its [Role.On] text on it.
	KindTone
	// KindSoft is a hue faded into the background, as a fill. The tone's base
	// color reads on it. See [Role.Soft].
	KindSoft
	// KindOn is text on a tone's base color used as a fill. See [Role.On].
	KindOn
)

// IsTone reports whether r is a tone: a hue with soft and on variants.
func (r Role) IsTone() bool { return r >= firstTone && r < numCore }

// Soft returns the tone's color faded into the background, as a fill. It
// returns r for anything but a tone's base role.
func (r Role) Soft() Role {
	if !r.IsTone() {
		return r
	}
	return r + numTones
}

// On returns the color for text on the tone's base color used as a fill. It
// returns r for anything but a tone's base role.
func (r Role) On() Role {
	if !r.IsTone() {
		return r
	}
	return r + 2*numTones
}

// Base returns the tone a soft or on role belongs to. Any other role is
// returned as is.
func (r Role) Base() Role {
	switch {
	case r >= numCore && r < numCore+numTones:
		return r - numTones
	case r >= numCore+numTones && r < numBuiltin:
		return r - 2*numTones
	}
	return r
}

// Kind returns how r is drawn.
func (r Role) Kind() Kind {
	switch {
	case r == Fg:
		return KindText
	case r <= Selection:
		return KindSurface
	case r < firstTone:
		return KindText
	case r < numCore:
		return KindTone
	case r < numCore+numTones:
		return KindSoft
	case r < numBuiltin:
		return KindOn
	}
	if s, ok := specOf(r); ok {
		return s.Kind
	}
	return KindText
}

// Hue returns the tone for one of the tint's chromatic slots: [Red] for
// [tint.SlotRed] or [tint.SlotBrightRed], and so on. Black, white and their
// bright variants have no tone, and return [Fg].
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
		return Fg
	}
}

var coreNames = [numCore]string{
	"bg", "fg", "surface", "raised", "overlay", "sunken", "hover", "selection",
	"muted", "subtle", "border", "border-strong",
	"accent", "accent2", "success", "warning", "danger", "info",
	"red", "green", "yellow", "blue", "magenta", "cyan",
}

// String returns the role's name, such as "danger-soft" or "on-accent".
func (r Role) String() string {
	switch r.Kind() {
	case KindSoft:
		if r < numBuiltin {
			return coreNames[r.Base()] + "-soft"
		}
	case KindOn:
		if r < numBuiltin {
			return "on-" + coreNames[r.Base()]
		}
	case KindSurface, KindText, KindTone:
	}
	if r < numCore {
		return coreNames[r]
	}
	if s, ok := specOf(r); ok {
		return s.Name
	}
	return "role(" + strconv.Itoa(int(r)) + ")"
}

// minFor is the contrast a core role keeps against the surface it's drawn on,
// when the guards are on. 0 is no floor.
func minFor(r Role) float64 {
	switch r { //nolint:exhaustive // tones and defined roles follow
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
	if s, ok := specOf(r); ok {
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
	// may use the core roles, and roles defined before this one.
	Derive func(p *Palette) tint.Color
	// Min is the contrast the role keeps against the surface it's drawn on,
	// when the guards are on: 4.5 for body text, 3 for marks. 0 is none. It
	// applies to [KindText] roles.
	Min float64
	// Distinct are roles this one shouldn't share a color with when it's
	// fitted to 256 colors, such as the surface a control sits on. It gives
	// way to Min: a readable color beats a distinct one. A role defined after
	// a palette was derived only keeps apart from roles derived with it, so
	// define roles before deriving, as from a package-level var.
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
// like them. Call it once per role, typically from a package-level var:
//
//	var Control = roles.Define(roles.Spec{
//		Name: "control",
//		Kind: roles.KindSurface,
//		Derive: func(p *roles.Palette) tint.Color {
//			fg := p.Color(roles.Fg)
//			return p.Color(roles.Raised).MixReadable(fg, fg, p.ByDark(0.22, 0.17), 4.5)
//		},
//		Distinct: []roles.Role{roles.Raised},
//	})
//
// It panics if Derive is nil.
func Define(s Spec) Role {
	if s.Derive == nil {
		panic("roles: Define with a nil Derive")
	}
	if s.Kind != KindSurface {
		s.Kind = KindText
	}
	s.Distinct = append([]Role(nil), s.Distinct...)
	s.Fallback.Slots = append([]tint.Slot(nil), s.Fallback.Slots...)
	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.specs = append(registry.specs, s)
	return numBuiltin + Role(len(registry.specs)-1) //nolint:gosec // bounded by the number of Define calls
}

func specOf(r Role) (Spec, bool) {
	if r < numBuiltin {
		return Spec{}, false
	}
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	if i := int(r - numBuiltin); i < len(registry.specs) {
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
