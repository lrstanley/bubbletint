// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"cmp"
	"image/color"
	"math"
	"slices"
	"strings"

	"github.com/charmbracelet/x/ansi"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Depth is how many colors the terminal can draw. Map whatever detects the
// terminal's colors to it; with github.com/charmbracelet/colorprofile, for
// example, TrueColor is [TrueColor], ANSI256 is [Colors256], ANSI is
// [Colors16], and Ascii or NoTTY are [NoColor]. A larger Depth is fewer colors,
// so max(detected, limit) caps a detected depth at a user's limit.
type Depth uint8

const (
	// TrueColor draws every color as derived.
	TrueColor Depth = iota
	// Colors256 draws from the fixed part of the 256-color palette.
	Colors256
	// Colors16 draws with the terminal's 16 ANSI colors.
	Colors16
	// NoColor draws no color, only text attributes.
	NoColor
)

// String returns the depth's name.
func (d Depth) String() string {
	switch d {
	case TrueColor:
		return "truecolor"
	case Colors256:
		return "256"
	case Colors16:
		return "16"
	case NoColor:
		return "none"
	}
	return "unknown"
}

// Attr is a set of text attributes, used where color can't carry a role: a
// [Muted] hint is faint where no gray reads, and the [Selection] is reversed
// with no color at all.
type Attr uint8

// Text attributes.
const (
	Bold Attr = 1 << iota
	Faint
	Italic
	Underline
	Reverse
	Strikethrough
)

var attrNames = [...]string{"bold", "faint", "italic", "underline", "reverse", "strikethrough"}

// String returns the attributes' names, joined by "|", or "none".
func (a Attr) String() string {
	if a == 0 {
		return "none"
	}
	var out []string
	for i, n := range attrNames {
		if a&(1<<i) != 0 {
			out = append(out, n)
		}
	}
	return strings.Join(out, "|")
}

// Target is the terminal a palette is fitted to.
type Target struct {
	// Depth is how many colors the terminal can draw.
	Depth Depth
	// ANSI is the terminal's 16 colors, where they're known, such as from an
	// OSC 4 query. With 16 colors, roles are fitted to slots that read on this
	// palette. nil assumes the terminal uses the tint's own colors.
	ANSI *[16]color.Color
	// Unpainted leaves [Bg] and [Fg] as the terminal's defaults (nil),
	// for apps that don't paint the background. Derive with [WithTerminal]
	// too, so the other roles are worked out against the terminal's colors.
	Unpainted bool
}

// Fitted is a [Palette] fitted to a terminal: every role as a color the
// terminal can draw, plus the attributes to add where color can't carry the
// role. It's immutable, and safe to share between goroutines.
type Fitted struct {
	p      *Palette
	target Target
	ansi   [16]tint.Color // the terminal's 16 colors, or our best guess
	colors []color.Color
	attrs  []Attr
	idx    []uint8 // with 256 colors, each role's index, from 16; 0 is not yet fitted

	// With 16 colors, what text is drawn on: the slot filling the palette's
	// ground, or the background, and whether it's filled.
	under   tint.Color
	painted bool
}

// Fit fits the palette to a terminal.
//
// With 256 colors, each role takes the nearest color that keeps its contrast
// against what it's drawn on, and that differs from the roles it sits beside
// (a dialog from the background, a hint from muted text), so no two roles
// collapse into one. With 16 colors, each role takes a slot from its family
// (grays for muted text, red's two slots for danger) that reads on the
// terminal's palette, or the terminal's default plus an [Attr].
func (p *Palette) Fit(tg Target) *Fitted {
	f := &Fitted{
		p:      p,
		target: tg,
		colors: make([]color.Color, len(p.colors)),
		attrs:  make([]Attr, len(p.colors)),
	}
	for i := range f.ansi {
		var known *tint.Color
		if tg.ANSI != nil {
			known = tint.FromColor(tg.ANSI[i])
		}
		switch own := p.t.ANSI(tint.Slot(i)); { //nolint:gosec // 0-15
		case known != nil:
			f.ansi[i] = *known
		case own != nil:
			f.ansi[i] = opaque(*own)
		default:
			// No better guess than the common defaults.
			f.ansi[i] = *tint.FromColor(ansi.BasicColor(i)) //nolint:gosec // 0-15
		}
	}
	switch tg.Depth {
	case TrueColor:
		for i, c := range p.colors {
			f.colors[i] = c
		}
	case Colors256:
		f.fit256()
	case Colors16:
		// The ground first, against the background, before anything is
		// drawn on it. Tones come before their on colors, which read them.
		f.under = p.colors[Bg]
		g := p.ground
		f.colors[g], f.attrs[g] = f.fit16(g)
		if s, ok := f.colors[g].(ansi.BasicColor); ok && g != Bg {
			f.under, f.painted = f.ansi[s], true
		}
		for i := range p.colors {
			if r := Role(i); r != g { //nolint:gosec // bounded by the number of roles
				f.colors[i], f.attrs[i] = f.fit16(r)
			}
		}
	case NoColor:
		for i := range p.colors {
			f.attrs[i] = noColorAttr(Role(i)) //nolint:gosec // bounded by the number of roles
		}
	}
	if tg.Unpainted {
		f.colors[Bg] = nil
		if p.ground == Bg {
			f.colors[Fg] = nil
		}
	}
	return f
}

// Color returns a role's color, fitted to the terminal: a [tint.Color], an
// [ansi.IndexedColor] or [ansi.BasicColor], or nil for the terminal's
// default. Renderers that know the ansi types, such as Lip Gloss and Bubble
// Tea, draw them as palette indexes, so the terminal's own 16 colors show
// through.
func (f *Fitted) Color(r Role) color.Color {
	if int(r) < len(f.colors) {
		return f.colors[r]
	}
	c, _ := f.fitLate(r)
	return c
}

// Attr returns the attributes to add to a role's color, where color can't
// carry the role. For a tone's [Role.On], it's what to add to a fill of the
// tone.
func (f *Fitted) Attr(r Role) Attr {
	if int(r) < len(f.attrs) {
		return f.attrs[r]
	}
	_, a := f.fitLate(r)
	return a
}

// fitLate fits a role defined after the palette was derived.
func (f *Fitted) fitLate(r Role) (color.Color, Attr) {
	switch f.target.Depth {
	case TrueColor:
		return f.p.Color(r), 0
	case Colors256:
		var taken []uint8
		if s, ok := specOf(r); ok {
			for _, o := range s.Distinct {
				if int(o) < len(f.idx) {
					taken = append(taken, f.idx[o])
				}
			}
		}
		return ansi.IndexedColor(f.nearest(f.p.Color(r), r, taken)), 0
	case Colors16:
		return f.fit16(r)
	case NoColor:
	}
	return nil, noColorAttr(r)
}

// On returns the palette for drawing on surface, fitted to the same
// terminal. See [Palette.On].
func (f *Fitted) On(surface Role) *Fitted { return f.p.On(surface).Fit(f.target) }

// Fit fits a color that isn't a role, such as a gradient's step or a mix of
// two roles, as if it were drawn like the role like. With 16 colors or none,
// it's like's color.
func (f *Fitted) Fit(c color.Color, like Role) color.Color {
	tc := tint.FromColor(c)
	if tc == nil {
		return nil
	}
	switch f.target.Depth {
	case TrueColor:
		return *tc
	case Colors256:
		return ansi.IndexedColor(f.nearest(*tc, like, nil))
	case Colors16, NoColor:
	}
	return f.Color(like)
}

// Depth returns the depth the palette was fitted to.
func (f *Fitted) Depth() Depth { return f.target.Depth }

// Palette returns the palette that was fitted.
func (f *Fitted) Palette() *Palette { return f.p }

// ---- 256 colors ----

// pairs are roles that sit beside each other, so must never share an index.
var pairs = [...][2]Role{
	{Bg, Surface},
	{Surface, Raised},
	{Raised, Overlay},
	{Bg, Hover},
	{Bg, Sunken},
	{Bg, Selection},
	{Hover, Selection},
	{Fg, Muted},
	{Muted, Subtle},
	{Subtle, Border},
	{Bg, Border},
	{Success, Warning},
	{Warning, Danger},
	{Accent, Accent2},
	{Bg, Accent.Soft()},
	{Bg, Accent2.Soft()},
	{Bg, Success.Soft()},
	{Bg, Warning.Soft()},
	{Bg, Danger.Soft()},
	{Bg, Info.Soft()},
	{Danger.Soft(), Warning.Soft()},
}

func (f *Fitted) fit256() {
	p := f.p
	partners := make(map[Role][]Role, len(pairs))
	link := func(a, b Role) {
		partners[a] = append(partners[a], b)
		partners[b] = append(partners[b], a)
	}
	for _, pr := range pairs {
		link(pr[0], pr[1])
	}
	for i, s := range defined() {
		for _, o := range s.Distinct {
			link(numBuiltin+Role(i), o) //nolint:gosec // bounded by the number of roles
		}
	}

	// Surfaces first, so the roles drawn on them are checked against their
	// fitted colors.
	order := make([]Role, len(p.colors))
	for i := range order {
		order[i] = Role(i) //nolint:gosec // bounded by the number of roles
	}
	slices.SortStableFunc(order, func(a, b Role) int {
		return cmp.Compare(f.drawnOn(a), f.drawnOn(b))
	})

	f.idx = make([]uint8, len(p.colors))
	for _, r := range order {
		var taken []uint8
		for _, o := range partners[r] {
			if int(o) < len(f.idx) && f.idx[o] >= 16 {
				taken = append(taken, f.idx[o])
			}
		}
		f.idx[r] = f.nearest(p.colors[r], r, taken)
		f.colors[r] = ansi.IndexedColor(f.idx[r])
	}
}

// drawnOn is 0 for what's drawn on (surfaces, soft fills, and the palette's
// ground), and 1 for what's drawn on them.
func (f *Fitted) drawnOn(r Role) int {
	if k := r.Kind(); k == KindSurface || k == KindSoft || r == f.p.ground {
		return 0
	}
	return 1
}

// floor256 is the contrast r keeps with 256 colors, and the luminance of what
// it's drawn on: text against the ground, a tone's on color against the
// tone. A floor of 0 is none.
func (f *Fitted) floor256(r Role) (floor, under float64) {
	p := f.p
	if !p.guards || r == p.ground {
		return 0, 0
	}
	against := p.ground
	switch r.Kind() {
	case KindText, KindTone:
		floor = minFor(r)
	case KindOn:
		floor, against = minColor, r.Base()
	case KindSurface, KindSoft:
	}
	switch {
	case floor == 0:
		return 0, 0
	case against == Bg && f.target.Unpainted:
		// The terminal draws its own background, not the fitted one.
		return floor, tint.Luminance(p.colors[Bg])
	case int(against) < len(f.idx) && f.idx[against] >= 16:
		return floor, lum256[f.idx[against]-16]
	}
	return 0, 0
}

// nearest is the 256-color index nearest c that keeps r's contrast against
// the ground, if it has a floor, and isn't taken. If none of the nearest
// few qualify, it's the nearest.
func (f *Fitted) nearest(c tint.Color, r Role, taken []uint8) uint8 {
	lab := c.OKLab()
	var cand [240]struct {
		i uint8
		d float64
	}
	for i := range cand {
		q := &lab256[i]
		cand[i].i = uint8(i + 16) //nolint:gosec // 16-255
		cand[i].d = (q.L-lab.L)*(q.L-lab.L) + (q.A-lab.A)*(q.A-lab.A) + (q.B-lab.B)*(q.B-lab.B)
	}
	slices.SortFunc(cand[:], func(a, b struct {
		i uint8
		d float64
	},
	) int {
		return cmp.Compare(a.d, b.d)
	})

	floor, under := f.floor256(r)
	reads := func(i uint8) bool { return floor == 0 || ratio(lum256[i-16], under) >= floor }
	// The nearest few that read and aren't taken, or failing that, the
	// nearest that reads, as a role that can't be read is worse than one that
	// looks like its neighbor.
	for _, k := range cand[:40] {
		if reads(k.i) && !slices.Contains(taken, k.i) {
			return k.i
		}
	}
	for i := range cand {
		if k := &cand[i]; reads(k.i) {
			return k.i
		}
	}
	return cand[0].i
}

func ratio(x, y float64) float64 {
	if x < y {
		x, y = y, x
	}
	return (x + 0.05) / (y + 0.05)
}

// lab256 and lum256 are the fixed part of the 256-color palette, from index
// 16 (the 6×6×6 cube and the gray ramp), in OKLab and as luminance.
var lab256, lum256 = func() (lab [240]tint.OKLab, lum [240]float64) {
	for i := range lab {
		c := tint.FromColor(ansi.IndexedColor(i + 16)) //nolint:gosec // 16-255
		lab[i], lum[i] = c.OKLab(), c.Luminance()
	}
	return lab, lum
}()

// ---- 16 colors ----

var grays = []tint.Slot{tint.SlotBrightBlack, tint.SlotWhite, tint.SlotBlack, tint.SlotBrightWhite}

// toneSlot is the slot a tone falls back to.
func (f *Fitted) toneSlot(r Role) tint.Slot {
	switch r { //nolint:exhaustive // only tones have a slot
	case Accent:
		return f.p.accentSlot
	case Accent2:
		return f.p.accent2Slot
	case Success:
		return tint.SlotGreen
	case Warning:
		return tint.SlotYellow
	case Danger:
		return tint.SlotRed
	case Info:
		return tint.SlotCyan
	case Red:
		return tint.SlotRed
	case Green:
		return tint.SlotGreen
	case Yellow:
		return tint.SlotYellow
	case Blue:
		return tint.SlotBlue
	case Magenta:
		return tint.SlotMagenta
	case Cyan:
		return tint.SlotCyan
	}
	return tint.SlotWhite
}

// fit16 is a role's slot (nil for the terminal's default) and attributes,
// with 16 colors.
func (f *Fitted) fit16(r Role) (color.Color, Attr) {
	p := f.p
	switch k := r.Kind(); k {
	case KindSurface:
		if r == Selection {
			a := p.accentSlot
			// Any accent slot that shows against the background.
			if s, ok := f.choose(p.colors[Bg], []tint.Slot{a.Bright(), a}, p.Color(r), 1.5); ok {
				return s, 0
			}
			return nil, Reverse
		}
		if s, ok := specOf(r); ok {
			return f.fallback(s, p.colors[Bg], p.Color(r), 0)
		}
		return nil, 0
	case KindSoft:
		return nil, 0
	case KindOn:
		s, ok := f.colors[r.Base()].(ansi.BasicColor)
		if !ok {
			return nil, Reverse | Bold
		}
		return f.blackOrWhite(f.ansi[s]), 0
	case KindText, KindTone:
	}

	ground, painted := f.under, f.painted
	floor := 0.0
	if p.guards {
		floor = minFor(r)
	}
	if s, ok := specOf(r); ok {
		return f.fallback(s, ground, p.Color(r), floor)
	}
	var family []tint.Slot
	var attr Attr
	switch {
	case r == Fg:
	case r == Muted || r == Subtle:
		// A gray that looks like body text wouldn't read as secondary; faint
		// body text does.
		family, attr = slices.DeleteFunc(slices.Clone(grays), func(s tint.Slot) bool {
			return f.ansi[s].Distance(p.colors[Fg]) < 0.04
		}), Faint
	case r == Border || r == BorderStrong:
		family, floor = grays, pickBool(p.guards, minBorder, 0)
	case r.Kind() == KindTone:
		s := f.toneSlot(r)
		family, attr = []tint.Slot{s.Bright(), s}, Bold
	}
	if s, ok := f.choose(ground, family, p.Color(r), floor); ok {
		return s, 0
	}
	if painted {
		// On a filled row the terminal's default text may not read.
		return f.blackOrWhite(ground), attr
	}
	return nil, attr
}

// blackOrWhite is whichever of the black and bright white slots reads better
// on fill.
func (f *Fitted) blackOrWhite(fill tint.Color) ansi.BasicColor {
	if tint.Contrast(f.ansi[tint.SlotBlack], fill) >= tint.Contrast(f.ansi[tint.SlotBrightWhite], fill) {
		return ansi.BasicColor(tint.SlotBlack)
	}
	return ansi.BasicColor(tint.SlotBrightWhite)
}

// fallback fits a defined role with 16 colors.
func (f *Fitted) fallback(s Spec, ground, want tint.Color, floor float64) (color.Color, Attr) {
	slots := s.Fallback.Slots
	if s.Fallback.Accent {
		a := f.p.accentSlot
		slots = append(slices.Clip(slots), a.Bright(), a)
	}
	if len(slots) == 0 {
		return nil, s.Fallback.Attr
	}
	if c, ok := f.choose(ground, slots, want, floor); ok {
		return c, 0
	}
	return nil, s.Fallback.Attr
}

// choose is the slot in family nearest want that contrasts with ground by
// floor.
func (f *Fitted) choose(ground tint.Color, family []tint.Slot, want tint.Color, floor float64) (color.Color, bool) {
	best, bestD := -1, math.MaxFloat64
	for _, s := range family {
		c := f.ansi[s]
		if floor > 0 && tint.Contrast(c, ground) < floor {
			continue
		}
		if d := c.Distance(want); d < bestD {
			best, bestD = int(s), d
		}
	}
	if best < 0 {
		return nil, false
	}
	return ansi.BasicColor(best), true //nolint:gosec // 0-15
}

// ---- no color ----

func noColorAttr(r Role) Attr {
	switch r { //nolint:exhaustive // the rest go by kind
	case Muted, Subtle:
		return Faint
	case Selection:
		return Reverse
	}
	switch r.Kind() {
	case KindTone:
		return Bold
	case KindOn:
		return Reverse | Bold
	case KindSurface, KindText, KindSoft:
	}
	if s, ok := specOf(r); ok {
		return s.Fallback.Attr
	}
	return 0
}
