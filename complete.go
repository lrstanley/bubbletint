// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"encoding/json"
	"fmt"
)

// Complete fills in what a hand-made tint, such as one from a config file,
// leaves out: each bright color from its normal one, and the display name
// from the ID. What's set is kept.
//
// Complete leaves Dark alone, as it can't tell an unset Dark from false.
// Where a config leaves "dark" out, set it from the background first, where
// there is one, with Bg.IsDark(), as [ParseJSON] does.
func (t *Tint) Complete() {
	for s := SlotBlack; s <= SlotWhite; s++ {
		if bright := t.slotField(s.Bright()); *bright == nil {
			*bright = *t.slotField(s)
		}
	}
	if t.DisplayName == "" {
		t.DisplayName = t.ID
	}
}

// Missing returns the JSON names of the colors a full palette needs that t
// doesn't set: its text and background, and the eight normal colors, which
// [Tint.Complete] fills the bright ones from. It's empty for a full palette.
func (t *Tint) Missing() []string {
	var out []string
	if t.Fg == nil {
		out = append(out, "fg")
	}
	if t.Bg == nil {
		out = append(out, "bg")
	}
	for s := SlotBlack; s <= SlotWhite; s++ {
		if *t.slotField(s) == nil {
			out = append(out, s.JSONName())
		}
	}
	return out
}

// ParseJSON reads a hand-made tint, such as one from a config file, in the
// JSON a [Tint] marshals to, and completes it with [Tint.Complete]. Where
// "dark" is left out, Dark is worked out from the background, by
// [Color.IsDark]; a light tint on a dark background says "dark": false. Use
// [Tint.Missing] to check it has a full palette.
func ParseJSON(data []byte) (*Tint, error) {
	var t struct {
		Tint

		Dark *bool `json:"dark"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, fmt.Errorf("parsing tint: %w", err)
	}
	switch {
	case t.Dark != nil:
		t.Tint.Dark = *t.Dark
	case t.Bg != nil:
		t.Tint.Dark = t.Bg.IsDark()
	}
	t.Complete()
	return &t.Tint, nil
}
