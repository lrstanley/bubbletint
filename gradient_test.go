// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"reflect"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestFromToHex(t *testing.T) {
	type args struct {
		from  lipgloss.TerminalColor
		to    lipgloss.TerminalColor
		steps int
	}
	tests := []struct {
		name         string
		args         args
		wantGradient []string
	}{
		{
			name: "simple-3",
			args: args{
				from:  lipgloss.Color("#000000"),
				to:    lipgloss.Color("#ffffff"),
				steps: 3,
			},
			wantGradient: []string{"#000000", "#4e4e4e", "#a2a2a2"},
		},
		{
			name: "simple-10",
			args: args{
				from:  lipgloss.Color("#000000"),
				to:    lipgloss.Color("#ffffff"),
				steps: 10,
			},
			wantGradient: []string{"#000000", "#1b1b1b", "#303030", "#474747", "#5e5e5e", "#777777", "#919191", "#ababab", "#c6c6c6", "#e2e2e2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGradient := FromToHex(tt.args.from, tt.args.to, tt.args.steps)
			gotGradientColor := FromTo(tt.args.from, tt.args.to, tt.args.steps)

			if !reflect.DeepEqual(gotGradient, tt.wantGradient) {
				t.Errorf("FromToHex() = %v, want %v", gotGradient, tt.wantGradient)
			}

			for i := range gotGradient {
				c := lipgloss.Color(gotGradient[i])
				cr, cg, cb, ca := c.RGBA()
				gr, gg, gb, ga := gotGradientColor[i].RGBA()

				if cr != gr || cg != gg || cb != gb || ca != ga {
					t.Errorf("FromTo() = %v, want %v", gotGradientColor[i], c)
				}
			}
		})
	}
}
