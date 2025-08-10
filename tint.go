// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

type CreditSource struct {
	// Name is the name of the credit source.
	Name string `json:"name,omitempty"`

	// Link is the link to the credit source.
	Link string `json:"link,omitempty"`
}

// Tint is a struct that represents each tint in this package.
type Tint struct {
	// DisplayName returns the display name of the tint.
	DisplayName string `json:"display_name,omitempty"`

	// ID returns the name of the tint (normalized, snakecase style).
	ID string `json:"id,omitempty"`

	// CreditSources returns the credit sources for the tint.
	CreditSources []*CreditSource `json:"credit_sources,omitempty"`

	// Dark returns whether the tint is dark (background color has a luminosity
	// less than 0.5).
	Dark bool `json:"dark,omitempty"`

	// Fg returns the recommended default foreground color for this tint.
	Fg *Color `json:"fg,omitempty"`

	// Bg returns the recommended default background color for this tint.
	Bg *Color `json:"bg,omitempty"`

	// SelectionBg returns the recommended background color for selected text. Note that this
	// is missing from most themes.
	SelectionBg *Color `json:"selection_bg,omitempty"`

	// Cursor returns the recommended color for the cursor. Note that this is missing from
	// most themes.
	Cursor *Color `json:"cursor,omitempty"`

	BrightBlack  *Color `json:"bright_black,omitempty"`
	BrightBlue   *Color `json:"bright_blue,omitempty"`
	BrightCyan   *Color `json:"bright_cyan,omitempty"`
	BrightGreen  *Color `json:"bright_green,omitempty"`
	BrightPurple *Color `json:"bright_purple,omitempty"`
	BrightRed    *Color `json:"bright_red,omitempty"`
	BrightWhite  *Color `json:"bright_white,omitempty"`
	BrightYellow *Color `json:"bright_yellow,omitempty"`

	Black  *Color `json:"black,omitempty"`
	Blue   *Color `json:"blue,omitempty"`
	Cyan   *Color `json:"cyan,omitempty"`
	Green  *Color `json:"green,omitempty"`
	Purple *Color `json:"purple,omitempty"`
	Red    *Color `json:"red,omitempty"`
	White  *Color `json:"white,omitempty"`
	Yellow *Color `json:"yellow,omitempty"`
}
