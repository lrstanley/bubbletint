// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

// CreditSource is an author or project credited with a tint.
type CreditSource struct {
	// Name is the author credit. Theme data stores the full author text here,
	// which may include a URL.
	Name string `json:"name,omitempty"`

	// Link is a URL for the credit source.
	//
	// Deprecated: theme data does not include a separate link. Any URL is part
	// of Name.
	Link string `json:"link,omitempty"`
}

// Tint is a struct that represents each tint in this package.
type Tint struct {
	// DisplayName returns the display name of the tint.
	DisplayName string `json:"display_name,omitempty"`

	// ID returns the name of the tint (normalized, snakecase style).
	ID string `json:"id,omitempty"`

	// Author is the theme author's credit line, when the YAML source provides one.
	Author string `json:"author,omitempty"`

	// CreditSources lists author credits for this tint. Built-in tints set Name
	// from Author when the YAML source provides one.
	CreditSources []*CreditSource `json:"credit_sources,omitempty"`

	// Dark reports whether the tint is a dark theme. Built-in tints use the
	// variant provided by the theme author when it is present. When the author
	// did not provide one, Dark is calculated from the background color's
	// luminosity (dark when it is below 0.5).
	Dark bool `json:"dark,omitempty"`

	// Hue is the dominant chromatic family of this tint's ANSI palette.
	// Built-in tints always set one. The zero value means it was not set.
	Hue Hue `json:"hue,omitempty"`

	// Tone is an aesthetic label such as muted or retro.
	// The zero value means uncategorized or not yet classified.
	Tone Tone `json:"tone,omitempty"`

	// Fg returns the recommended default foreground color for this tint.
	Fg *Color `json:"fg,omitempty"`

	// Bg returns the recommended default background color for this tint.
	Bg *Color `json:"bg,omitempty"`

	// SelectionBg returns the recommended background color for selected text. Note that this
	// is missing from most themes.
	SelectionBg *Color `json:"selection_bg,omitempty"`

	// SelectionText is the foreground color for selected text, when the theme
	// author set one.
	SelectionText *Color `json:"selection_text,omitempty"`

	// Bold is the color for bold text, when the theme author set one.
	Bold *Color `json:"bold,omitempty"`

	// Cursor returns the recommended color for the cursor. Note that this is missing from
	// most themes.
	Cursor *Color `json:"cursor,omitempty"`

	// CursorText is the color for text drawn on the cursor, when the theme
	// author set one.
	CursorText *Color `json:"cursor_text,omitempty"`

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
