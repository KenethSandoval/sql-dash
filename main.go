package main

import (
	"adminmsyql/ui"
	"fmt"
	"os"
)

func main() {
	if err := ui.NewProgram().Start(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
