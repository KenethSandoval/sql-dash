package navigation

import (
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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

var Navigation = []string{"Users", "Tables"}

// Model id tab navigation
type Model struct {
	CurrentId int
	spinner   spinner.Model
}

func NewModel() Model {
	m := Model{
		CurrentId: 0,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var items []string

	for i, nav := range Navigation {
		if m.CurrentId == i {
			items = append(items, activeTab.Render(nav))
		} else {
			items = append(items, tab.Render(nav))
		}
	}

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		items...,
	)

	gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-4)))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap, " ")

	return lipgloss.JoinHorizontal(lipgloss.Top, row, "\n\n")
}

// NthTab is placed in the currentid of the tab
func (m *Model) NthTab(nth int) {
	if nth > len(Navigation) {
		nth = len(Navigation)
	} else if nth < 1 {
		nth = 1
	}
	m.CurrentId = nth - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
