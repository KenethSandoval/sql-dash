package ui

import (
	"adminmsyql/ui/navigation"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	FirstTab  key.Binding
	SecondTab key.Binding
	// more tab
	Quit key.Binding
}

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

type Model struct {
	keymap KeyMap
	nav    navigation.Model
	debug  bool
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			m.debug = true
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.FirstTab):
			m.nav.NthTab(1)
			fmt.Println("1")
			return m, nil
		case key.Matches(msg, DefaultKeyMap.SecondTab):
			m.nav.NthTab(2)
			fmt.Println("2")
			return m, nil
		}
	}

	nav, cmd := m.nav.Update(msg)
	m.nav = nav
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.debug {
		return "Bye\n"
	}

	s := strings.Builder{}
	s.WriteString(m.nav.View() + "\n\n")
	return s.String()
}

func NewProgram() *tea.Program {
	m := Model{}
	return tea.NewProgram(m)
}
