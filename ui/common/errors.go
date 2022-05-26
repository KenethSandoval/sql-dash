package common

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	width  = 100
	height = 25
)

var (
	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)

	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)
)

func ErrorDialog(error string) {
	doc := strings.Builder{}

	errorMessage := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(error)
	ui := lipgloss.JoinVertical(lipgloss.Center, errorMessage)
	dialog := lipgloss.Place(width, height,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	doc.WriteString(dialog)

	fmt.Println(docStyle.Render(doc.String()))
}
