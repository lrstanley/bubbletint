// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"image/color"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
	tint "github.com/lrstanley/bubbletint/v2"
)

func TestAttrApply(t *testing.T) {
	t.Parallel()
	all := Bold | Faint | Italic | Underline | Reverse | Strikethrough
	s := all.Apply(lipgloss.NewStyle())
	if !s.GetBold() || !s.GetFaint() || !s.GetItalic() || !s.GetUnderline() || !s.GetReverse() || !s.GetStrikethrough() {
		t.Errorf("Apply(%v) didn't set every attribute", all)
	}

	// Attributes that aren't in a are left unset, so they're inherited.
	parent := lipgloss.NewStyle().Bold(true).Reverse(true)
	child := Faint.Apply(lipgloss.NewStyle()).Inherit(parent)
	if !child.GetFaint() || !child.GetBold() || !child.GetReverse() {
		t.Errorf("Faint.Apply().Inherit() = faint %v, bold %v, reverse %v, want all true",
			child.GetFaint(), child.GetBold(), child.GetReverse())
	}
}

func TestStyles(t *testing.T) {
	t.Parallel()
	pal := Derive(tint.TintTokyoNight)
	sel := Bg.Selected()
	for _, d := range []colorprofile.Profile{colorprofile.TrueColor, colorprofile.ANSI256, colorprofile.ANSI, colorprofile.ASCII} {
		f := pal.Fit(Target{Profile: d})
		row := f.On(sel)
		tests := []struct {
			name   string
			style  lipgloss.Style
			fg, bg color.Color
			attr   Attr
		}{
			{name: "text", style: f.Text(Muted), fg: f.Color(Muted), attr: f.Attr(Muted)},
			{name: "tone", style: f.Fill(Danger), fg: f.On(Danger).Color(Fg), bg: f.Color(Danger), attr: f.On(Danger).Attr(Fg) | f.fillOf(Danger).a},
			{name: "on", style: row.Text(Muted), fg: row.Color(Muted), attr: row.Attr(Muted) | f.Attr(sel)},
			{name: "span", style: row.Span(Muted), fg: row.Color(Muted), bg: f.Color(sel), attr: row.Attr(Muted) | f.Attr(sel)},
			{name: "soft", style: f.Fill(Danger.Soft()), fg: f.On(Danger.Soft()).Color(Danger), bg: f.Color(Danger.Soft()), attr: f.On(Danger.Soft()).Attr(Danger)},
			{name: "selected", style: f.Fill(sel), fg: row.Color(Fg), bg: f.Color(sel), attr: f.Attr(sel) | row.Attr(Fg)},
			{name: "hover", style: f.Fill(Overlay.Hover()), fg: f.On(Overlay.Hover()).Color(Fg), bg: f.Color(Overlay.Hover()), attr: f.Attr(Overlay.Hover())},
			{name: "control", style: f.On(Raised).Fill(Control), fg: f.On(Raised).On(Control).Color(Fg), bg: f.On(Raised).Color(Control), attr: f.On(Raised).Attr(Control)},
			{name: "defined", style: f.Fill(testShadow), fg: f.On(testShadow).Color(Fg), bg: f.Color(testShadow), attr: f.Attr(testShadow) | f.On(testShadow).Attr(Fg)},
		}
		for _, tt := range tests {
			if tt.fg != nil && tt.style.GetForeground() != tt.fg {
				t.Errorf("%v %s: foreground %v, want %v", d, tt.name, tt.style.GetForeground(), tt.fg)
			}
			if tt.bg != nil && tt.style.GetBackground() != tt.bg {
				t.Errorf("%v %s: background %v, want %v", d, tt.name, tt.style.GetBackground(), tt.bg)
			}
			if got := attrOf(tt.style); got != tt.attr {
				t.Errorf("%v %s: attributes %v, want %v", d, tt.name, got, tt.attr)
			}
		}
		// Text sets no background, on the page or on a fill.
		for _, st := range []lipgloss.Style{f.Text(Muted), row.Text(Muted), f.On(Danger).Text(Fg)} {
			if _, ok := st.GetBackground().(lipgloss.NoColor); !ok {
				t.Errorf("%v: Text has background %v", d, st.GetBackground())
			}
		}
	}

	// With no color, the attributes carry each role.
	f := pal.Fit(Target{Profile: colorprofile.ASCII})
	if !f.Text(Muted).GetFaint() || !f.Fill(sel).GetReverse() || !f.Fill(Danger).GetReverse() || !f.On(sel).Text(Fg).GetReverse() {
		t.Error("with no color, muted text isn't faint, or a fill or the text on it isn't reversed")
	}
	if !f.Fill(Bg.Hover()).GetUnderline() || !f.Fill(Control).GetReverse() {
		t.Error("with no color, a hovered row isn't underlined, or a control isn't reversed")
	}
}

// A fill's text reads on it, for a text role too, and on a palette from On.
func TestFillTextRoles(t *testing.T) {
	t.Parallel()
	f := Derive(tint.TintTokyoNight).Fit(Target{Profile: colorprofile.TrueColor})
	for _, tt := range []struct {
		name  string
		st    lipgloss.Style
		floor float64
	}{
		{name: "muted", st: f.Fill(Muted), floor: minText},
		{name: "muted on fg", st: f.On(Fg).Fill(Muted), floor: minText},
		{name: "control on raised", st: f.On(Raised).Fill(Control), floor: minText},
		// A pill's text reads as well as the tone does on the page.
		{name: "accent", st: f.Fill(Accent), floor: minColor},
		{name: "neutral", st: f.Fill(Neutral), floor: minColor},
	} {
		if got := tint.Contrast(tt.st.GetForeground(), tt.st.GetBackground()); got < tt.floor {
			t.Errorf("%s: text contrast %.2f < %.1f", tt.name, got, tt.floor)
		}
	}
}

// Fill(Fg) is reverse video, so it reads with the terminal's own colors too.
func TestFillFg(t *testing.T) {
	t.Parallel()
	p := Derive(tint.TintTokyoNight)
	for _, tg := range []Target{
		{Profile: colorprofile.TrueColor},
		{Profile: colorprofile.TrueColor, Unpainted: true},
		{Profile: colorprofile.ANSI},
		{Profile: colorprofile.ASCII},
	} {
		f := p.Fit(tg)
		st := f.Fill(Fg)
		if !st.GetReverse() {
			t.Errorf("%v unpainted %v: Fill(Fg) isn't reversed", tg.Profile, tg.Unpainted)
		}
		if fg, bg := f.Color(Fg), f.Color(Bg); fg != nil && bg != nil && tint.Contrast(fg, bg) < minText {
			t.Errorf("unpainted %v: Fill(Fg) contrast %.2f", tg.Unpainted, tint.Contrast(fg, bg))
		}
	}
}

// attrOf returns the attributes set on a style.
func attrOf(s lipgloss.Style) Attr {
	var a Attr
	for flag, set := range map[Attr]bool{
		Bold: s.GetBold(), Faint: s.GetFaint(), Italic: s.GetItalic(),
		Underline: s.GetUnderline(), Reverse: s.GetReverse(), Strikethrough: s.GetStrikethrough(),
	} {
		if set {
			a |= flag
		}
	}
	return a
}

func TestOnCached(t *testing.T) {
	t.Parallel()
	f := Derive(tint.TintTokyoNight).Fit(Target{Profile: colorprofile.ANSI256})
	if f.On(Bg) != f {
		t.Error("On(Bg) isn't the palette itself")
	}
	row := f.On(Bg.Selected())
	if f.On(Bg.Selected()) != row || row.On(Bg.Selected()) != row {
		t.Error("On(selected) isn't cached, or On of its own ground isn't itself")
	}
	// On of a layer is shared, however it's reached.
	if row.On(Raised) != f.On(Raised) || row.On(Bg) != f {
		t.Error("On of a layer isn't shared")
	}
	// A control on a dialog is the dialog's, not the page's.
	dialog := f.On(Raised)
	button := dialog.On(Control)
	if button == f.On(Control) || dialog.On(Control) != button {
		t.Error("On(control) on a dialog is the page's, or isn't cached")
	}
	// Span on the page draws the page's background; Text doesn't.
	if _, ok := f.Text(Muted).GetBackground().(lipgloss.NoColor); !ok {
		t.Errorf("Text on Bg has background %v", f.Text(Muted).GetBackground())
	}
	if f.Span(Muted).GetBackground() != f.Color(Bg) {
		t.Error("Span on Bg isn't on the background")
	}
	if unpainted := f.Palette().Fit(Target{Profile: colorprofile.TrueColor, Unpainted: true}); unpainted.Span(Muted).GetBackground() != (lipgloss.NoColor{}) {
		t.Error("Span on an unpainted background paints it")
	}
}

// With 16 colors, muted text on a surface never takes the gray body text is
// drawn in, unless it's faint, which would make it read as body text.
func TestMutedOnSurface16(t *testing.T) {
	t.Parallel()
	for _, tn := range tint.DefaultTints() {
		base := Derive(tn).Fit(Target{Profile: colorprofile.ANSI})
		for _, surface := range []Role{Bg, Raised, Bg.Selected()} {
			f := base.On(surface)
			body := f.p.termFg
			if s, ok := f.Color(Fg).(ansi.BasicColor); ok {
				body = f.ansi[s]
			}
			for _, r := range []Role{Muted, Subtle} {
				// Faint body text is fine: that's the fallback.
				if s, ok := f.Color(r).(ansi.BasicColor); ok && f.ansi[s].Distance(body) < 0.04 && f.Attr(r)&Faint == 0 {
					t.Errorf("%s on %s: %s is slot %d, the same as body text", tn.ID, surface, r, s)
				}
			}
		}
	}
}

// A surface's fill has text that reads on the fill as fitted, not the page's
// text: with 16 colors the selection is a slot the terminal's default text
// may not read on.
func TestFillSurfaceText16(t *testing.T) {
	t.Parallel()
	f := Derive(tint.TintTokyoNight).Fit(Target{Profile: colorprofile.ANSI})
	st := f.Fill(Bg.Selected())
	bg, ok1 := st.GetBackground().(ansi.BasicColor)
	fg, ok2 := st.GetForeground().(ansi.BasicColor)
	if !ok1 || !ok2 {
		t.Fatalf("Fill(selected) = %v on %v, want two slots", st.GetForeground(), st.GetBackground())
	}
	if got := tint.Contrast(f.ansi[fg], f.ansi[bg]); got < minText {
		t.Errorf("Fill(selected) text contrast %.2f < %.1f", got, minText)
	}
}
