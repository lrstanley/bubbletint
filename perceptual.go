// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"image/color"
	"math"
)

// The color science here follows published specifications, so results match
// other implementations:
//
//   - sRGB's transfer function, between gamma-encoded and linear light, is from
//     IEC 61966-2-1, as given in CSS Color 4:
//     https://www.w3.org/TR/css-color-4/#color-conversion-code
//   - OKLab, and its matrices, are Björn Ottosson's:
//     https://bottosson.github.io/posts/oklab/
//   - Relative luminance and contrast ratio are WCAG 2.2's:
//     https://www.w3.org/TR/WCAG22/#dfn-relative-luminance and
//     https://www.w3.org/TR/WCAG22/#dfn-contrast-ratio

// sRGB transfer function constants (IEC 61966-2-1).
const (
	srgbThreshold    = 0.04045   // encoded value below which the curve is linear
	srgbLinearThresh = 0.0031308 // the same point, in linear light
	srgbSlope        = 12.92     // slope of the linear segment
	srgbOffset       = 0.055     // offset of the power segment
	srgbGamma        = 2.4       // exponent of the power segment
)

// WCAG 2.2 relative luminance: the Rec. 709 primaries' weights, and the flare
// added to both luminances in the contrast ratio.
const (
	lumR, lumG, lumB = 0.2126, 0.7152, 0.0722
	contrastFlare    = 0.05
)

// darkLuminance is the luminance below which white text contrasts more than
// black: where (1+f)/(L+f) = (L+f)/f, so L = sqrt(f*(1+f)) - f, about 0.179.
var darkLuminance = math.Sqrt(contrastFlare*(1+contrastFlare)) - contrastFlare

// gamutEpsilon is how far outside 0-1 a linear channel may be and still count
// as in gamut, to absorb floating-point error in the round trip.
const gamutEpsilon = 1e-6

// Tuning values, not from a specification.
const (
	// chromaStep is how much chroma is scaled by on each step toward the
	// gamut, by [OKLCh.Color] and [Color.EnsureContrast]. Smaller steps keep
	// more chroma; larger ones take fewer.
	chromaStep = 0.94
	// minChroma is where chroma stops being reduced: the color is gray.
	minChroma = 0.001
	// lightnessStep is how far [Color.EnsureContrast] moves OKLab lightness on
	// each step, fine enough that it overshoots the ratio by little.
	lightnessStep = 0.01
	// mixStep is how far [Color.MixReadable] backs off the mix on each step.
	mixStep = 0.01
)

// OKLab is a color in the OKLab space, where equal distances look about equally
// different. L runs from 0 (black) to 1 (white).
type OKLab struct{ L, A, B float64 }

// OKLCh is [OKLab] in polar form: lightness, chroma (0 is grey), and hue in
// degrees.
type OKLCh struct{ L, C, H float64 }

// toLinear decodes an sRGB channel to linear light, from 0 to 1.
func toLinear(v uint8) float64 {
	f := float64(v) / 255
	if f <= srgbThreshold {
		return f / srgbSlope
	}
	return math.Pow((f+srgbOffset)/(1+srgbOffset), srgbGamma)
}

// fromLinear encodes linear light as an sRGB channel, clipping to 0-255.
func fromLinear(v float64) uint8 {
	if v <= srgbLinearThresh {
		v *= srgbSlope
	} else {
		v = (1+srgbOffset)*math.Pow(v, 1/srgbGamma) - srgbOffset
	}
	return uint8(math.Round(clamp(v, 0, 1) * 255))
}

// OKLab returns c in OKLab. The matrices are from
// https://bottosson.github.io/posts/oklab/: linear sRGB to LMS cone response,
// a cube root, then LMS to Lab.
func (c Color) OKLab() OKLab {
	r, g, b := toLinear(c.R), toLinear(c.G), toLinear(c.B)
	l := math.Cbrt(0.4122214708*r + 0.5363325363*g + 0.0514459929*b)
	m := math.Cbrt(0.2119034982*r + 0.6806995451*g + 0.1073969566*b)
	s := math.Cbrt(0.0883024619*r + 0.2817188376*g + 0.6299787005*b)
	return OKLab{
		L: 0.2104542553*l + 0.7936177850*m - 0.0040720468*s,
		A: 1.9779984951*l - 2.4285922050*m + 0.4505937099*s,
		B: 0.0259040371*l + 0.7827717662*m - 0.8086757660*s,
	}
}

// OKLCh returns c in OKLCh.
func (c Color) OKLCh() OKLCh { return c.OKLab().LCh() }

// linear is the inverse of [Color.OKLab], before encoding: Lab to LMS, cubed,
// then LMS to linear sRGB.
func (o OKLab) linear() (r, g, b float64) {
	l := o.L + 0.3963377774*o.A + 0.2158037573*o.B
	m := o.L - 0.1055613458*o.A - 0.0638541728*o.B
	s := o.L - 0.0894841775*o.A - 1.2914855480*o.B
	l, m, s = l*l*l, m*m*m, s*s*s
	return 4.0767416621*l - 3.3077115913*m + 0.2309699292*s,
		-1.2684380046*l + 2.6097574011*m - 0.3413193965*s,
		-0.0041960863*l - 0.7034186147*m + 1.7076147010*s
}

// InGamut reports whether o is a displayable sRGB color.
func (o OKLab) InGamut() bool {
	r, g, b := o.linear()
	lo, hi := -gamutEpsilon, 1+gamutEpsilon
	return r >= lo && r <= hi && g >= lo && g <= hi && b >= lo && b <= hi
}

// Color converts o to sRGB, clipping each channel.
func (o OKLab) Color() Color {
	r, g, b := o.linear()
	return Color{R: fromLinear(r), G: fromLinear(g), B: fromLinear(b), A: 0xff}
}

// LCh returns o in polar form.
func (o OKLab) LCh() OKLCh {
	h := math.Atan2(o.B, o.A) * 180 / math.Pi
	if h < 0 {
		h += 360
	}
	return OKLCh{L: o.L, C: math.Hypot(o.A, o.B), H: h}
}

// Lab returns o in rectangular form.
func (o OKLCh) Lab() OKLab {
	r := o.H * math.Pi / 180
	return OKLab{L: o.L, A: o.C * math.Cos(r), B: o.C * math.Sin(r)}
}

// Color converts o to sRGB. Out of gamut, chroma is reduced, keeping lightness
// and hue, rather than clipping channels.
func (o OKLCh) Color() Color {
	for !o.Lab().InGamut() && o.C > minChroma {
		o.C *= chromaStep
	}
	return o.Lab().Color()
}

// Mix moves c toward o by t, from 0 (c) to 1 (o), in OKLab, so halfway looks
// halfway. Unlike [Color.Lighten] and [Color.Darken], it keeps hue.
func (c Color) Mix(o Color, t float64) Color {
	switch {
	case t <= 0:
		return c
	case t >= 1:
		return o
	}
	a, b := c.OKLab(), o.OKLab()
	return OKLab{L: a.L + (b.L-a.L)*t, A: a.A + (b.A-a.A)*t, B: a.B + (b.B-a.B)*t}.Color()
}

// Mix is [Color.Mix] for any [color.Color]. It returns nil if either is nil.
func Mix(a, b color.Color, t float64) color.Color {
	if a == nil || b == nil {
		return nil
	}
	return colorOf(a).Mix(colorOf(b), t)
}

// Distance is the perceptual distance between two colors (Euclidean in OKLab,
// or deltaEOK): about 0.02 is a just-noticeable difference, the value CSS
// Color 4's gamut mapping uses
// (https://www.w3.org/TR/css-color-4/#binsearch).
func (c Color) Distance(o Color) float64 {
	a, b := c.OKLab(), o.OKLab()
	return math.Sqrt((a.L-b.L)*(a.L-b.L) + (a.A-b.A)*(a.A-b.A) + (a.B-b.B)*(a.B-b.B))
}

// Luminance is c's WCAG relative luminance, from 0 (black) to 1 (white). See
// https://www.w3.org/TR/WCAG22/#dfn-relative-luminance.
func (c Color) Luminance() float64 {
	return lumR*toLinear(c.R) + lumG*toLinear(c.G) + lumB*toLinear(c.B)
}

// Luminance is [Color.Luminance] for any [color.Color]. nil is black.
func Luminance(c color.Color) float64 { return colorOf(c).Luminance() }

// Contrast is the WCAG contrast ratio between c and o, from 1 (the same) to 21
// (black and white). WCAG's AA level asks 4.5 of body text, and 3 of large text
// and of controls' outlines. See
// https://www.w3.org/TR/WCAG22/#dfn-contrast-ratio.
func (c Color) Contrast(o Color) float64 {
	x, y := c.Luminance(), o.Luminance()
	if x < y {
		x, y = y, x
	}
	return (x + contrastFlare) / (y + contrastFlare)
}

// Contrast is [Color.Contrast] for any [color.Color]. nil is black.
func Contrast(a, b color.Color) float64 { return colorOf(a).Contrast(colorOf(b)) }

// IsDark reports whether c is dark: whether white text on it contrasts more
// than black text does.
func (c Color) IsDark() bool { return c.Luminance() < darkLuminance }

// IsDark is [Color.IsDark] for any [color.Color]. nil is black, so dark.
func IsDark(c color.Color) bool { return colorOf(c).IsDark() }

// EnsureContrast returns c, moved in lightness (keeping its hue) until it
// contrasts with against by at least ratio. It moves away from against first,
// and the other way if that runs out of room. A color that already passes is
// returned unchanged.
func (c Color) EnsureContrast(against Color, ratio float64) Color {
	if c.Contrast(against) >= ratio {
		return c
	}
	dir := -1.0
	if against.IsDark() {
		dir = 1
	}
	best, ok := walkContrast(c, against, ratio, dir)
	if !ok {
		if other, ok2 := walkContrast(c, against, ratio, -dir); ok2 || other.Contrast(against) > best.Contrast(against) {
			best = other
		}
	}
	return best
}

// EnsureContrast is [Color.EnsureContrast] for any [color.Color].
func EnsureContrast(c, against color.Color, ratio float64) color.Color {
	if c == nil || against == nil {
		return c
	}
	return colorOf(c).EnsureContrast(colorOf(against), ratio)
}

func walkContrast(c, against Color, ratio, dir float64) (Color, bool) {
	lch := c.OKLCh()
	out := c
	for range int(1 / lightnessStep) {
		lch.L = clamp(lch.L+dir*lightnessStep, 0, 1)
		// Chroma stays reduced once a step falls out of gamut, so the walk
		// doesn't swing back to a more saturated color.
		for !lch.Lab().InGamut() && lch.C > minChroma {
			lch.C *= chromaStep
		}
		out = lch.Lab().Color()
		if out.Contrast(against) >= ratio {
			return out, true
		}
		if lch.L <= 0 || lch.L >= 1 {
			break
		}
	}
	return out, false
}

// MixReadable is the strongest mix of c toward toward, at most t, on which
// text still contrasts by ratio: a fill that stands out as much as it can
// without losing the text on it. The mix is reduced in steps of 0.01, down to
// c itself.
func (c Color) MixReadable(toward, text Color, t, ratio float64) Color {
	for t = min(t, 1); t > mixStep/2; t -= mixStep {
		if out := c.Mix(toward, t); text.Contrast(out) >= ratio {
			return out
		}
	}
	return c
}

// MixReadable is [Color.MixReadable] for any [color.Color].
func MixReadable(base, toward, text color.Color, t, ratio float64) color.Color {
	if base == nil || toward == nil || text == nil {
		return base
	}
	return colorOf(base).MixReadable(colorOf(toward), colorOf(text), t, ratio)
}
