// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles_test

import (
	"fmt"

	"github.com/lrstanley/bubbletint/roles/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

func ExampleDerive() {
	pal := roles.Derive(tint.TintTokyoNight)
	for _, r := range []roles.Role{roles.Bg, roles.Raised, roles.Muted, roles.Accent, roles.Danger, roles.Danger.Soft()} {
		fmt.Printf("%-12s %s\n", r, pal.Color(r).Hex())
	}
	fmt.Println("accent from", pal.AccentSlot())
	// Output:
	// bg           #1a1b26
	// raised       #272937
	// muted        #7f85a4
	// accent       #79a2f7
	// danger       #f7768e
	// danger-soft  #3b2a36
	// accent from blue
}

func ExamplePalette_Fit() {
	pal := roles.Derive(tint.TintTokyoNight)
	for _, d := range []roles.Depth{roles.TrueColor, roles.Colors256, roles.Colors16, roles.NoColor} {
		f := pal.Fit(roles.Target{Depth: d})
		var c any = f.Color(roles.Muted)
		if tc, ok := c.(tint.Color); ok {
			c = tc.Hex()
		}
		fmt.Printf("%-9s %T %v, attr %v\n", d, f.Color(roles.Muted), c, f.Attr(roles.Muted))
	}
	// Output:
	// truecolor tint.Color #7f85a4, attr none
	// 256       ansi.IndexedColor 103, attr none
	// 16        ansi.BasicColor 7, attr none
	// none      <nil> <nil>, attr faint
}

func ExamplePalette_On() {
	pal := roles.Derive(tint.TintCatppuccinLatte)
	row := pal.On(roles.Selection)
	fmt.Println("subtle on the page:", pal.Color(roles.Subtle).Hex())
	fmt.Println("subtle on the selected row:", row.Color(roles.Subtle).Hex())
	// Output:
	// subtle on the page: #9194a2
	// subtle on the selected row: #828593
}

func ExampleDefine() {
	control := roles.Define(roles.Spec{
		Name: "control",
		Kind: roles.KindSurface,
		Derive: func(p *roles.Palette) tint.Color {
			fg := p.Color(roles.Fg)
			return p.Color(roles.Raised).MixReadable(fg, fg, p.ByDark(0.22, 0.17), 4.5)
		},
		Distinct: []roles.Role{roles.Raised},
		Fallback: roles.Fallback{Slots: []tint.Slot{tint.SlotBrightBlack}},
	})
	pal := roles.Derive(tint.TintTokyoNight)
	fmt.Println(control, pal.Color(control).Hex())
	// Output:
	// control #45495c
}
