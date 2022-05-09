package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	FirstTab  key.Binding
	SecondTab key.Binding
	// more tab
	Quit key.Binding
}

var Navigation = []string{"Tab1", "Tab2"}

var DefaultKeyMap = KeyMap{
	FirstTab: key.NewBinding(
		key.WithKeys("f1"),
		key.WithHelp("f1", "First tab"),
	),
	SecondTab: key.NewBinding(
		key.WithKeys("f2"),
		key.WithHelp("f2", "Second tab"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+q"),
		key.WithHelp("q/Q", "quit"),
	),
}

var (
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlight).
		Padding(0, 1)

	activeTab = tab.Copy().Border(activeTabBorder, true)

	tabGap = tab.Copy().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)
)

const width = 96

type Model struct {
	keymap KeyMap
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	var items []string

	for _, nav := range Navigation {
		items = append(items, tab.Render(nav))
	}

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		items...,
	)

	gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

	return lipgloss.JoinHorizontal(lipgloss.Top, row, "\n\n")
}

func NewProgram() *tea.Program {
	m := Model{}
	return tea.NewProgram(m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
