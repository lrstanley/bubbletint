// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"image/color"
	"slices"

	tint "github.com/lrstanley/bubbletint/v2"
)

// Minimum contrast against what a role is drawn on, enforced by the guards, as
// WCAG contrast ratios (see [tint.Color.Contrast]). Only minText is a WCAG
// level; WCAG sets no minimum for hints, disabled text or decorative lines, so
// the rest are tuned against the built-in tints, to keep a tint's look while
// staying legible.
const (
	// minText is for body and muted text: WCAG 2.2 AA for normal text, success
	// criterion 1.4.3 (https://www.w3.org/WAI/WCAG22/Understanding/contrast-minimum).
	// Fg gets fgMargin more, so muted text can sit between it and the floor.
	minText  = 4.5
	fgMargin = 0.3
	// minColor is for accents and status colors used as text and marks. It's
	// WCAG's 3:1 for large text and UI components (success criterion 1.4.11,
	// https://www.w3.org/WAI/WCAG22/Understanding/non-text-contrast), with some
	// margin, as terminal text is never large.
	minColor = 3.5
	// minSubtle is for hints, placeholders and disabled text, which WCAG
	// exempts from 1.4.3. It's tuned to stay readable while plainly fainter
	// than muted text.
	minSubtle = 2.6
	// minBorder is for outlines and dividers, which WCAG exempts as
	// decorative: enough to see a line, short of competing with text.
	minBorder = 1.45
	// minSunken is how far an input's well stands off the background: just
	// visible as a separate area.
	minSunken = 1.1
	// minStep is how far, in OKLab (see [tint.Color.Distance]), a state must
	// move from its role at rest to show: about a just-noticeable difference.
	minStep = 0.02

	// dangerHueGap is how close, in OKLCh hue degrees, a slot can be to the
	// tint's red before it can't be the accent, so the accent doesn't read as
	// an error. Tuned: about the distance between a red and an orange.
	dangerHueGap = 22.0
)

// accentOrder is the order slots are tried in for the accent.
var accentOrder = [...]tint.Slot{tint.SlotBlue, tint.SlotMagenta, tint.SlotCyan, tint.SlotGreen, tint.SlotYellow, tint.SlotRed}

// Option changes how a palette is derived.
type Option func(*options)

type options struct {
	bg, fg      *tint.Color
	accentSlot  tint.Slot
	accentSet   bool
	accentColor *tint.Color
	flat        bool
	noGuards    bool
	stateAttr   [3]Attr
}

// WithTerminal derives against the terminal's own background and text, rather
// than the tint's. Terminals report them to OSC 11 and OSC 10 queries; Bubble
// Tea, for example, delivers them as tea.BackgroundColorMsg and
// tea.ForegroundColorMsg. Use it when the app leaves the terminal's background
// unpainted. Either may be nil, to keep the tint's.
func WithTerminal(bg, fg color.Color) Option {
	return func(o *options) { o.bg, o.fg = tint.FromColor(bg), tint.FromColor(fg) }
}

// WithAccentSlot uses one of the tint's chromatic slots as the accent, rather
// than letting the tint's hue choose it. Either variant of a slot may be
// given; the bright one is used unless it's grey.
func WithAccentSlot(s tint.Slot) Option {
	return func(o *options) { o.accentSlot, o.accentSet = s.Normal(), true }
}

// WithAccentColor uses a fixed color as the accent. With 16 colors it falls
// back to the slot nearest in hue. A nil color keeps the tint's.
func WithAccentColor(c color.Color) Option {
	return func(o *options) { o.accentColor = tint.FromColor(c) }
}

// WithFlatSurfaces makes [Surface], [Raised], [Overlay] and [Sunken] the
// background, for a flat look. Controls and states still show.
func WithFlatSurfaces() Option { return func(o *options) { o.flat = true } }

// WithoutGuards turns off the contrast guards: text, marks and lines keep
// the tint's colors even where they read poorly. It's meant for a hand-made
// tint whose author wants exactly their colors.
func WithoutGuards() Option { return func(o *options) { o.noGuards = true } }

// WithStateAttr sets the attributes that mark a role in state s where its
// color can't show the state: with 16 colors, where hover and pressed have
// no step to draw, or with none. The defaults are [Underline] for
// [StateHover], [Underline] and [Bold] for [StatePressed], and [Reverse] for
// [StateSelected]. s is a single state.
func WithStateAttr(s State, a Attr) Option {
	return func(o *options) {
		if i, ok := stateIndex(s); ok {
			o.stateAttr[i] = a
		}
	}
}

// defaultStateAttr is by state: selected, hover, pressed.
var defaultStateAttr = [3]Attr{Reverse, Underline, Underline | Bold}

func stateIndex(s State) (int, bool) {
	switch s {
	case StateSelected:
		return 0, true
	case StateHover:
		return 1, true
	case StatePressed:
		return 2, true
	}
	return 0, false
}

// Palette is every role's color for one tint, in true color, before it's
// fitted to a terminal with [Palette.Fit]. It's immutable, and safe to share
// between goroutines.
//
// A palette is for drawing on one surface, its [Palette.Ground]: the
// background, for the palette [Derive] returns, or the surface given to
// [Palette.On]. Text, lines, [Control], [Sunken] and soft fills are worked out
// from the ground; the layers ([Bg] to [Overlay]) are the same on every ground.
type Palette struct {
	t      *tint.Tint
	colors []tint.Color // by role at rest: the builtins, then the defined roles

	root, parent *Palette // nil for the palette Derive returns
	ground       Role
	groundColor  tint.Color

	pageFg       tint.Color // the background's text
	bgSelected   tint.Color // the background's selection, Bg.Selected()
	termFg       tint.Color // the terminal's default text color, as best known
	guards, dark bool
	flat         bool
	stateAttr    [3]Attr

	accentSlot, accent2Slot tint.Slot
}

// Derive works out every role's color for t. A tint's missing colors are
// filled in, so hand-made tints work too.
func Derive(t *tint.Tint, opts ...Option) *Palette {
	o := options{stateAttr: defaultStateAttr}
	for _, opt := range opts {
		opt(&o)
	}
	if t == nil {
		t = &tint.Tint{}
	}
	bg := orColor(o.bg, t.Bg, tint.Color{A: 0xff})
	p := &Palette{
		t:           t,
		colors:      make([]tint.Color, numBuiltin, int(numBuiltin)+len(defined())),
		ground:      Bg,
		groundColor: bg,
		guards:      !o.noGuards,
		flat:        o.flat,
		dark:        tint.IsDark(bg),
		stateAttr:   o.stateAttr,
	}
	white := tint.Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	fg := orColor(o.fg, t.Fg, pickBool(p.dark, white, tint.Color{A: 0xff}))
	p.termFg = fg
	d := deriver{p: p, t: t, o: &o, bg: bg}
	d.derive(fg)
	// Before relate, as defined roles may read Bg.Selected().
	p.bgSelected = p.pageSelection()
	p.relate()
	return p
}

type deriver struct {
	p  *Palette
	t  *tint.Tint
	o  *options
	bg tint.Color
}

func (d *deriver) set(r Role, c tint.Color) { d.p.colors[r] = c }

// lift is [tint.Color.EnsureContrast] against the background, when the
// guards are on.
func (d *deriver) lift(c tint.Color, ratio float64) tint.Color {
	if !d.p.guards {
		return c
	}
	return c.EnsureContrast(d.bg, ratio)
}

// derive works out the roles that are the same on every surface: the layers
// and the tones. The rest are worked out from the ground, by relate.
func (d *deriver) derive(fg tint.Color) {
	p, bg := d.p, d.bg
	fg = d.lift(fg, minFor(Fg))
	p.pageFg = fg
	d.set(Bg, bg)
	accent, accent2 := d.accents(fg)
	d.status()
	d.surfaces(fg)
	d.set(Neutral, d.lift(fg.Mix(bg, 0.45), minColor))
	d.set(Accent, d.lift(accent, minColor))
	d.set(Accent2, d.lift(accent2, minColor))
}

// pick is a slot's color: the bright variant, unless it's grey, as some
// tints' brights are (Solarized's are its neutrals).
func (d *deriver) pick(s tint.Slot) tint.Color {
	n, b := d.t.ANSI(s.Normal()), d.t.ANSI(s.Bright())
	switch {
	case b != nil && (n == nil || chroma(*b) >= 0.06):
		return opaque(*b)
	case n != nil:
		return opaque(*n)
	}
	return d.p.pageFg
}

// accents picks the accent, the tint's own hue unless it reads as danger,
// and a second accent at least 50 degrees away.
func (d *deriver) accents(fg tint.Color) (accent, accent2 tint.Color) {
	p, o := d.p, d.o
	order := accentOrder[:]
	if s, ok := tint.SlotOf(d.t.Hue); ok {
		order = append([]tint.Slot{s}, without(accentOrder[:], s)...)
	}
	found := false
	switch {
	case o.accentColor != nil:
		accent, p.accentSlot, found = *o.accentColor, d.nearestSlot(*o.accentColor), true
	case o.accentSet:
		accent, p.accentSlot, found = d.pick(o.accentSlot), o.accentSlot, true
	}
	dangerHue := hue(d.pick(tint.SlotRed))
	for _, s := range order {
		if found {
			break
		}
		c := d.pick(s)
		switch {
		case chroma(c) < 0.05:
		case p.guards && (s == tint.SlotRed || hueDist(hue(c), dangerHue) < dangerHueGap):
		default:
			accent, p.accentSlot, found = c, s, true
		}
	}
	if !found {
		p.accentSlot = order[0]
		accent = d.pick(p.accentSlot)
	}
	p.accent2Slot = p.accentSlot
	for _, s := range order {
		if s == p.accentSlot || s == tint.SlotRed {
			continue
		}
		if c := d.pick(s); chroma(c) >= 0.05 && hueDist(hue(c), hue(accent)) >= 50 {
			p.accent2Slot = s
			return accent, c
		}
	}
	return accent, accent.Mix(fg, 0.4)
}

// status sets the status and hue colors.
func (d *deriver) status() {
	success, danger, warning := d.pick(tint.SlotGreen), d.pick(tint.SlotRed), d.pick(tint.SlotYellow)
	d.set(Success, d.lift(success, minColor))
	d.set(Warning, d.lift(warning, minColor))
	d.set(Danger, d.lift(danger, minColor))
	d.set(Info, d.lift(d.pick(tint.SlotCyan), minColor))
	for _, s := range tint.ChromaticSlots {
		d.set(Hue(s), d.lift(d.pick(s), minColor))
	}
}

// surfaces are the background's shades: mixes toward the text, or toward
// black for a well on a dark tint.
func (d *deriver) surfaces(fg tint.Color) {
	p, bg := d.p, d.bg
	for _, r := range [...]Role{Surface, Raised, Overlay} {
		d.set(r, bg)
	}
	if d.o.flat {
		return
	}
	d.set(Surface, bg.Mix(fg, p.ByDark(0.05, 0.04)))
	raised := bg.Mix(fg, p.ByDark(0.095, 0.065))
	d.set(Raised, raised)
	// Raised's step again, smaller where text would stop reading on it.
	if p.guards {
		d.set(Overlay, raised.MixReadable(fg, fg, p.ByDark(0.095, 0.065), minColor))
	} else {
		d.set(Overlay, raised.Mix(fg, p.ByDark(0.095, 0.065)))
	}
}

// sunken is the well inputs and controls sit in, sunk into g: toward black
// on a dark ground, and toward the darker of the background and its text on
// a light one, as far as it takes to show. A dark ground already near black
// has no room below it, so the well rises toward the text instead, past a
// row under the pointer, so a control still shows on one.
func (p *Palette) sunken(g tint.Color) tint.Color {
	if p.flat {
		return g
	}
	if !g.IsDark() {
		darker := p.pageFg
		if p.colors[Bg].Luminance() < darker.Luminance() {
			darker = p.colors[Bg]
		}
		in := g.Mix(darker, 0.07)
		for amt := 0.08; amt < 0.3 && tint.Contrast(in, g) < minSunken; amt += 0.01 {
			in = g.Mix(darker, amt)
		}
		return in
	}
	if in := g.Mix(tint.Color{A: 0xff}, 0.32); tint.Contrast(in, g) >= minSunken {
		return in
	}
	hover := g.Mix(p.pageFg, p.hoverStep())
	in := hover
	for amt := hoverStepDark; amt < 0.4 && tint.Contrast(in, hover) < minSunken; amt += 0.01 {
		in = g.Mix(p.pageFg, amt)
	}
	return in
}

// nearestSlot is the chromatic slot closest in hue to c.
func (d *deriver) nearestSlot(c tint.Color) tint.Slot {
	best, bestD := tint.SlotBlue, 999.0
	for _, s := range accentOrder {
		if dist := hueDist(hue(d.pick(s)), hue(c)); dist < bestD {
			best, bestD = s, dist
		}
	}
	return best
}

// relate works out the roles that depend on the ground: text and lines,
// tones kept readable on it, soft fills, controls, and the defined roles.
func (p *Palette) relate() {
	root, g := p.rootOf(), p.groundColor
	fg := p.textOn(g)
	p.colors[Fg] = fg
	p.colors[Muted] = p.guard(fg.Mix(g, 0.36), g, p.floor(Muted))
	p.colors[Subtle] = p.guard(fg.Mix(g, 0.6), g, p.floor(Subtle))
	p.colors[Border] = p.guard(fg.Mix(g, 0.78), g, p.floor(Border))
	p.colors[BorderStrong] = fg.Mix(g, 0.5)
	for r := firstTone; r < numCore; r++ {
		tone := root.colors[r]
		if r != p.ground {
			p.colors[r] = p.guard(tone, g, p.floor(r))
		}
		amt := p.ByDark(0.17, 0.14)
		if r == Accent || r == Accent2 || r == Neutral {
			amt = p.ByDark(0.2, 0.16)
		}
		p.colors[r.Soft()] = g.Mix(tone, amt)
	}
	p.colors[Control] = p.control(g, fg)
	p.colors[Sunken] = p.sunken(g)
	p.colors = p.colors[:numBuiltin]
	for _, s := range defined() {
		p.colors = append(p.colors, p.custom(s))
	}
}

// guard is [tint.Color.EnsureContrast] against under, when the guards are on.
func (p *Palette) guard(c, under tint.Color, ratio float64) tint.Color {
	if !p.guards || ratio <= 0 {
		return c
	}
	return c.EnsureContrast(under, ratio)
}

// bestText is the background's text, or the background itself, whichever
// reads better on fill: the text on a surface, and on a tone's fill whichever
// stands out from it.
func (p *Palette) bestText(fill tint.Color) tint.Color {
	if p.colors[Bg].Contrast(fill) > p.pageFg.Contrast(fill) {
		return p.colors[Bg]
	}
	return p.pageFg
}

// textOn is body text for fill, kept readable on it by the guards.
func (p *Palette) textOn(fill tint.Color) tint.Color {
	return p.guard(p.bestText(fill), fill, p.floor(Fg))
}

// toneFloor scales the floors on a tone's fill, such as a pill: its text only
// has to read as well as the tone does on the page. Body text's floor there
// is the tone's own. Higher, light text on a mid-toned fill would be walked
// all the way to black.
const toneFloor = minColor / (minText + fgMargin)

// floor is the contrast r keeps against the palette's ground, when the
// guards are on: [minFor]'s, scaled by toneFloor on a tone's fill.
func (p *Palette) floor(r Role) float64 { return p.scaled(minFor(r)) }

func (p *Palette) scaled(m float64) float64 {
	if p.ground.Kind() == KindTone {
		return m * toneFloor
	}
	return m
}

// control is a neutral button on g: well clear of g, while its text still
// reads, and never so close that it vanishes, even where the tint's text
// barely reads.
func (p *Palette) control(g, fg tint.Color) tint.Color {
	amt := p.ByDark(0.22, 0.17)
	if !p.guards {
		return g.Mix(fg, amt)
	}
	c, least := g.MixReadable(fg, fg, amt, minText), g.Mix(fg, 0.08)
	if c.Contrast(g) < least.Contrast(g) {
		return least
	}
	return c
}

// custom works out a defined role's color.
func (p *Palette) custom(s Spec) tint.Color {
	c := opaque(s.Derive(p))
	if s.Kind != KindSurface {
		c = p.guard(c, p.groundColor, p.scaled(s.Min))
	}
	return c
}

// Hover and pressed steps, toward the background's text.
const (
	hoverStepDark, hoverStepLight     = 0.1, 0.08
	pressedStepDark, pressedStepLight = 0.18, 0.14
)

func (p *Palette) hoverStep() float64 { return p.ByDark(hoverStepDark, hoverStepLight) }

// state works out a role in a state from its color at rest.
func (p *Palette) state(r Role) tint.Color {
	rest := r.Rest()
	c := p.Color(rest)
	if r&selectedBit != 0 {
		c = p.selected(rest, c)
	}
	if r&(hoverBit|pressedBit) == 0 {
		return c
	}
	c, toward := p.step(c, p.hoverStep())
	if r&pressedBit != 0 {
		// A step past hover, the same way, so the two never meet.
		c = p.stepToward(c, toward, p.ByDark(pressedStepDark-hoverStepDark, pressedStepLight-hoverStepLight))
	}
	return c
}

// step moves c toward the background's text, by amt of the way from the
// background to its text, so a step is as large on a tone as on the page:
// lighter on a dark tint, darker on a light one. Where c is too near the
// text's color for a pressed step past it, as a light accent on a dark tint
// can be, it steps toward the background instead. It returns the step, and
// which way it went.
func (p *Palette) step(c tint.Color, amt float64) (out, toward tint.Color) {
	toward = p.pageFg
	room := p.ByDark(pressedStepDark, pressedStepLight) * p.colors[Bg].Distance(p.pageFg)
	if d := c.Distance(p.pageFg); d < room && c.Distance(p.colors[Bg]) > d {
		toward = p.colors[Bg]
	}
	if out = p.stepToward(c, toward, amt); out.Distance(c) >= minStep || toward == p.colors[Bg] {
		return out, toward
	}
	return p.stepToward(c, p.colors[Bg], amt), p.colors[Bg]
}

// stepToward moves c toward toward, by amt of the way from the background to
// its text. It keeps the text on c as readable as it was, where that still
// leaves a step that shows; text on the state, from [Palette.On], is kept
// readable either way.
func (p *Palette) stepToward(c, toward tint.Color, amt float64) tint.Color {
	t := 0.0
	if d := c.Distance(toward); d > 0 {
		t = min(1, amt*p.colors[Bg].Distance(p.pageFg)/d)
	}
	if p.guards {
		text := p.bestText(c)
		if out := c.MixReadable(toward, text, t, min(minText, text.Contrast(c))); out.Distance(c) >= minStep {
			return out
		}
	}
	return c.Mix(toward, t)
}

// selected is base tinted toward the accent, for a selected row: the tint's
// own selection on its background, if text stays readable on it. On a fill
// the accent can't tint, such as the accent itself, it's a pressed step.
func (p *Palette) selected(rest Role, base tint.Color) tint.Color {
	if rest == Bg {
		return p.bgSelected
	}
	return p.tinted(base)
}

// pageSelection is the background's selection: the tint's own, if text
// stays readable on it, or else the background tinted toward the accent.
func (p *Palette) pageSelection() tint.Color {
	t, bg := p.t, p.colors[Bg]
	if t.SelectionBg != nil && t.Bg != nil && sameRGB(bg, *t.Bg) {
		sel := opaque(*t.SelectionBg)
		if !p.guards || (tint.Contrast(p.pageFg, sel) >= minText && tint.Contrast(sel, bg) >= 1.12) {
			return sel
		}
	}
	return p.tinted(bg)
}

// tinted is base tinted toward the accent, as far as text still reads on it.
func (p *Palette) tinted(base tint.Color) tint.Color {
	accent, text := p.rootOf().colors[Accent], p.bestText(base)
	floor := min(minText, text.Contrast(base))
	amt := p.ByDark(0.27, 0.2)
	sel := base.Mix(accent, amt)
	// Fade the accent until text reads on it. Where the tint's own text is
	// barely readable, that can take it all the way to base.
	for ; p.guards && amt > 0 && tint.Contrast(text, sel) < floor; amt -= 0.02 {
		sel = base.Mix(accent, max(amt, 0))
	}
	if sel.Distance(base) < minStep {
		out, _ := p.step(base, p.ByDark(pressedStepDark, pressedStepLight))
		return out
	}
	return sel
}

func (p *Palette) rootOf() *Palette {
	if p.root != nil {
		return p.root
	}
	return p
}

// Color returns a role's color, in its states. The ground is what it was
// where it was drawn, and the rest are worked out for drawing on it; a text
// role used as a fill, such as [Fg], is text that reads on that fill.
func (p *Palette) Color(r Role) tint.Color {
	if r == p.ground && r.Kind() != KindText {
		return p.groundColor
	}
	if r&stateMask != 0 {
		return p.state(r)
	}
	if int(r) < len(p.colors) {
		return p.colors[r]
	}
	if s, ok := specOf(r); ok {
		// Defined after the palette was derived.
		return p.custom(s)
	}
	return p.colors[Fg]
}

// Tint returns the tint the palette was derived from.
func (p *Palette) Tint() *tint.Tint { return p.t }

// Dark reports whether the background is dark.
func (p *Palette) Dark() bool { return p.dark }

// Ground returns the surface the palette draws on: [Bg], or the surface
// given to [Palette.On].
func (p *Palette) Ground() Role { return p.ground }

// GroundColor returns the color of the surface the palette draws on. It's
// [Palette.Color] of [Palette.Ground], but for a text role used as a fill,
// such as [Fg], whose Color is the text that reads on it.
func (p *Palette) GroundColor() tint.Color { return p.groundColor }

// Mono reports whether the status tones can't be told apart by hue: a grey
// success or danger, or a tint whose hues sit close together. [Fitted.Attr]
// marks the status tones with attributes then, and wherever the terminal's
// colors can't tell them apart.
func (p *Palette) Mono() bool {
	return mono(p.colors[Success], p.colors[Warning], p.colors[Danger])
}

// ByDark returns ifDark on a dark background and ifLight on a light one. It's
// for a [Spec]'s Derive, where steps usually want to be a little larger on
// dark backgrounds.
func (p *Palette) ByDark(ifDark, ifLight float64) float64 { return pickBool(p.dark, ifDark, ifLight) }

// On returns the palette for drawing on surface rather than on the palette's
// ground. surface may be any role used as a fill, in any state: a layer such
// as [Raised], a tone such as [Danger] for a pill, a soft fill, [Control], or
// a selected row, Bg.Selected().
//
// Text and lines are worked out again for surface: body text is whatever
// reads on it, the background's text on a dialog, or the background itself
// on a light accent, and the rest are mixed from that toward surface, as they
// are toward the background. Tones are kept as far from surface as the
// guards keep them from the background; [Control], [Sunken] and soft fills
// step off surface, so an input on a dialog sinks into the dialog. The
// layers don't change:
//
//	row := pal.On(roles.Bg.Selected())
//	hint := row.Color(roles.Subtle) // reads on the selected row
//	pill := pal.On(roles.Danger)
//	text := pill.Color(roles.Fg) // text on a danger pill
func (p *Palette) On(surface Role) *Palette {
	if surface == p.ground {
		return p
	}
	root := p.rootOf()
	if surface == Bg {
		return root
	}
	if p != root && isLayer(surface) {
		return root.On(surface)
	}
	out := *p
	out.root, out.parent = root, p
	out.ground, out.groundColor = surface, p.Color(surface)
	out.colors = slices.Clone(root.colors[:numBuiltin])
	out.relate()
	return &out
}

// LerpPalette returns a palette partway from a to b, for crossfading between
// tints: t runs from 0 (a) to 1 (b). Fit it like any palette, or use [Lerp],
// which fits it too. The palettes Derive returned are mixed; a palette from
// [Palette.On] is taken as the one it came from. The layers and tones are
// mixed, and the rest worked out again from the mix.
//
// Between a dark and a light tint, the background passes through gray. Text
// stays on the side of it that reads, so it flips from light to dark once,
// where the background turns light, rather than fading out and back in. Right
// at gray, body text can't quite keep its contrast: no color has more than
// about 4.6:1 against it.
func LerpPalette(a, b *Palette, t float64) *Palette {
	if a == nil || b == nil {
		return b
	}
	a, b = a.rootOf(), b.rootOf()
	switch {
	case t >= 1:
		return b
	case t <= 0:
		return a
	}
	out := *b
	out.colors = slices.Clone(b.colors)
	// The layers and the tones; relate works out the rest from them.
	for _, r := range absolute {
		out.colors[r] = a.colors[r].Mix(b.colors[r], t)
	}
	for r := firstTone; r < numCore; r++ {
		out.colors[r] = a.colors[r].Mix(b.colors[r], t)
	}
	bg := out.colors[Bg]
	out.groundColor = bg
	out.dark = tint.IsDark(bg)
	out.termFg = a.termFg.Mix(b.termFg, t)
	out.bgSelected = a.bgSelected.Mix(b.bgSelected, t)
	out.pageFg = out.guard(a.pageFg.Mix(b.pageFg, t), bg, minFor(Fg))
	for r := firstTone; r < numCore; r++ {
		out.colors[r] = out.guard(out.colors[r], bg, minColor)
	}
	out.relate()
	return &out
}
