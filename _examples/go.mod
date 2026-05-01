module github.com/lrstanley/bubbletint/examples

go 1.24.4

replace (
	github.com/lrstanley/bubbletint/chromatint/v2 => ../chromatint
	github.com/lrstanley/bubbletint/v2 => ../
)

require (
	charm.land/bubbles/v2 v2.0.0
	charm.land/bubbletea/v2 v2.0.0
	charm.land/lipgloss/v2 v2.0.0
	github.com/alecthomas/chroma/v2 v2.24.0
	github.com/lrstanley/bubbletint/chromatint/v2 v2.0.0-alpha.1
	github.com/lrstanley/bubbletint/v2 v2.0.0-alpha.10
	github.com/lrstanley/bubblezone/v2 v2.0.0-alpha.3
)

require (
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/charmbracelet/colorprofile v0.4.2 // indirect
	github.com/charmbracelet/harmonica v0.2.0 // indirect
	github.com/charmbracelet/ultraviolet v0.0.0-20260223171050-89c142e4aa73 // indirect
	github.com/charmbracelet/x/ansi v0.11.6 // indirect
	github.com/charmbracelet/x/term v0.2.2 // indirect
	github.com/charmbracelet/x/termios v0.1.1 // indirect
	github.com/charmbracelet/x/windows v0.2.2 // indirect
	github.com/clipperhouse/displaywidth v0.11.0 // indirect
	github.com/clipperhouse/uax29/v2 v2.7.0 // indirect
	github.com/dlclark/regexp2 v1.12.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.3.0 // indirect
	github.com/mattn/go-runewidth v0.0.20 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
)
