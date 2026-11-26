// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles_test

import (
	"fmt"

	"github.com/charmbracelet/colorprofile"
	"github.com/lrstanley/bubbletint/roles/v2"
	tint "github.com/lrstanley/bubbletint/v2"
)

func ExampleDerive() {
	pal := roles.Derive(tint.TintTokyoNight)
	for _, r := range []roles.Role{roles.Bg, roles.Raised, roles.Muted, roles.Accent, roles.Danger, roles.Danger.Soft()} {
		fmt.Printf("%-12s %s\n", r, pal.Color(r).Hex())
	}
	// Output:
	// bg           #1a1b26
	// raised       #272937
	// muted        #7f85a4
	// accent       #79a2f7
	// danger       #f7768e
	// danger-soft  #3b2a36
}

func ExamplePalette_Fit() {
	pal := roles.Derive(tint.TintTokyoNight)
	for _, p := range []colorprofile.Profile{colorprofile.TrueColor, colorprofile.ANSI256, colorprofile.ANSI, colorprofile.ASCII} {
		f := pal.Fit(roles.Target{Profile: p})
		var c any = f.Color(roles.Muted)
		if tc, ok := c.(tint.Color); ok {
			c = tc.Hex()
		}
		fmt.Printf("%-9s %T %v, attr %v\n", p, f.Color(roles.Muted), c, f.Attr(roles.Muted))
	}
	// Output:
	// TrueColor tint.Color #7f85a4, attr none
	// ANSI256   ansi.IndexedColor 103, attr none
	// ANSI      ansi.BasicColor 7, attr none
	// Ascii     <nil> <nil>, attr faint
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

func ExampleFitted_Fill() {
	f := roles.Derive(tint.TintTokyoNight).Fit(roles.Target{Profile: colorprofile.ANSI})
	pill := f.Fill(roles.Danger)
	hint := f.Text(roles.Subtle)
	fmt.Println("danger pill:", pill.GetBackground(), "on", pill.GetForeground())
	fmt.Println("subtle hint:", hint.GetForeground())

	f = f.Palette().Fit(roles.Target{Profile: colorprofile.ASCII})
	fmt.Println("with no color, the pill is reversed:", f.Fill(roles.Danger).GetReverse())
	// Output:
	// danger pill: 9 on 0
	// subtle hint: 7
	// with no color, the pill is reversed: true
}
