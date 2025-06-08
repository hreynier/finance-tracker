package models

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
