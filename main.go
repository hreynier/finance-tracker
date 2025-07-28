package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func main() {
	// debugFlag := flag.Bool("debug", false, "Enables Debug logging")
	// flag.Parse()
	// debug := *debugFlag

	config, err := GetConfiguration()
	if err != nil {
		log.Fatal("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	if err := setupLogging(); err != nil {
		log.Fatal("Failed to setup logging", "error", err)
		os.Exit(1)
	}

	log.Info("Starting finance tracker", "user", config.UserName)
	l := NewLedger(config)
	p := tea.NewProgram(l, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal("Failed to start TUI!", err)
		os.Exit(1)
	}
	log.Info("Bye Bye! 👋")
}

func setupLogging() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get configuration directory: %w", err)
	}

	logFile := filepath.Join(configDir, "finance-tracker.log")

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	log.SetOutput(file)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	log.SetReportTimestamp(true)

	log.Info("Logging initialized", "log_file", logFile)
	return nil
}
