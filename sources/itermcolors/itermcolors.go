// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

// Package itermcolors is a [tint.Source] for [iTerm2] color presets
// (.itermcolors files), an interchange format that many terminals and theme
// collections use, and the one bubbletint's own built-in tints come from.
//
// Unlike the other sources, it isn't tied to a terminal's config, so it needs to
// be told where to read from. Given a file, it provides a single tint, "current".
// Given a directory, it provides a tint for every .itermcolors file in it, with
// the ID of the file's name, sorted by name, ignoring case. In that case no tint is more active
// than another, so you'll likely want to attach the source with
// [tint.WithFollow](false).
//
// [iTerm2]: https://iterm2.com
package itermcolors

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/lrstanley/bubbletint/sources/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Name is the name of the source.
const Name = "iterm2"

// Extension is the file extension of color presets.
const Extension = ".itermcolors"

// New returns a source that reads the .itermcolors file, or directory of them,
// at path.
func New(path string, opts ...sources.Option) *sources.FileSource {
	opts = append([]sources.Option{sources.WithPath(path)}, opts...)

	return sources.New(sources.Spec{
		Name: Name,
		Load: func(path string) (tints []*tint.Tint, files []string, err error) {
			if !sources.IsDir(path) {
				t, loadErr := LoadFile(path, "current")
				if loadErr != nil {
					return nil, []string{path}, loadErr
				}
				return []*tint.Tint{t}, []string{path}, nil
			}

			entries, err := os.ReadDir(path)
			if err != nil {
				return nil, nil, err
			}

			sort.Slice(entries, func(i, j int) bool {
				return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
			})
			for _, e := range entries {
				if e.IsDir() || !strings.EqualFold(filepath.Ext(e.Name()), Extension) {
					continue
				}

				file := filepath.Join(path, e.Name())
				files = append(files, file)

				id := sources.ID(strings.TrimSuffix(e.Name(), filepath.Ext(e.Name())))
				t, loadErr := LoadFile(file, id)
				if loadErr != nil {
					continue // One broken preset shouldn't hide the rest.
				}
				tints = append(tints, t)
			}

			if len(tints) == 0 {
				return nil, files, fmt.Errorf("%w: no %s files in %s", tint.ErrNotFound, Extension, path)
			}
			return tints, files, nil
		},
	}, opts...)
}

// LoadFile reads the .itermcolors file at path as a tint with the provided ID.
// The tint's display name is the file's name.
func LoadFile(path, id string) (*tint.Tint, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	t, err := Parse(f, id, name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return t, nil
}

// Parse reads a .itermcolors (XML property list) document.
func Parse(r io.Reader, id, displayName string) (*tint.Tint, error) {
	root, err := parsePlist(r)
	if err != nil {
		return nil, err
	}

	dict, ok := root.(map[string]any)
	if !ok {
		return nil, errors.New("not a color preset: top-level value is not a dictionary")
	}

	// Newer presets carry separate "(Dark)"/"(Light)" variants of some colors,
	// in addition to the plain ones. Prefer the plain ones, then dark.
	get := func(key string) *tint.Color {
		for _, k := range []string{key, key + " (Dark)"} {
			if c, isDict := dict[k].(map[string]any); isDict {
				return component(c)
			}
		}
		return nil
	}

	var p sources.Palette
	p.Foreground, p.Background = get("Foreground Color"), get("Background Color")
	p.Cursor, p.CursorText = get("Cursor Color"), get("Cursor Text Color")
	p.SelectionBackground, p.SelectionForeground = get("Selection Color"), get("Selected Text Color")
	p.Bold = get("Bold Color")
	for i := range p.ANSI {
		p.SetANSI(i, get("Ansi "+strconv.Itoa(i)+" Color"))
	}

	if p.Empty() {
		return nil, errors.New("not a color preset: no colors found")
	}
	return p.Tint(id, displayName), nil
}

// component converts an iTerm2 color dictionary (components from 0 to 1) to a
// color.
func component(d map[string]any) *tint.Color {
	r, rok := d["Red Component"].(float64)
	g, gok := d["Green Component"].(float64)
	b, bok := d["Blue Component"].(float64)
	if !rok || !gok || !bok {
		return nil
	}

	to8 := func(v float64) uint8 { return uint8(math.Round(math.Min(1, math.Max(0, v)) * 255)) }
	return &tint.Color{R: to8(r), G: to8(g), B: to8(b), A: 0xff}
}

// parsePlist decodes an XML property list into maps, slices, strings, float64s
// and bools.
func parsePlist(r io.Reader) (any, error) {
	d := xml.NewDecoder(r)

	for {
		tok, err := d.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil, errors.New("not a property list: no content")
			}
			return nil, err
		}

		start, ok := tok.(xml.StartElement)
		if !ok || start.Name.Local == "plist" {
			continue
		}
		return parseValue(d, start)
	}
}

func parseValue(d *xml.Decoder, start xml.StartElement) (any, error) {
	switch start.Name.Local {
	case "dict":
		m := make(map[string]any)
		var key string
		for {
			tok, err := d.Token()
			if err != nil {
				return nil, err
			}
			switch t := tok.(type) {
			case xml.EndElement:
				return m, nil
			case xml.StartElement:
				if t.Name.Local == "key" {
					if key, err = text(d); err != nil {
						return nil, err
					}
					continue
				}
				v, valueErr := parseValue(d, t)
				if valueErr != nil {
					return nil, valueErr
				}
				m[key] = v
			}
		}
	case "array":
		var list []any
		for {
			tok, err := d.Token()
			if err != nil {
				return nil, err
			}
			switch t := tok.(type) {
			case xml.EndElement:
				return list, nil
			case xml.StartElement:
				v, valueErr := parseValue(d, t)
				if valueErr != nil {
					return nil, valueErr
				}
				list = append(list, v)
			}
		}
	case "true", "false":
		return start.Name.Local == "true", d.Skip()
	case "real", "integer":
		s, err := text(d)
		if err != nil {
			return nil, err
		}
		v, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q", s)
		}
		return v, nil
	default: // string, date, data
		return text(d)
	}
}

// text returns the character data of the current element, and consumes its end.
func text(d *xml.Decoder) (string, error) {
	var sb strings.Builder
	for {
		tok, err := d.Token()
		if err != nil {
			return "", err
		}
		switch t := tok.(type) {
		case xml.CharData:
			sb.Write(t)
		case xml.EndElement:
			return sb.String(), nil
		case xml.StartElement:
			if skipErr := d.Skip(); skipErr != nil {
				return "", skipErr
			}
		}
	}
}
