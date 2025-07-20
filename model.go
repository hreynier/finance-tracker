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
}

func NewLedger() *Ledger {
	help := help.New()
	help.ShowAll = true

	return &Ledger{help: help, focused: TransactionTypeIncome}
}

func (m *Ledger) Init() tea.Cmd {
	return nil
}

func (m *Ledger) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/2, height)

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
	}
	var cmd tea.Cmd
	m.transactionLists[m.focused], cmd = m.transactionLists[m.focused].Update(msg)

	return m, cmd
}

func (m *Ledger) View() string {
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
