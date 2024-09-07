// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/lucasb-eyer/go-colorful"
)

type M map[string]any

const (
	tintPath       = "defaulttints/%s.go"
	tintSVGPath    = "defaulttints/%s.svg"
	tintReadmePath = "DEFAULT_TINTS.md"
	registryPath   = "default_registry.gen.go"
)

var (
	TintUrls = []string{
		"https://github.com/atomcorp/themes/raw/master/app/src/custom-colour-schemes.json",
		"https://github.com/atomcorp/themes/raw/master/app/src/backupthemes.json",
	}
	CredUrls = []string{
		"https://github.com/atomcorp/themes/raw/master/app/src/credits.json",
	}
	ColorMapSpecial = []string{
		"SelectionBg", "Cursor",
	}
	ColorMap = []string{
		"Fg", "Bg",
		"Black", "BrightBlack",
		"Blue", "BrightBlue",
		"Cyan", "BrightCyan",
		"Green", "BrightGreen",
		"Purple", "BrightPurple",
		"Red", "BrightRed",
		"White", "BrightWhite",
		"Yellow", "BrightYellow",
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
		"dynamiccolor": func(t Tint, field string) string {
			r := reflect.ValueOf(&t)
			f := reflect.Indirect(r).FieldByName(field)
			return f.String()
		},
		"rgba": func(hex string) string {
			c, err := colorful.Hex(hex)
			if err != nil {
				return "rgba(0, 0, 0, 0.0)"
			}

			return fmt.Sprintf("rgba(%d, %d, %d, 1)", int(c.R*255), int(c.G*255), int(c.B*255))
		},
		"isbright": func(c string) bool {
			return strings.Contains(c, "Bright")
		},
		"shortbright": func(c string) string {
			return strings.Replace(c, "Bright", "Bright\n", 1)
		},
	}
	tintTmpl = template.Must(
		template.New("tint.gotmpl").
			Funcs(funcMap).
			ParseFiles("cmd/tintgen/templates/tint.gotmpl"),
	)
	tintSVGTmpl = template.Must(
		template.New("tint_svg.gotmpl").
			Funcs(funcMap).
			ParseFiles("cmd/tintgen/templates/tint_svg.gotmpl"),
	)
	tintReadmeTmpl = template.Must(
		template.New("tint_readme.gotmpl").
			Funcs(funcMap).
			ParseFiles("cmd/tintgen/templates/tint_readme.gotmpl"),
	)
	registryTmpl = template.Must(
		template.New("default_registry.gotmpl").
			Funcs(funcMap).
			ParseFiles("cmd/tintgen/templates/default_registry.gotmpl"),
	)

	header = `// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
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
		generateTintSVG(tint)
	}

	generateTintReadme(tints)
	generateRegistry(tints)
}

var reStripTrailingComma = regexp.MustCompile(`,\s*\]\s*$`)

func fetchTints() (tints []Tint) {
	for _, url := range TintUrls {
		fmt.Printf("fetching %s\n", url) //nolint:all
		resp, err := http.Get(url)       //nolint:gosec,noctx
		if err != nil {
			panic(err)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()

		b = reStripTrailingComma.ReplaceAll(b, []byte("]"))

		var rawTints []Tint
		err = json.Unmarshal(b, &rawTints)
		if err != nil {
			panic(err)
		}

		tints = append(tints, rawTints...)
	}

	return tints
}

func fetchCredits() map[string][]CreditSource {
	sourceMap := map[string][]CreditSource{}

	for _, url := range CredUrls {
		resp, err := http.Get(url) //nolint:gosec,noctx
		if err != nil {
			panic(err)
		}

		var rawCredits []Credit
		err = json.NewDecoder(resp.Body).Decode(&rawCredits)
		if err != nil {
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

	err = tintTmpl.Execute(f, tint)
	if err != nil {
		panic(err)
	}

	fmt.Println("generated", fn) //nolint:all
}

func generateTintSVG(tint TintTemplate) {
	fn := fmt.Sprintf(tintSVGPath, tint.NameNormalized)

	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tintSVGTmpl.Execute(f, M{
		"TintTemplate":    tint,
		"ColorMap":        ColorMap,
		"ColorMapSpecial": ColorMapSpecial,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("generated", fn) //nolint:all
}

func generateTintReadme(tints []TintTemplate) {
	f, err := os.Create(tintReadmePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tintReadmeTmpl.Execute(f, tints)
	if err != nil {
		panic(err)
	}

	fmt.Printf("generated %s\n", tintReadmePath) //nolint:all
}

func generateRegistry(tints []TintTemplate) {
	f, err := os.Create(registryPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = registryTmpl.Execute(f, tints)
	if err != nil {
		panic(err)
	}

	fmt.Printf("generated %s\n", registryPath) //nolint:all
}
