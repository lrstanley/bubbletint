// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package glamourtint

import (
	"strings"
	"testing"

	"charm.land/glamour/v2"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
	tint "github.com/lrstanley/bubbletint/v2"
)

const sampleMD = `# Tint

## Section

A paragraph with [a link](https://github.com/lrstanley/bubbletint) and ` + "`inline`" + ` code.

- first
- [x] done

> quoted

---

` + "```go\nfunc main() {}\n```\n"

func TestStyleConfigNil(t *testing.T) {
	cfg := StyleConfig(nil)
	if cfg.Document.Color != nil || cfg.CodeBlock.Theme != "" || cfg.CodeBlock.Chroma != nil {
		t.Fatalf("nil tint: got document color %v theme %q chroma %v", cfg.Document.Color, cfg.CodeBlock.Theme, cfg.CodeBlock.Chroma)
	}
}

func TestStyleConfigSparse(t *testing.T) {
	cfg := StyleConfig(&tint.Tint{Dark: true})
	if cfg.Document.Color != nil {
		t.Fatalf("sparse tint document color = %v", *cfg.Document.Color)
	}
	if cfg.H2.Prefix != headingPrefixH2 {
		t.Fatalf("H2 prefix = %q", cfg.H2.Prefix)
	}
	if cfg.Document.Margin == nil || *cfg.Document.Margin != defaultMargin {
		t.Fatal("document margin should be the default")
	}
	if cfg.H1.Bold == nil || !*cfg.H1.Bold {
		t.Fatal("H1 bold should be set")
	}
	if cfg.Document.BackgroundColor != nil {
		t.Fatal("document background should be unset")
	}
}

func TestStyleConfig(t *testing.T) {
	tests := []struct {
		name     string
		tint     *tint.Tint
		heading  string
		h1Bg     string
		linkText string
	}{
		{
			name:     "dark",
			tint:     tint.TintDraculaPlus,
			heading:  tint.TintDraculaPlus.BrightBlue.Hex(),
			h1Bg:     tint.TintDraculaPlus.Blue.Hex(),
			linkText: tint.TintDraculaPlus.BrightCyan.Hex(),
		},
		{
			name:     "light",
			tint:     tint.TintCatppuccinLatte,
			heading:  tint.TintCatppuccinLatte.Blue.Hex(),
			h1Bg:     tint.TintCatppuccinLatte.BrightBlue.Hex(),
			linkText: tint.TintCatppuccinLatte.Cyan.Hex(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := StyleConfig(tt.tint)

			if got := deref(cfg.Document.Color); got != tt.tint.Fg.Hex() {
				t.Fatalf("document color = %s, want %s", got, tt.tint.Fg.Hex())
			}
			if cfg.Document.BackgroundColor != nil {
				t.Fatal("document background should be unset")
			}
			if got := deref(cfg.Heading.Color); got != tt.heading {
				t.Fatalf("heading color = %s, want %s", got, tt.heading)
			}
			if cfg.Heading.Bold == nil || !*cfg.Heading.Bold {
				t.Fatal("heading should be bold")
			}
			if got := deref(cfg.H1.Color); got != tt.tint.BrightYellow.Hex() {
				t.Fatalf("h1 color = %s, want %s", got, tt.tint.BrightYellow.Hex())
			}
			if got := deref(cfg.H1.BackgroundColor); got != tt.h1Bg {
				t.Fatalf("h1 background = %s, want %s", got, tt.h1Bg)
			}
			shift := func(amount float64) string {
				if tt.tint.Dark {
					amount = -amount
				}
				return toHex(step(tt.tint.Fg, nil, amount))
			}
			wantH2 := shift(h2Fade)
			if got := deref(cfg.H2.Color); got != wantH2 {
				t.Fatalf("h2 color = %s, want %s", got, wantH2)
			}
			if cfg.H2.Prefix != headingPrefixH2 {
				t.Fatalf("h2 prefix = %q", cfg.H2.Prefix)
			}
			wantH6 := shift(h6Fade)
			if got := deref(cfg.H6.Color); got != wantH6 || got == wantH2 {
				t.Fatalf("h6 color = %s, want %s distinct from h2", got, wantH6)
			}
			if cfg.H6.Bold == nil || *cfg.H6.Bold {
				t.Fatal("h6 bold should be false")
			}
			if cfg.Document.Margin == nil || *cfg.Document.Margin != defaultMargin {
				t.Fatal("document margin should be the default")
			}
			if got := deref(cfg.Code.BackgroundColor); got == "" || got == tt.tint.Bg.Hex() {
				t.Fatalf("inline code background = %q", got)
			}
			if cfg.Link.Underline == nil || !*cfg.Link.Underline {
				t.Fatal("link should be underlined")
			}
			if got := deref(cfg.LinkText.Color); got != tt.linkText {
				t.Fatalf("link text = %s, want %s", got, tt.linkText)
			}
			if cfg.CodeBlock.Chroma != nil {
				t.Fatal("code block chroma should be unset so glamour uses Theme")
			}
			if cfg.CodeBlock.Theme == "" {
				t.Fatal("code block theme should be registered")
			}

			keyword := styles.Get(cfg.CodeBlock.Theme).Get(chroma.Keyword).Colour
			if !keyword.IsSet() {
				t.Fatal("keyword color unset")
			}
		})
	}
}

func TestCodeThemeFollowsTint(t *testing.T) {
	dark := StyleConfig(tint.TintDraculaPlus)
	light := StyleConfig(tint.TintCatppuccinLatte)
	if dark.CodeBlock.Theme == light.CodeBlock.Theme {
		t.Fatalf("theme name %q used for both tints", dark.CodeBlock.Theme)
	}

	darkKeyword := styles.Get(dark.CodeBlock.Theme).Get(chroma.Keyword).Colour
	lightKeyword := styles.Get(light.CodeBlock.Theme).Get(chroma.Keyword).Colour
	if !darkKeyword.IsSet() || !lightKeyword.IsSet() || darkKeyword == lightKeyword {
		t.Fatalf("keyword colors = %s and %s", darkKeyword, lightKeyword)
	}
}

func TestStyleConfigRender(t *testing.T) {
	dark, err := render(tint.TintDraculaPlus)
	if err != nil {
		t.Fatalf("render dark: %v", err)
	}
	light, err := render(tint.TintCatppuccinLatte)
	if err != nil {
		t.Fatalf("render light: %v", err)
	}

	for _, out := range []string{dark, light} {
		if !strings.Contains(out, "Tint") || !strings.Contains(out, "func") {
			t.Fatalf("render missing content:\n%s", out)
		}
		if !strings.Contains(out, "\x1b[") {
			t.Fatalf("render missing ANSI color:\n%s", out)
		}
	}
	if dark == light {
		t.Fatal("dark and light renders should differ")
	}
}

func render(t *tint.Tint) (string, error) {
	r, err := glamour.NewTermRenderer(
		glamour.WithStyles(StyleConfig(t)),
		glamour.WithWordWrap(60),
		glamour.WithChromaFormatter("terminal16m"),
	)
	if err != nil {
		return "", err
	}
	return r.Render(sampleMD)
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
