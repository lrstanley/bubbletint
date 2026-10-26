// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

//go:build !windows

package terminal

// hostProcess is only implemented on Windows, where a console can be hosted by
// a terminal without that being visible in the environment.
func hostProcess() (string, error) { return "", nil }
