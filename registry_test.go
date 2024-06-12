// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"reflect"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestNewRegistry(t *testing.T) {
	type args struct {
		defaultTint Tint
		tints       []Tint
	}
	tests := []struct {
		name string
		args args
		want *Registry
	}{
		{
			name: "no default",
			args: args{defaultTint: nil, tints: []Tint{TintDraculaPlus}},
			want: &Registry{
				currentTint: "",
				tints: map[string]Tint{
					TintDraculaPlus.ID(): TintDraculaPlus,
				},
			},
		},
		{
			name: "no default and no tints",
			args: args{defaultTint: nil, tints: nil},
			want: &Registry{
				currentTint: "",
				tints:       map[string]Tint{},
			},
		},
		{
			name: "with default and tints",
			args: args{defaultTint: Tint3024Day, tints: []Tint{TintDraculaPlus}},
			want: &Registry{
				currentTint: Tint3024Day.ID(),
				tints: map[string]Tint{
					TintDraculaPlus.ID(): TintDraculaPlus,
					Tint3024Day.ID():     Tint3024Day,
				},
			},
		},
		{
			name: "with default and no tints",
			args: args{defaultTint: Tint3024Day, tints: nil},
			want: &Registry{
				currentTint: Tint3024Day.ID(),
				tints: map[string]Tint{
					Tint3024Day.ID(): Tint3024Day,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegistry(tt.args.defaultTint, tt.args.tints...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegistry_Unregister(t *testing.T) {
	reg := NewRegistry(nil, Tint3024Day, TintDraculaPlus)

	if len(reg.Tints()) != 2 {
		t.Errorf("expected 2 tints, got %d", len(reg.Tints()))
	}

	reg.Unregister(Tint3024Day)

	if len(reg.Tints()) != 1 {
		t.Errorf("expected 1 tint, got %d", len(reg.Tints()))
	}

	if other := reg.Tints()[0]; other != TintDraculaPlus {
		t.Errorf("expected %v, got %v", TintDraculaPlus, other)
	}

	reg.Register(Tint3024Day)
	if len(reg.Tints()) != 2 {
		t.Errorf("expected 2 tints, got %d", len(reg.Tints()))
	}

	reg.UnregisterIDs(TintDraculaPlus.ID())
	if len(reg.Tints()) != 1 {
		t.Errorf("expected 1 tint, got %d", len(reg.Tints()))
	}

	if other := reg.Tints()[0]; other.ID() != Tint3024Day.ID() {
		t.Errorf("expected %v, got %v", Tint3024Day, other)
	}

	reg = NewRegistry(nil, Tint3024Day, TintDraculaPlus)
	reg.UnregisterAll()

	if len(reg.Tints()) != 0 {
		t.Errorf("expected 0 tints, got %d", len(reg.Tints()))
	}
}

func TestRegistry_SetTint(t *testing.T) {
	reg := NewRegistry(nil, Tint3024Day, TintDraculaPlus)

	reg.SetTint(Tint3024Day)

	if reg.GetCurrentTint().ID() != Tint3024Day.ID() {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.GetCurrentTint())
	}

	reg.SetTintID(TintDraculaPlus.ID())

	if reg.GetCurrentTint().ID() != TintDraculaPlus.ID() {
		t.Errorf("expected %v, got %v", TintDraculaPlus, reg.GetCurrentTint())
	}

	if ok := reg.SetTintID("invalid"); ok {
		t.Errorf("expected false, got true")
	}
}

func TestRegistry_PreviousNextTint(t *testing.T) {
	// Intentionally adding tints in the list out of order.
	reg := NewRegistry(TintDraculaPlus, TintDraculaPlus, Tint3024Day, TintSynthwave)

	if reg.GetCurrentTint().ID() != TintDraculaPlus.ID() {
		t.Errorf("expected %v, got %v", TintDraculaPlus, reg.GetCurrentTint())
	}

	reg.NextTint()

	if reg.GetCurrentTint().ID() != TintSynthwave.ID() {
		t.Errorf("expected %v, got %v", TintSynthwave, reg.GetCurrentTint())
	}

	reg.NextTint()

	if reg.GetCurrentTint().ID() != TintSynthwave.ID() {
		t.Errorf("expected %v, got %v", TintSynthwave, reg.GetCurrentTint())
	}

	reg.PreviousTint()

	if reg.GetCurrentTint().ID() != TintDraculaPlus.ID() {
		t.Errorf("expected %v, got %v", TintDraculaPlus, reg.GetCurrentTint())
	}

	reg.PreviousTint()

	if reg.GetCurrentTint().ID() != Tint3024Day.ID() {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.GetCurrentTint())
	}

	reg.PreviousTint()

	if reg.GetCurrentTint().ID() != Tint3024Day.ID() {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.GetCurrentTint())
	}

	// If there currently isn't a default, then the first tint in the list is used.
	reg = NewRegistry(nil, TintDraculaPlus, Tint3024Day, TintSynthwave)
	reg.NextTint()

	if reg.GetCurrentTint().ID() != Tint3024Day.ID() {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.GetCurrentTint())
	}

	reg = NewRegistry(nil, TintDraculaPlus, Tint3024Day, TintSynthwave)
	reg.PreviousTint()

	if reg.GetCurrentTint().ID() != Tint3024Day.ID() {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.GetCurrentTint())
	}
}

func TestRegistry_AllColors(t *testing.T) {
	NewDefaultRegistry()

	ids := DefaultRegistry.TintIDs()

	for _, id := range ids {
		t.Run(id, func(t *testing.T) {
			DefaultRegistry.SetTintID(id)

			colors := []lipgloss.TerminalColor{
				Fg(),
				Bg(),
				SelectionBg(),
				Cursor(),
				BrightBlack(),
				BrightBlue(),
				BrightCyan(),
				BrightGreen(),
				BrightPurple(),
				BrightRed(),
				BrightWhite(),
				BrightYellow(),
				Black(),
				Blue(),
				Cyan(),
				Green(),
				Purple(),
				Red(),
				White(),
				Yellow(),

				DefaultRegistry.Fg(),
				DefaultRegistry.Bg(),
				DefaultRegistry.SelectionBg(),
				DefaultRegistry.Cursor(),
				DefaultRegistry.BrightBlack(),
				DefaultRegistry.BrightBlue(),
				DefaultRegistry.BrightCyan(),
				DefaultRegistry.BrightGreen(),
				DefaultRegistry.BrightPurple(),
				DefaultRegistry.BrightRed(),
				DefaultRegistry.BrightWhite(),
				DefaultRegistry.BrightYellow(),
				DefaultRegistry.Black(),
				DefaultRegistry.Blue(),
				DefaultRegistry.Cyan(),
				DefaultRegistry.Green(),
				DefaultRegistry.Purple(),
				DefaultRegistry.Red(),
				DefaultRegistry.White(),
				DefaultRegistry.Yellow(),
			}

			for _, c := range colors {
				if c == nil {
					t.Errorf("expected color, got nil")
				}
			}
		})
	}
}
