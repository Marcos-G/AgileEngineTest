package ports

import "example.com/test_task/core/domain"

type AccountsRepository interface {
	Get(id string) (*domain.Account,error)
	Save(*domain.Account) error
}

type TransactionsRepository interface {
	Get(id string) (*domain.Transaction,error)
	GetByAccountID(accountId string) ([]domain.Transaction,error)
	Save(transaction *domain.Transaction) error
}