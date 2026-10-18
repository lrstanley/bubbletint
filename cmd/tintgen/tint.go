// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"github.com/lucasb-eyer/go-colorful"
)

type TintTemplate struct {
	StructName     string
	NameNormalized string
	Hue            string
	Tone           string

	Tint Tint
}

type Tint struct {
	Name   string `json:"name" validate:"required"`
	Author string `json:"-"`

	Fg            string `json:"foreground" validate:"required,hexcolor"`
	Bg            string `json:"background" validate:"required,hexcolor"`
	SelectionBg   string `json:"selection"`
	SelectionText string `json:"-"`
	Bold          string `json:"-"`
	Cursor        string `json:"cursor"`
	CursorText    string `json:"-"`

	// Variant is the author-provided "dark" or "light" label from the YAML
	// source. Empty means the author did not set one.
	Variant string `json:"-"`

	BrightBlack  string `json:"brightBlack" validate:"required,hexcolor"`
	BrightBlue   string `json:"brightBlue" validate:"required,hexcolor"`
	BrightCyan   string `json:"brightCyan" validate:"required,hexcolor"`
	BrightGreen  string `json:"brightGreen" validate:"required,hexcolor"`
	BrightPurple string `json:"brightPurple" validate:"required,hexcolor"`
	BrightRed    string `json:"brightRed" validate:"required,hexcolor"`
	BrightWhite  string `json:"brightWhite" validate:"required,hexcolor"`
	BrightYellow string `json:"brightYellow" validate:"required,hexcolor"`

	Black  string `json:"black" validate:"required,hexcolor"`
	Blue   string `json:"blue" validate:"required,hexcolor"`
	Cyan   string `json:"cyan" validate:"required,hexcolor"`
	Green  string `json:"green" validate:"required,hexcolor"`
	Purple string `json:"purple" validate:"required,hexcolor"`
	Red    string `json:"red" validate:"required,hexcolor"`
	White  string `json:"white" validate:"required,hexcolor"`
	Yellow string `json:"yellow" validate:"required,hexcolor"`
}

func (t *Tint) IsDark() bool {
	switch t.Variant {
	case variantDark:
		return true
	case variantLight:
		return false
	}

	bg, err := colorful.Hex(t.Bg)
	if err != nil {
		return false
	}
	_, _, l := bg.Hsl()
	return l < 0.5
}
