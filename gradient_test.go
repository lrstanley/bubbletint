// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func mustHex(t *testing.T, hex string) color.Color {
	t.Helper()

	c, err := colorful.Hex(hex)
	if err != nil {
		panic(err)
	}

	return c
}

func TestFromToHex(t *testing.T) {
	tests := []struct {
		name         string
		from         color.Color
		to           color.Color
		steps        int
		wantGradient []string
	}{
		{
			name:         "simple-3",
			from:         mustHex(t, "#000000"),
			to:           mustHex(t, "#ffffff"),
			steps:        3,
			wantGradient: []string{"#000000", "#4e4e4e", "#a2a2a2"},
		},
		{
			name:         "simple-10",
			from:         mustHex(t, "#000000"),
			to:           mustHex(t, "#ffffff"),
			steps:        10,
			wantGradient: []string{"#000000", "#1b1b1b", "#303030", "#474747", "#5e5e5e", "#777777", "#919191", "#ababab", "#c6c6c6", "#e2e2e2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGradient := FromToHex(tt.from, tt.to, tt.steps)

			if !reflect.DeepEqual(gotGradient, tt.wantGradient) {
				t.Errorf("FromToHex() = %v, want %v", gotGradient, tt.wantGradient)
			}
		})
	}
}
