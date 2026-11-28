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
	row := pal.On(roles.Bg.Selected())
	fmt.Println("subtle on the page:", pal.Color(roles.Subtle).Hex())
	fmt.Println("subtle on the selected row:", row.Color(roles.Subtle).Hex())
	pill := pal.On(roles.Accent)
	fmt.Println("text on an accent pill:", pill.Color(roles.Fg).Hex(), "on", pill.Color(roles.Accent).Hex())
	// Output:
	// subtle on the page: #9194a2
	// subtle on the selected row: #7a85a0
	// text on an accent pill: #eff1f5 on #3878f6
}

func ExampleRole_Hover() {
	pal := roles.Derive(tint.TintTokyoNight)
	for _, r := range []roles.Role{roles.Overlay, roles.Overlay.Hover(), roles.Overlay.Pressed(), roles.Bg.Selected(), roles.Bg.Selected().Hover()} {
		fmt.Printf("%-18s %s\n", r, pal.Color(r).Hex())
	}
	// Output:
	// overlay            #343647
	// overlay:hover      #43465a
	// overlay:pressed    #50536a
	// bg:selected        #33467c
	// bg:selected:hover  #415489
}

func ExampleDefine() {
	// Shadow is a role an app adds: the shadow under a dragged pane.
	shadow := roles.Define(roles.Spec{
		Name: "shadow",
		Kind: roles.KindSurface,
		Derive: func(p *roles.Palette) tint.Color {
			return p.Color(roles.Bg).Mix(tint.Color{A: 0xff}, 0.5)
		},
		Fallback: roles.Fallback{Slots: []tint.Slot{tint.SlotBlack}},
	})
	pal := roles.Derive(tint.TintTokyoNight)
	fmt.Println(shadow, pal.Color(shadow).Hex())
	// Output:
	// shadow #040508
}

func ExampleFitted_Fill() {
	f := roles.Derive(tint.TintTokyoNight).Fit(roles.Target{Profile: colorprofile.ANSI})
	pill := f.Fill(roles.Danger)
	hint := f.Text(roles.Subtle)
	fmt.Println("danger pill:", pill.GetBackground(), "on", pill.GetForeground())
	fmt.Println("subtle hint:", hint.GetForeground())
	fmt.Println("a row under the pointer is underlined:", f.Fill(roles.Bg.Hover()).GetUnderline())

	f = f.Palette().Fit(roles.Target{Profile: colorprofile.ASCII})
	fmt.Println("with no color, the pill is reversed:", f.Fill(roles.Danger).GetReverse())
	// Output:
	// danger pill: 9 on 0
	// subtle hint: 7
	// a row under the pointer is underlined: true
	// with no color, the pill is reversed: true
}

func ExampleFitted_Distinct() {
	pal := roles.Derive(tint.TintTokyoNight)
	for _, p := range []colorprofile.Profile{colorprofile.TrueColor, colorprofile.ANSI256, colorprofile.ANSI} {
		f := pal.Fit(roles.Target{Profile: p})
		// Outline a dialog where its surface won't show against the page.
		fmt.Printf("%-9s outline dialogs: %v\n", p, !f.Distinct(roles.Raised, roles.Bg))
	}
	// Output:
	// TrueColor outline dialogs: false
	// ANSI256   outline dialogs: false
	// ANSI      outline dialogs: true
}

func ExampleParseAccent() {
	for _, s := range []string{"auto", "purple", "#ff8800", "teal"} {
		a, err := roles.ParseAccent(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		pal := roles.Derive(tint.TintTokyoNight, a.Option())
		fmt.Printf("%-7s accent %s\n", a, pal.Color(roles.Accent).Hex())
	}
	// Output:
	// auto    accent #79a2f7
	// magenta accent #bb9af7
	// #ff8800 accent #ff8800
	// accent "teal": want auto, one of the tint's colors (red, green, yellow, blue, purple, cyan, or a bright_ one), or #rrggbb
}
