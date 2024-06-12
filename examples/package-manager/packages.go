// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Example modified from the following:
//  * https://github.com/charmbracelet/bubbletea/blob/test/examples/package-manager/main.go

package main

import (
	"fmt"
	"math/rand"
)

var packages = []string{
	"vegeutils",
	"libgardening",
	"currykit",
	"spicerack",
	"fullenglish",
	"eggy",
	"bad-kitty",
	"chai",
	"hojicha",
	"libtacos",
	"babys-monads",
	"libpurring",
	"currywurst-devel",
	"xmodmeow",
	"licorice-utils",
	"cashew-apple",
	"rock-lobster",
	"standmixer",
	"coffee-CUPS",
	"libesszet",
	"zeichenorientierte-benutzerschnittstellen",
	"schnurrkit",
	"old-socks-devel",
	"jalapeño",
	"molasses-utils",
	"xkohlrabi",
	"party-gherkin",
	"snow-peas",
	"libyuzu",
}

func getPackages() []string {
	pkgs := packages
	copy(pkgs, packages)

	rand.Shuffle(len(pkgs), func(i, j int) {
		pkgs[i], pkgs[j] = pkgs[j], pkgs[i]
	})

	for k := range pkgs {
		pkgs[k] += fmt.Sprintf("-%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10)) // nolint:gosec
	}
	return pkgs
}
