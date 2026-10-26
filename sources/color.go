// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package sources

import (
	"fmt"
	"strconv"
	"strings"

	tint "github.com/lrstanley/bubbletint/v2"
)

// ParseColor parses a color as written in terminal configs. It supports
// "#rgb", "#rrggbb" (a trailing alpha channel is ignored), "rrggbb", "0xrrggbb",
// and the X11 "rgb:r/g/b" form (1-4 hex digits per channel). Color names are not
// supported.
func ParseColor(s string) (*tint.Color, bool) {
	s = strings.Trim(strings.TrimSpace(s), `"'`)
	if s == "" {
		return nil, false
	}

	if rest, ok := strings.CutPrefix(strings.ToLower(s), "rgb:"); ok {
		return parseX11(rest)
	}

	switch {
	case s[0] == '#':
		s = s[1:]
	case len(s) > 2 && (s[:2] == "0x" || s[:2] == "0X"):
		s = s[2:]
	}

	switch len(s) {
	case 3:
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	case 6:
	case 8:
		s = s[:6]
	default:
		return nil, false
	}

	v, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return nil, false
	}
	return &tint.Color{R: uint8(v >> 16 & 0xff), G: uint8(v >> 8 & 0xff), B: uint8(v & 0xff), A: 0xff}, true
}

func parseX11(s string) (*tint.Color, bool) {
	parts := strings.Split(s, "/")
	if len(parts) != 3 {
		return nil, false
	}

	var rgb [3]uint8
	for i, p := range parts {
		if len(p) < 1 || len(p) > 4 {
			return nil, false
		}
		v, err := strconv.ParseUint(p, 16, 16)
		if err != nil {
			return nil, false
		}
		// Scale to 8 bits, e.g. "f" -> 0xff, "ff" -> 0xff, "ffff" -> 0xff.
		maxV := uint64(1)<<(4*uint(len(p))) - 1
		rgb[i] = uint8(min((v*255+maxV/2)/maxV, 255))
	}
	return &tint.Color{R: rgb[0], G: rgb[1], B: rgb[2], A: 0xff}, true
}

// ID converts a name, such as a theme name, to a tint ID: lowercase, with
// runs of anything other than letters and digits replaced with an underscore
// (e.g. "One Half Dark" -> "one_half_dark").
func ID(name string) string {
	var b strings.Builder
	underscore := false
	for _, r := range strings.ToLower(strings.TrimSpace(name)) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r > 127 {
			if underscore && b.Len() > 0 {
				b.WriteByte('_')
			}
			underscore = false
			b.WriteRune(r)
			continue
		}
		underscore = true
	}
	return b.String()
}

// Palette is a terminal's colors, as read from its config. Every field is
// optional. Use [Palette.Tint] to turn it into a [tint.Tint].
type Palette struct {
	Foreground          *tint.Color
	Background          *tint.Color
	SelectionForeground *tint.Color
	SelectionBackground *tint.Color
	Cursor              *tint.Color
	CursorText          *tint.Color
	Bold                *tint.Color

	// ANSI holds the 16 ANSI colors: black, red, green, yellow, blue, magenta,
	// cyan, white, followed by their bright variants.
	ANSI [16]*tint.Color
}

// PaletteOf returns the colors of an existing tint.
func PaletteOf(t *tint.Tint) Palette {
	return Palette{
		Foreground:          t.Fg,
		Background:          t.Bg,
		SelectionForeground: t.SelectionText,
		SelectionBackground: t.SelectionBg,
		Cursor:              t.Cursor,
		CursorText:          t.CursorText,
		Bold:                t.Bold,
		ANSI: [16]*tint.Color{
			t.Black, t.Red, t.Green, t.Yellow, t.Blue, t.Purple, t.Cyan, t.White,
			t.BrightBlack, t.BrightRed, t.BrightGreen, t.BrightYellow,
			t.BrightBlue, t.BrightPurple, t.BrightCyan, t.BrightWhite,
		},
	}
}

// Empty reports whether the palette doesn't define any colors at all.
func (p *Palette) Empty() bool {
	for _, c := range []*tint.Color{
		p.Foreground, p.Background, p.SelectionForeground, p.SelectionBackground,
		p.Cursor, p.CursorText, p.Bold,
	} {
		if c != nil {
			return false
		}
	}
	for _, c := range p.ANSI {
		if c != nil {
			return false
		}
	}
	return true
}

// SetANSI sets the ANSI color at index i (0-15), ignoring out of range indexes.
func (p *Palette) SetANSI(i int, c *tint.Color) {
	if i >= 0 && i < len(p.ANSI) && c != nil {
		p.ANSI[i] = c
	}
}

// defaultANSI is used for any ANSI colors a config doesn't set: the xterm
// defaults, which most terminals' own defaults are close to.
var defaultANSI = [16]string{
	"#000000", "#cd0000", "#00cd00", "#cdcd00", "#0000ee", "#cd00cd", "#00cdcd", "#e5e5e5",
	"#7f7f7f", "#ff0000", "#00ff00", "#ffff00", "#5c5cff", "#ff00ff", "#00ffff", "#ffffff",
}

// Tint converts the palette to a tint. A config often doesn't set every color,
// so anything missing is filled in: the foreground and background with whatever
// contrasts with the other (or a dark scheme if neither is set), and the ANSI
// colors with the xterm defaults. Optional colors that aren't set (selection,
// cursor, bold) are left nil, as they are for most built-in tints.
func (p *Palette) Tint(id, displayName string) *tint.Tint {
	bg, fg := p.Background, p.Foreground

	switch {
	case bg == nil && fg == nil:
		bg, _ = ParseColor("#000000")
		fg, _ = ParseColor("#e5e5e5")
	case bg == nil:
		bg = contrast(fg)
	case fg == nil:
		fg = contrast(bg)
	}

	ansi := p.ANSI
	for i := range ansi {
		if ansi[i] == nil {
			ansi[i], _ = ParseColor(defaultANSI[i])
		}
	}

	return &tint.Tint{
		ID:            id,
		DisplayName:   displayName,
		Dark:          luma(bg) < 0.5,
		Fg:            fg,
		Bg:            bg,
		SelectionBg:   p.SelectionBackground,
		SelectionText: p.SelectionForeground,
		Bold:          p.Bold,
		Cursor:        p.Cursor,
		CursorText:    p.CursorText,
		Black:         ansi[0],
		Red:           ansi[1],
		Green:         ansi[2],
		Yellow:        ansi[3],
		Blue:          ansi[4],
		Purple:        ansi[5],
		Cyan:          ansi[6],
		White:         ansi[7],
		BrightBlack:   ansi[8],
		BrightRed:     ansi[9],
		BrightGreen:   ansi[10],
		BrightYellow:  ansi[11],
		BrightBlue:    ansi[12],
		BrightPurple:  ansi[13],
		BrightCyan:    ansi[14],
		BrightWhite:   ansi[15],
	}
}

// luma returns the perceived brightness of c, from 0 (black) to 1 (white).
func luma(c *tint.Color) float64 {
	return (0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B)) / 255
}

// contrast returns black or white, whichever is easier to read on c.
func contrast(c *tint.Color) *tint.Color {
	if luma(c) < 0.5 {
		return &tint.Color{R: 0xe5, G: 0xe5, B: 0xe5, A: 0xff}
	}
	return &tint.Color{A: 0xff}
}

// ErrNoColors returns an error wrapping [tint.ErrNotFound], for configs that
// exist but don't set any colors, so there's nothing to follow.
func ErrNoColors(terminal, path string) error {
	return fmt.Errorf("%w: %s does not configure any %s colors", tint.ErrNotFound, path, terminal)
}
