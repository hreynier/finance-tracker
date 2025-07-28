package main

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Ledger struct {
	help             help.Model
	loaded           bool
	focused          TransactionType
	transactionLists []list.Model
	quitting         bool
	cfg              *Config
}

func NewLedger(cfg *Config) *Ledger {
	help := help.New()
	help.ShowAll = true

	return &Ledger{help: help, focused: TransactionTypeIncome, cfg: cfg}
}

func (m *Ledger) Init() tea.Cmd {
	return nil
}

func (m *Ledger) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/2, height)
	defaultList.SetShowHelp(false)

	m.transactionLists = []list.Model{defaultList, defaultList}
	// Init Income
	m.transactionLists[TransactionTypeIncome].Title = "Income"
	m.transactionLists[TransactionTypeIncome].SetItems([]list.Item{

		NewIncome(350.46, "Salary", "Salary", time.Now()),
		NewIncome(250, "Side Hustle", "Side Hustle", time.Now()),
	})
	// Init Expense
	m.transactionLists[TransactionTypeExpense].Title = "Expense"
	m.transactionLists[TransactionTypeExpense].SetItems([]list.Item{
		NewExpense(3.75, "Starbucks", "Coffee", time.Now()),
		NewExpense(3.75, "Cafe Nero", "Coffee", time.Now()),
		NewExpense(3.75, "Costa", "Coffee", time.Now()),
		NewExpense(3.75, "Blank Street", "Coffee", time.Now()),
	})
}

func (m *Ledger) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.transactionLists[m.focused], cmd = m.transactionLists[m.focused].Update(msg)

	return m, cmd
}

func (m *Ledger) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.transactionLists[TransactionTypeIncome].View(),
			m.transactionLists[TransactionTypeExpense].View(),
		)
	} else {
		return "Loading..."
	}
}
