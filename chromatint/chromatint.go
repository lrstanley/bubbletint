// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package chromatint

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/alecthomas/chroma/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// toHex converts a [color.Color] to a hex string.
func toHex(c color.Color) string {
	if c == nil {
		return ""
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}

// StaticStyle is an implementation of [StyleCompat] that can be used to create
// a [chroma.StyleEntries].
type StaticStyle struct {
	Fg        color.Color
	Bg        color.Color
	Italic    bool
	Bold      bool
	Underline bool
	Faint     bool
}

func (s *StaticStyle) GetForeground() color.Color {
	return s.Fg
}

func (s *StaticStyle) GetBackground() color.Color {
	return s.Bg
}

func (s *StaticStyle) GetItalic() bool {
	return s.Italic
}

func (s *StaticStyle) GetBold() bool {
	return s.Bold
}

func (s *StaticStyle) GetUnderline() bool {
	return s.Underline
}

func (s *StaticStyle) GetFaint() bool {
	return s.Faint
}

// StyleCompat is an interface that can be used to create a [chroma.StyleEntries],
// using either a [lipgloss.Style] or a [StaticStyle].
type StyleCompat interface {
	GetForeground() color.Color
	GetBackground() color.Color
	GetItalic() bool
	GetBold() bool
	GetUnderline() bool
	GetFaint() bool
}

// FromStyle converts a [StyleCompat] (either a [lipgloss.Style] or a [StaticStyle])
// to a string that can be used in a [chroma.StyleEntries].
func FromStyle(in StyleCompat, setBg bool) string {
	var s strings.Builder

	if v := in.GetForeground(); v != nil {
		if in.GetFaint() {
			v = tint.Darken(v, 0.3)
		}
		s.WriteString(toHex(v) + " ")
	}
	if v := in.GetBackground(); v != nil && setBg {
		s.WriteString("bg:" + toHex(v) + " ")
	}
	if in.GetItalic() {
		s.WriteString("italic ")
	}
	if in.GetBold() {
		s.WriteString("bold ")
	}
	if in.GetUnderline() {
		s.WriteString("underline ")
	}
	return strings.TrimSpace(s.String())
}

// StyleEntry converts a [tint.Tint] to a [chroma.StyleEntries], with pre-configured
// color matching to different token types. This may not work perfectly for all
// tints, however, you can always call this, then modify the result, or use this
// implementation as a starting point for your own.
func StyleEntry(t *tint.Tint, setBg bool) chroma.StyleEntries {
	if t == nil {
		return chroma.StyleEntries{}
	}

	adapt := func(light, dark color.Color) color.Color {
		if t.Dark {
			return dark
		}
		return light
	}

	compat := func(in *StaticStyle) string {
		return FromStyle(in, setBg)
	}

	return chroma.StyleEntries{ //nolint:exhaustive
		chroma.Other:                    compat(&StaticStyle{Fg: t.Fg}),
		chroma.Error:                    compat(&StaticStyle{Fg: t.Fg}),
		chroma.Background:               compat(&StaticStyle{Bg: t.Bg}),
		chroma.Keyword:                  compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.KeywordConstant:          compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.KeywordDeclaration:       compat(&StaticStyle{Fg: t.Cyan, Italic: true}),
		chroma.KeywordNamespace:         compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.KeywordPseudo:            compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.KeywordReserved:          compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.KeywordType:              compat(&StaticStyle{Fg: t.Cyan}),
		chroma.Name:                     compat(&StaticStyle{Fg: t.Fg}),
		chroma.NameAttribute:            compat(&StaticStyle{Fg: t.Green}),
		chroma.NameBuiltin:              compat(&StaticStyle{Fg: t.Cyan, Italic: true}),
		chroma.NameBuiltinPseudo:        compat(&StaticStyle{Fg: t.Fg}),
		chroma.NameClass:                compat(&StaticStyle{Fg: t.Green}),
		chroma.NameConstant:             compat(&StaticStyle{Fg: t.Fg}),
		chroma.NameDecorator:            compat(&StaticStyle{Fg: t.Fg}),
		chroma.NameEntity:               compat(&StaticStyle{Fg: t.Fg}),
		chroma.NameException:            compat(&StaticStyle{Fg: t.Fg}),
		chroma.NameFunction:             compat(&StaticStyle{Fg: t.Green}),
		chroma.NameLabel:                compat(&StaticStyle{Fg: t.Cyan, Italic: true}),
		chroma.NameNamespace:            compat(&StaticStyle{Fg: t.Fg}),
		chroma.NameOther:                compat(&StaticStyle{Fg: t.Fg}),
		chroma.NameTag:                  compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.NameVariable:             compat(&StaticStyle{Fg: t.Cyan, Italic: true}),
		chroma.NameVariableClass:        compat(&StaticStyle{Fg: t.Cyan, Italic: true}),
		chroma.NameVariableGlobal:       compat(&StaticStyle{Fg: t.Cyan, Italic: true}),
		chroma.NameVariableInstance:     compat(&StaticStyle{Fg: t.Cyan, Italic: true}),
		chroma.Literal:                  compat(&StaticStyle{Fg: t.Fg}),
		chroma.LiteralDate:              compat(&StaticStyle{Fg: t.Fg}),
		chroma.LiteralString:            compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringBacktick:    compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringChar:        compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringDoc:         compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringDouble:      compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringEscape:      compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringHeredoc:     compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringInterpol:    compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringOther:       compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringRegex:       compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringSingle:      compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralStringSymbol:      compat(&StaticStyle{Fg: t.Yellow}),
		chroma.LiteralNumber:            compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4))}),
		chroma.LiteralNumberBin:         compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4))}),
		chroma.LiteralNumberFloat:       compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4))}),
		chroma.LiteralNumberHex:         compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4))}),
		chroma.LiteralNumberInteger:     compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4))}),
		chroma.LiteralNumberIntegerLong: compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4))}),
		chroma.LiteralNumberOct:         compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4))}),
		chroma.Operator:                 compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.OperatorWord:             compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.Punctuation:              compat(&StaticStyle{Fg: t.BrightWhite}),
		chroma.Comment:                  compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4)), Faint: true}),
		chroma.CommentHashbang:          compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4)), Faint: true}),
		chroma.CommentMultiline:         compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4)), Faint: true}),
		chroma.CommentSingle:            compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4)), Faint: true}),
		chroma.CommentSpecial:           compat(&StaticStyle{Fg: adapt(tint.Darken(t.Purple, 0.3), tint.Lighten(t.Purple, 0.4)), Faint: true}),
		chroma.CommentPreproc:           compat(&StaticStyle{Fg: adapt(tint.Darken(t.BrightPurple, 0.3), tint.Lighten(t.BrightPurple, 0.2))}),
		chroma.Generic:                  compat(&StaticStyle{Fg: t.Fg}),
		chroma.GenericDeleted:           compat(&StaticStyle{Fg: t.Red}),
		chroma.GenericEmph:              compat(&StaticStyle{Fg: t.Fg, Underline: true}),
		chroma.GenericError:             compat(&StaticStyle{Fg: t.Fg}),
		chroma.GenericHeading:           compat(&StaticStyle{Fg: t.Fg, Bold: true}),
		chroma.GenericInserted:          compat(&StaticStyle{Fg: t.Green, Bold: true}),
		chroma.GenericOutput:            compat(&StaticStyle{Fg: t.Fg, Faint: true}),
		chroma.GenericPrompt:            compat(&StaticStyle{Fg: t.Fg}),
		chroma.GenericStrong:            compat(&StaticStyle{Fg: t.Fg}),
		chroma.GenericSubheading:        compat(&StaticStyle{Fg: t.Fg, Bold: true}),
		chroma.GenericTraceback:         compat(&StaticStyle{Fg: t.Fg}),
		chroma.GenericUnderline:         compat(&StaticStyle{Underline: true}),
		chroma.Text:                     compat(&StaticStyle{Fg: t.Fg}),
		chroma.TextWhitespace:           compat(&StaticStyle{Fg: t.Fg}),
		chroma.TextPunctuation:          compat(&StaticStyle{Fg: t.BrightWhite}),
	}
}
