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
// background, for a flat look. [Hover] and [Selection] still show.
func WithFlatSurfaces() Option { return func(o *options) { o.flat = true } }

// WithoutGuards turns off the contrast guards: text, marks and lines keep
// the tint's colors even where they read poorly. It's meant for a hand-made
// tint whose author wants exactly their colors.
func WithoutGuards() Option { return func(o *options) { o.noGuards = true } }

// Palette is every role's color for one tint, in true color, before it's
// fitted to a terminal with [Palette.Fit]. It's immutable, and safe to share
// between goroutines.
type Palette struct {
	t      *tint.Tint
	colors []tint.Color // by role: the builtins, then the roles defined when it was derived
	ground Role         // what text is guarded against: Bg, or the surface given to On
	guards bool
	dark   bool
	mono   bool

	accentSlot, accent2Slot tint.Slot
}

// Derive works out every role's color for t. A tint's missing colors are
// filled in, so hand-made tints work too.
func Derive(t *tint.Tint, opts ...Option) *Palette {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	if t == nil {
		t = &tint.Tint{}
	}
	bg := orColor(o.bg, t.Bg, tint.Color{A: 0xff})
	p := &Palette{
		t:      t,
		colors: make([]tint.Color, numBuiltin, int(numBuiltin)+len(defined())),
		ground: Bg,
		guards: !o.noGuards,
		dark:   tint.IsDark(bg),
	}
	white := tint.Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	fg := orColor(o.fg, t.Fg, pickBool(p.dark, white, tint.Color{A: 0xff}))
	d := deriver{p: p, t: t, o: &o, bg: bg}
	d.derive(fg)
	for _, s := range defined() {
		p.colors = append(p.colors, p.custom(s))
	}
	return p
}

type deriver struct {
	p  *Palette
	t  *tint.Tint
	o  *options
	bg tint.Color
}

func (d *deriver) set(r Role, c tint.Color) { d.p.colors[r] = c }

func (d *deriver) get(r Role) tint.Color { return d.p.colors[r] }

// lift is [tint.Color.EnsureContrast] against the background, when the
// guards are on.
func (d *deriver) lift(c tint.Color, ratio float64) tint.Color {
	if !d.p.guards {
		return c
	}
	return c.EnsureContrast(d.bg, ratio)
}

func (d *deriver) derive(fg tint.Color) {
	p, bg := d.p, d.bg
	fg = d.lift(fg, minFor(Fg))
	d.set(Bg, bg)
	d.set(Fg, fg)
	accent, accent2 := d.accents(fg)
	d.status()
	d.surfaces(fg)
	d.set(Muted, d.lift(fg.Mix(bg, 0.36), minFor(Muted)))
	d.set(Subtle, d.lift(fg.Mix(bg, 0.6), minFor(Subtle)))
	d.set(Border, d.lift(fg.Mix(bg, 0.78), minFor(Border)))
	d.set(BorderStrong, fg.Mix(bg, 0.5))
	d.set(Accent, d.lift(accent, minColor))
	d.set(Accent2, d.lift(accent2, minColor))

	for r := firstTone; r < numCore; r++ {
		amt := p.ByDark(0.17, 0.14)
		if r == Accent || r == Accent2 {
			amt = p.ByDark(0.2, 0.16)
		}
		d.set(r.Soft(), bg.Mix(d.get(r), amt))
		d.set(r.On(), onColor(bg, fg, d.get(r)))
	}
	d.set(Selection, d.selection(fg))
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
	return d.get(Fg)
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

// status sets the status and hue colors, and marks a single-hue tint, where
// hue can't tell them apart.
func (d *deriver) status() {
	success, danger, warning := d.pick(tint.SlotGreen), d.pick(tint.SlotRed), d.pick(tint.SlotYellow)
	d.p.mono = chroma(success) < 0.04 || chroma(danger) < 0.04 || hueDist(hue(success), hue(danger)) < 45 ||
		(hueDist(hue(warning), hue(danger)) < 12 && hueDist(hue(success), hue(warning)) < 30)
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
	for _, r := range [...]Role{Surface, Raised, Overlay, Sunken} {
		d.set(r, bg)
	}
	if !d.o.flat {
		d.set(Surface, bg.Mix(fg, p.ByDark(0.05, 0.04)))
		d.set(Raised, bg.Mix(fg, p.ByDark(0.095, 0.065)))
		// Raised's step again, smaller where text would stop reading on it.
		raised := d.get(Raised)
		if p.guards {
			d.set(Overlay, raised.MixReadable(fg, fg, p.ByDark(0.095, 0.065), minColor))
		} else {
			d.set(Overlay, raised.Mix(fg, p.ByDark(0.095, 0.065)))
		}
	}
	d.set(Hover, bg.Mix(fg, p.ByDark(0.12, 0.09)))
	if !d.o.flat {
		d.set(Sunken, d.sunken(fg))
	}
}

// sunken is the well inputs and controls sit in: sunk below the background,
// toward black on a dark tint. A background already near black has no room
// below it, so the well rises toward the text instead, past the hover
// shade, so a control still shows on a row under the pointer.
func (d *deriver) sunken(fg tint.Color) tint.Color {
	bg := d.bg
	if !d.p.dark {
		return bg.Mix(fg, 0.07)
	}
	if in := bg.Mix(tint.Color{A: 0xff}, 0.32); tint.Contrast(in, bg) >= minSunken {
		return in
	}
	hover := d.get(Hover)
	in := hover
	for amt := 0.12; amt < 0.4 && tint.Contrast(in, hover) < minSunken; amt += 0.01 {
		in = bg.Mix(fg, amt)
	}
	return in
}

// selection is the tint's own selection if text stays readable on it,
// otherwise the accent faded into the background.
func (d *deriver) selection(fg tint.Color) tint.Color {
	t, bg, guards := d.t, d.bg, d.p.guards
	if t.SelectionBg != nil && t.Bg != nil && sameRGB(bg, *t.Bg) {
		sel := opaque(*t.SelectionBg)
		if !guards || (tint.Contrast(fg, sel) >= minText && tint.Contrast(sel, bg) >= 1.12) {
			return sel
		}
	}
	accent := d.get(Accent)
	amt := d.p.ByDark(0.27, 0.2)
	sel := bg.Mix(accent, amt)
	// Fade the accent until text reads on it. Where the tint's own text is
	// barely readable, that can take it all the way to the background.
	for ; guards && amt > 0 && tint.Contrast(fg, sel) < minText; amt -= 0.02 {
		sel = bg.Mix(accent, max(amt, 0))
	}
	return sel
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

// custom works out a defined role's color.
func (p *Palette) custom(s Spec) tint.Color {
	c := opaque(s.Derive(p))
	if p.guards && s.Kind != KindSurface && s.Min > 0 {
		c = c.EnsureContrast(p.Color(p.ground), s.Min)
	}
	return c
}

// Color returns a role's color.
func (p *Palette) Color(r Role) tint.Color {
	if int(r) < len(p.colors) {
		return p.colors[r]
	}
	if s, ok := specOf(r); ok {
		// Defined after the palette was derived.
		return p.custom(s)
	}
	return p.colors[Fg]
}

// Tone returns a tone's three colors. For any other role, all three are its
// color.
func (p *Palette) Tone(r Role) Tone {
	if !r.IsTone() {
		c := p.Color(r)
		return Tone{Base: c, Soft: c, On: c}
	}
	return Tone{Base: p.Color(r), Soft: p.Color(r.Soft()), On: p.Color(r.On())}
}

// Tone is a hue's three uses.
type Tone struct {
	// Base is the hue as text and marks, readable on the background.
	Base tint.Color
	// Soft is a fill: the hue faded into the background. Base reads on it.
	Soft tint.Color
	// On is text on Base used as a solid fill.
	On tint.Color
}

// Tint returns the tint the palette was derived from.
func (p *Palette) Tint() *tint.Tint { return p.t }

// Dark reports whether the background is dark.
func (p *Palette) Dark() bool { return p.dark }

// Guarded reports whether the contrast guards are on.
func (p *Palette) Guarded() bool { return p.guards }

// Mono reports whether the tint's status hues are too alike to tell apart,
// as in a single-hue tint such as amber or green CRT. Show status with
// weight or an icon too.
func (p *Palette) Mono() bool { return p.mono }

// AccentSlot returns the slot the accent came from, or for a custom accent
// color, the slot nearest it in hue.
func (p *Palette) AccentSlot() tint.Slot { return p.accentSlot }

// Accent2Slot returns the slot the second accent came from. It's the
// accent's when no other slot was far enough from it in hue.
func (p *Palette) Accent2Slot() tint.Slot { return p.accent2Slot }

// ByDark returns ifDark on a dark background and ifLight on a light one. It's
// for a [Spec]'s Derive, where steps usually want to be a little larger on
// dark backgrounds.
func (p *Palette) ByDark(ifDark, ifLight float64) float64 { return pickBool(p.dark, ifDark, ifLight) }

// On returns the palette for drawing on surface rather than on [Bg], with
// every text, line and tone color kept as far from surface as the guards keep
// it from the background. Use it for text on a highlighted row, where a hint
// that reads on the page could vanish into the highlight:
//
//	row := pal.On(roles.Selection)
//	hint := row.Color(roles.Subtle) // reads on the selection, not only on Bg
//
// Surfaces, soft fills and on colors are unchanged. With the guards off, it
// only changes what [Palette.Fit] treats as the surface.
func (p *Palette) On(surface Role) *Palette {
	out := *p
	out.ground = surface
	out.colors = slices.Clone(p.colors)
	if p.guards {
		out.guard(p.Color(surface))
	}
	return &out
}

// guard keeps every text, line and tone color as far from under as the
// guards keep it from the background. The ground itself is left alone: it's
// what under is.
func (p *Palette) guard(under tint.Color) {
	for i := range p.colors {
		r := Role(i) //nolint:gosec // bounded by the number of roles
		if k := r.Kind(); r == p.ground || (k != KindText && k != KindTone) {
			continue
		}
		if m := minFor(r); m > 0 {
			p.colors[i] = p.colors[i].EnsureContrast(under, m)
		}
	}
}

// Lerp returns a palette partway from a to b, for crossfading between tints:
// t runs from 0 (a) to 1 (b). Fit it like any palette. Between a dark and a
// light tint there's no readable middle, so it returns b. Otherwise the mixed
// colors are guarded again, and text on each tone is chosen again for the
// mixed tone.
func Lerp(a, b *Palette, t float64) *Palette {
	if a == nil || t >= 1 || a.dark != b.dark {
		return b
	}
	if t <= 0 {
		return a
	}
	out := *b
	out.colors = slices.Clone(b.colors)
	for i := range min(len(a.colors), len(b.colors)) {
		out.colors[i] = a.colors[i].Mix(b.colors[i], t)
	}
	if out.guards {
		out.guard(out.Color(out.ground))
	}
	bg, fg := out.colors[Bg], out.colors[Fg]
	for r := firstTone; r < numCore; r++ {
		out.colors[r.On()] = onColor(bg, fg, out.colors[r])
	}
	return &out
}

// onColor is text for a fill: the background or text color, whichever reads
// better, or black or white when neither does.
func onColor(bg, fg, fill tint.Color) tint.Color {
	best := bg
	if tint.Contrast(fg, fill) > tint.Contrast(bg, fill) {
		best = fg
	}
	if tint.Contrast(best, fill) >= minColor {
		return best
	}
	black, white := tint.Color{A: 0xff}, tint.Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	if tint.Contrast(black, fill) > tint.Contrast(white, fill) {
		return black
	}
	return white
}
