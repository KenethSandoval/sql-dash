package tables

import (
	"adminmsyql/dash/models"
	"adminmsyql/ui/uictx"
	"fmt"
	"math"

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
	ctx      *uictx.Ctx

	focused    int
	focusables [2]tea.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func NewModel(ctx *uictx.Ctx) Model {

	m := Model{
		keymap:  DefaultKeyMap,
		focused: 0,
	}

	m.list = list.New(m.items, list.NewDefaultDelegate(), 0, 0)
	m.list.Title = "Tables"
	m.ctx = ctx

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
		case key.Matches(msg, m.keymap.Select):
			_, ok := m.list.SelectedItem().(models.Tables)
			if ok {
				m.viewport.SetContent(m.renderViewport())
				return m, nil
			}
		}

	case tea.WindowSizeMsg:
		listWidth := int(math.Floor(float64(m.ctx.Content[0]) / 4.0))
		listHeigh := m.ctx.Content[1] - 1
		viewportWidth := m.ctx.Content[0] - listWidth - 4
		viewportHeight := m.ctx.Content[1] - 1

		listStyle.Width(listWidth)
		listStyle.Height(listHeigh)
		m.list.SetSize(listWidth-2, listHeigh-2)

		viewportStyle.Width(viewportWidth)
		viewportStyle.Height(viewportHeight)
		m.viewport = viewport.New(viewportWidth-4, viewportHeight-4)
		m.viewport.Width = viewportWidth - 4
		m.viewport.Height = viewportHeight - 4

	case []list.Item:
		m.items = msg
		m.list.SetItems(m.items)
		m.ctx.Loading = false
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

		tables, err := (*m.ctx.Client).ListTables()
		if err != nil {
			fmt.Printf("%s", err)
		}

		for _, table := range tables {
			items = append(items, table)
		}

		return items
	}
}

func (m *Model) renderViewport() string {
	table, err := (*m.ctx.Client).DescribeTables()
	if err != nil {
		fmt.Printf("%s", err)
	}
	var (
		vp string = ""
	)

	vp = fmt.Sprintf(
		"Field",
	)

	vp = fmt.Sprintf(
		"%s      Type",
		vp,
	)
	vp = fmt.Sprintf(
		"%s        Null",
		vp,
	)

	vp = fmt.Sprintf("%s\n\n", vp)

	for _, t := range table {
		vp = fmt.Sprintf(
			"%s%s      %s       %s\n",
			vp,
			t.Field,
			t.Type,
			t.Null,
		)
	}

	return vp
}
