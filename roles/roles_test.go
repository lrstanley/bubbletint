// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"image/color"
	"testing"

	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
	tint "github.com/lrstanley/bubbletint/v2"
)

// testShadow is a defined surface, as an app would add one.
var testShadow = Define(Spec{
	Name: "test-shadow",
	Kind: KindSurface,
	Derive: func(p *Palette) tint.Color {
		return p.Color(Bg).Mix(tint.Color{A: 0xff}, 0.5)
	},
	Fallback: Fallback{Slots: []tint.Slot{tint.SlotBlack}},
})

// testHint is a defined text role with a floor.
var testHint = Define(Spec{
	Name:     "test-hint",
	Kind:     KindTone, // taken as KindText
	Derive:   func(p *Palette) tint.Color { return p.Color(Fg).Mix(p.Color(p.Ground()), 0.7) },
	Min:      3,
	Fallback: Fallback{Slots: []tint.Slot{tint.SlotBrightBlack, tint.SlotWhite}, Attr: Italic},
})

// textRoles are the roles with a contrast floor.
var textRoles = []Role{
	Fg, Muted, Subtle, Border, Neutral, Accent, Accent2, Success, Warning, Danger, Info,
	Red, Green, Yellow, Blue, Magenta, Cyan, testHint,
}

// fills are roles drawn on, in their states, for checking what's drawn on
// them.
var fills = []Role{
	Raised, Overlay.Hover(), Sunken.Pressed(), Bg.Selected(), Bg.Selected().Hover(),
	Control, Accent, Danger, Warning.Soft(), Neutral,
}

var (
	black = tint.Color{A: 0xff}
	white = tint.Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
)

// reachable is the most contrast any color has against c.
func reachable(c tint.Color) float64 { return max(tint.Contrast(black, c), tint.Contrast(white, c)) }

func TestRoleNames(t *testing.T) {
	t.Parallel()
	tests := []struct {
		r    Role
		want string
		kind Kind
	}{
		{r: Bg, want: "bg", kind: KindSurface},
		{r: Control, want: "control", kind: KindSurface},
		{r: BorderStrong, want: "border-strong", kind: KindText},
		{r: Neutral, want: "neutral", kind: KindTone},
		{r: Danger, want: "danger", kind: KindTone},
		{r: Danger.Soft(), want: "danger-soft", kind: KindSoft},
		{r: Overlay.Hover(), want: "overlay:hover", kind: KindSurface},
		{r: Bg.Selected().Hover(), want: "bg:selected:hover", kind: KindSurface},
		{r: Accent.Pressed(), want: "accent:pressed", kind: KindTone},
		{r: Danger.Soft().Hover(), want: "danger-soft:hover", kind: KindSoft},
		{r: testShadow, want: "test-shadow", kind: KindSurface},
		{r: testHint.Hover(), want: "test-hint:hover", kind: KindText},
		{r: Role(4000), want: "role(4000)", kind: KindText},
	}
	for _, tt := range tests {
		if got := tt.r.String(); got != tt.want {
			t.Errorf("String() = %q, want %q", got, tt.want)
		}
		if got := tt.r.Kind(); got != tt.kind {
			t.Errorf("%s: Kind() = %v, want %v", tt.r, got, tt.kind)
		}
	}
	if Muted.Soft() != Muted || Bg.Base() != Bg || Raised.Hover().Soft() != Raised.Hover() {
		t.Error("Soft and Base of a non-tone should return it")
	}
	if Accent.Soft().Base() != Accent || Accent.Hover().Soft() != Accent.Soft().Hover() {
		t.Error("Base() doesn't undo Soft(), or Soft() drops the state")
	}
	for _, s := range tint.ChromaticSlots {
		if Hue(s) != Hue(s.Bright()) || !Hue(s).IsTone() {
			t.Errorf("Hue(%v) = %v", s, Hue(s))
		}
	}
	if Hue(tint.SlotWhite) != Neutral {
		t.Error("Hue(white) != Neutral")
	}
}

func TestStates(t *testing.T) {
	t.Parallel()
	tests := []struct {
		r    Role
		want State
	}{
		{r: Bg, want: 0},
		{r: Bg.Hover(), want: StateHover},
		{r: Bg.Hover().Pressed(), want: StatePressed},
		{r: Bg.Pressed().Hover(), want: StateHover},
		{r: Bg.Selected().Pressed(), want: StateSelected | StatePressed},
		{r: Bg.In(StateSelected | StateHover), want: StateSelected | StateHover},
		{r: Bg.In(StateHover | StatePressed), want: StatePressed},
	}
	for _, tt := range tests {
		if got := tt.r.State(); got != tt.want {
			t.Errorf("%s: State() = %v, want %v", tt.r, got, tt.want)
		}
		if tt.r.Rest() != Bg {
			t.Errorf("%s: Rest() = %v", tt.r, tt.r.Rest())
		}
	}
	if s := (StateSelected | StateHover).String(); s != "selected:hover" {
		t.Errorf("State.String() = %q", s)
	}
}

func TestDeriveFloors(t *testing.T) {
	t.Parallel()
	var fails, checked int
	fail := func(format string, args ...any) {
		fails++
		if fails < 10 {
			t.Errorf(format, args...)
		}
	}
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		bg := p.Color(Bg)
		for _, r := range textRoles {
			checked++
			if got := tint.Contrast(p.Color(r), bg); got < minFor(r)-0.01 {
				fail("%s: %s contrast %.2f < %.2f", tn.ID, r, got, minFor(r))
			}
		}
		// Text on every fill reads, as far as any color can.
		for _, g := range fills {
			on, under := p.On(g), p.Color(g)
			for _, r := range []Role{Fg, Muted, Subtle} {
				checked++
				if got := tint.Contrast(on.Color(r), under); got < min(on.floor(r), reachable(under))-0.05 {
					fail("%s: %s on %s contrast %.2f", tn.ID, r, g, got)
				}
			}
		}
		// Text on a tone's fill reads at least as well as the tone does.
		for r := firstTone; r < numCore; r++ {
			checked++
			if got := tint.Contrast(p.On(r).Color(Fg), p.Color(r)); got < minColor {
				fail("%s: fg contrast %.2f on %s", tn.ID, got, r)
			}
		}
	}
	t.Logf("checked %d contrasts", checked)
}

// A state shows: it moves off the role at rest, toward the text on a surface,
// and keeps the text on it readable.
func TestStateColors(t *testing.T) {
	t.Parallel()
	var fails int
	fail := func(format string, args ...any) {
		fails++
		if fails < 10 {
			t.Errorf(format, args...)
		}
	}
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		for _, r := range []Role{Bg, Surface, Raised, Overlay, Sunken, Control, Accent, Danger, Danger.Soft(), Neutral} {
			rest, hover, pressed, sel := p.Color(r), p.Color(r.Hover()), p.Color(r.Pressed()), p.Color(r.Selected())
			for name, c := range map[string]tint.Color{"hover": hover, "pressed": pressed, "selected": sel} {
				if c.Distance(rest) < minStep-0.001 {
					fail("%s: %s %s is %.3f from rest", tn.ID, r, name, c.Distance(rest))
				}
			}
			if pressed.Distance(hover) < minStep/2 {
				fail("%s: %s pressed is hover", tn.ID, r)
			}
			if r.Kind() == KindSurface && hover.Contrast(p.pageFg) >= rest.Contrast(p.pageFg) {
				fail("%s: %s hover isn't lifted toward the text", tn.ID, r)
			}
		}
		// A well sinks into whatever it's drawn on, on light tints too, where
		// it once matched a dialog.
		for _, g := range []Role{Bg, Surface, Raised, Overlay} {
			if c := tint.Contrast(p.On(g).Color(Sunken), p.Color(g)); c < minSunken-0.01 {
				fail("%s: a well on %s: contrast %.3f", tn.ID, g, c)
			}
		}
		// A menu's row under the pointer is a step off the menu, not the
		// page's hover shade.
		if over, row := p.Color(Overlay), p.Color(Overlay.Hover()); row.Contrast(p.pageFg) >= over.Contrast(p.pageFg) {
			fail("%s: overlay hover %s under overlay %s", tn.ID, row.Hex(), over.Hex())
		}
	}
}

func TestDeriveOptions(t *testing.T) {
	t.Parallel()
	tn := tint.TintTokyoNight
	t.Run("flat", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithFlatSurfaces())
		for _, r := range []Role{Surface, Raised, Overlay, Sunken} {
			if p.Color(r) != p.Color(Bg) {
				t.Errorf("%s != bg", r)
			}
		}
		if p.Color(Bg.Hover()) == p.Color(Bg) || p.Color(Control) == p.Color(Bg) {
			t.Error("hover or control == bg")
		}
	})
	t.Run("accent-slot", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithAccentSlot(tint.SlotBrightGreen))
		if p.accentSlot != tint.SlotGreen {
			t.Errorf("accent slot = %v", p.accentSlot)
		}
		if p.accent2Slot == tint.SlotGreen {
			t.Error("accent2 shares the accent's slot")
		}
	})
	t.Run("accent-color", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithAccentColor(tint.FromHex("#ff8800")))
		if p.accentSlot != tint.SlotYellow && p.accentSlot != tint.SlotRed {
			t.Errorf("accent slot = %v, want the slot nearest orange", p.accentSlot)
		}
	})
	t.Run("terminal", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithTerminal(color.Black, nil))
		if p.Color(Bg) != black {
			t.Errorf("Bg = %s", p.Color(Bg).Hex())
		}
		if !p.Dark() {
			t.Error("not dark on black")
		}
	})
	t.Run("without-guards", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithoutGuards())
		if p.guards {
			t.Error("guards on")
		}
		if p.Color(Fg) != *tn.Fg {
			t.Errorf("Fg = %s, want the tint's", p.Color(Fg).Hex())
		}
	})
	t.Run("nil-tint", func(t *testing.T) {
		t.Parallel()
		p := Derive(nil)
		if !p.Dark() || tint.Contrast(p.Color(Fg), p.Color(Bg)) < minText {
			t.Error("an empty tint should derive white on black")
		}
	})
	t.Run("state-attr", func(t *testing.T) {
		t.Parallel()
		f := Derive(tn, WithStateAttr(StateHover, Bold), WithStateAttr(State(0), Italic)).Fit(Target{Profile: colorprofile.ANSI})
		if got := f.Attr(Bg.Hover()); got != Bold {
			t.Errorf("hover attr = %v, want bold", got)
		}
		if got := f.Attr(Bg.Pressed()); got != Underline|Bold {
			t.Errorf("pressed attr = %v, want the default", got)
		}
	})
}

func TestPaletteOn(t *testing.T) {
	t.Parallel()
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		for _, g := range fills {
			on := p.On(g)
			if on.Ground() != g || on.Color(g) != p.Color(g) {
				t.Fatalf("%s: On(%s) has ground %s, %s", tn.ID, g, on.Ground(), on.Color(g).Hex())
			}
			for _, r := range absolute {
				if on.Color(r) != p.Color(r) {
					t.Fatalf("%s: On(%s) changed %s", tn.ID, g, r)
				}
			}
		}
		// A control and a soft fill step off what they're drawn on. Where
		// the tint's text barely reads, a control on the page and on a
		// dialog can both stop where text stops reading.
		dialog := p.On(Raised)
		if dialog.Color(Danger.Soft()) == p.Color(Danger.Soft()) {
			t.Errorf("%s: soft fill isn't worked out from the dialog", tn.ID)
		}
		if raised, c := p.Color(Raised), dialog.Color(Control); c.Contrast(raised) < raised.Mix(p.Color(Fg), 0.08).Contrast(raised)-0.01 {
			t.Errorf("%s: control %s too close to the dialog %s", tn.ID, c.Hex(), raised.Hex())
		}
		if p.On(Bg) != p || dialog.On(Bg) != p || dialog.On(Raised) != dialog {
			t.Errorf("%s: On of the ground isn't the palette itself", tn.ID)
		}
		// On of a layer is the same wherever it's asked from.
		if a, b := dialog.On(Overlay), p.On(Overlay); a.Color(Muted) != b.Color(Muted) {
			t.Errorf("%s: On(overlay) differs on a dialog", tn.ID)
		}
	}
	// A defined role is worked out from the ground.
	p := Derive(tint.TintTokyoNight)
	if p.On(Raised).Color(testHint) == p.Color(testHint) {
		t.Error("a defined role isn't worked out from the ground")
	}
}

func TestLerp(t *testing.T) {
	t.Parallel()
	a, b := Derive(tint.TintTokyoNight), Derive(tint.TintDracula)
	if LerpPalette(a, b, 0) != a || LerpPalette(a, b, 1) != b || LerpPalette(nil, b, 0.5) != b {
		t.Error("LerpPalette endpoints")
	}
	mid := LerpPalette(a, b, 0.5)
	if mid.Color(Bg) == a.Color(Bg) || mid.Color(Bg) == b.Color(Bg) {
		t.Error("LerpPalette didn't mix")
	}
	for _, r := range textRoles {
		if got := tint.Contrast(mid.Color(r), mid.Color(Bg)); got < minFor(r)-0.01 {
			t.Errorf("%s: contrast %.2f mid-crossfade", r, got)
		}
	}

	// Dark to light and back: text keeps reading on every frame, but for body
	// text right at gray, where nothing reaches its floor.
	for _, pr := range [][2]*tint.Tint{
		{tint.TintTokyoNight, tint.TintTokyoNightDay},
		{tint.TintCatppuccinLatte, tint.TintDracula},
		{tint.TintTokyoNight, tint.TintCatppuccinLatte},
	} {
		from, to := Derive(pr[0]), Derive(pr[1])
		flips := 0
		prev := from.Dark()
		for i := 1; i < 40; i++ {
			frame := LerpPalette(from, to, float64(i)/40)
			if frame.Dark() != prev {
				flips++
				prev = frame.Dark()
			}
			for _, r := range textRoles {
				floor := minFor(r) - 0.01
				if r == Fg {
					floor = 4.5
				}
				if got := tint.Contrast(frame.Color(r), frame.Color(Bg)); got < floor {
					t.Errorf("%s to %s at %d/40: %s contrast %.2f", pr[0].ID, pr[1].ID, i, r, got)
				}
			}
			for r := firstTone; r < numCore; r++ {
				if got := tint.Contrast(frame.On(r).Color(Fg), frame.Color(r)); got < minColor {
					t.Errorf("%s to %s at %d/40: fg contrast %.2f on %s", pr[0].ID, pr[1].ID, i, got, r)
				}
			}
		}
		if flips != 1 {
			t.Errorf("%s to %s: Dark() flipped %d times, want once", pr[0].ID, pr[1].ID, flips)
		}
	}
}

func TestLerpFitted(t *testing.T) {
	t.Parallel()
	a, b := Derive(tint.TintTokyoNight), Derive(tint.TintDracula)
	for _, tt := range []struct {
		profile colorprofile.Profile
		mixes   bool
	}{
		{profile: colorprofile.TrueColor, mixes: true},
		{profile: colorprofile.ANSI256, mixes: true},
		{profile: colorprofile.ANSI},
		{profile: colorprofile.ASCII},
	} {
		tg := Target{Profile: tt.profile}
		fa, fb := a.Fit(tg), b.Fit(tg)
		if Lerp(fa, fb, 0) != fa || Lerp(fa, fb, 1) != fb || Lerp(nil, fb, 0.5) != fb {
			t.Errorf("%v: Lerp endpoints", tt.profile)
		}
		mid := Lerp(fa, fb, 0.5)
		if mixes := mid != fb; mixes != tt.mixes {
			t.Errorf("%v: Lerp mixes = %v, want %v", tt.profile, mixes, tt.mixes)
		}
		if on := Lerp(fa.On(Raised), fb.On(Raised), 0.5); on.Ground() != Raised {
			t.Errorf("%v: Lerp of On(raised) has ground %v", tt.profile, on.Ground())
		}
	}
	// Different terminals: nothing between.
	fa, fb := a.Fit(Target{Profile: colorprofile.TrueColor}), b.Fit(Target{Profile: colorprofile.ANSI256})
	if Lerp(fa, fb, 0.5) != fb {
		t.Error("Lerp across terminals should be b")
	}
}

func TestDefineLate(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintTokyoNight)
	late := Define(Spec{Name: "late", Derive: func(p *Palette) tint.Color { return p.Color(Accent) }})
	if p.Color(late) != p.Color(Accent) {
		t.Error("a role defined after Derive isn't derived on demand")
	}
	if f := p.Fit(Target{Profile: colorprofile.ANSI256}); f.Color(late) == nil {
		t.Error("a role defined after Fit isn't fitted on demand")
	}

	// A surface defined late can be drawn on, at every profile.
	surface := Define(Spec{Name: "late-surface", Kind: KindSurface, Derive: func(p *Palette) tint.Color { return p.Color(Raised) }})
	for _, pr := range []colorprofile.Profile{colorprofile.TrueColor, colorprofile.ANSI256, colorprofile.ANSI, colorprofile.ASCII} {
		f := p.Fit(Target{Profile: pr})
		fill, text := f.Fill(surface), f.On(surface).Span(Muted)
		if pr >= colorprofile.ANSI256 && (fill.GetBackground() != f.Color(surface) || text.GetBackground() != f.Color(surface)) {
			t.Errorf("%v: a surface defined late isn't fitted", pr)
		}
	}
}

// pairCollapses counts pairs fitted to the same color, though derived apart.
func pairCollapses(p *Palette, f *Fitted) (n int, which []string) {
	for _, pr := range pairs {
		a, b := pr[0], pr[1]
		if a == groundMark {
			a = Bg
		}
		if f.Color(a) == f.Color(b) && p.Color(a).Distance(p.Color(b)) > 0.004 {
			n++
			which = append(which, a.String()+"="+b.String())
		}
	}
	if f.Color(testShadow) == f.Color(Bg) {
		n++
		which = append(which, "test-shadow=bg")
	}
	return n, which
}

func TestFit256(t *testing.T) {
	t.Parallel()
	var collapses, floors, states int
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		f := p.Fit(Target{Profile: colorprofile.ANSI256})
		n, which := pairCollapses(p, f)
		collapses += n
		if n > 0 {
			t.Logf("%s: %v", tn.ID, which)
		}
		bg := f.Color(Bg)
		for _, r := range textRoles {
			c, ok := f.Color(r).(ansi.IndexedColor)
			if !ok || c < 16 {
				t.Fatalf("%s: %s = %#v, want an index from 16", tn.ID, r, f.Color(r))
			}
			if got := tint.Contrast(c, bg); got < minFor(r)-0.05 {
				floors++
				t.Logf("%s: %s contrast %.2f < %.2f", tn.ID, r, got, minFor(r))
			}
		}
		// Every state stands apart from the role at rest, and a selected row
		// from a hovered one, on the page and in a dialog.
		dialog := f.On(Raised)
		for _, g := range []*Fitted{f, dialog} {
			for _, r := range []Role{Bg, Surface, Raised, Overlay, Sunken, Control, Accent, Danger.Soft()} {
				for _, s := range []Role{r.Hover(), r.Pressed(), r.Selected(), r.Selected().Hover()} {
					if g.Color(s) == g.Color(r) {
						states++
						t.Logf("%s on %s: %s shares %v with %s", tn.ID, g.Ground(), s, g.Color(s), r)
					}
				}
				if g.Color(r.Selected()) == g.Color(r.Hover()) || g.Color(r.Pressed()) == g.Color(r.Hover()) {
					states++
					t.Logf("%s on %s: %s's states collapse", tn.ID, g.Ground(), r)
				}
			}
			if g.Color(Control) == g.Color(g.Ground()) {
				states++
				t.Logf("%s: control shares %s's color", tn.ID, g.Ground())
			}
		}
	}
	// A tint or two has no 256-color index that keeps every floor and
	// pair; nearly all do.
	if collapses > 3 || floors > 3 || states > 0 {
		t.Errorf("collapses %d, broken floors %d, collapsed states %d", collapses, floors, states)
	}

	// A side panel's selected and hovered rows, and a dialog over the page,
	// once snapped to the surface under them on this light tint.
	f := Derive(tint.TintTokyoNightDay).Fit(Target{Profile: colorprofile.ANSI256})
	for _, pr := range [][2]Role{{Surface, Surface.Selected()}, {Surface, Surface.Hover()}, {Bg, Raised}, {Bg, Overlay}} {
		if f.Color(pr[0]) == f.Color(pr[1]) {
			t.Errorf("tokyo_night_day: %s and %s share %v", pr[0], pr[1], f.Color(pr[0]))
		}
	}
}

func TestFit16(t *testing.T) {
	t.Parallel()
	var attrs int
	sel := Bg.Selected()
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		f := p.Fit(Target{Profile: colorprofile.ANSI})
		bg := p.Color(Bg)
		ansi16 := tn.ANSI16()
		for _, r := range append(textRoles, sel) {
			c := f.Color(r)
			if c == nil {
				if r != Fg && r != Border && r != Neutral && f.Attr(r) == 0 {
					t.Errorf("%s: %s is the default with no attribute", tn.ID, r)
				}
				attrs++
				continue
			}
			s, ok := c.(ansi.BasicColor)
			if !ok {
				t.Fatalf("%s: %s = %#v, want a BasicColor", tn.ID, r, c)
			}
			if r == sel {
				continue
			}
			if got := tint.Contrast(*ansi16[s], bg); got < minFor(r)-0.01 {
				t.Errorf("%s: %s slot %d contrast %.2f < %.2f", tn.ID, r, s, got, minFor(r))
			}
			if r.IsTone() && r != Neutral {
				want := f.toneSlot(r)
				if tint.Slot(s).Normal() != want {
					t.Errorf("%s: %s took slot %d, want %v's", tn.ID, r, s, want)
				}
			}
		}
		if f.Color(Bg) != nil || f.Color(Fg) != nil || f.Color(Raised) != nil {
			t.Errorf("%s: surfaces painted with 16 colors", tn.ID)
		}
		// A hovered row has no slot to step to; it's underlined.
		if f.Color(Bg.Hover()) != f.Color(Bg) || f.Attr(Bg.Hover())&Underline == 0 {
			t.Errorf("%s: hover = %v, %v", tn.ID, f.Color(Bg.Hover()), f.Attr(Bg.Hover()))
		}
		// Text on the selection reads against the selection's slot.
		if slot, ok := f.Color(sel).(ansi.BasicColor); ok {
			on := f.On(sel)
			for _, r := range []Role{Fg, Muted} {
				s, isBasic := on.Color(r).(ansi.BasicColor)
				if !isBasic {
					t.Errorf("%s: %s on the selection is %#v", tn.ID, r, on.Color(r))
					continue
				}
				if s == slot {
					t.Errorf("%s: %s on the selection is the selection's slot", tn.ID, r)
				}
			}
		} else if f.Attr(sel)&Reverse == 0 {
			t.Errorf("%s: a selection with no slot isn't reversed", tn.ID)
		}
	}
	t.Logf("%d roles fell back to an attribute", attrs)
}

func TestFitOther(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintGruvboxDark)

	tc := p.Fit(Target{Profile: colorprofile.TrueColor})
	if tc.Color(Accent) != p.Color(Accent) || tc.Attr(Accent) != 0 {
		t.Error("true color should be the derived color")
	}
	if tc.Fit(color.White, Fg) != white {
		t.Error("Fit() at true color")
	}
	if tc.Profile() != colorprofile.TrueColor || p.Fit(Target{}).Profile() != colorprofile.TrueColor {
		t.Error("Profile()")
	}

	unpainted := p.Fit(Target{Profile: colorprofile.TrueColor, Unpainted: true})
	if unpainted.Color(Bg) != nil || unpainted.Color(Fg) != nil || unpainted.Color(Surface) == nil {
		t.Error("Unpainted should leave only bg and fg to the terminal")
	}

	none := p.Fit(Target{Profile: colorprofile.ASCII})
	for i := range numBuiltin {
		if none.Color(i) != nil {
			t.Fatalf("%s has a color with no color", i)
		}
	}
	for r, want := range map[Role]Attr{ //nolint:exhaustive // a sample
		Muted: Faint, Bg.Selected(): Reverse, Bg.Hover(): Underline, Danger: Bold | Underline,
		Warning: Underline, Success: Bold, Control: Reverse, Neutral: 0, testHint: Italic, Bg: 0,
	} {
		if got := none.Attr(r); got != want {
			t.Errorf("no color %s: Attr() = %v, want %v", r, got, want)
		}
	}
	if !none.Mono() {
		t.Error("with no color, status can't be told apart by hue")
	}

	f256 := p.Fit(Target{Profile: colorprofile.ANSI256})
	if got := f256.Fit(p.Color(Accent).Mix(p.Color(Raised), 0.3), Accent); got == nil {
		t.Error("Fit() at 256")
	}
	if f256.Palette() != p {
		t.Error("Palette()")
	}
	// A small step off a surface doesn't snap back to the surface, and it's
	// kept apart from what it's asked to be.
	var snaps int
	for _, tn := range tint.DefaultTints() {
		pal := Derive(tn)
		f := pal.Fit(Target{Profile: colorprofile.ANSI256})
		for _, r := range []Role{Surface, Raised, Overlay, Sunken} {
			if f.Fit(pal.Color(r).Mix(pal.Color(Fg), 0.06), r) == f.Color(r) {
				snaps++
			}
		}
		if f.Fit(pal.Color(Bg), Bg, Raised) == f.Color(Raised) {
			t.Errorf("%s: Fit() isn't kept apart", tn.ID)
		}
	}
	if snaps > 0 {
		t.Errorf("%d steps snapped back to their surface", snaps)
	}
	if f256.Fit(p.Color(Raised), Raised) != f256.Color(Raised) {
		t.Error("Fit() of a role's own color isn't its color")
	}

	// A terminal palette that isn't the tint's: everything comes from it.
	var term [16]color.Color
	for i := range term {
		term[i] = ansi.BasicColor(i) //nolint:gosec // 0-15
	}
	f16 := p.Fit(Target{Profile: colorprofile.ANSI, ANSI: &term})
	if s, ok := f16.Color(Danger).(ansi.BasicColor); !ok || tint.Slot(s).Normal() != tint.SlotRed {
		t.Errorf("Danger = %#v", f16.Color(Danger))
	}
	if got := f16.Fit(color.White, Danger); got != f16.Color(Danger) {
		t.Error("Fit() at 16 should be like's color")
	}
}

func TestDistinct(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintTokyoNight)
	tests := []struct {
		name string
		f    *Fitted
		a, b Role
		want bool
	}{
		{name: "true color raised", f: p.Fit(Target{}), a: Raised, b: Bg, want: true},
		{name: "flat raised", f: Derive(tint.TintTokyoNight, WithFlatSurfaces()).Fit(Target{}), a: Raised, b: Bg},
		{name: "16 raised", f: p.Fit(Target{Profile: colorprofile.ANSI}), a: Raised, b: Bg},
		{name: "16 hover", f: p.Fit(Target{Profile: colorprofile.ANSI}), a: Bg.Hover(), b: Bg, want: true},
		{name: "none raised", f: p.Fit(Target{Profile: colorprofile.ASCII}), a: Raised, b: Bg},
		{name: "256 menu hover", f: p.Fit(Target{Profile: colorprofile.ANSI256}), a: Overlay.Hover(), b: Overlay, want: true},
	}
	for _, tt := range tests {
		if got := tt.f.Distinct(tt.a, tt.b); got != tt.want {
			t.Errorf("%s: Distinct(%s, %s) = %v, want %v", tt.name, tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMono(t *testing.T) {
	t.Parallel()
	var monos int
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		f := p.Fit(Target{})
		if p.Mono() != f.Mono() {
			t.Errorf("%s: Palette.Mono %v, Fitted.Mono %v", tn.ID, p.Mono(), f.Mono())
		}
		if !f.Mono() {
			if f.Attr(Danger) != 0 {
				t.Errorf("%s: danger has attributes in true color", tn.ID)
			}
			continue
		}
		monos++
		if f.Attr(Danger) != Bold|Underline || f.Attr(Warning) != Underline || f.Attr(Success) != Bold ||
			f.On(Raised).Attr(Danger) != Bold|Underline {
			t.Errorf("%s: mono, but danger %v, warning %v, success %v", tn.ID, f.Attr(Danger), f.Attr(Warning), f.Attr(Success))
		}
	}
	if !Derive(tint.TintAmberCrtRetro).Mono() || Derive(tint.TintTokyoNight).Mono() {
		t.Error("amber_crt_retro should be mono, and tokyo_night not")
	}
	t.Logf("%d mono tints", monos)
}

func BenchmarkDerive(b *testing.B) {
	for b.Loop() {
		Derive(tint.TintTokyoNight)
	}
}

func BenchmarkFit256(b *testing.B) {
	p := Derive(tint.TintTokyoNight)
	for b.Loop() {
		p.Fit(Target{Profile: colorprofile.ANSI256})
	}
}

// A crossfade's frame: mix, fit, and the states and surfaces a screen draws.
func BenchmarkLerpFrame256(b *testing.B) {
	tg := Target{Profile: colorprofile.ANSI256}
	from, to := Derive(tint.TintTokyoNight).Fit(tg), Derive(tint.TintDracula).Fit(tg)
	for b.Loop() {
		f := Lerp(from, to, 0.5)
		_ = f.Color(Bg.Hover())
		_ = f.Color(Bg.Selected())
		_ = f.On(Raised).Color(Control)
		_ = f.On(Bg.Selected()).Color(Muted)
	}
}

func TestFitEdgeCases(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintTokyoNight)

	// Text on a solid tone. The tone itself stays as it was, and with 16
	// colors gets the same slot as on the background.
	if got := p.On(Danger).Color(Danger); got != p.Color(Danger) {
		t.Errorf("On(Danger) changed danger to %s", got.Hex())
	}
	on, base := p.Fit(Target{Profile: colorprofile.ANSI}).On(Danger), p.Fit(Target{Profile: colorprofile.ANSI})
	if on.Color(Fg) == nil {
		t.Error("fg on a danger fill is the terminal's default")
	}
	if on.Color(Danger) != base.Color(Danger) {
		t.Errorf("danger %v on its own fill, want %v", on.Color(Danger), base.Color(Danger))
	}

	// A nil *tint.Color behind the color.Color interface.
	var none *tint.Color
	var term [16]color.Color
	term[3] = none
	if f := p.Fit(Target{Profile: colorprofile.ANSI, ANSI: &term}); f.Color(Warning) == nil {
		t.Error("Warning fell back to the default with a typed nil slot")
	}
	if got := p.Fit(Target{Profile: colorprofile.TrueColor}).Fit(none, Fg); got != nil {
		t.Errorf("Fit(typed nil) = %v, want nil", got)
	}

	// Text on a tone keeps its contrast with 256 colors, and across a fade.
	f := p.Fit(Target{Profile: colorprofile.ANSI256})
	mid := LerpPalette(p, Derive(tint.TintDracula), 0.5)
	for _, r := range []Role{Neutral, Accent, Success, Warning, Danger, Info} {
		if got := tint.Contrast(f.On(r).Color(Fg), f.Color(r)); got < minColor {
			t.Errorf("256: fg contrast %.2f on %s", got, r)
		}
		if got := tint.Contrast(mid.On(r).Color(Fg), mid.Color(r)); got < minColor {
			t.Errorf("Lerp: fg contrast %.2f on %s", got, r)
		}
	}

	// A palette from On, fitted on its own, matches the one Fitted.On gives.
	for _, g := range []Role{Raised, Bg.Selected(), Danger, Control} {
		a, b := p.On(g).Fit(Target{Profile: colorprofile.ANSI256}), f.On(g)
		for _, r := range []Role{g, Fg, Muted, Control, Danger.Soft()} {
			if a.Color(r) != b.Color(r) {
				t.Errorf("On(%s).Fit: %s = %v, want %v", g, r, a.Color(r), b.Color(r))
			}
		}
	}
}

// A text role used as a fill gets text that reads on it, at rest or in a
// state, and the fill itself is what it was where it was drawn.
func TestTextRoleAsGround(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintTokyoNight)
	for _, g := range []Role{Fg, Fg.Hover(), Muted, Control.Pressed()} {
		on := p.On(g)
		if on.groundColor != p.Color(g) {
			t.Errorf("On(%s): ground %s, want %s", g, on.groundColor.Hex(), p.Color(g).Hex())
		}
		if g.Kind() != KindText && on.Color(g) != p.Color(g) {
			t.Errorf("On(%s): Color(ground) %s, want the ground", g, on.Color(g).Hex())
		}
		if c := tint.Contrast(on.Color(Fg), p.Color(g)); c < min(minFor(Fg), reachable(p.Color(g)))-0.05 {
			t.Errorf("On(%s): text contrast %.2f", g, c)
		}
	}
}

// A defined role worked out from the ground sees the ground itself, not the
// role worked out again for drawing on it.
func TestDefinedSeesGround(t *testing.T) {
	t.Parallel()
	echo := Define(Spec{Name: "echo-ground", Kind: KindSurface, Derive: func(p *Palette) tint.Color {
		return p.Color(p.Ground())
	}})
	fill := Define(Spec{Name: "echo-ground-color", Kind: KindSurface, Derive: func(p *Palette) tint.Color {
		return p.GroundColor()
	}})
	p := Derive(tint.TintTokyoNight)
	on := p.On(Control)
	if on.Color(echo) != p.Color(Control) {
		t.Errorf("Derive saw %s on a control, want the control %s", on.Color(echo).Hex(), p.Color(Control).Hex())
	}
	// On a text role used as a fill, GroundColor is the fill, not the text.
	if onFg := p.On(Fg); onFg.Color(fill) != p.Color(Fg) {
		t.Errorf("GroundColor on fg is %s, want fg %s", onFg.Color(fill).Hex(), p.Color(Fg).Hex())
	}
}

// A crossfade between palettes for drawing on nested surfaces keeps the
// nesting on every frame.
func TestLerpNested(t *testing.T) {
	t.Parallel()
	tg := Target{Profile: colorprofile.TrueColor}
	a := Derive(tint.TintTokyoNight).Fit(tg).On(Raised).On(Control)
	b := Derive(tint.TintDracula).Fit(tg).On(Raised).On(Control)
	mid := Lerp(a, b, 0.5)
	if mid.Ground() != Control || mid.p.parent == nil || mid.p.parent.ground != Raised {
		t.Fatalf("Lerp lost the nesting: ground %s", mid.Ground())
	}
	near := Lerp(a, b, 0.999)
	if d := tint.FromColor(near.Color(Control)).Distance(*tint.FromColor(b.Color(Control))); d > 0.01 {
		t.Errorf("the control jumps at the end of the fade: %.3f", d)
	}
}

// With no color, the status tones' fills are told apart, as their text is.
func TestStatusFillsNoColor(t *testing.T) {
	t.Parallel()
	f := Derive(tint.TintTokyoNight).Fit(Target{Profile: colorprofile.ASCII})
	seen := map[Attr]Role{} //nolint:exhaustive // filled as it goes
	for _, r := range []Role{Success, Warning, Danger} {
		a := attrOf(f.Fill(r))
		if o, dup := seen[a]; dup {
			t.Errorf("%s and %s fills are both %v", r, o, a)
		}
		seen[a] = r
	}
}

// A crossfade from a palette to itself changes nothing, however the palettes
// for drawing on nested surfaces were fitted.
func TestLerpSelf(t *testing.T) {
	t.Parallel()
	tg := Target{Profile: colorprofile.ANSI256}
	for _, tn := range []*tint.Tint{tint.TintTokyoNight, tint.TintCatppuccinLatte, tint.TintGruvboxDark} {
		for _, p := range []*Palette{Derive(tn), Derive(tn, WithFlatSurfaces(), WithTerminal(color.Black, nil))} {
			f := p.Fit(tg)
			for _, chain := range [][]Role{{Raised, Sunken}, {Raised, Control}, {Overlay.Hover(), Control}} {
				b := f
				for _, g := range chain {
					b = b.On(g)
				}
				mid := Lerp(b, b, 0.5)
				for _, r := range []Role{chain[len(chain)-1], Fg, Muted, Sunken, Control} {
					if mid.Color(r) != b.Color(r) {
						t.Errorf("%s %v: %s is %v mid-fade, want %v", tn.ID, chain, r, mid.Color(r), b.Color(r))
					}
				}
			}
		}
	}
}

// On a single-hue tint, a status fill in a state keeps its status's mark.
func TestMonoStateFills(t *testing.T) {
	t.Parallel()
	for _, pr := range []colorprofile.Profile{colorprofile.TrueColor, colorprofile.ANSI256} {
		f := Derive(tint.TintAmberCrtRetro).Fit(Target{Profile: pr})
		seen := map[Attr]Role{} //nolint:exhaustive // filled as it goes
		for _, r := range []Role{Success, Warning, Danger} {
			a := attrOf(f.Fill(r.Hover()))
			if o, dup := seen[a]; dup {
				t.Errorf("%v: hovered %s and %s fills are both %v", pr, r, o, a)
			}
			seen[a] = r
		}
	}
}

// With 256 colors, a color kept apart from the ground isn't the ground's.
func TestFitApartFromGround(t *testing.T) {
	t.Parallel()
	for _, tn := range tint.DefaultTints()[:200] {
		on := Derive(tn).Fit(Target{Profile: colorprofile.ANSI256}).On(Control)
		c := tint.FromColor(on.Color(Control))
		if got := on.Fit(*c, Bg, Control); got == on.Color(Control) {
			t.Errorf("%s: Fit kept apart from the control returned its index %v", tn.ID, got)
		}
	}
}

// A crossfade between tints with the same background and selections of
// their own passes through the selections between, not straight to the
// destination's.
func TestLerpSelection(t *testing.T) {
	t.Parallel()
	ta, tb := *tint.TintTokyoNight, *tint.TintTokyoNight
	ta.SelectionBg, tb.SelectionBg = tint.FromHex("#2a3a6a"), tint.FromHex("#4a2a5a")
	a, b := Derive(&ta), Derive(&tb)
	if a.Color(Bg.Selected()) != *ta.SelectionBg || b.Color(Bg.Selected()) != *tb.SelectionBg {
		t.Fatalf("the tints' own selections aren't used: %s, %s", a.Color(Bg.Selected()).Hex(), b.Color(Bg.Selected()).Hex())
	}
	mid := LerpPalette(a, b, 0.5).Color(Bg.Selected())
	if mid == b.Color(Bg.Selected()) || mid == a.Color(Bg.Selected()) {
		t.Errorf("mid-fade selection is %s, an endpoint", mid.Hex())
	}
	// The endpoints are the palettes Derive returned, for a palette from On
	// too, as every frame between is.
	if LerpPalette(a, b.On(Raised), 1) != b || LerpPalette(a.On(Raised), b, 0) != a {
		t.Error("LerpPalette's endpoints aren't the root palettes")
	}
}

// A defined role can read the background's selection while it's derived.
func TestDefinedReadsSelection(t *testing.T) {
	t.Parallel()
	sel := Define(Spec{Name: "copy-selection", Kind: KindSurface, Derive: func(p *Palette) tint.Color {
		return p.Color(Bg.Selected())
	}})
	p := Derive(tint.TintTokyoNight)
	if p.Color(sel) != p.Color(Bg.Selected()) || p.Color(sel) == (tint.Color{A: 0xff}) {
		t.Errorf("a defined role read the selection as %s, want %s", p.Color(sel).Hex(), p.Color(Bg.Selected()).Hex())
	}
}
