// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Example modified from the following:
//  * https://github.com/charmbracelet/bubbletea/blob/test/examples/package-manager/main.go

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	tint "github.com/lrstanley/bubbletint"
	"github.com/muesli/termenv"
)

type model struct {
	packages []string
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
}

func newModel() model {
	p := progress.New(
		progress.WithColorProfile(termenv.TrueColor),
		progress.WithScaledGradient(tint.Hex(tint.BrightPurple()), tint.Hex(tint.BrightBlue())),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(tint.Fg())
	return model{
		packages: getPackages(),
		spinner:  s,
		progress: p,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(downloadAndInstall(m.packages[m.index]), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case installedPkgMsg:
		if m.index >= len(m.packages)-1 {
			// Everything's been installed. We're done!
			m.done = true
			return m, tea.Quit
		}

		// Update progress bar
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.packages)-1))

		m.index++

		checkMark := lipgloss.NewStyle().Foreground(tint.Green()).SetString("✓")

		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s", checkMark, m.packages[m.index]), // print success message above our program
			downloadAndInstall(m.packages[m.index]),             // download the next package
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	currentPkgNameStyle := lipgloss.NewStyle().Foreground(tint.BrightPurple())
	doneStyle := lipgloss.NewStyle().Margin(1, 2)

	n := len(m.packages)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done! Installed %d packages.\n", n))
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n-1)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+pkgCount))

	pkgName := currentPkgNameStyle.Render(m.packages[m.index])
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Installing " + pkgName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + pkgCount
}

type installedPkgMsg string

func downloadAndInstall(pkg string) tea.Cmd {
	// This is where you'd do i/o stuff to download and install packages. In
	// our case we're just pausing for a moment to simulate the process.
	d := time.Millisecond * time.Duration(rand.Intn(500))
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return installedPkgMsg(pkg)
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tint.NewDefaultRegistry()
	tint.SetTint(tint.TintICOrangePPL)

	rand.Seed(time.Now().Unix())

	if err := tea.NewProgram(newModel()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
