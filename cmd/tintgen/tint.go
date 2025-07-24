// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import "github.com/lucasb-eyer/go-colorful"

type TintTemplate struct {
	StructName     string
	NameNormalized string

	Tint          Tint
	CreditSources []CreditSource
}

type Tint struct {
	Name string `json:"name" validate:"required"`

	Fg          string `json:"foreground" validate:"required,hexcolor"`
	Bg          string `json:"background" validate:"required,hexcolor"`
	SelectionBg string `json:"selectionBackground"`
	Cursor      string `json:"cursorColor"`

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
	bg, err := colorful.Hex(t.Bg)
	if err != nil {
		return false
	}
	_, _, l := bg.Hsl()
	return l < 0.5
}

type Credit struct {
	Tints   []string       `json:"themeNames" validate:"required,min=1"`
	Sources []CreditSource `json:"sources" validate:"required,min=1"`
}

type CreditSource struct {
	Name string `json:"name" validate:"required"`
	Link string `json:"link" validate:"required,url"`
}
