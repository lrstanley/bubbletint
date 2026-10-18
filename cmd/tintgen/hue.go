// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	hueNeutral = "neutral"
	hueMixed   = "mixed"

	// hueMinSaturation drops near-gray colors so they cannot decide the hue.
	hueMinSaturation = 0.15
	// hueMinLightness and hueMaxLightness drop colors that are almost black
	// or almost white, whose hue is unstable.
	hueMinLightness = 0.08
	hueMaxLightness = 0.92
	// hueMinShare is the fraction of saturation weight the leading bucket
	// must hold. Below that, the palette is mixed.
	hueMinShare = 0.45
)

// hueStop is one chromatic bucket. Until is the exclusive upper bound in
// degrees. Stops are in ascending order, which is also the tie break: the
// earlier stop wins. The first stop also owns the wrap from the last bound
// up to 360°.
type hueStop struct {
	Until float64
	Name  string
}

// hueStops is the hue catalog. Adding a chromatic hue is one entry here,
// then regenerate. neutral and mixed are outcomes, not stops.
var hueStops = []hueStop{
	{Until: 15, Name: "red"},
	{Until: 45, Name: "orange"},
	{Until: 75, Name: "yellow"},
	{Until: 165, Name: "green"},
	{Until: 195, Name: "cyan"},
	{Until: 255, Name: "blue"},
	{Until: 315, Name: "purple"},
	{Until: 345, Name: "pink"},
}

func hueNames() []string {
	names := make([]string, 0, len(hueStops)+2)
	for _, stop := range hueStops {
		names = append(names, stop.Name)
	}
	return append(names, hueNeutral, hueMixed)
}

// classifyHue reports the dominant hue of the twelve chromatic ANSI colors.
// Black, white, their bright variants, background, and foreground are ignored.
// Colors that are too gray, too dark, or too light are ignored. The rest are
// weighted by saturation.
func classifyHue(t Tint) string {
	weights := make(map[string]float64, len(hueStops))
	for _, hex := range chromaticColors(t) {
		c, err := colorful.Hex(hex)
		if err != nil {
			continue
		}
		h, s, l := c.Hsl()
		if s < hueMinSaturation || l < hueMinLightness || l > hueMaxLightness {
			continue
		}
		weights[hueBin(h)] += s
	}
	return hueFromWeights(weights)
}

func chromaticColors(t Tint) []string {
	return []string{
		t.Red, t.BrightRed,
		t.Green, t.BrightGreen,
		t.Yellow, t.BrightYellow,
		t.Blue, t.BrightBlue,
		t.Purple, t.BrightPurple,
		t.Cyan, t.BrightCyan,
	}
}

// hueBin maps a hue angle in degrees onto a chromatic bucket.
// Buckets are half-open on the upper bound. Angles at or past the last stop
// belong to the first stop, which is what lets red wrap across 0°.
func hueBin(angle float64) string {
	angle = math.Mod(angle, 360)
	if angle < 0 {
		angle += 360
	}
	for _, stop := range hueStops {
		if angle < stop.Until {
			return stop.Name
		}
	}
	return hueStops[0].Name
}

// hueFromWeights returns the leading bucket when it holds at least hueMinShare
// of the weight. Equal leaders resolve to the earlier hue stop.
func hueFromWeights(weights map[string]float64) string {
	var total float64
	for _, stop := range hueStops {
		total += weights[stop.Name]
	}
	if total == 0 {
		return hueNeutral
	}

	leadName := hueStops[0].Name
	lead := weights[leadName]
	for _, stop := range hueStops[1:] {
		if weights[stop.Name] > lead {
			lead = weights[stop.Name]
			leadName = stop.Name
		}
	}
	if lead/total < hueMinShare {
		return hueMixed
	}
	return leadName
}
