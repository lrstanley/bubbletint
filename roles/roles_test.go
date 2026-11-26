// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"image/color"
	"testing"

	"github.com/charmbracelet/x/ansi"
	tint "github.com/lrstanley/bubbletint/v2"
)

// testControl is a defined role, as an app would add one.
var testControl = Define(Spec{
	Name: "test-control",
	Kind: KindSurface,
	Derive: func(p *Palette) tint.Color {
		fg := p.Color(Fg)
		return p.Color(Raised).MixReadable(fg, fg, p.ByDark(0.22, 0.17), 4.5)
	},
	Distinct: []Role{Raised},
	Fallback: Fallback{Slots: []tint.Slot{tint.SlotBrightBlack}},
})

// testHint is a defined text role with a floor.
var testHint = Define(Spec{
	Name:     "test-hint",
	Kind:     KindTone, // taken as KindText
	Derive:   func(p *Palette) tint.Color { return p.Color(Fg).Mix(p.Color(Bg), 0.7) },
	Min:      3,
	Fallback: Fallback{Slots: []tint.Slot{tint.SlotBrightBlack, tint.SlotWhite}, Attr: Italic},
})

// textRoles are the roles with a contrast floor.
var textRoles = []Role{Fg, Muted, Subtle, Border, Accent, Accent2, Success, Warning, Danger, Info, Red, Green, Yellow, Blue, Magenta, Cyan, testHint}

func TestRoleNames(t *testing.T) {
	t.Parallel()
	tests := []struct {
		r    Role
		want string
		kind Kind
	}{
		{r: Bg, want: "bg", kind: KindSurface},
		{r: Selection, want: "selection", kind: KindSurface},
		{r: BorderStrong, want: "border-strong", kind: KindText},
		{r: Danger, want: "danger", kind: KindTone},
		{r: Danger.Soft(), want: "danger-soft", kind: KindSoft},
		{r: Danger.On(), want: "on-danger", kind: KindOn},
		{r: Cyan.On(), want: "on-cyan", kind: KindOn},
		{r: testControl, want: "test-control", kind: KindSurface},
		{r: testHint, want: "test-hint", kind: KindText},
		{r: Role(60000), want: "role(60000)", kind: KindText},
	}
	for _, tt := range tests {
		if got := tt.r.String(); got != tt.want {
			t.Errorf("String() = %q, want %q", got, tt.want)
		}
		if got := tt.r.Kind(); got != tt.kind {
			t.Errorf("%s: Kind() = %v, want %v", tt.r, got, tt.kind)
		}
	}
	if Muted.Soft() != Muted || Muted.On() != Muted || Bg.Base() != Bg {
		t.Error("Soft, On and Base of a non-tone should return it")
	}
	if Accent.Soft().Base() != Accent || Accent.On().Base() != Accent {
		t.Error("Base() doesn't undo Soft() and On()")
	}
	for _, s := range tint.ChromaticSlots {
		if Hue(s) != Hue(s.Bright()) || !Hue(s).IsTone() {
			t.Errorf("Hue(%v) = %v", s, Hue(s))
		}
	}
	if Hue(tint.SlotWhite) != Fg {
		t.Error("Hue(white) != Fg")
	}
}

func TestDeriveFloors(t *testing.T) {
	t.Parallel()
	var fails, checked int
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		bg := p.Color(Bg)
		for _, r := range textRoles {
			checked++
			if got := tint.Contrast(p.Color(r), bg); got < minFor(r)-0.01 {
				fails++
				if fails < 10 {
					t.Errorf("%s: %s contrast %.2f < %.2f", tn.ID, r, got, minFor(r))
				}
			}
		}
		for _, r := range []Role{Accent, Success, Warning, Danger, Info} {
			checked++
			if got := tint.Contrast(p.Color(r.On()), p.Color(r)); got < minColor {
				fails++
				t.Errorf("%s: %s contrast %.2f on %s", tn.ID, r.On(), got, r)
			}
		}
		// Text reads on the selection, and as well on a control as on the
		// dialog it sits on.
		checked += 2
		if got := tint.Contrast(p.Color(Fg), p.Color(Selection)); got < minText-0.01 {
			fails++
			t.Errorf("%s: fg contrast %.2f on the selection", tn.ID, got)
		}
		want := min(minText, tint.Contrast(p.Color(Fg), p.Color(Raised)))
		if got := tint.Contrast(p.Color(Fg), p.Color(testControl)); got < want-0.01 {
			fails++
			t.Errorf("%s: fg contrast %.2f on test-control, want %.2f", tn.ID, got, want)
		}
	}
	t.Logf("checked %d contrasts", checked)
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
		if p.Color(Hover) == p.Color(Bg) {
			t.Error("hover == bg")
		}
	})
	t.Run("accent-slot", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithAccentSlot(tint.SlotBrightGreen))
		if p.AccentSlot() != tint.SlotGreen {
			t.Errorf("AccentSlot() = %v", p.AccentSlot())
		}
		if p.Accent2Slot() == tint.SlotGreen {
			t.Error("accent2 shares the accent's slot")
		}
	})
	t.Run("accent-color", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithAccentColor(tint.FromHex("#ff8800")))
		if p.AccentSlot() != tint.SlotYellow && p.AccentSlot() != tint.SlotRed {
			t.Errorf("AccentSlot() = %v, want the slot nearest orange", p.AccentSlot())
		}
	})
	t.Run("terminal", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithTerminal(color.Black, nil))
		if p.Color(Bg) != (tint.Color{A: 0xff}) {
			t.Errorf("Bg = %s", p.Color(Bg).Hex())
		}
		if !p.Dark() {
			t.Error("not dark on black")
		}
	})
	t.Run("without-guards", func(t *testing.T) {
		t.Parallel()
		p := Derive(tn, WithoutGuards())
		if p.Guarded() {
			t.Error("Guarded()")
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
	t.Run("mono", func(t *testing.T) {
		t.Parallel()
		if Derive(tn).Mono() {
			t.Error("tokyo night is not mono")
		}
		if !Derive(&tint.Tint{Bg: tint.FromHex("#000000"), Fg: tint.FromHex("#ffb000")}).Mono() {
			t.Error("a tint with no hues should be mono")
		}
	})
}

func TestPaletteOn(t *testing.T) {
	t.Parallel()
	var fails int
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		for _, s := range []Role{Hover, Selection, Raised} {
			on := p.On(s)
			under := p.Color(s)
			// Some surfaces are too mid-toned for even black or white to
			// reach the floor; there, text goes as far as it can.
			best := max(tint.Contrast(tint.Color{A: 0xff}, under), tint.Contrast(tint.Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, under))
			for _, r := range textRoles {
				if got := tint.Contrast(on.Color(r), under); got < min(minFor(r), best)-0.01 {
					fails++
					if fails < 10 {
						t.Errorf("%s: %s on %s contrast %.2f", tn.ID, r, s, got)
					}
				}
			}
			if on.Color(s) != under || on.Color(Danger.Soft()) != p.Color(Danger.Soft()) {
				t.Fatalf("%s: On(%s) changed a surface", tn.ID, s)
			}
		}
	}
}

func TestLerp(t *testing.T) {
	t.Parallel()
	a, b := Derive(tint.TintTokyoNight), Derive(tint.TintDracula)
	light := Derive(tint.TintCatppuccinLatte)
	if Lerp(a, light, 0.5) != light {
		t.Error("Lerp across dark and light should return b")
	}
	if Lerp(a, b, 0) != a || Lerp(a, b, 1) != b || Lerp(nil, b, 0.5) != b {
		t.Error("Lerp endpoints")
	}
	mid := Lerp(a, b, 0.5)
	if mid.Color(Bg) == a.Color(Bg) || mid.Color(Bg) == b.Color(Bg) {
		t.Error("Lerp didn't mix")
	}
	for _, r := range textRoles {
		if got := tint.Contrast(mid.Color(r), mid.Color(Bg)); got < minFor(r)-0.01 {
			t.Errorf("%s: contrast %.2f mid-crossfade", r, got)
		}
	}
}

func TestDefineLate(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintTokyoNight)
	late := Define(Spec{Name: "late", Derive: func(p *Palette) tint.Color { return p.Color(Accent) }})
	if p.Color(late) != p.Color(Accent) {
		t.Error("a role defined after Derive isn't derived on demand")
	}
	if f := p.Fit(Target{Depth: Colors256}); f.Color(late) == nil {
		t.Error("a role defined after Fit isn't fitted on demand")
	}
}

// pairCollapses counts pairs fitted to the same color, though derived apart.
func pairCollapses(p *Palette, f *Fitted) (n int, which []string) {
	for _, pr := range pairs {
		if f.Color(pr[0]) == f.Color(pr[1]) && p.Color(pr[0]).Distance(p.Color(pr[1])) > 0.004 {
			n++
			which = append(which, pr[0].String()+"="+pr[1].String())
		}
	}
	if f.Color(testControl) == f.Color(Raised) {
		n++
		which = append(which, "test-control=raised")
	}
	return n, which
}

func TestFit256(t *testing.T) {
	t.Parallel()
	var collapses, floors int
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		f := p.Fit(Target{Depth: Colors256})
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
	}
	// A tint or two has no 256-color index that keeps every floor and
	// pair; nearly all do.
	if collapses > 3 || floors > 3 {
		t.Errorf("collapses %d, broken floors %d", collapses, floors)
	}
}

func TestFit16(t *testing.T) {
	t.Parallel()
	var attrs int
	for _, tn := range tint.DefaultTints() {
		p := Derive(tn)
		f := p.Fit(Target{Depth: Colors16})
		bg := p.Color(Bg)
		ansi16 := tn.ANSI16()
		for _, r := range append(textRoles, Selection) {
			c := f.Color(r)
			if c == nil {
				if r != Fg && r != Border && f.Attr(r) == 0 {
					t.Errorf("%s: %s is the default with no attribute", tn.ID, r)
				}
				attrs++
				continue
			}
			s, ok := c.(ansi.BasicColor)
			if !ok {
				t.Fatalf("%s: %s = %#v, want a BasicColor", tn.ID, r, c)
			}
			if r == Selection {
				continue
			}
			if got := tint.Contrast(*ansi16[s], bg); got < minFor(r)-0.01 {
				t.Errorf("%s: %s slot %d contrast %.2f < %.2f", tn.ID, r, s, got, minFor(r))
			}
			if r.IsTone() {
				want := f.toneSlot(r)
				if tint.Slot(s).Normal() != want {
					t.Errorf("%s: %s took slot %d, want %v's", tn.ID, r, s, want)
				}
			}
		}
		if f.Color(Bg) != nil || f.Color(Fg) != nil || f.Color(Raised) != nil {
			t.Errorf("%s: surfaces painted with 16 colors", tn.ID)
		}
		// Text on the selection reads against the selection's slot.
		if sel, ok := f.Color(Selection).(ansi.BasicColor); ok {
			on := f.On(Selection)
			for _, r := range []Role{Fg, Muted} {
				s, isBasic := on.Color(r).(ansi.BasicColor)
				if !isBasic {
					t.Errorf("%s: %s on the selection is %#v", tn.ID, r, on.Color(r))
					continue
				}
				if s == sel {
					t.Errorf("%s: %s on the selection is the selection's slot", tn.ID, r)
				}
			}
		}
	}
	t.Logf("%d roles fell back to an attribute", attrs)
}

func TestFitOther(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintGruvboxDark)

	tc := p.Fit(Target{Depth: TrueColor})
	if tc.Color(Accent) != p.Color(Accent) || tc.Attr(Accent) != 0 {
		t.Error("true color should be the derived color")
	}
	if tc.Fit(color.White, Fg) != (tint.Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}) {
		t.Error("Fit() at true color")
	}

	unpainted := p.Fit(Target{Depth: TrueColor, Unpainted: true})
	if unpainted.Color(Bg) != nil || unpainted.Color(Fg) != nil || unpainted.Color(Surface) == nil {
		t.Error("Unpainted should leave only bg and fg to the terminal")
	}

	none := p.Fit(Target{Depth: NoColor})
	for i := range numBuiltin {
		if none.Color(i) != nil {
			t.Fatalf("%s has a color with NoColor", i)
		}
	}
	for r, want := range map[Role]Attr{Muted: Faint, Selection: Reverse, Danger: Bold, Danger.On(): Reverse | Bold, testHint: Italic, Bg: 0} { //nolint:exhaustive // a sample
		if got := none.Attr(r); got != want {
			t.Errorf("NoColor %s: Attr() = %v, want %v", r, got, want)
		}
	}

	f256 := p.Fit(Target{Depth: Colors256})
	if got := f256.Fit(p.Color(Accent).Mix(p.Color(Raised), 0.3), Accent); got == nil {
		t.Error("Fit() at 256")
	}
	if f256.Depth() != Colors256 || f256.Palette() != p {
		t.Error("Depth() or Palette()")
	}

	// A terminal palette that isn't the tint's: everything comes from it.
	var term [16]color.Color
	for i := range term {
		term[i] = ansi.BasicColor(i) //nolint:gosec // 0-15
	}
	f16 := p.Fit(Target{Depth: Colors16, ANSI: &term})
	if s, ok := f16.Color(Danger).(ansi.BasicColor); !ok || tint.Slot(s).Normal() != tint.SlotRed {
		t.Errorf("Danger = %#v", f16.Color(Danger))
	}
	if got := f16.Fit(color.White, Danger); got != f16.Color(Danger) {
		t.Error("Fit() at 16 should be like's color")
	}
}

func TestDepthString(t *testing.T) {
	t.Parallel()
	for d, want := range map[Depth]string{TrueColor: "truecolor", Colors256: "256", Colors16: "16", NoColor: "none", 9: "unknown"} {
		if got := d.String(); got != want {
			t.Errorf("String() = %q, want %q", got, want)
		}
	}
}

func BenchmarkDerive(b *testing.B) {
	for b.Loop() {
		Derive(tint.TintTokyoNight)
	}
}

func BenchmarkFit256(b *testing.B) {
	p := Derive(tint.TintTokyoNight)
	for b.Loop() {
		p.Fit(Target{Depth: Colors256})
	}
}

func TestFitEdgeCases(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintTokyoNight)

	// Text on a solid tone, rather than a surface. The tone itself stays as
	// it was, and with 16 colors gets the same slot as on the background.
	if got := p.On(Danger).Color(Danger); got != p.Color(Danger) {
		t.Errorf("On(Danger) changed danger to %s", got.Hex())
	}
	on, base := p.On(Danger).Fit(Target{Depth: Colors16}), p.Fit(Target{Depth: Colors16})
	if on.Color(Fg) == nil {
		t.Error("fg on a danger fill is the terminal's default")
	}
	if on.Color(Danger) != base.Color(Danger) || on.Color(Danger.On()) != base.Color(Danger.On()) {
		t.Errorf("danger %v / %v on its own fill, want %v / %v", on.Color(Danger), on.Color(Danger.On()), base.Color(Danger), base.Color(Danger.On()))
	}

	// A nil *tint.Color behind the color.Color interface.
	var none *tint.Color
	var term [16]color.Color
	term[3] = none
	if f := p.Fit(Target{Depth: Colors16, ANSI: &term}); f.Color(Warning) == nil {
		t.Error("Warning fell back to the default with a typed nil slot")
	}
	if got := p.Fit(Target{Depth: TrueColor}).Fit(none, Fg); got != nil {
		t.Errorf("Fit(typed nil) = %v, want nil", got)
	}

	// Text on a tone keeps its contrast with 256 colors, and across a fade.
	f := p.Fit(Target{Depth: Colors256})
	mid := Lerp(p, Derive(tint.TintDracula), 0.5)
	for _, r := range []Role{Accent, Success, Warning, Danger, Info} {
		if got := tint.Contrast(f.Color(r.On()), f.Color(r)); got < minColor {
			t.Errorf("256: %s contrast %.2f on %s", r.On(), got, r)
		}
		if got := tint.Contrast(mid.Color(r.On()), mid.Color(r)); got < minColor {
			t.Errorf("Lerp: %s contrast %.2f on %s", r.On(), got, r)
		}
	}
}
