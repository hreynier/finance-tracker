package storage

import (
	"github.com/hreynier/finance-tracker/internal/models"
)

type StorageInterface interface {
	AddTransaction(transaction *models.Transaction) error
	GetTransaction(id string) (*models.Transaction, error)
	GetAllTransactions() ([]*models.Transaction, error)
	DeleteTransaction(id string) error
	UpdateTransaction(id string, transaction *models.Transaction) error
}
