// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"charm.land/lipgloss/v2"
)

// Apply adds the attributes in a to s. Only those in a are set, so a style
// that inherits from the result, with [lipgloss.Style.Inherit], still picks up
// the rest from its parent.
func (a Attr) Apply(s lipgloss.Style) lipgloss.Style {
	if a&Bold != 0 {
		s = s.Bold(true)
	}
	if a&Faint != 0 {
		s = s.Faint(true)
	}
	if a&Italic != 0 {
		s = s.Italic(true)
	}
	if a&Underline != 0 {
		s = s.Underline(true)
	}
	if a&Reverse != 0 {
		s = s.Reverse(true)
	}
	if a&Strikethrough != 0 {
		s = s.Strikethrough(true)
	}
	return s
}

// Text returns a style for text in a role's color, such as [Muted] text or a
// [Danger] mark, with the attributes that stand in for the color where the
// terminal can't draw it. On a palette for drawing on a surface, from
// [Fitted.On], the text is on the surface's color too, so it doesn't leave a
// hole in the fill around it.
func (f *Fitted) Text(r Role) lipgloss.Style {
	s := lipgloss.NewStyle().Foreground(f.Color(r))
	a := f.Attr(r)
	if g := f.p.ground; g != Bg {
		s = s.Background(f.Color(g))
		a |= f.fillAttr(g)
	}
	return a.Apply(s)
}

// Fill returns a style filled with a role's color, with the text that reads
// on it:
//
//   - a tone, such as a [Danger] pill: the tone's [Role.On] text.
//   - a soft fill, such as Danger.Soft(), or an on color: the tone itself,
//     as text.
//   - a surface: [Fg], fitted to the surface as [Fitted.On] fits it. For
//     other roles' text on the surface, use On's [Fitted.Text].
//
// The attributes that stand in for either color are added. It's meant for
// surfaces and tones; a text role's fill, such as [BorderStrong] for a
// scrollbar thumb, is treated as a surface, and [Fg]'s is reverse video.
func (f *Fitted) Fill(r Role) lipgloss.Style {
	if r == Fg {
		// Reverse video: it reads even where both are the terminal's defaults.
		s := lipgloss.NewStyle().Foreground(f.Color(Fg)).Background(f.Color(Bg))
		return (f.Attr(Fg) | Reverse).Apply(s)
	}
	on, text := f, Fg
	switch r.Kind() {
	case KindTone:
		text = r.On()
	case KindSoft, KindOn:
		text = r.Base()
	case KindSurface, KindText:
		// The fill and its text both come from the palette for drawing on r.
		on = f.On(r)
	}
	s := lipgloss.NewStyle().Background(on.Color(r)).Foreground(on.Color(text))
	return (on.Attr(text) | on.fillAttr(r)).Apply(s)
}

// fillAttr is what a fill of r adds where its color can't be drawn. A tone's
// is its on color's, which [Fitted.Attr] gives for a fill of the tone.
func (f *Fitted) fillAttr(r Role) Attr {
	if r.Kind() == KindTone {
		return f.Attr(r.On())
	}
	return f.Attr(r)
}
