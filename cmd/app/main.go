package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hreynier/finance-tracker/internal/config"
	"github.com/hreynier/finance-tracker/internal/models"
)

type model struct {
	input string
	done  bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.input != "" {
				m.done = true
				return m, tea.Quit
			}
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			m.input += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.done {
		return fmt.Sprintf("Hello %s!\n\nPress any key to exit...", m.input)
	}

	return fmt.Sprintf("What's your name?\n\n%s\n\n(Press Enter to continue or Esc to quit)", m.input)
}

func main() {
	txn := models.Transaction{}
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	fmt.Println(cfg.UserName)
	fmt.Println("Transactions: %w", txn)

	transaction, err := models.NewExpense("25.567", "Coffee At Starbucks", "Food & Drink")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println(transaction)
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
