package ports

import "example.com/test_task/core/domain"

type AccountsService interface {
	Get(id string) (*domain.Account,[]domain.Transaction,error)
	SetBalance(id ,mode string, balance float64) (*domain.Account,error)
}

type TransactionsService interface {
	Get(id string) (*domain.Transaction,error)
	Authorize(transaction *domain.Transaction) (*domain.Transaction,error)
}