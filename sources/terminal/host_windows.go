// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package terminal

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

const (
	gaRootOwner                    = 3      // GA_ROOTOWNER
	processQueryLimitedInformation = 0x1000 // PROCESS_QUERY_LIMITED_INFORMATION
)

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	user32                         = syscall.NewLazyDLL("user32.dll")
	procGetConsoleWindow           = kernel32.NewProc("GetConsoleWindow")
	procQueryFullProcessImageNameW = kernel32.NewProc("QueryFullProcessImageNameW")
	procGetAncestor                = user32.NewProc("GetAncestor")
	procGetWindowThreadProcessID   = user32.NewProc("GetWindowThreadProcessId")
)

// hostProcess returns the lowercase executable name (e.g. "windowsterminal.exe")
// of the process that owns the window of the console we're attached to.
//
// Environment variables only tell which terminal started a process. When
// Windows Terminal is the default terminal application, a program started from
// anywhere else (Explorer, the Run dialog, another program) is hosted by Windows
// Terminal, but inherits its starter's environment, so has no WT_SESSION. The
// console window of a pseudoconsole session is owned by the terminal's window,
// which doesn't depend on any of that.
func hostProcess() (string, error) {
	console, _, _ := procGetConsoleWindow.Call()
	if console == 0 {
		return "", errors.New("not attached to a console window")
	}

	owner, _, callErr := procGetAncestor.Call(console, gaRootOwner)
	if owner == 0 {
		return "", fmt.Errorf("console window %#x has no owner: %w", console, callErr)
	}

	var pid uint32
	_, _, _ = procGetWindowThreadProcessID.Call(owner, uintptr(unsafe.Pointer(&pid)))
	if pid == 0 {
		return "", fmt.Errorf("no process for owner window %#x of console window %#x", owner, console)
	}

	handle, err := syscall.OpenProcess(processQueryLimitedInformation, false, pid)
	if err != nil {
		return "", fmt.Errorf("opening process %d owning the console window: %w", pid, err)
	}
	defer syscall.CloseHandle(handle) //nolint:errcheck

	const bufLen = 1024
	buf := make([]uint16, bufLen)
	size := uint32(bufLen)
	ok, _, callErr := procQueryFullProcessImageNameW.Call(
		uintptr(handle), 0, uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&size)),
	)
	if ok == 0 {
		return "", fmt.Errorf("reading the name of process %d owning the console window: %w", pid, callErr)
	}
	return strings.ToLower(filepath.Base(syscall.UTF16ToString(buf[:size]))), nil
}
