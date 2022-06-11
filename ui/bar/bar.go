package bar

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

const height = 1

// ColorConfig
type ColorConfig struct {
	Foreground lipgloss.AdaptiveColor
	Background lipgloss.AdaptiveColor
}

// Bubble represents the properties of the statusbar.
type Bubble struct {
	Width             int
	Height            int
	FirstColumn       string
	SecondColumn      string
	ThirdColumn       string
	FourthColumn      string
	FirstColumnColor  ColorConfig
	SecondColumnColor ColorConfig
	ThirdColumnColor  ColorConfig
	FourthColumnColor ColorConfig
}

// New create a new instance of the statusbar
func New(firstColumnColor, secondColumnColor, thirdColumnColor, fourthColumnColor ColorConfig) Bubble {
	return Bubble{
		FirstColumnColor:  firstColumnColor,
		SecondColumnColor: secondColumnColor,
		ThirdColumnColor:  thirdColumnColor,
		FourthColumnColor: fourthColumnColor,
	}
}

// SetSize set the width of statusbar
func (b *Bubble) SetSize(width int) {
	b.Width = width
}

// Update update the size for statusbar
func (b Bubble) Update(msg tea.Msg) (Bubble, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.SetSize(msg.Width)
	}
	return b, nil
}

// SetContent set the content for statusbar
func (b *Bubble) SetContent(firstColumn, secondColumn, thridColumn, fourthColum string) {
	b.FirstColumn = firstColumn
	b.SecondColumn = secondColumn
	b.ThirdColumn = thridColumn
	b.FourthColumn = fourthColum
}

// SetColor set the color for statusbar the 4 colums
func (b *Bubble) SetColor(firstColumnColor, secondColumnColor, thridColumnColor, fourthColumColor ColorConfig) {
	b.FirstColumnColor = firstColumnColor
	b.SecondColumnColor = secondColumnColor
	b.ThirdColumnColor = thridColumnColor
	b.FourthColumnColor = fourthColumColor
}

func (b Bubble) View() string {
	width := lipgloss.Width
	firstColumn := lipgloss.NewStyle().
		Foreground(b.FirstColumnColor.Foreground).
		Background(b.FirstColumnColor.Background).
		Padding(0, 1).
		Height(height).
		Render(truncate.StringWithTail(b.FirstColumn, 30, "..."))

	thirdColumn := lipgloss.NewStyle().
		Foreground(b.ThirdColumnColor.Foreground).
		Background(b.ThirdColumnColor.Background).
		Align(lipgloss.Right).
		Padding(0, 1).
		Height(height).
		Render(b.ThirdColumn)

	fourthColumn := lipgloss.NewStyle().
		Foreground(b.FourthColumnColor.Foreground).
		Background(b.FourthColumnColor.Background).
		Padding(0, 1).
		Height(height).
		Render(b.FourthColumn)

	secondColumn := lipgloss.NewStyle().
		Foreground(b.SecondColumnColor.Foreground).
		Background(b.SecondColumnColor.Background).
		Padding(0, 1).
		Height(height).
		Width(b.Width - width(firstColumn) - width(thirdColumn) - width(fourthColumn)).
		Render(truncate.StringWithTail(
			b.SecondColumn,
			uint(b.Width-width(firstColumn)-width(thirdColumn)-width(fourthColumn)-3),
			"..."),
		)

	return lipgloss.JoinHorizontal(lipgloss.Top, firstColumn, secondColumn, thirdColumn, fourthColumn)
}
