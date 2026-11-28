// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import "testing"

func TestSlotNames(t *testing.T) {
	t.Parallel()
	for s := SlotBlack; s <= SlotBrightWhite; s++ {
		for _, name := range []string{s.String(), s.JSONName()} {
			if got, ok := ParseSlot(name); !ok || got != s {
				t.Errorf("ParseSlot(%q) = %v, %v", name, got, ok)
			}
		}
	}
	tests := []struct {
		in   string
		want Slot
		ok   bool
	}{
		{in: "purple", want: SlotMagenta, ok: true},
		{in: "bright_purple", want: SlotBrightMagenta, ok: true},
		{in: "red", want: SlotRed, ok: true},
		{in: "orange"},
		{in: ""},
	}
	for _, tt := range tests {
		if got, ok := ParseSlot(tt.in); got != tt.want || ok != tt.ok {
			t.Errorf("ParseSlot(%q) = %v, %v, want %v, %v", tt.in, got, ok, tt.want, tt.ok)
		}
	}
	if SlotBrightMagenta.JSONName() != "bright_purple" || SlotRed.JSONName() != "red" || Slot(21).JSONName() != "unknown" {
		t.Errorf("JSONName() = %q, %q", SlotBrightMagenta.JSONName(), SlotRed.JSONName())
	}
	if Slot(16).String() != "unknown" {
		t.Errorf("Slot(16).String() = %q", Slot(16).String())
	}
}

func TestSlotVariants(t *testing.T) {
	t.Parallel()
	tests := []struct {
		s, bright, normal Slot
	}{
		{s: SlotRed, bright: SlotBrightRed, normal: SlotRed},
		{s: SlotBrightRed, bright: SlotBrightRed, normal: SlotRed},
		{s: SlotBlack, bright: SlotBrightBlack, normal: SlotBlack},
		{s: SlotBrightWhite, bright: SlotBrightWhite, normal: SlotWhite},
	}
	for _, tt := range tests {
		if got := tt.s.Bright(); got != tt.bright {
			t.Errorf("%v.Bright() = %v, want %v", tt.s, got, tt.bright)
		}
		if got := tt.s.Normal(); got != tt.normal {
			t.Errorf("%v.Normal() = %v, want %v", tt.s, got, tt.normal)
		}
		if got := tt.s.IsBright(); got != (tt.s == tt.bright) {
			t.Errorf("%v.IsBright() = %v", tt.s, got)
		}
	}
}

func TestTintANSI(t *testing.T) {
	t.Parallel()
	tn := TintTokyoNight
	want := [16]*Color{
		tn.Black, tn.Red, tn.Green, tn.Yellow, tn.Blue, tn.Purple, tn.Cyan, tn.White,
		tn.BrightBlack, tn.BrightRed, tn.BrightGreen, tn.BrightYellow, tn.BrightBlue, tn.BrightPurple, tn.BrightCyan, tn.BrightWhite,
	}
	if got := tn.ANSI16(); got != want {
		t.Error("ANSI16() isn't in ANSI order")
	}
	if tn.ANSI(Slot(16)) != nil {
		t.Error("ANSI(16) != nil")
	}
}

func TestSlotOf(t *testing.T) {
	t.Parallel()
	tests := []struct {
		h    Hue
		want Slot
		ok   bool
	}{
		{h: HueRed, want: SlotRed, ok: true},
		{h: HueOrange, want: SlotYellow, ok: true},
		{h: HueYellow, want: SlotYellow, ok: true},
		{h: HueGreen, want: SlotGreen, ok: true},
		{h: HueCyan, want: SlotCyan, ok: true},
		{h: HueBlue, want: SlotBlue, ok: true},
		{h: HuePurple, want: SlotMagenta, ok: true},
		{h: HuePink, want: SlotMagenta, ok: true},
		{h: HueNeutral},
		{h: HueMixed},
		{h: ""},
	}
	for _, tt := range tests {
		if got, ok := SlotOf(tt.h); got != tt.want || ok != tt.ok {
			t.Errorf("SlotOf(%q) = %v, %v, want %v, %v", tt.h, got, ok, tt.want, tt.ok)
		}
	}
}
