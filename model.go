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

func (l *Ledger) Next() {
	if l.focused == TransactionTypeExpense {
		l.focused = TransactionTypeIncome
	} else {
		l.focused++
	}
}

func (l *Ledger) Prev() {
	if l.focused == TransactionTypeIncome {
		l.focused = TransactionTypeExpense
	} else {
		l.focused--
	}
}

func (m *Ledger) Init() tea.Cmd {
	return nil
}

func (m *Ledger) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height/2)
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
			focusedStyle.Width(msg.Width / divisor)
			unfocusedStyle.Width(msg.Width / divisor)
			focusedStyle.Height(msg.Height - divisor)
			unfocusedStyle.Height(msg.Height - divisor)
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit
		case "right", "l", "tab":
			m.Next()
		case "left", "h", "shift+tab":
			m.Prev()
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
		incomeView := m.transactionLists[TransactionTypeIncome].View()
		expenseView := m.transactionLists[TransactionTypeExpense].View()

		switch m.focused {
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(incomeView),
				unfocusedStyle.Render(expenseView),
			)
		case TransactionTypeExpense:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				unfocusedStyle.Render(incomeView),
				focusedStyle.Render(expenseView),
			)
		}
	} else {
		return "Loading..."
	}
}
