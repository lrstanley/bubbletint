// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package windowsterminal

import (
	"encoding/json"
	"errors"
	"testing"

	tint "github.com/lrstanley/bubbletint/v2"
)

const testSettings = "\xef\xbb\xbf" + `// Windows Terminal settings.
{
    "$help": "https://aka.ms/terminal-documentation",
    "defaultProfile": "{61c54bbd-c2c6-5271-96e7-009a87ff44bf}",
    /* block
       comment */
    "profiles": {
        "defaults": { "colorScheme": "Gruvbox", },
        "list": [
            { "guid": "{61c54bbd-c2c6-5271-96e7-009a87ff44bf}", "name": "PowerShell" },
            { "guid": "{574e775e-4f2a-5b96-ac1e-a2962a402336}", "colorScheme": "One Half Dark" },
            { "guid": "{b453ae62-4e3d-5e58-b989-0a998ec441b8}", "colorScheme": { "light": "Gruvbox", "dark": "Campbell" } },
            { "guid": "{11111111-1111-1111-1111-111111111111}", "colorScheme": "No Such Scheme" },
        ],
    },
    "schemes": [
        {
            "name": "Gruvbox", // trailing comma and "//" in strings below
            "foreground": "#ebdbb2", "background": "#282828", "cursorColor": "#ebdbb2",
            "selectionBackground": "#665c54",
            "black": "#282828", "red": "#cc241d", "green": "#98971a", "yellow": "#d79921",
            "blue": "#458588", "purple": "#b16286", "cyan": "#689d6a", "white": "#a89984",
            "brightBlack": "#928374", "brightRed": "#fb4934", "brightGreen": "#b8bb26",
            "brightYellow": "#fabd2f", "brightBlue": "#83a598", "brightPurple": "#d3869b",
            "brightCyan": "#8ec07c", "brightWhite": "#ebdbb2",
        },
    ],
}`

func ids(tints []*tint.Tint) []string {
	var out []string
	for _, t := range tints {
		out = append(out, t.ID)
	}
	return out
}

func TestParse(t *testing.T) {
	t.Parallel()

	// The default profile has no scheme of its own, so it gets the default one.
	tints, err := Parse([]byte(testSettings), "")
	if err != nil {
		t.Fatal(err)
	}
	if len(tints) != 2 || tints[0].ID != "current" || tints[1].ID != "gruvbox" {
		t.Fatalf("unexpected tints: %v", ids(tints))
	}
	if tints[0].Bg.Hex() != "#282828" || tints[0].DisplayName != "Gruvbox" || !tints[0].Dark {
		t.Errorf("unexpected current tint: %s %s", tints[0].DisplayName, tints[0].Bg.Hex())
	}
	if tints[1].Purple.Hex() != "#b16286" || tints[1].BrightWhite.Hex() != "#ebdbb2" || tints[1].Cursor.Hex() != "#ebdbb2" {
		t.Errorf("unexpected scheme colors")
	}

	// A profile that names a built-in scheme gets bubbletint's version of it.
	tints, err = Parse([]byte(testSettings), "{574E775E-4F2A-5B96-AC1E-A2962A402336}")
	if err != nil {
		t.Fatal(err)
	}
	if tints[0].ID != "current" || tints[0].DisplayName != "One Half Dark" || tints[0].Bg.Hex() != tint.TintOneHalfDark.Bg.Hex() {
		t.Errorf("unexpected built-in scheme: %s %s", tints[0].DisplayName, tints[0].Bg.Hex())
	}

	// Light/dark pairs use the dark scheme, which here is the default scheme.
	tints, err = Parse([]byte(testSettings), "b453ae62-4e3d-5e58-b989-0a998ec441b8")
	if err != nil {
		t.Fatal(err)
	}
	if tints[0].DisplayName != "Campbell" || tints[0].Bg.Hex() != "#0c0c0c" {
		t.Errorf("unexpected scheme %s %s", tints[0].DisplayName, tints[0].Bg.Hex())
	}

	// A scheme that doesn't exist anywhere is an error, not a guess.
	if _, err = Parse([]byte(testSettings), "{11111111-1111-1111-1111-111111111111}"); !errors.Is(err, tint.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestParse_Defaults(t *testing.T) {
	t.Parallel()

	// Nothing configured at all means Windows Terminal's default scheme.
	tints, err := Parse([]byte(`{}`), "")
	if err != nil {
		t.Fatal(err)
	}
	if len(tints) != 1 || tints[0].DisplayName != "Campbell" {
		t.Fatalf("unexpected tints: %v", ids(tints))
	}

	if _, err = Parse([]byte(`{`), ""); err == nil {
		t.Fatal("expected an error for invalid JSON")
	}
}

func TestStripJSONC(t *testing.T) {
	t.Parallel()

	in := "\xef\xbb\xbf{\n" +
		`  "url": "http://example.com/*not a comment*/", // real comment` + "\n" +
		`  "esc": "quote \" // still string, }",` + "\n" +
		`  "list": [1, 2, /* inline */ 3,],` + "\n" +
		`  "obj": {"a": 1,},` + "\n" +
		"}"

	var got struct {
		URL  string         `json:"url"`
		Esc  string         `json:"esc"`
		List []int          `json:"list"`
		Obj  map[string]int `json:"obj"`
	}
	if err := json.Unmarshal(StripJSONC([]byte(in)), &got); err != nil {
		t.Fatalf("result is not valid JSON: %v\n%s", err, StripJSONC([]byte(in)))
	}
	if got.URL != "http://example.com/*not a comment*/" || got.Esc != `quote " // still string, }` ||
		len(got.List) != 3 || got.Obj["a"] != 1 {
		t.Errorf("unexpected result: %+v", got)
	}
}
