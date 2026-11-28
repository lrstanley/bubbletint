// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package roles

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/colorprofile"
	tint "github.com/lrstanley/bubbletint/v2"
)

// Settings an app reads from its config file. Each parses itself from text,
// so it can be a field of a config struct, and has a JSON schema, for an
// app's own config schema.

// AccentChoice is a user's choice of accent: "auto" to follow the tint's own
// hue, one of the tint's colors by name ("blue", "purple", or "bright_cyan",
// say), or a fixed color, "#rgb" or "#rrggbb". The zero value is auto.
type AccentChoice struct {
	slot  tint.Slot
	color *tint.Color
	isSet bool // a slot is chosen
}

// AccentSchema is the JSON schema for an [AccentChoice].
const AccentSchema = `{
  "type": "string",
  "description": "auto follows the tint's own hue; or one of its colors (red, green, yellow, blue, purple, magenta, cyan, or their bright_ variants), or a color as \"#rrggbb\"",
  "anyOf": [
    {"enum": ["auto", "red", "green", "yellow", "blue", "purple", "magenta", "cyan", "bright_red", "bright_green", "bright_yellow", "bright_blue", "bright_purple", "bright_magenta", "bright_cyan"]},
    {"pattern": "^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"}
  ]
}`

// ErrInvalidAccent is returned for text that isn't an [AccentChoice].
var ErrInvalidAccent = errors.New("want auto, one of the tint's colors (red, green, yellow, blue, purple, cyan, or a bright_ one), or #rrggbb")

// ParseAccent parses an [AccentChoice]. An empty string is auto.
func ParseAccent(s string) (AccentChoice, error) {
	s = strings.TrimSpace(s)
	switch {
	case s == "" || strings.EqualFold(s, autoName):
		return AccentChoice{}, nil
	case strings.HasPrefix(s, "#"):
		if c := tint.FromHex(s); c != nil {
			return AccentChoice{color: c}, nil
		}
	default:
		if slot, ok := tint.ParseSlot(strings.ToLower(s)); ok && Hue(slot) != Neutral {
			return AccentChoice{slot: slot, isSet: true}, nil
		}
	}
	return AccentChoice{}, fmt.Errorf("accent %q: %w", s, ErrInvalidAccent)
}

// Auto reports whether the accent follows the tint's own hue.
func (a AccentChoice) Auto() bool { return a.color == nil && !a.isSet }

// Slot returns the tint's color the accent is, and whether it's one: false
// for auto, or a fixed color.
func (a AccentChoice) Slot() (tint.Slot, bool) { return a.slot, a.isSet }

// Option returns the option that derives with the accent: [WithAccentSlot]
// or [WithAccentColor], or one that changes nothing, for auto.
func (a AccentChoice) Option() Option {
	switch {
	case a.color != nil:
		return WithAccentColor(*a.color)
	case a.isSet:
		return WithAccentSlot(a.slot)
	}
	return func(*options) {}
}

// String returns the accent as [ParseAccent] reads it: "auto", a slot's name,
// or "#rrggbb". A slot is named as a [tint.Tint]'s JSON names it, by
// [tint.Slot.JSONName].
func (a AccentChoice) String() string {
	switch {
	case a.color != nil:
		return a.color.Hex()
	case a.isSet:
		return a.slot.JSONName()
	}
	return autoName
}

// Accents returns the accents to offer as choices: auto, then each of the
// tint's colors that has a hue, normal then bright. A fixed color is also an
// accent, but isn't listed.
func Accents() []AccentChoice {
	out := []AccentChoice{{}}
	for s := tint.SlotBlack; s <= tint.SlotBrightWhite; s++ {
		if Hue(s) != Neutral {
			out = append(out, AccentChoice{slot: s, isSet: true})
		}
	}
	return out
}

// MarshalText implements [encoding.TextMarshaler].
func (a AccentChoice) MarshalText() ([]byte, error) { return []byte(a.String()), nil }

// UnmarshalText implements [encoding.TextUnmarshaler].
func (a *AccentChoice) UnmarshalText(b []byte) error {
	v, err := ParseAccent(string(b))
	if err != nil {
		return err
	}
	*a = v
	return nil
}

// ColorLimit is the most colors a user wants drawn, whatever the terminal
// can do: "auto" for all it can, "256", "16", or "none" for attributes only.
// The zero value is auto, so a field tagged omitzero leaves auto out when
// it's marshaled, a config's defaults included.
type ColorLimit uint8

// Color limits.
const (
	ColorsAuto ColorLimit = iota
	Colors256
	Colors16
	ColorsNone
)

// ColorLimits returns every [ColorLimit], from auto to none, to offer as
// choices.
func ColorLimits() []ColorLimit {
	out := make([]ColorLimit, len(colorLimitNames))
	for i := range out {
		out[i] = ColorLimit(i) //nolint:gosec // bounded by the names
	}
	return out
}

// ColorLimitSchema is the JSON schema for a [ColorLimit].
const ColorLimitSchema = `{
  "type": "string",
  "description": "the most colors to use; auto uses what the terminal supports, and none only bold, faint, underline and reverse",
  "enum": ["auto", "256", "16", "none"]
}`

// autoName is the setting that follows the tint, or the terminal.
const autoName = "auto"

var colorLimitNames = [...]string{ColorsAuto: autoName, Colors256: "256", Colors16: "16", ColorsNone: "none"}

// ErrInvalidColorLimit is returned for text that isn't a [ColorLimit].
var ErrInvalidColorLimit = errors.New("want auto, 256, 16 or none")

// ParseColorLimit parses a [ColorLimit]. An empty string is auto.
func ParseColorLimit(s string) (ColorLimit, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return ColorsAuto, nil
	}
	for i, n := range colorLimitNames {
		if n == s {
			return ColorLimit(i), nil //nolint:gosec // bounded by the names
		}
	}
	return ColorsAuto, fmt.Errorf("colors %q: %w", s, ErrInvalidColorLimit)
}

// Cap returns the profile to fit to: detected, with no more colors than the
// limit allows. Unknown is taken as TrueColor, and NoTTY as ASCII, as
// [Palette.Fit] takes them.
func (l ColorLimit) Cap(detected colorprofile.Profile) colorprofile.Profile {
	p := normalize(Target{Profile: detected}).Profile
	switch l {
	case Colors256:
		return min(p, colorprofile.ANSI256)
	case Colors16:
		return min(p, colorprofile.ANSI)
	case ColorsNone:
		return colorprofile.ASCII
	case ColorsAuto:
	}
	return p
}

// String returns the limit as [ParseColorLimit] reads it.
func (l ColorLimit) String() string {
	if int(l) < len(colorLimitNames) {
		return colorLimitNames[l]
	}
	return autoName
}

// MarshalText implements [encoding.TextMarshaler].
func (l ColorLimit) MarshalText() ([]byte, error) { return []byte(l.String()), nil }

// UnmarshalText implements [encoding.TextUnmarshaler].
func (l *ColorLimit) UnmarshalText(b []byte) error {
	v, err := ParseColorLimit(string(b))
	if err != nil {
		return err
	}
	*l = v
	return nil
}
