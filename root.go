package main

import (
	"flag"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Enables Debug logging")
	flag.Parse()
	debug := *debugFlag

	p := tea.NewProgram(InitialiseModel(debug))

	if _, err := p.Run(); err != nil {
		log.Fatal("Failed to start TUI!", err)
		os.Exit(1)
	}
	log.Info("Bye Bye! 👋")
}
