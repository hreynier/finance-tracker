package models

import (
	"fmt"
	"time"

	d "github.com/shopspring/decimal"
)

type Transaction struct {
	ID          int64           `json:"id" validate:"required"`
	Amount      d.Decimal       `json:"amount" validate:"required,gt=0"`
	Description string          `json:"description" validate:"required,min=1,max=255"`
	Category    string          `json:"category" validate:"required"`
	Date        time.Time       `json:"date" validate:"required"`
	Type        TransactionType `json:"type" validate:"required,oneof=income expense"`
}

func (t Transaction) String() string {
	return fmt.Sprintf("%s: £%s - %s (%s)", t.Date.Format("02/01/2006"), t.Amount.StringFixed(2), t.Description, t.Category)
}

func NewTransaction(amount string, description string, category string, txnType TransactionType) (*Transaction, error) {
	amt, err := d.NewFromString(amount)
	if err != nil {
		return nil, fmt.Errorf("Invalid amount: %w", err)
	}

	return &Transaction{
		ID:          0,
		Amount:      amt,
		Description: description,
		Category:    category,
		Date:        time.Now(),
		Type:        txnType,
	}, nil
}

func NewExpense(amount, description, category string) (*Transaction, error) {
	return NewTransaction(amount, description, category, TransactionTypeExpense)
}

func NewIncome(amount, description, category string) (*Transaction, error) {
	return NewTransaction(amount, description, category, TransactionTypeIncome)
}
