package main

import (
	"fmt"
	"time"

	d "github.com/shopspring/decimal"
)

type TransactionType int

const (
	TransactionTypeIncome TransactionType = iota
	TransactionTypeExpense
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
	DateString string          `json:"datestring"`
}

func NewTransaction(amount float32, name string, category string, date time.Time, tt TransactionType) Transaction {
	amountd := d.NewFromFloat32(amount)
	return Transaction{Amount: amountd, Name: name, Category: category, Date: date, Type: tt}
}

func NewExpense(amount float32, name string, category string, date time.Time) Transaction {
	return NewTransaction(amount, name, category, date, TransactionTypeExpense)
}

func NewIncome(amount float32, name string, category string, date time.Time) Transaction {
	return NewTransaction(amount, name, category, date, TransactionTypeExpense)
}

// list.Item interface

func (t Transaction) FilterValue() string {
	return t.Name
}

func (t Transaction) Title() string {
	return fmt.Sprintf("%s – £%s", t.Name, t.Amount.StringFixed(2))
}

func (t Transaction) Description() string {
	return fmt.Sprintf("%s | %s | £%s", t.Date.Format("2006-01-02"), t.Category, t.Amount.StringFixed(2))
}
