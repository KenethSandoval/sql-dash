package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KenethSandoval/tuidb/dash"
	"github.com/KenethSandoval/tuidb/ui"
	"github.com/KenethSandoval/tuidb/ui/uictx"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var provider = "mysql"

	flag.StringVar(&provider, "c", "mysql", "db connect for inspect")
	flag.Parse()

	if provider == "" {
		flag.Usage()
		os.Exit(1)
	}

	client, err := dash.New(&provider)
	if err != nil {
		panic(err)
	}

	ctx := uictx.New(&client)

	if err := tea.NewProgram(ui.NewModel(&ctx), tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
