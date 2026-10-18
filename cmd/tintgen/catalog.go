// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	categoryGoName     = "category.gen.go"
	categoriesSchemaID = "https://github.com/lrstanley/bubbletint/raw/master/categories.schema.json"
)

// categoryTmpl renders the public Hue and Tone constants from the catalogs.
var categoryTmpl = template.Must(template.New("category.gotmpl").ParseFiles("templates/category.gotmpl"))

type categoryIdent struct {
	Func string
	Name string
}

// writeCategorySources writes the public category constants and the tone
// cache schema from the hue and tone catalogs.
func writeCategorySources(dir string) error {
	goSrc, err := renderCategoryGo()
	if err != nil {
		return err
	}
	if writeErr := writeSource(filepath.Join(dir, categoryGoName), goSrc); writeErr != nil {
		return writeErr
	}
	return writeSource(filepath.Join(dir, "categories.schema.json"), renderSchema())
}

func renderCategoryGo() ([]byte, error) {
	hues, err := idents(hueNames())
	if err != nil {
		return nil, err
	}
	toneList, err := idents(toneNames())
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = categoryTmpl.Execute(&buf, struct {
		Header string
		Hues   []categoryIdent
		Tones  []categoryIdent
	}{
		Header: header,
		Hues:   hues,
		Tones:  toneList,
	})
	if err != nil {
		return nil, fmt.Errorf("render category constants: %w", err)
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format category constants: %w", err)
	}
	return src, nil
}

func idents(names []string) ([]categoryIdent, error) {
	seenName := make(map[string]struct{}, len(names))
	seenFunc := make(map[string]struct{}, len(names))
	out := make([]categoryIdent, len(names))
	for i, name := range names {
		if name == "" {
			return nil, errors.New("empty category name")
		}
		if _, ok := seenName[name]; ok {
			return nil, fmt.Errorf("duplicate category name %q", name)
		}
		fn := strcase.ToCamel(name)
		if _, ok := seenFunc[fn]; ok {
			return nil, fmt.Errorf("duplicate category identifier %s", fn)
		}
		seenName[name] = struct{}{}
		seenFunc[fn] = struct{}{}
		out[i] = categoryIdent{Func: fn, Name: name}
	}
	return out, nil
}

func renderSchema() []byte {
	names := cacheToneNames()
	required := make([]string, len(names))
	props := make([]string, len(names))
	for i, name := range names {
		quoted := strconv.Quote(name)
		required[i] = "        " + quoted
		props[i] = "        " + quoted + `: { "$ref": "#/$defs/names" }`
	}
	var b strings.Builder
	b.WriteString(`{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "` + categoriesSchemaID + `",
  "title": "Bubbletint tone cache",
  "description": "Cached aesthetic tone for each built-in tint.",
  "type": "object",
  "additionalProperties": false,
  "required": ["$schema", "tone"],
  "properties": {
    "$schema": {
      "type": "string"
    },
    "tone": {
      "type": "object",
      "additionalProperties": false,
      "required": [
`)
	b.WriteString(strings.Join(required, ",\n"))
	b.WriteString(`
      ],
      "properties": {
`)
	b.WriteString(strings.Join(props, ",\n"))
	b.WriteString(`
      }
    }
  },
  "$defs": {
    "names": {
      "description": "Upstream theme display names, sorted.",
      "type": "array",
      "items": {
        "type": "string",
        "minLength": 1
      },
      "uniqueItems": true
    }
  }
}
`)
	return []byte(b.String())
}

func writeSource(path string, body []byte) error {
	if err := os.WriteFile(path, body, 0o644); err != nil { //nolint:gosec // Generated source is repository content.
		return fmt.Errorf("write %s: %w", filepath.Base(path), err)
	}
	logger.Info("generated file", "file", path)
	return nil
}
