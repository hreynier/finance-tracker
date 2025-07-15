package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	d "github.com/shopspring/decimal"
)

type state int

const (
	stateMenu state = iota
	stateAddTransaction
	stateViewTransactions
	stateAddCategory
	stateAddDate
	stateAddDescription
	stateAddAmount
)

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "income"
	TransactionTypeExpense TransactionType = "expense"
)

func (tt TransactionType) String() string {
	switch tt {
	case TransactionTypeIncome:
		return "Income"
	case TransactionTypeExpense:
		return "Expense"
	default:
		return "Unknown"
	}
}

type Transaction struct {
	Amount     d.Decimal       `json:"amount" validate:"required,gt=0"`
	Name       string          `json:"name" validate:"required,min=1,max=255"`
	Category   string          `json:"category" validate:"required"`
	Date       time.Time       `json:"date" validate:"required"`
	Type       TransactionType `json:"type" validate:"required,oneof=income expense"`
	DateString string          `json:"datestring" validate:"required"`
}

func (t Transaction) FilterValue() string { return t.Name }
func (t Transaction) Title() string       { return t.Name }
func (t Transaction) Description() string {
	return fmt.Sprintf("%s | %s | %.2f %s", t.DateString, t.Category, t.Amount, "GBP")
}

type Model struct {
	state        state
	config       *Config
	transactions []Transaction

	categoryList    list.Model
	transactionList list.Model
	textInput       textinput.Model

	currentTransaction Transaction
	categoryIndex      int

	message string
	err     error
}

func InitialiseModel(debug bool) Model {
	config, err := GetConfiguration()
	if err != nil {
		log.Error("Failed to load configuration", err)
	}

	model, logger := getModelAndLogging(*config, debug)
	if logger != nil {
		defer logger.Close()
	}

	return model
}

func createModel(config Config) Model {
	log.Info("Initialising transaction tracker", "user", config.UserName)

	categoryItems := make([]list.Item, len(config.Categories))
	for i, cat := range config.Categories {
		categoryItems[i] = categoryItem{name: cat}
	}

	categoryList := list.New(categoryItems, categoryDelegate{}, 30, 10)
	categoryList.Title = "Select Category"
	categoryList.SetShowStatusBar(false)
	categoryList.SetFilteringEnabled(false)

	// Setup transaction list
	transactionList := list.New([]list.Item{}, transactionDelegate{}, 60, 15)
	transactionList.Title = "Transactions"
	transactionList.SetShowStatusBar(false)
	transactionList.SetFilteringEnabled(false)

	// Setup text input
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30

	log.Debug("Model initialized successfully")

	return Model{
		state:           stateMenu,
		config:          &config,
		transactions:    []Transaction{},
		categoryList:    categoryList,
		transactionList: transactionList,
		textInput:       ti,
	}

}

func getModelAndLogging(config Config, debug bool) (Model, *os.File) {
	var loggerFile *os.File

	if debug {
		var fileErr error
		newLoggerFile, fileErr := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if fileErr == nil {
			log.SetOutput(newLoggerFile)
			log.SetTimeFormat(time.Kitchen)
			log.SetReportCaller(true)
			log.SetLevel(log.DebugLevel)
			log.Debug("Logging debug messages to debug.log")
		} else {
			loggerFile, _ = tea.LogToFile("debug.log", "debug")
			log.Error("Failed to set up logging", fileErr)
		}
	} else {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.FatalLevel)
	}

	return createModel(config), loggerFile
}

// TODO: Create InitialiseModel fn
