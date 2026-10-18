// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package glamourtint builds a Glamour stylesheet from a bubbletint palette.
//
// The layout matches Glamour's built-in dark style: an H1 banner, markdown
// prefixes on the remaining heading levels, and fenced-code colors taken from
// the same palette roles as chromatint. The document background is left unset
// so the host application can paint its own.
package glamourtint

import (
	"sync"

	"charm.land/glamour/v2/ansi"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/lrstanley/bubbletint/chromatint/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

const (
	defaultMargin = 2
	listIndent    = 2

	headingPrefixH2 = "## "
	headingPrefixH3 = "### "
	headingPrefixH4 = "#### "
	headingPrefixH5 = "##### "
	headingPrefixH6 = "###### "

	taskTicked      = "[✓] "
	taskUnticked    = "[ ] "
	imageFormat     = "Image: {{.text}} →"
	horizontalRule  = "\n--------\n"
	definitionArrow = "\n🠶 "
	quoteToken      = "│ "
	itemPrefix      = "• "
	enumPrefix      = ". "
	codePad         = "\u00a0"
)

// themeMu serializes writes to the chroma style registry. Chroma's registry is a
// process-global map, and its reads are not synchronized with writes. Call
// StyleConfig from one goroutine, the same way Bubble Tea runs Update.
var themeMu sync.Mutex

// StyleConfig converts t into a Glamour stylesheet for [glamour.WithStyles].
//
// Fenced code is highlighted with a chroma style registered as
// "bubbletint-<tint id>". Glamour's Chroma field is left unset: when that
// field is set, Glamour copies it into one shared style named "charm" and
// keeps the first registration, so a later tint would keep the first tint's
// code colors.
//
// The document background is left unset so the host can paint its own. Set
// Document.BackgroundColor on the returned value to paint one. A nil tint
// returns the zero StyleConfig. Colors missing from t are skipped.
func StyleConfig(t *tint.Tint) ansi.StyleConfig {
	if t == nil {
		return ansi.StyleConfig{}
	}

	heading := pick(t.Dark, t.Blue, t.BrightBlue)
	if heading == nil {
		heading = firstColor(t.Cyan, t.Fg)
	}

	h1Bg := pick(t.Dark, t.BrightBlue, t.Blue)
	if h1Bg == nil {
		h1Bg = firstColor(t.Purple, t.Fg)
	}

	linkText := pick(t.Dark, t.Cyan, t.BrightCyan)
	if linkText == nil {
		linkText = firstColor(t.Cyan, t.Fg)
	}

	return ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix: "\n",
				BlockSuffix: "\n",
				Color:       colorPtr(t.Fg),
			},
			Margin: new(uint(defaultMargin)),
		},
		BlockQuote: ansi.StyleBlock{
			Indent:      new(uint(1)),
			IndentToken: new(quoteToken),
		},
		List: ansi.StyleList{
			LevelIndent: listIndent,
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockSuffix: "\n",
				Color:       colorPtr(heading),
				Bold:        new(true),
			},
		},
		H1: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix:          " ",
				Suffix:          " ",
				Color:           colorPtr(firstColor(t.BrightYellow, t.BrightWhite, t.Fg)),
				BackgroundColor: colorPtr(h1Bg),
				Bold:            new(true),
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{Prefix: headingPrefixH2},
		},
		H3: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{Prefix: headingPrefixH3},
		},
		H4: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{Prefix: headingPrefixH4},
		},
		H5: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{Prefix: headingPrefixH5},
		},
		H6: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: headingPrefixH6,
				Color:  colorPtr(firstColor(t.Purple, t.Fg)),
				Bold:   new(false),
			},
		},
		Strikethrough: ansi.StylePrimitive{CrossedOut: new(true)},
		Emph:          ansi.StylePrimitive{Italic: new(true)},
		Strong:        ansi.StylePrimitive{Bold: new(true)},
		HorizontalRule: ansi.StylePrimitive{
			Color:  colorPtr(t.BrightBlack),
			Format: horizontalRule,
		},
		Item:        ansi.StylePrimitive{BlockPrefix: itemPrefix},
		Enumeration: ansi.StylePrimitive{BlockPrefix: enumPrefix},
		Task: ansi.StyleTask{
			Ticked:   taskTicked,
			Unticked: taskUnticked,
		},
		Link: ansi.StylePrimitive{
			Color:     colorPtr(firstColor(t.Cyan, t.Fg)),
			Underline: new(true),
		},
		LinkText: ansi.StylePrimitive{
			Color: colorPtr(linkText),
			Bold:  new(true),
		},
		Image: ansi.StylePrimitive{
			Color:     colorPtr(firstColor(t.BrightPurple, t.Purple, t.Fg)),
			Underline: new(true),
		},
		ImageText: ansi.StylePrimitive{
			Color:  colorPtr(t.BrightBlack),
			Format: imageFormat,
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix:          codePad,
				Suffix:          codePad,
				Color:           colorPtr(firstColor(t.Red, t.Fg)),
				BackgroundColor: colorPtr(codeChip(t.Dark, t.Bg, t.BrightBlack)),
			},
		},
		CodeBlock: ansi.StyleCodeBlock{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Color: colorPtr(firstColor(t.BrightBlack, t.Fg)),
				},
				Margin: new(uint(defaultMargin)),
			},
			Theme: registerTheme(t),
		},
		DefinitionDescription: ansi.StylePrimitive{
			BlockPrefix: definitionArrow,
		},
	}
}

// canHighlight reports whether t has every color chromatint.StyleEntry reads.
// A missing color is a nil *tint.Color, and StyleEntry panics on those.
func canHighlight(t *tint.Tint) bool {
	return t.Fg != nil &&
		t.BrightPurple != nil &&
		t.Cyan != nil &&
		t.Green != nil &&
		t.Yellow != nil &&
		t.Purple != nil &&
		t.BrightWhite != nil &&
		t.Red != nil
}

func themeName(t *tint.Tint) string {
	if t.ID != "" {
		return "bubbletint-" + t.ID
	}
	return "bubbletint-" + toHex(t.Fg) + "-" + toHex(t.Bg)
}

// registerTheme installs a chroma style for t and returns its name.
// Glamour looks the name up on each fenced block, so a later call for the
// same tint replaces the previous registration.
func registerTheme(t *tint.Tint) string {
	if !canHighlight(t) {
		return ""
	}

	name := themeName(t)
	style, err := chroma.NewStyle(name, chromatint.StyleEntry(t, false))
	if err != nil {
		return ""
	}

	themeMu.Lock()
	styles.Register(style)
	themeMu.Unlock()
	return name
}
