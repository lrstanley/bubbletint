// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"image/color"
	"math"
	"testing"
)

var (
	black = Color{A: 0xff}
	white = Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
)

func TestContrast(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		a, b color.Color
		want float64
	}{
		{name: "black-white", a: black, b: white, want: 21},
		{name: "white-black", a: white, b: black, want: 21},
		{name: "same", a: Color{R: 0x80, G: 0x40, B: 0x20, A: 0xff}, b: Color{R: 0x80, G: 0x40, B: 0x20, A: 0xff}, want: 1},
		{name: "gray-white", a: Color{R: 0x76, G: 0x76, B: 0x76, A: 0xff}, b: white, want: 4.54},
		{name: "pointer-and-stdlib", a: &Color{A: 0xff}, b: color.White, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Contrast(tt.a, tt.b); math.Abs(got-tt.want) > 0.01 {
				t.Errorf("Contrast() = %.3f, want %.3f", got, tt.want)
			}
		})
	}
}

func TestOKLabRoundTrip(t *testing.T) {
	t.Parallel()
	for r := 0; r < 256; r += 15 {
		for g := 0; g < 256; g += 15 {
			for b := 0; b < 256; b += 15 {
				c := Color{R: uint8(r), G: uint8(g), B: uint8(b), A: 0xff}
				lab := c.OKLab()
				if !lab.InGamut() {
					t.Fatalf("%s: OKLab() not in gamut", c.Hex())
				}
				if got := lab.Color(); got != c {
					t.Fatalf("OKLab().Color() = %s, want %s", got.Hex(), c.Hex())
				}
				if got := c.OKLCh().Lab().Color(); got != c {
					t.Fatalf("OKLCh().Lab().Color() = %s, want %s", got.Hex(), c.Hex())
				}
			}
		}
	}
	if l := white.OKLab().L; math.Abs(l-1) > 1e-6 {
		t.Errorf("white L = %f, want 1", l)
	}
	if c := (Color{R: 0x80, G: 0x80, B: 0x80, A: 0xff}).OKLCh().C; c > 1e-6 {
		t.Errorf("gray chroma = %f, want 0", c)
	}
}

func TestOKLChColorReducesChroma(t *testing.T) {
	t.Parallel()
	// Far out of gamut: chroma is reduced, keeping lightness.
	c := OKLCh{L: 0.7, C: 0.5, H: 150}.Color()
	if l := c.OKLab().L; math.Abs(l-0.7) > 0.01 {
		t.Errorf("L = %.3f, want 0.7", l)
	}
}

func TestMix(t *testing.T) {
	t.Parallel()
	red := Color{R: 0xff, A: 0xff}
	tests := []struct {
		name string
		t    float64
		want Color
	}{
		{name: "start", t: 0, want: red},
		{name: "below", t: -1, want: red},
		{name: "end", t: 1, want: white},
		{name: "above", t: 2, want: white},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := red.Mix(white, tt.t); got != tt.want {
				t.Errorf("Mix() = %s, want %s", got.Hex(), tt.want.Hex())
			}
		})
	}

	// Halfway in OKLab is halfway in lightness.
	mid := black.Mix(white, 0.5).OKLab().L
	if math.Abs(mid-0.5) > 0.01 {
		t.Errorf("black.Mix(white, 0.5) L = %.3f, want 0.5", mid)
	}
	if got := Mix(nil, white, 0.5); got != nil {
		t.Errorf("Mix(nil, ...) = %v, want nil", got)
	}
	expectColorMatches(t, Mix(color.Black, color.White, 1), white)
}

func TestEnsureContrast(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		c, bg     Color
		ratio     float64
		unchanged bool
	}{
		{name: "passes", c: white, bg: black, ratio: 4.5, unchanged: true},
		{name: "dark-bg", c: Color{R: 0x30, G: 0x30, B: 0x60, A: 0xff}, bg: Color{R: 0x1a, G: 0x1b, B: 0x26, A: 0xff}, ratio: 4.5},
		{name: "light-bg", c: Color{R: 0xe0, G: 0xc0, B: 0x40, A: 0xff}, bg: Color{R: 0xfa, G: 0xfa, B: 0xfa, A: 0xff}, ratio: 3.5},
		{name: "saturated", c: Color{B: 0xff, A: 0xff}, bg: Color{R: 0x10, G: 0x10, B: 0x40, A: 0xff}, ratio: 7},
		{name: "mid-gray-bg", c: Color{R: 0x77, G: 0x77, B: 0x77, A: 0xff}, bg: Color{R: 0x70, G: 0x70, B: 0x70, A: 0xff}, ratio: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.c.EnsureContrast(tt.bg, tt.ratio)
			if tt.unchanged && got != tt.c {
				t.Errorf("EnsureContrast() = %s, want unchanged %s", got.Hex(), tt.c.Hex())
			}
			if r := Contrast(got, tt.bg); r < tt.ratio {
				t.Errorf("EnsureContrast() = %s, contrast %.2f < %.2f", got.Hex(), r, tt.ratio)
			}
			// Hue is kept, where there's chroma left to keep.
			if a, b := tt.c.OKLCh(), got.OKLCh(); a.C > 0.05 && b.C > 0.05 {
				d := math.Abs(a.H - b.H)
				if d > 180 {
					d = 360 - d
				}
				if d > 10 {
					t.Errorf("hue moved %.1f degrees", d)
				}
			}
		})
	}
	if got := EnsureContrast(nil, white, 3); got != nil {
		t.Errorf("EnsureContrast(nil) = %v, want nil", got)
	}
}

func TestMixReadable(t *testing.T) {
	t.Parallel()
	bg := Color{R: 0x1a, G: 0x1b, B: 0x26, A: 0xff}
	fg := Color{R: 0xa9, G: 0xb1, B: 0xd6, A: 0xff}
	got := bg.MixReadable(fg, fg, 0.6, 4.5)
	if r := Contrast(fg, got); r < 4.5 {
		t.Errorf("text contrast %.2f < 4.5", r)
	}
	if got == bg {
		t.Error("MixReadable() didn't move at all")
	}
	if more := bg.Mix(fg, 0.6); Contrast(fg, more) >= 4.5 {
		t.Error("test case should need the mix reduced")
	}
	expectColorMatches(t, MixReadable(bg, fg, fg, 0.1, 4.5), bg.Mix(fg, 0.1))
}

func TestFromColor(t *testing.T) {
	t.Parallel()
	if FromColor(nil) != nil {
		t.Error("FromColor(nil) != nil")
	}
	var p *Color
	if FromColor(p) != nil {
		t.Error("FromColor((*Color)(nil)) != nil")
	}
	if got := FromColor(color.NRGBA{R: 0x10, G: 0x20, B: 0x30, A: 0x00}); *got != (Color{R: 0x00, G: 0x00, B: 0x00, A: 0xff}) {
		// A transparent NRGBA premultiplies to black.
		t.Errorf("FromColor(transparent) = %+v", *got)
	}
	if got := FromColor(Color{R: 1, G: 2, B: 3}); *got != (Color{R: 1, G: 2, B: 3, A: 0xff}) {
		t.Errorf("FromColor(Color) = %+v, want opaque", *got)
	}
}

func TestIsDark(t *testing.T) {
	t.Parallel()
	if !IsDark(Color{R: 0x1a, G: 0x1b, B: 0x26, A: 0xff}) || IsDark(white) {
		t.Error("IsDark() wrong for tokyo night bg or white")
	}
}

func TestDistance(t *testing.T) {
	t.Parallel()
	if d := white.Distance(white); d != 0 {
		t.Errorf("Distance(self) = %f", d)
	}
	if d := black.Distance(white); math.Abs(d-1) > 1e-3 {
		t.Errorf("Distance(black, white) = %f, want 1", d)
	}
}

func TestColorMethods(t *testing.T) {
	t.Parallel()
	for _, c := range []Color{black, white, {R: 0x1a, G: 0x1b, B: 0x26, A: 0xff}, {R: 0x76, G: 0x76, B: 0x76, A: 0xff}} {
		if got, want := c.Luminance(), Luminance(c); got != want {
			t.Errorf("%s: Luminance() = %f, want %f", c.Hex(), got, want)
		}
		if got, want := c.Contrast(white), Contrast(&c, color.White); got != want {
			t.Errorf("%s: Contrast() = %f, want %f", c.Hex(), got, want)
		}
		if got, want := c.IsDark(), IsDark(c); got != want {
			t.Errorf("%s: IsDark() = %v, want %v", c.Hex(), got, want)
		}
	}
}

func TestIsDarkCrossover(t *testing.T) {
	t.Parallel()
	// On every gray, IsDark is whether white text contrasts more than black.
	for v := range 256 {
		c := Color{R: uint8(v), G: uint8(v), B: uint8(v), A: 0xff}
		if want := c.Contrast(white) > c.Contrast(black); c.IsDark() != want {
			t.Errorf("%s: IsDark() = %v, want %v", c.Hex(), c.IsDark(), want)
		}
	}
}
