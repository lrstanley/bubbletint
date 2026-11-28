// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"math"
	"slices"

	tint "github.com/lrstanley/bubbletint/v2"
)

// orColor is the first color that's set, or def.
func orColor(a, b *tint.Color, def tint.Color) tint.Color {
	switch {
	case a != nil:
		return opaque(*a)
	case b != nil:
		return opaque(*b)
	}
	return def
}

// opaque is c with full alpha: terminals draw no other kind.
func opaque(c tint.Color) tint.Color {
	c.A = 0xff
	return c
}

func sameRGB(a, b tint.Color) bool { return a.R == b.R && a.G == b.G && a.B == b.B }

func pickBool[T any](cond bool, ifTrue, ifFalse T) T {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func chroma(c tint.Color) float64 { return c.OKLCh().C }

func hue(c tint.Color) float64 { return c.OKLCh().H }

// hueDist is the distance between two hues in degrees, from 0 to 180.
func hueDist(a, b float64) float64 {
	d := math.Mod(math.Abs(a-b), 360)
	if d > 180 {
		d = 360 - d
	}
	return d
}

func without(list []tint.Slot, s tint.Slot) []tint.Slot {
	return slices.DeleteFunc(slices.Clone(list), func(x tint.Slot) bool { return x == s })
}

// mono reports whether status can't be told apart by hue: a grey success or
// danger, or hues that sit close together. Red and green need to be well
// apart, as they're the pair most often confused.
func mono(success, warning, danger tint.Color) bool {
	s, w, d := success.OKLCh(), warning.OKLCh(), danger.OKLCh()
	return s.C < 0.04 || d.C < 0.04 || hueDist(s.H, d.H) < 45 ||
		(hueDist(w.H, d.H) < 12 && hueDist(s.H, w.H) < 30)
}
