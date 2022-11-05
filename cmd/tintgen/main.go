// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/iancoleman/strcase"
)

const (
	tintPath = "defaulttints/%s.go"
)

var (
	TintUrls = []string{
		"https://github.com/atomcorp/themes/raw/master/app/src/custom-colour-schemes.json",
		"https://github.com/atomcorp/themes/raw/master/app/src/backupthemes.json",
	}
	CredUrls = []string{
		"https://github.com/atomcorp/themes/raw/master/app/src/credits.json",
	}
)

var (
	funcMap = template.FuncMap{
		"header": func() string { return header },
		"urlenc": url.QueryEscape,
		"color": func(c string) string {
			if c == "" {
				return "lipgloss.NoColor{}"
			}

			return fmt.Sprintf(`lipgloss.Color("%s")`, c) //nolint:all
		},
	}
	tintTmpl = template.Must(
		template.New("tint.gotmpl").
			Funcs(sprig.FuncMap()).
			Funcs(funcMap).
			ParseFiles("cmd/tintgen/templates/tint.gotmpl"),
	)
	registerTmpl = template.Must(
		template.New("register.gotmpl").
			Funcs(sprig.FuncMap()).
			Funcs(funcMap).
			ParseFiles("cmd/tintgen/templates/register.gotmpl"),
	)

	header = `// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by tintgen. DO NOT EDIT.`

	tintReplacer = strings.NewReplacer(
		"+", " Plus",
		"(", "",
		")", "",
	)
)

func main() {
	rawTints := fetchTints()
	sourceMap := fetchCredits()

	sort.SliceStable(rawTints, func(i, j int) bool {
		return rawTints[i].Name < rawTints[j].Name
	})

	tintNames := []string{}
	tints := []TintTemplate{}
	for _, tint := range rawTints { //nolint:all
		name := tintReplacer.Replace(tint.Name)

		t := TintTemplate{
			StructName:     strcase.ToCamel(name),
			NameNormalized: strcase.ToSnake(name),
			Tint:           tint,
		}

		var exists bool
		for _, name := range tintNames {
			if name == t.NameNormalized || name == t.StructName {
				exists = true
				break
			}
		}

		// If the tint name already exists, skip it.
		if exists {
			continue
		}

		tintNames = append(tintNames, t.NameNormalized, t.StructName)

		if _, ok := sourceMap[tint.Name]; ok {
			t.CreditSources = sourceMap[tint.Name]
		}

		tints = append(tints, t)
	}

	for _, tint := range tints { //nolint:all
		generateTint(tint)
	}

	generateRegistry(tints)
}

func fetchTints() (tints []Tint) {
	for _, url := range TintUrls {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		var rawTints []Tint
		if err := json.NewDecoder(resp.Body).Decode(&rawTints); err != nil {
			panic(err)
		}
		resp.Body.Close()

		tints = append(tints, rawTints...)
	}

	return tints
}

func fetchCredits() map[string][]CreditSource {
	sourceMap := map[string][]CreditSource{}

	for _, url := range CredUrls {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		var rawCredits []Credit
		if err := json.NewDecoder(resp.Body).Decode(&rawCredits); err != nil {
			panic(err)
		}
		resp.Body.Close()

		for _, credit := range rawCredits {
			for _, tint := range credit.Tints {
				if _, ok := sourceMap[tint]; !ok {
					sourceMap[tint] = []CreditSource{}
				}

				for _, source := range credit.Sources {
					// Check to see if the credit was already added.
					exists := false
					for _, origSource := range sourceMap[tint] {
						if origSource.Link == source.Link {
							exists = true
							break
						}
					}
					if !exists {
						sourceMap[tint] = append(sourceMap[tint], source)
					}
				}
			}
		}
	}

	return sourceMap
}

func generateTint(tint TintTemplate) {
	fn := fmt.Sprintf(tintPath, tint.NameNormalized)

	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := tintTmpl.Execute(f, tint); err != nil {
		panic(err)
	}

	fmt.Println("generated", fn)
}

func generateRegistry(tints []TintTemplate) {
	f, err := os.Create("default_registry.gen.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := registerTmpl.Execute(f, tints); err != nil {
		panic(err)
	}

	fmt.Println("generated default_registry.gen.go")
}
