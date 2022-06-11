package rg

import (
	"adminmsyql/dash/models"
	"adminmsyql/ui/bar"
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
	bar      bar.Bubble

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
	m.list.Title = "Users"
	m.ctx = ctx

	return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Refresh):
			m.ctx.Loading = true
			cmds = append(cmds, m.refresh())

		case key.Matches(msg, m.keymap.SwitchFocus):
			m.focused++
			if m.focused >= len(m.focusables) {
				m.focused = 0
			}

		case key.Matches(msg, m.keymap.Select):
			i, ok := m.list.SelectedItem().(models.Credential)
			if ok {
				m.viewport.SetContent(m.renderViewport(&i))
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

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, listStyle.Render(m.list.View()),
			viewportStyle.Render(m.viewport.View()),
			m.bar.View(),
		),
	)
}

func (m *Model) refresh() tea.Cmd {
	return func() tea.Msg {
		var items []list.Item

		users, err := (*m.ctx.Client).ListProfile()
		if err != nil {
			fmt.Printf("%s", err)
		}

		for _, user := range users {
			items = append(items, user)
		}

		return items
	}
}

func (m *Model) renderViewport(user *models.Credential) string {
	var (
		vp         string = ""
		insertPriv string = ""
		selectPriv string = ""
		updatePriv string = ""
		deletePriv string = ""
		createPriv string = ""
		dropPriv   string = ""
		grantPriv  string = ""
		indexPriv  string = ""
		alterPriv  string = ""
	)

	if user.SelectPriv == "Y" {
		selectPriv = "Sentencias SELECT"
	}

	if user.InsertPriv == "Y" {
		insertPriv = "Sentencias INSERT"
	}

	if user.UpdatePriv == "Y" {
		updatePriv = "Sentencias UPDATE"
	}

	if user.DeletePriv == "Y" {
		deletePriv = "Sentencias DELETE"
	}

	if user.CreatePriv == "Y" {
		createPriv = "Sentencias CREATE"
	}

	if user.DropPriv == "Y" {
		dropPriv = "Sentencias DROP"
	}

	if user.GrantPriv == "Y" {
		grantPriv = "Sentencias GRANT"
	}

	if user.IndexPriv == "Y" {
		indexPriv = "Sentencias INDEX"
	}

	if user.AlterPriv == "Y" {
		alterPriv = "Sentencias ALTER"
	}

	vp = fmt.Sprintf(
		"%s\n\nPrivilegios \n",
		"",
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		selectPriv,
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		insertPriv,
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		updatePriv,
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		deletePriv,
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		createPriv,
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		dropPriv,
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		grantPriv,
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		indexPriv,
	)
	vp = fmt.Sprintf(
		"%s\n     %s\n",
		vp,
		alterPriv,
	)

	return vp
}
