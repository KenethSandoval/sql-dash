package main

import (
	"fmt"
	"os"

	"adminmsyql/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if err := tea.NewProgram(ui.NewModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
