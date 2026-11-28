// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

// Slot is one of the 16 ANSI colors, by its terminal index: SlotRed is 1,
// SlotBrightRed is 9. It is the indexed view of a [Tint]'s Black..BrightWhite
// fields.
type Slot uint8

const (
	SlotBlack Slot = iota
	SlotRed
	SlotGreen
	SlotYellow
	SlotBlue
	SlotMagenta // the Tint field is named Purple
	SlotCyan
	SlotWhite
	SlotBrightBlack
	SlotBrightRed
	SlotBrightGreen
	SlotBrightYellow
	SlotBrightBlue
	SlotBrightMagenta
	SlotBrightCyan
	SlotBrightWhite
)

// ChromaticSlots are the six slots with a hue, normal variants, in ANSI order.
var ChromaticSlots = [...]Slot{SlotRed, SlotGreen, SlotYellow, SlotBlue, SlotMagenta, SlotCyan}

var slotNames = [16]string{
	"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white",
	"bright_black", "bright_red", "bright_green", "bright_yellow", "bright_blue", "bright_magenta", "bright_cyan", "bright_white",
}

// String returns the slot's name, such as "bright_red".
func (s Slot) String() string {
	if s > SlotBrightWhite {
		return "unknown"
	}
	return slotNames[s]
}

// JSONName returns the name a [Tint]'s JSON gives the slot's color: its
// [Slot.String], but purple for magenta, such as "bright_purple".
func (s Slot) JSONName() string {
	if s > SlotBrightWhite || s.Normal() != SlotMagenta {
		return s.String()
	}
	if s.IsBright() {
		return "bright_purple"
	}
	return "purple"
}

// ParseSlot parses a slot name as returned by [Slot.String] or
// [Slot.JSONName].
func ParseSlot(name string) (Slot, bool) {
	switch name {
	case "purple":
		return SlotMagenta, true
	case "bright_purple":
		return SlotBrightMagenta, true
	}
	for i, n := range slotNames {
		if n == name {
			return Slot(i), true
		}
	}
	return 0, false
}

// IsBright reports whether s is one of the bright variants (8-15).
func (s Slot) IsBright() bool { return s >= SlotBrightBlack }

// Bright returns the bright variant of s. Bright slots are returned as is.
func (s Slot) Bright() Slot { return s&7 | 8 }

// Normal returns the normal variant of s. Normal slots are returned as is.
func (s Slot) Normal() Slot { return s & 7 }

// SlotOf returns the normal slot for a hue bucket: orange maps to yellow, pink
// to magenta. Neutral, mixed and unset hues report false.
func SlotOf(h Hue) (Slot, bool) {
	switch h { //nolint:exhaustive
	case HueRed:
		return SlotRed, true
	case HueOrange, HueYellow:
		return SlotYellow, true
	case HueGreen:
		return SlotGreen, true
	case HueCyan:
		return SlotCyan, true
	case HueBlue:
		return SlotBlue, true
	case HuePurple, HuePink:
		return SlotMagenta, true
	}
	return 0, false
}

// ANSI returns the tint's color for slot s, which may be nil for a
// hand-built tint.
func (t *Tint) ANSI(s Slot) *Color {
	if s > SlotBrightWhite {
		return nil
	}
	return *t.slotField(s)
}

// slotField is the field holding slot s, which must be 0-15.
func (t *Tint) slotField(s Slot) **Color {
	return [16]**Color{
		&t.Black, &t.Red, &t.Green, &t.Yellow, &t.Blue, &t.Purple, &t.Cyan, &t.White,
		&t.BrightBlack, &t.BrightRed, &t.BrightGreen, &t.BrightYellow,
		&t.BrightBlue, &t.BrightPurple, &t.BrightCyan, &t.BrightWhite,
	}[s]
}

// ANSI16 returns all 16 slots in index order.
func (t *Tint) ANSI16() (out [16]*Color) {
	for i := range out {
		out[i] = t.ANSI(Slot(i))
	}
	return out
}
