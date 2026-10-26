// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package windowsterminal is a [tint.Source] for the colors of [Windows Terminal].
//
// It reads settings.json and provides the color scheme of the active profile as
// the tint "current", followed by every scheme defined in settings.json, under
// the ID of its name (e.g. "one_half_dark"). The active profile is the one in
// $WT_PROFILE_ID, which Windows Terminal sets for the shells it starts, or else
// the default profile. If its scheme is not defined in settings.json it is one
// of the schemes built into Windows Terminal, which are taken from bubbletint's
// built-in tints, with Windows Terminal's default "Campbell" included. When a
// profile sets a scheme for both light and dark mode, the dark one is used.
//
// Locations searched, first match wins (each under %LOCALAPPDATA%):
//
//   - Packages\Microsoft.WindowsTerminal_8wekyb3d8bbwe\LocalState (Store, stable)
//   - Packages\Microsoft.WindowsTerminalPreview_8wekyb3d8bbwe\LocalState (Store, preview)
//   - Microsoft\Windows Terminal (unpackaged)
//
// When running in WSL, the Windows user's %LOCALAPPDATA% is searched too, via
// /mnt/c/Users.
//
// [Windows Terminal]: https://learn.microsoft.com/windows/terminal/
package windowsterminal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Name is the name of the source.
const Name = "windowsterminal"

// defaultScheme is the scheme Windows Terminal uses when none is configured.
const defaultScheme = "Campbell"

// New returns a source that reads Windows Terminal's settings.
func New(opts ...sources.Option) *sources.FileSource {
	return sources.New(sources.Spec{
		Name:       Name,
		Candidates: Candidates,
		Load: func(path string) ([]*tint.Tint, []string, error) {
			tints, err := Load(path)
			return tints, nil, err
		},
	}, opts...)
}

var settingsPaths = []string{
	filepath.Join("Packages", "Microsoft.WindowsTerminal_8wekyb3d8bbwe", "LocalState", "settings.json"),
	filepath.Join("Packages", "Microsoft.WindowsTerminalPreview_8wekyb3d8bbwe", "LocalState", "settings.json"),
	filepath.Join("Microsoft", "Windows Terminal", "settings.json"),
}

// Candidates returns the paths Windows Terminal's settings are searched for, in
// order.
func Candidates() []string {
	var roots []string
	if dir := os.Getenv("LOCALAPPDATA"); dir != "" {
		roots = append(roots, dir)
	}

	// Under WSL, Windows Terminal's settings are on the Windows side.
	if os.Getenv("WSL_DISTRO_NAME") != "" || os.Getenv("WSL_INTEROP") != "" {
		matches, _ := filepath.Glob("/mnt/c/Users/*/AppData/Local")
		roots = append(roots, matches...)
	}

	var paths []string
	for _, root := range roots {
		for _, p := range settingsPaths {
			paths = append(paths, filepath.Join(root, p))
		}
	}
	return paths
}

type scheme struct {
	Name                string `json:"name"`
	Foreground          string `json:"foreground"`
	Background          string `json:"background"`
	CursorColor         string `json:"cursorColor"`
	SelectionBackground string `json:"selectionBackground"`
	Black               string `json:"black"`
	Red                 string `json:"red"`
	Green               string `json:"green"`
	Yellow              string `json:"yellow"`
	Blue                string `json:"blue"`
	Purple              string `json:"purple"`
	Cyan                string `json:"cyan"`
	White               string `json:"white"`
	BrightBlack         string `json:"brightBlack"`
	BrightRed           string `json:"brightRed"`
	BrightGreen         string `json:"brightGreen"`
	BrightYellow        string `json:"brightYellow"`
	BrightBlue          string `json:"brightBlue"`
	BrightPurple        string `json:"brightPurple"`
	BrightCyan          string `json:"brightCyan"`
	BrightWhite         string `json:"brightWhite"`
}

func (s *scheme) tint(id string) *tint.Tint {
	var p sources.Palette
	p.Foreground, _ = sources.ParseColor(s.Foreground)
	p.Background, _ = sources.ParseColor(s.Background)
	p.Cursor, _ = sources.ParseColor(s.CursorColor)
	p.SelectionBackground, _ = sources.ParseColor(s.SelectionBackground)

	for i, v := range []string{
		s.Black, s.Red, s.Green, s.Yellow, s.Blue, s.Purple, s.Cyan, s.White,
		s.BrightBlack, s.BrightRed, s.BrightGreen, s.BrightYellow,
		s.BrightBlue, s.BrightPurple, s.BrightCyan, s.BrightWhite,
	} {
		c, _ := sources.ParseColor(v)
		p.SetANSI(i, c)
	}
	return p.Tint(id, s.Name)
}

// campbell is Windows Terminal's default scheme. Always defined, as it's the one
// scheme that has no built-in tint to fall back on.
var campbell = scheme{
	Name: "Campbell", Foreground: "#CCCCCC", Background: "#0C0C0C", CursorColor: "#FFFFFF",
	SelectionBackground: "#FFFFFF",
	Black:               "#0C0C0C", Red: "#C50F1F", Green: "#13A10E", Yellow: "#C19C00",
	Blue: "#0037DA", Purple: "#881798", Cyan: "#3A96DD", White: "#CCCCCC",
	BrightBlack: "#767676", BrightRed: "#E74856", BrightGreen: "#16C60C", BrightYellow: "#F9F1A5",
	BrightBlue: "#3B78FF", BrightPurple: "#B4009E", BrightCyan: "#61D6D6", BrightWhite: "#F2F2F2",
}

// colorScheme is a profile's "colorScheme", which is either a scheme name, or
// an object naming a scheme for each of light and dark mode.
type colorScheme string

func (c *colorScheme) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err == nil {
		*c = colorScheme(name)
		return nil
	}

	var pair struct {
		Light string `json:"light"`
		Dark  string `json:"dark"`
	}
	if err := json.Unmarshal(data, &pair); err != nil {
		return err
	}
	*c = colorScheme(pair.Dark)
	if pair.Dark == "" {
		*c = colorScheme(pair.Light)
	}
	return nil
}

type profile struct {
	GUID        string      `json:"guid"`
	ColorScheme colorScheme `json:"colorScheme"`
}

type settings struct {
	DefaultProfile string `json:"defaultProfile"`
	Profiles       struct {
		Defaults profile   `json:"defaults"`
		List     []profile `json:"list"`
	} `json:"profiles"`
	Schemes []scheme `json:"schemes"`
}

// Load reads the Windows Terminal settings at path, and returns the tints
// described in the package documentation.
func Load(path string) ([]*tint.Tint, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(data, os.Getenv("WT_PROFILE_ID"))
}

// Parse returns the tints for settings.json contents (which may contain comments
// and trailing commas). profileID is the GUID of the profile whose scheme should
// be the "current" tint; when empty, the default profile is used.
func Parse(data []byte, profileID string) ([]*tint.Tint, error) {
	var s settings
	if err := json.Unmarshal(StripJSONC(data), &s); err != nil {
		return nil, fmt.Errorf("parsing settings.json: %w", err)
	}

	if profileID == "" {
		profileID = s.DefaultProfile
	}
	profileID = normalizeGUID(profileID)

	active := string(s.Profiles.Defaults.ColorScheme)
	for _, p := range s.Profiles.List {
		if profileID != "" && normalizeGUID(p.GUID) == profileID && p.ColorScheme != "" {
			active = string(p.ColorScheme)
			break
		}
	}
	if active == "" {
		active = defaultScheme
	}

	var tints []*tint.Tint
	var current *tint.Tint
	for i := range s.Schemes {
		sc := &s.Schemes[i]
		if sc.Name == "" {
			continue
		}
		t := sc.tint(sources.ID(sc.Name))
		tints = append(tints, t)
		if current == nil && strings.EqualFold(sc.Name, active) {
			current = copyAs(t, "current")
		}
	}

	if current == nil {
		if strings.EqualFold(active, campbell.Name) {
			current = campbell.tint("current")
		} else {
			for _, t := range tint.DefaultTints() {
				if strings.EqualFold(t.DisplayName, active) {
					current = copyAs(t, "current")
					break
				}
			}
		}
	}
	if current == nil {
		return nil, fmt.Errorf("%w: color scheme %q is not defined in settings.json, and is not a built-in scheme", tint.ErrNotFound, active)
	}

	return append([]*tint.Tint{current}, tints...), nil
}

func copyAs(t *tint.Tint, id string) *tint.Tint {
	cp := *t
	cp.ID = id
	return &cp
}

func normalizeGUID(s string) string {
	return strings.ToLower(strings.Trim(strings.TrimSpace(s), "{}"))
}

// StripJSONC removes the comments, trailing commas and byte order mark that
// Windows Terminal allows in settings.json, leaving valid JSON.
func StripJSONC(data []byte) []byte {
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
	return stripTrailingCommas(stripComments(data))
}

// copyString copies the string literal starting at data[i] (the opening quote)
// to out, returning the index of its closing quote.
func copyString(out, data []byte, i int) (result []byte, end int) {
	out = append(out, data[i])
	for i++; i < len(data); i++ {
		out = append(out, data[i])
		switch data[i] {
		case '\\':
			if i+1 < len(data) {
				i++
				out = append(out, data[i])
			}
		case '"':
			return out, i
		}
	}
	return out, i
}

func stripComments(data []byte) []byte {
	out := make([]byte, 0, len(data))

	for i := 0; i < len(data); i++ {
		switch {
		case data[i] == '"':
			out, i = copyString(out, data, i)
		case data[i] == '/' && i+1 < len(data) && data[i+1] == '/':
			for i < len(data) && data[i] != '\n' {
				i++
			}
			out = append(out, '\n')
		case data[i] == '/' && i+1 < len(data) && data[i+1] == '*':
			i += 2
			for i+1 < len(data) && (data[i] != '*' || data[i+1] != '/') {
				i++
			}
			i++ // Skip the closing "/".
			out = append(out, ' ')
		default:
			out = append(out, data[i])
		}
	}
	return out
}

func stripTrailingCommas(data []byte) []byte {
	out := make([]byte, 0, len(data))

	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '"':
			out, i = copyString(out, data, i)
		case ',':
			j := i + 1
			for j < len(data) && (data[j] == ' ' || data[j] == '\t' || data[j] == '\r' || data[j] == '\n') {
				j++
			}
			if j < len(data) && (data[j] == '}' || data[j] == ']') {
				continue
			}
			out = append(out, data[i])
		default:
			out = append(out, data[i])
		}
	}
	return out
}
