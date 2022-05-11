package rg

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	listStyle = lipgloss.NewStyle().
			Margin(0, 0, 0, 0).
			Padding(1, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	viewportStyle = lipgloss.NewStyle().
			Margin(0, 0, 0, 0).
			Padding(1, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
)

type KeyMap struct {
	Refresh     key.Binding
	Select      key.Binding
	SwitchFocus key.Binding
}

var DefaultKeyMap = KeyMap{
	Refresh: key.NewBinding(
		key.WithKeys("r", "R"),
		key.WithHelp("r/R", "refresh"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	SwitchFocus: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch focus"),
	),
}

type Model struct {
	keymap   KeyMap
	list     list.Model
	items    []list.Item
	viewport viewport.Model

	focused    int
	focusables [2]tea.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func NewModel() Model {
	m := Model{
		keymap:  DefaultKeyMap,
		focused: 0,
	}

	m.list = list.New(m.items, list.NewDefaultDelegate(), 0, 0)
	m.list.Title = "Users"

	return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Refresh):
			cmds = append(cmds, m.refresh())

		case key.Matches(msg, m.keymap.SwitchFocus):
			m.focused++
			if m.focused >= len(m.focusables) {
				m.focused = 0
			}
		}
	}

	var cmd tea.Cmd

	if m.focused == 0 {
		listStyle.BorderForeground(lipgloss.Color("#FFFFFF"))
		viewportStyle.BorderForeground(lipgloss.Color("#874BFD"))
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.focused == 1 {
		listStyle.BorderForeground(lipgloss.Color("#874BFD"))
		viewportStyle.BorderForeground(lipgloss.Color("#FFFFFF"))
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var view string

	view = lipgloss.JoinHorizontal(
		lipgloss.Top,
		listStyle.Render(m.list.View()),
		viewportStyle.Render(m.viewport.View()),
	)

	return view
}

func (m *Model) refresh() tea.Cmd {
	return func() tea.Msg {
		var items []list.Item
		fmt.Println("refresh")
		return items
	}
}
