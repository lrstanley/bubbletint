// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"encoding/json"

	"github.com/lucasb-eyer/go-colorful"
)

type TintTemplate struct {
	StructName     string
	NameNormalized string

	Tint Tint
}

type Tint struct {
	Name string   `json:"name" validate:"required"`
	Meta TintMeta `json:"meta"`

	Fg          string `json:"foreground" validate:"required,hexcolor"`
	Bg          string `json:"background" validate:"required,hexcolor"`
	SelectionBg string `json:"selection"`
	Cursor      string `json:"cursor"`

	BrightBlack  string `json:"brightBlack"   validate:"required,hexcolor"`
	BrightBlue   string `json:"brightBlue"    validate:"required,hexcolor"`
	BrightCyan   string `json:"brightCyan"    validate:"required,hexcolor"`
	BrightGreen  string `json:"brightGreen"   validate:"required,hexcolor"`
	BrightPurple string `json:"brightMagenta" validate:"required,hexcolor"`
	BrightRed    string `json:"brightRed"     validate:"required,hexcolor"`
	BrightWhite  string `json:"brightWhite"   validate:"required,hexcolor"`
	BrightYellow string `json:"brightYellow"  validate:"required,hexcolor"`

	Black  string `json:"black"   validate:"required,hexcolor"`
	Blue   string `json:"blue"    validate:"required,hexcolor"`
	Cyan   string `json:"cyan"    validate:"required,hexcolor"`
	Green  string `json:"green"   validate:"required,hexcolor"`
	Purple string `json:"magenta" validate:"required,hexcolor"`
	Red    string `json:"red"     validate:"required,hexcolor"`
	White  string `json:"white"   validate:"required,hexcolor"`
	Yellow string `json:"yellow"  validate:"required,hexcolor"`
}

type TintMeta struct {
	IsDark  bool    `json:"isDark"`
	Credits Credits `json:"credits"`
}

func (t *Tint) IsDark() bool {
	bg, err := colorful.Hex(t.Bg)
	if err != nil {
		return false
	}
	_, _, l := bg.Hsl()
	return l < 0.5
}

type Credits []CreditSource

func (c *Credits) UnmarshalJSON(b []byte) error {
	var cs []CreditSource
	if err := json.Unmarshal(b, &cs); err != nil {
		var single CreditSource
		if serr := json.Unmarshal(b, &single); serr != nil {
			return err // Return original error since array is the default.
		}
		cs = append(cs, single)
	}
	*c = append(*c, cs...)
	return nil
}

type CreditSource struct {
	Name string `json:"name" validate:"required"`
	Link string `json:"link" validate:"required,url"`
}
