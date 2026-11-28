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
// [Fitted.On], the attributes that mark the surface are added too, such as
// the reverse video of a selected row with no color, so the text doesn't
// break the fill.
//
// It sets no background: what's drawn on keeps its own, as it does in a
// renderer that keeps the background under text, such as bloom's. Where text
// is rendered on its own and then placed on a fill, as with nested
// [lipgloss.Style.Render] calls, use [Fitted.Span], or the reset after the
// text leaves a hole in the fill.
func (f *Fitted) Text(r Role) lipgloss.Style {
	s := lipgloss.NewStyle().Foreground(f.Color(r))
	a := f.Attr(r)
	if f.p.ground != Bg {
		a |= f.ground.a
	}
	return a.Apply(s)
}

// Span returns [Fitted.Text] drawn on the ground's color too, for text
// rendered on its own and placed on a fill. On the background it's the
// background's color, unless that's the terminal's own.
func (f *Fitted) Span(r Role) lipgloss.Style {
	s := f.Text(r)
	if f.ground.c != nil {
		s = s.Background(f.ground.c)
	}
	return s
}

// Fill returns a style filled with a role's color, with the text that reads
// on it:
//
//   - a surface, a tone, or a role in a state, such as [Raised], a [Danger]
//     pill or Bg.Selected(): [Fg], as [Fitted.On] works it out for the fill.
//   - a soft fill, such as Danger.Soft(): the tone itself, as text.
//
// The attributes that stand in for either color are added. For other roles'
// text on the fill, use On's [Fitted.Text]. A text role's fill, such as
// [BorderStrong] for a scrollbar thumb, is treated as a surface, and [Fg]'s
// is reverse video.
func (f *Fitted) Fill(r Role) lipgloss.Style {
	if r == Fg {
		// Reverse video: it reads even where both are the terminal's defaults.
		s := lipgloss.NewStyle().Foreground(f.Color(Fg)).Background(f.Color(Bg))
		return (f.Attr(Fg) | Reverse).Apply(s)
	}
	on, text := f.On(r), Fg
	if r.Kind() == KindSoft {
		text = r.Base().Rest()
	}
	s := lipgloss.NewStyle().Background(on.ground.c).Foreground(on.Color(text))
	return (on.Attr(text) | on.ground.a).Apply(s)
}
