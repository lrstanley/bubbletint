// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Pair is a tint for dark terminals and one for light ones, by ID, for an
// app that follows the terminal's background. In a config file it's a tint
// ID, for both, or {"dark": ID, "light": ID}. [PairSchema] is its JSON
// schema.
type Pair struct {
	Dark  string `json:"dark"`
	Light string `json:"light"`
}

// PairSchema is the JSON schema for a [Pair], for an app's own config schema.
const PairSchema = `{
  "description": "a tint ID, or {\"dark\": ID, \"light\": ID} to follow the terminal's background",
  "anyOf": [
    {"type": "string", "minLength": 1},
    {
      "type": "object",
      "properties": {
        "dark": {"type": "string", "minLength": 1},
        "light": {"type": "string", "minLength": 1}
      },
      "additionalProperties": false
    }
  ]
}`

// ErrInvalidPair is returned for JSON that isn't a [Pair].
var ErrInvalidPair = errors.New(`want a tint ID, or {"dark": ID, "light": ID}`)

// For returns the tint ID for a dark terminal, or a light one. A terminal's
// background is dark by [IsDark].
func (p Pair) For(dark bool) string {
	if dark {
		return p.Dark
	}
	return p.Light
}

// MarshalJSON implements [json.Marshaler]: a single ID where both are the
// same.
func (p Pair) MarshalJSON() ([]byte, error) {
	if p.Dark == p.Light {
		return json.Marshal(p.Dark)
	}
	type plain Pair
	return json.Marshal(plain(p))
}

// UnmarshalJSON implements [json.Unmarshaler]. An object may leave out
// either side, which keeps the one already set, as from an app's defaults.
func (p *Pair) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err == nil {
		p.Dark, p.Light = id, id
		return nil
	}
	var pair struct {
		Dark  *string `json:"dark"`
		Light *string `json:"light"`
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&pair); err != nil {
		return ErrInvalidPair
	}
	if pair.Dark != nil {
		p.Dark = *pair.Dark
	}
	if pair.Light != nil {
		p.Light = *pair.Light
	}
	return nil
}
