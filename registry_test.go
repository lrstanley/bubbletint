// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"reflect"
	"testing"
)

func TestNewRegistry(t *testing.T) {
	type args struct {
		defaultTint *Tint
		tints       []*Tint
	}
	tests := []struct {
		name string
		args args
		want *Registry
	}{
		{
			name: "no default",
			args: args{defaultTint: nil, tints: []*Tint{TintDraculaPlus}},
			want: &Registry{
				currentTint: "",
				tints: map[string]*Tint{
					TintDraculaPlus.ID: TintDraculaPlus,
				},
			},
		},
		{
			name: "no default and no tints",
			args: args{defaultTint: nil, tints: nil},
			want: &Registry{
				currentTint: "",
				tints:       map[string]*Tint{},
			},
		},
		{
			name: "with default and tints",
			args: args{defaultTint: Tint3024Day, tints: []*Tint{TintDraculaPlus}},
			want: &Registry{
				currentTint: Tint3024Day.ID,
				tints: map[string]*Tint{
					TintDraculaPlus.ID: TintDraculaPlus,
					Tint3024Day.ID:     Tint3024Day,
				},
			},
		},
		{
			name: "with default and no tints",
			args: args{defaultTint: Tint3024Day, tints: nil},
			want: &Registry{
				currentTint: Tint3024Day.ID,
				tints: map[string]*Tint{
					Tint3024Day.ID: Tint3024Day,
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

	reg.UnregisterIDs(TintDraculaPlus.ID)
	if len(reg.Tints()) != 1 {
		t.Errorf("expected 1 tint, got %d", len(reg.Tints()))
	}

	if other := reg.Tints()[0]; other.ID != Tint3024Day.ID {
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

	if reg.Current().ID != Tint3024Day.ID {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.Current())
	}

	reg.SetTintID(TintDraculaPlus.ID)

	if reg.Current().ID != TintDraculaPlus.ID {
		t.Errorf("expected %v, got %v", TintDraculaPlus, reg.Current())
	}

	if ok := reg.SetTintID("invalid"); ok {
		t.Errorf("expected false, got true")
	}
}

func TestRegistry_PreviousNextTint(t *testing.T) {
	// Intentionally adding tints in the list out of order.
	reg := NewRegistry(TintDraculaPlus, TintDraculaPlus, Tint3024Day, TintSynthwave)

	if reg.Current().ID != TintDraculaPlus.ID {
		t.Errorf("expected %v, got %v", TintDraculaPlus, reg.Current())
	}

	reg.NextTint()

	if reg.Current().ID != TintSynthwave.ID {
		t.Errorf("expected %v, got %v", TintSynthwave, reg.Current())
	}

	reg.NextTint()

	if reg.Current().ID != TintSynthwave.ID {
		t.Errorf("expected %v, got %v", TintSynthwave, reg.Current())
	}

	reg.PreviousTint()

	if reg.Current().ID != TintDraculaPlus.ID {
		t.Errorf("expected %v, got %v", TintDraculaPlus, reg.Current())
	}

	reg.PreviousTint()

	if reg.Current().ID != Tint3024Day.ID {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.Current())
	}

	reg.PreviousTint()

	if reg.Current().ID != Tint3024Day.ID {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.Current())
	}

	// If there currently isn't a default, then the first tint in the list is used.
	reg = NewRegistry(nil, TintDraculaPlus, Tint3024Day, TintSynthwave)
	reg.NextTint()

	if reg.Current().ID != Tint3024Day.ID {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.Current())
	}

	reg = NewRegistry(nil, TintDraculaPlus, Tint3024Day, TintSynthwave)
	reg.PreviousTint()

	if reg.Current().ID != Tint3024Day.ID {
		t.Errorf("expected %v, got %v", Tint3024Day, reg.Current())
	}
}
