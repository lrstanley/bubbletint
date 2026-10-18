// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package tint

// Hue is the dominant chromatic family of a tint's ANSI palette.
// The zero value means the tint did not set a hue. Built-in tints always set one.
type Hue string

// Tone is an aesthetic label for a tint, such as muted or retro.
// The zero value means the tint is uncategorized or has not been classified.
type Tone string
