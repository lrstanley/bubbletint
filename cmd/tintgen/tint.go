// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

type TintTemplate struct {
	StructName     string
	NameNormalized string

	Tint          Tint
	CreditSources []CreditSource
}

type Tint struct {
	Name string `json:"name"`

	Fg          string `json:"foreground"`
	Bg          string `json:"background"`
	SelectionBg string `json:"selectionBackground"`
	Cursor      string `json:"cursorColor"`

	BrightBlack  string `json:"brightBlack"`
	BrightBlue   string `json:"brightBlue"`
	BrightCyan   string `json:"brightCyan"`
	BrightGreen  string `json:"brightGreen"`
	BrightPurple string `json:"brightPurple"`
	BrightRed    string `json:"brightRed"`
	BrightWhite  string `json:"brightWhite"`
	BrightYellow string `json:"brightYellow"`

	Black  string `json:"black"`
	Blue   string `json:"blue"`
	Cyan   string `json:"cyan"`
	Green  string `json:"green"`
	Purple string `json:"purple"`
	Red    string `json:"red"`
	White  string `json:"white"`
	Yellow string `json:"yellow"`
}

type Credit struct {
	Tints   []string       `json:"themeNames"`
	Sources []CreditSource `json:"sources"`
}

type CreditSource struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
