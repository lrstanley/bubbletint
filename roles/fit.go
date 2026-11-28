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
	"sync"

	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Attr is a set of text attributes, used where color can't carry a role: a
// [Muted] hint is faint where no gray reads, and a selected row is reversed
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
	// Profile is how many colors the terminal can draw, such as Bubble Tea
	// detects and sends as a tea.ColorProfileMsg. Unknown is taken as
	// TrueColor, and NoTTY as ASCII: no color, only attributes. See
	// [ColorLimit] to cap it at a user's setting.
	Profile colorprofile.Profile
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
	colors []color.Color  // by role at rest
	attrs  []Attr
	idx    []uint8 // with 256 colors, each role's index, from 16; 0 is not yet fitted
	ground fill    // the ground, as the palette it was drawn on fitted it

	// With 16 colors, what text is drawn on: the slot filling the ground, or
	// the background, and whether it's filled.
	under   tint.Color
	painted bool
	mono    bool

	root   *Fitted  // the palette Fit returned, or nil for this one
	parent *Fitted  // the palette On was called on, or nil for this one
	on     sync.Map // On, by surface
	states sync.Map // roles in a state, as fills
}

// fill is a role fitted as a fill: its color, what to add where color can't
// carry it, and its 256-color index, if it has one.
type fill struct {
	c   color.Color
	a   Attr
	idx uint8
}

// Fit fits the palette to a terminal.
//
// With 256 colors, each role takes the nearest color that keeps its contrast
// against what it's drawn on, and that differs from the roles it sits beside
// (a dialog from the background, a hint from muted text, a row under the
// pointer from the row at rest), so no two collapse into one. With 16
// colors, each role takes a slot from its family (grays for muted text, red's
// two slots for danger) that reads on the terminal's palette, or the
// terminal's default plus an [Attr].
func (p *Palette) Fit(tg Target) *Fitted {
	if p.parent != nil {
		return fit(p, tg, p.parent.Fit(tg))
	}
	return fit(p, tg, nil)
}

func normalize(tg Target) Target {
	switch tg.Profile { //nolint:exhaustive // the rest are as given
	case colorprofile.Unknown:
		tg.Profile = colorprofile.TrueColor
	case colorprofile.NoTTY:
		tg.Profile = colorprofile.ASCII
	}
	return tg
}

// absolute are the roles that are the same on every ground, which a palette
// for drawing on a surface takes as the palette it was drawn on fitted them.
var absolute = [...]Role{Bg, Surface, Raised, Overlay}

// isLayer reports whether r is one of the layers, in any state: the same on
// every ground, and so the same to draw on wherever it's reached from.
func isLayer(r Role) bool { return slices.Contains(absolute[:], r.Rest()) }

// fit fits p, to draw on parent's fill of p's ground, or on the background
// for a palette Derive returned.
func fit(p *Palette, tg Target, parent *Fitted) *Fitted {
	f := &Fitted{
		p:      p,
		target: normalize(tg),
		colors: make([]color.Color, len(p.colors)),
		attrs:  make([]Attr, len(p.colors)),
	}
	preset := make([]bool, len(p.colors))
	if parent != nil {
		f.root, f.parent = parent.rootOf(), parent
		f.ansi = parent.ansi
		f.ground = parent.fillOf(p.ground)
		f.mono = parent.mono
		root := f.root
		if root.idx != nil {
			f.idx = make([]uint8, len(p.colors))
		}
		for _, r := range absolute {
			f.colors[r], f.attrs[r], preset[r] = root.colors[r], root.attrs[r], true
			if f.idx != nil {
				f.idx[r] = root.idx[r]
			}
		}
	} else {
		f.ansi = ansiOf(p, f.target)
	}

	switch f.target.Profile { //nolint:exhaustive // Unknown and NoTTY are mapped by normalize
	case colorprofile.TrueColor:
		for i, c := range p.colors {
			f.colors[i] = c
		}
	case colorprofile.ANSI256:
		f.fit256(preset)
	case colorprofile.ANSI:
		f.under = p.colors[Bg]
		if s, ok := f.ground.c.(ansi.BasicColor); ok && parent != nil {
			f.under, f.painted = f.ansi[s], true
		}
		for i := range p.colors {
			if !preset[i] {
				f.colors[i], f.attrs[i] = f.fit16(Role(i)) //nolint:gosec // bounded by the number of roles
			}
		}
	default:
		for i := range p.colors {
			f.attrs[i] = noColorAttr(Role(i)) //nolint:gosec // bounded by the number of roles
		}
	}
	if f.target.Unpainted {
		f.colors[Bg] = nil
		if p.ground == Bg {
			f.colors[Fg] = nil
		}
	}
	if parent == nil {
		f.ground = fill{c: f.colors[Bg]}
		if f.idx != nil {
			f.ground.idx = f.idx[Bg]
		}
		f.mono = f.detectMono()
	}
	for _, r := range [...]Role{Success, Warning, Danger} {
		f.attrs[r] = f.marked(r, f.attrs[r])
	}
	return f
}

// marked is a, the attributes of the status tone r as text or a fill, with
// what tells status apart where hue can't in place of the weight any tone
// takes without color: weight and an underline for danger, an underline for
// warning, and weight for success.
func (f *Fitted) marked(r Role, a Attr) Attr {
	if !f.mono {
		return a
	}
	a &^= Bold | Underline
	switch r { //nolint:exhaustive // only status is marked
	case Danger:
		return a | Bold | Underline
	case Warning:
		return a | Underline
	case Success:
		return a | Bold
	}
	return a
}

// ansiOf is the terminal's 16 colors: as given, else the tint's, else the
// common defaults.
func ansiOf(p *Palette, tg Target) (out [16]tint.Color) {
	for i := range out {
		var known *tint.Color
		if tg.ANSI != nil {
			known = tint.FromColor(tg.ANSI[i])
		}
		switch own := p.t.ANSI(tint.Slot(i)); { //nolint:gosec // 0-15
		case known != nil:
			out[i] = *known
		case own != nil:
			out[i] = opaque(*own)
		default:
			// No better guess than the common defaults.
			out[i] = *tint.FromColor(ansi.BasicColor(i)) //nolint:gosec // 0-15
		}
	}
	return out
}

func (f *Fitted) rootOf() *Fitted {
	if f.root != nil {
		return f.root
	}
	return f
}

// fillOf is r fitted as a fill.
func (f *Fitted) fillOf(r Role) fill {
	if r&stateMask != 0 {
		return f.state(r)
	}
	out := fill{c: f.Color(r), a: f.Attr(r)}
	if r.Kind() == KindTone {
		out.a = f.marked(r, f.toneFill(r))
	}
	if i, ok := out.c.(ansi.IndexedColor); ok {
		out.idx = uint8(i)
	}
	return out
}

// toneFill is what a fill of the tone r adds where its color can't be drawn.
func (f *Fitted) toneFill(r Role) Attr {
	switch f.target.Profile { //nolint:exhaustive // the rest draw it
	case colorprofile.ANSI:
		if _, ok := f.Color(r).(ansi.BasicColor); ok {
			return 0
		}
		return Reverse | Bold
	case colorprofile.ASCII:
		return Reverse | Bold
	}
	return 0
}

// detectMono reports whether the status tones, as fitted, can't be told
// apart by hue.
func (f *Fitted) detectMono() bool {
	return mono(f.colorOf(f.colors[Success]), f.colorOf(f.colors[Warning]), f.colorOf(f.colors[Danger]))
}

// colorOf is the color the terminal draws for c, as best known.
func (f *Fitted) colorOf(c color.Color) tint.Color {
	switch v := c.(type) {
	case nil:
		return f.p.termFg
	case ansi.BasicColor:
		return f.ansi[v&15]
	}
	if tc := tint.FromColor(c); tc != nil {
		return *tc
	}
	return f.p.termFg
}

// Color returns a role's color, fitted to the terminal: a [tint.Color], an
// [ansi.IndexedColor] or [ansi.BasicColor], or nil for the terminal's
// default. Renderers that know the ansi types, such as Lip Gloss, draw them
// as palette indexes, so the terminal's own 16 colors show through. Add
// [Fitted.Attr] where color can't carry the role, or use [Fitted.Text] and
// [Fitted.Fill], which add it.
func (f *Fitted) Color(r Role) color.Color {
	if f.root != nil && r == f.p.ground && r.Kind() != KindText {
		return f.ground.c
	}
	if r&stateMask != 0 {
		return f.state(r).c
	}
	if int(r) < len(f.colors) {
		return f.colors[r]
	}
	c, _ := f.fitLate(r)
	return c
}

// Attr returns the attributes to add to a role's color, where color can't
// carry the role. A role in a state is a fill, and its attributes are what
// marks the state, such as [Underline] for a row under the pointer with 16
// colors.
func (f *Fitted) Attr(r Role) Attr {
	if r&stateMask != 0 {
		return f.state(r).a
	}
	if int(r) < len(f.attrs) {
		return f.attrs[r]
	}
	_, a := f.fitLate(r)
	return a
}

// fitLate fits a role defined after the palette was derived.
func (f *Fitted) fitLate(r Role) (color.Color, Attr) {
	switch f.target.Profile { //nolint:exhaustive // Unknown and NoTTY are mapped by Fit
	case colorprofile.TrueColor:
		return f.p.Color(r), 0
	case colorprofile.ANSI256:
		var taken []uint8
		if s, ok := specOf(r); ok {
			for _, o := range s.Distinct {
				taken = f.take(taken, o)
			}
			if s.Kind == KindSurface {
				taken = f.take(taken, groundMark)
			}
		}
		floor, under := f.floor256(r)
		return ansi.IndexedColor(f.nearest(f.p.Color(r), floor, under, taken)), 0
	case colorprofile.ANSI:
		return f.fit16(r)
	}
	return nil, noColorAttr(r)
}

// Distinct reports whether a and b look different on the terminal: in color,
// or in the attributes that stand in for it. Use it to decide on a design,
// such as outlining a dialog where its surface wouldn't show against the
// page: f.Distinct(roles.Raised, roles.Bg) is false with flat surfaces, with
// 16 colors, and with none.
func (f *Fitted) Distinct(a, b Role) bool {
	return f.Color(a) != f.Color(b) || f.Attr(a) != f.Attr(b)
}

// Profile returns how many colors the palette was fitted to: never Unknown
// or NoTTY, which [Palette.Fit] takes as TrueColor and ASCII.
func (f *Fitted) Profile() colorprofile.Profile { return f.target.Profile }

// Target returns the terminal the palette was fitted to.
func (f *Fitted) Target() Target { return f.target }

// Ground returns the surface the palette draws on: [Bg], or the surface
// given to [Fitted.On].
func (f *Fitted) Ground() Role { return f.p.ground }

// Mono reports whether the status tones, as fitted, can't be told apart by
// hue, so [Success], [Warning] and [Danger] are each marked with attributes,
// as text and as fills: on a tint whose hues sit close together, or with no
// color.
func (f *Fitted) Mono() bool { return f.mono }

// On returns the palette for drawing on surface, fitted to the same terminal.
// See [Palette.On]. Its [Fitted.Text] styles add the attributes surface's
// fill needs, and [Fitted.Span] draws on surface's color. A palette for
// drawing on a layer, or on anything else that's the same wherever it's
// drawn, is shared; it's worked out once per surface, so it's cheap to call
// while rendering.
func (f *Fitted) On(surface Role) *Fitted {
	if surface == f.p.ground {
		return f
	}
	root := f.rootOf()
	if surface == Bg {
		return root
	}
	// Only *Fitted is stored. A miss is fitted outside any lock, and the
	// first to store it wins.
	if v, ok := f.on.Load(surface); ok {
		on, _ := v.(*Fitted)
		return on
	}
	var on *Fitted
	if f != root && isLayer(surface) {
		on = root.On(surface) // the same as on the page, so shared with it
	} else {
		on = fit(f.p.On(surface), f.target, f)
	}
	v, _ := f.on.LoadOrStore(surface, on)
	on, _ = v.(*Fitted)
	return on
}

// Fit fits a color that isn't a role, such as a gradient's step or a mix of
// two roles, as if it were drawn like the role like. With 256 colors, it's
// kept apart from like's color, unless it is like's color, and from the
// roles in apart. With 16 colors or none, it's like's color.
func (f *Fitted) Fit(c color.Color, like Role, apart ...Role) color.Color {
	tc := tint.FromColor(c)
	if tc == nil {
		return nil
	}
	switch f.target.Profile { //nolint:exhaustive // the rest are like's color
	case colorprofile.TrueColor:
		return *tc
	case colorprofile.ANSI256:
		var taken []uint8
		if tc.Distance(f.p.Color(like)) > 0.004 {
			taken = f.take(taken, like)
		}
		for _, r := range apart {
			taken = f.take(taken, r)
		}
		floor, under := f.floor256(like)
		return ansi.IndexedColor(f.nearest(*tc, floor, under, taken))
	}
	return f.Color(like)
}

// Palette returns the palette that was fitted.
func (f *Fitted) Palette() *Palette { return f.p }

// Lerp returns a palette partway from a to b, fitted to b's terminal, for
// crossfading between tints: t runs from 0 (a) to 1 (b). See
// [LerpPalette]. Below 256 colors there's nothing between, so it's b, as it
// is when a and b were fitted to different terminals. For a palette from
// [Fitted.On], it's the crossfade's palette for drawing on the same surfaces,
// in the same order.
func Lerp(a, b *Fitted, t float64) *Fitted {
	switch {
	case a == nil || b == nil || t >= 1:
		return b
	case t <= 0:
		return a
	case a.target != b.target || b.target.Profile < colorprofile.ANSI256:
		return b
	}
	out := LerpPalette(a.p, b.p, t).Fit(b.target)
	var grounds []Role
	for g := b; g.parent != nil; g = g.parent {
		grounds = append(grounds, g.p.ground)
	}
	for _, g := range slices.Backward(grounds) {
		out = out.On(g)
	}
	return out
}

// ---- states ----

// state is r, in a state, fitted as a fill. It's worked out once per role.
func (f *Fitted) state(r Role) fill {
	if v, ok := f.states.Load(r); ok {
		s, _ := v.(fill)
		return s
	}
	v, _ := f.states.LoadOrStore(r, f.fitState(r))
	s, _ := v.(fill)
	return s
}

// stateAttr is what marks r's states where color can't.
func (f *Fitted) stateAttr(r Role) Attr {
	var a Attr
	if r&selectedBit != 0 {
		a |= f.p.stateAttr[0]
	}
	if r&hoverBit != 0 {
		a |= f.p.stateAttr[1]
	}
	if r&pressedBit != 0 {
		a |= f.p.stateAttr[2]
	}
	return a
}

// beside are the roles r, in a state, sits beside: its role at rest, a row
// under the pointer beside a selected one, and the state it steps from.
func beside(r Role) []Role {
	rest := r.Rest()
	out := []Role{rest}
	if r&selectedBit != 0 {
		out = append(out, rest.Hover())
		if r&(hoverBit|pressedBit) != 0 {
			out = append(out, rest.Selected())
		}
	}
	if r&pressedBit != 0 {
		out = append(out, r.Hover())
	}
	return out
}

func (f *Fitted) fitState(r Role) fill {
	p, rest := f.p, r.Rest()
	switch f.target.Profile { //nolint:exhaustive // Unknown and NoTTY are mapped by Fit
	case colorprofile.TrueColor:
		return fill{c: p.Color(r), a: f.marked(rest, 0)}
	case colorprofile.ANSI256:
		var taken []uint8
		for _, o := range beside(r) {
			taken = f.take(taken, o)
		}
		i := f.nearest(p.Color(r), 0, 0, taken)
		return fill{c: ansi.IndexedColor(i), a: f.marked(rest, 0), idx: i}
	case colorprofile.ANSI:
		base := f.fillOf(rest)
		out := fill{c: base.c, a: base.a}
		if r&selectedBit != 0 {
			under := p.colors[Bg]
			if s, ok := base.c.(ansi.BasicColor); ok {
				under = f.ansi[s]
			}
			a := p.accentSlot
			if s, ok := f.choose(under, []tint.Slot{a.Bright(), a}, p.Color(rest.Selected()), 1.5); ok && s != base.c {
				out.c = s
			} else {
				out.a |= p.stateAttr[0]
			}
		}
		out.a |= f.stateAttr(r &^ selectedBit)
		return out
	}
	return fill{a: f.fillOf(rest).a | f.stateAttr(r)}
}

// ---- 256 colors ----

// groundMark stands for the ground among the roles one is kept apart from.
const groundMark Role = 1 << 15

// pairs are roles that sit beside each other, so must never share an index.
// A dialog or menu floats over the page as well as over a panel. Text, lines,
// wells and fills that step off the ground are kept apart from it.
var pairs = [...][2]Role{
	{Bg, Surface},
	{Surface, Raised},
	{Raised, Overlay},
	{Bg, Raised},
	{Bg, Overlay},
	{Surface, Overlay},
	{groundMark, Sunken},
	{groundMark, Control},
	{Fg, Muted},
	{Muted, Subtle},
	{Subtle, Border},
	{groundMark, Border},
	{Success, Warning},
	{Warning, Danger},
	{Accent, Accent2},
	{groundMark, Neutral.Soft()},
	{groundMark, Accent.Soft()},
	{groundMark, Accent2.Soft()},
	{groundMark, Success.Soft()},
	{groundMark, Warning.Soft()},
	{groundMark, Danger.Soft()},
	{groundMark, Info.Soft()},
	{Danger.Soft(), Warning.Soft()},
}

// take adds o's 256-color index to taken, if it has one: the ground's for
// groundMark.
func (f *Fitted) take(taken []uint8, o Role) []uint8 {
	switch {
	case o == groundMark, f.root != nil && o == f.p.ground && o.Kind() != KindText:
		if i := f.groundIdx(); i >= 16 {
			return append(taken, i)
		}
	case o&stateMask != 0:
		if i := f.state(o).idx; i >= 16 {
			return append(taken, i)
		}
	case int(o) < len(f.idx) && f.idx[o] >= 16:
		return append(taken, f.idx[o])
	}
	return taken
}

// groundIdx is the ground's 256-color index, or 0 if it hasn't one yet.
func (f *Fitted) groundIdx() uint8 {
	if f.root == nil {
		if int(f.p.ground) < len(f.idx) {
			return f.idx[f.p.ground]
		}
		return 0
	}
	return f.ground.idx
}

func (f *Fitted) fit256(preset []bool) {
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
		r := firstDefined + Role(i) //nolint:gosec // bounded by the number of roles
		for _, o := range s.Distinct {
			if o&stateMask == 0 {
				link(r, o)
			}
		}
		if s.Kind == KindSurface {
			link(r, groundMark)
		}
	}

	// Surfaces first, so the roles drawn on them are checked against their
	// fitted colors.
	order := make([]Role, 0, len(p.colors))
	for i := range p.colors {
		if !preset[i] {
			order = append(order, Role(i)) //nolint:gosec // bounded by the number of roles
		}
	}
	slices.SortStableFunc(order, func(a, b Role) int {
		return cmp.Compare(f.drawnOn(a), f.drawnOn(b))
	})

	if f.idx == nil {
		f.idx = make([]uint8, len(p.colors))
	}
	for _, r := range order {
		var taken []uint8
		for _, o := range partners[r] {
			if o != r {
				taken = f.take(taken, o)
			}
		}
		floor, under := f.floor256(r)
		f.idx[r] = f.nearest(p.colors[r], floor, under, taken)
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

// floor256 is the contrast r keeps with 256 colors, and the luminance of the
// ground it's drawn on. A floor of 0 is none.
func (f *Fitted) floor256(r Role) (floor, under float64) {
	p := f.p
	if !p.guards || r == p.ground {
		return 0, 0
	}
	if k := r.Kind(); k != KindText && k != KindTone {
		return 0, 0
	}
	floor = p.floor(r)
	switch i := f.groundIdx(); {
	case floor == 0:
		return 0, 0
	case p.ground == Bg && f.target.Unpainted:
		// The terminal draws its own background, not the fitted one.
		return floor, tint.Luminance(p.groundColor)
	case i >= 16:
		return floor, lum256[i-16]
	}
	return 0, 0
}

// nearest is the 256-color index nearest c that contrasts with a ground of
// luminance under by floor, if floor isn't 0, and isn't taken. If none of the
// nearest few qualify, it's the nearest that reads, as a role that can't be
// read is worse than one that looks like its neighbor, and failing that, the
// nearest.
//
// It runs for every role on every fit, and so on every frame of a crossfade,
// so it scans the palette rather than sorting it.
func (f *Fitted) nearest(c tint.Color, floor, under float64, taken []uint8) uint8 {
	// How far down the list, by distance, a role may go to stay distinct.
	const few = 40

	lab := c.OKLab()
	var dist [240]float64
	// The nearest index, the nearest that reads, and the nearest that reads
	// and isn't taken, with their distances. 0 is none yet.
	var best, reads, free uint8
	bestD, readsD, freeD := math.Inf(1), math.Inf(1), math.Inf(1)
	for i := range lab256 {
		q, idx := &lab256[i], uint8(i+16) //nolint:gosec // 16-255
		d := (q.L-lab.L)*(q.L-lab.L) + (q.A-lab.A)*(q.A-lab.A) + (q.B-lab.B)*(q.B-lab.B)
		dist[i] = d
		if d < bestD {
			best, bestD = idx, d
		}
		if floor > 0 && ratio(lum256[i], under) < floor {
			continue
		}
		if d < readsD {
			reads, readsD = idx, d
		}
		if d < freeD && !slices.Contains(taken, idx) {
			free, freeD = idx, d
		}
	}
	if free != 0 {
		// Only if it's among the nearest few: count those nearer.
		nearer := 0
		for i := range dist {
			if dist[i] < freeD {
				nearer++
			}
		}
		if nearer < few {
			return free
		}
	}
	if reads != 0 {
		return reads
	}
	return best
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
	return tint.SlotBrightBlack
}

// fit16 is a role's slot (nil for the terminal's default) and attributes,
// with 16 colors.
func (f *Fitted) fit16(r Role) (color.Color, Attr) {
	p := f.p
	switch r.Kind() {
	case KindSurface:
		if s, ok := specOf(r); ok {
			return f.fallback(s, f.under, p.Color(r), 0)
		}
		if r == Control {
			// A gray that stands off what it's on, filled with text that
			// reads on it.
			if s, ok := f.choose(f.under, []tint.Slot{tint.SlotBrightBlack, tint.SlotWhite}, p.Color(r), 1.5); ok {
				return s, 0
			}
			return nil, Reverse
		}
		return nil, 0
	case KindSoft:
		return nil, 0
	case KindText, KindTone:
	}

	ground, painted := f.under, f.painted
	floor := 0.0
	if p.guards {
		floor = p.floor(r)
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
		// body text does. Body text is the terminal's default, or the slot
		// it's drawn in on a filled ground. Fg is fitted before Muted.
		family, attr = f.unlikeBody(), Faint
	case r == Border || r == BorderStrong:
		family, floor = grays, pickBool(p.guards, p.floor(Border), 0)
	case r == Neutral:
		family = f.unlikeBody()
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

// unlikeBody are the grays that don't look like body text.
func (f *Fitted) unlikeBody() []tint.Slot {
	body := f.p.termFg
	if s, ok := f.colors[Fg].(ansi.BasicColor); ok {
		body = f.ansi[s]
	}
	return slices.DeleteFunc(slices.Clone(grays), func(s tint.Slot) bool {
		return f.ansi[s].Distance(body) < 0.04
	})
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
	case Control:
		return Reverse
	case Neutral:
		return 0
	}
	if r.Kind() == KindTone {
		return Bold
	}
	if s, ok := specOf(r); ok {
		return s.Fallback.Attr
	}
	return 0
}
