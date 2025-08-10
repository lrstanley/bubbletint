module github.com/lrstanley/bubbletint/examples

go 1.24.4

replace (
	github.com/lrstanley/bubbletint/chromatint/v2 => ../chromatint
	github.com/lrstanley/bubbletint/v2 => ../
)

require (
	github.com/alecthomas/chroma/v2 v2.20.0
	github.com/charmbracelet/bubbles/v2 v2.0.0-beta.1
	github.com/charmbracelet/bubbletea/v2 v2.0.0-beta.4
	github.com/charmbracelet/lipgloss/v2 v2.0.0-beta.3
	github.com/lrstanley/bubbletint/chromatint/v2 v2.0.0-00010101000000-000000000000
	github.com/lrstanley/bubbletint/v2 v2.0.0-alpha.7
	github.com/lrstanley/bubblezone/v2 v2.0.0-alpha.2
)

require (
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/charmbracelet/colorprofile v0.3.1 // indirect
	github.com/charmbracelet/harmonica v0.2.0 // indirect
	github.com/charmbracelet/x/ansi v0.9.3 // indirect
	github.com/charmbracelet/x/cellbuf v0.0.14-0.20250505150409-97991a1f17d1 // indirect
	github.com/charmbracelet/x/input v0.3.7 // indirect
	github.com/charmbracelet/x/term v0.2.1 // indirect
	github.com/charmbracelet/x/windows v0.2.1 // indirect
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
)
