// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

type CreditSource struct {
	// Name is the name of the credit source.
	Name string

	// Link is the link to the credit source.
	Link string
}

// Tint is a struct that represents each tint in this package.
type Tint struct {
	// DisplayName returns the display name of the tint.
	DisplayName string

	// ID returns the name of the tint (normalized, snakecase style).
	ID string

	// CreditSources returns the credit sources for the tint.
	CreditSources []*CreditSource

	// Fg returns the recommended default foreground color for this tint.
	Fg *Color

	// Bg returns the recommended default background color for this tint.
	Bg *Color

	// SelectionBg returns the recommended background color for selected text. Note that this
	// is missing from most themes.
	SelectionBg *Color

	// Cursor returns the recommended color for the cursor. Note that this is missing from
	// most themes.
	Cursor *Color

	BrightBlack  *Color
	BrightBlue   *Color
	BrightCyan   *Color
	BrightGreen  *Color
	BrightPurple *Color
	BrightRed    *Color
	BrightWhite  *Color
	BrightYellow *Color

	Black  *Color
	Blue   *Color
	Cyan   *Color
	Green  *Color
	Purple *Color
	Red    *Color
	White  *Color
	Yellow *Color
}
