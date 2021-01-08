package ports

import "example.com/test_task/AgileEngineTest/core/domain"

type AccountsService interface {
	Get(id string) (*domain.Account,error)
	SetBalance(id ,mode string, balance float64) (*domain.Account,error)
}

type DebitService interface {
	Get(id string) (*domain.Transaction,error)
	Authorize(transaction *domain.Transaction) (*domain.Transaction,error)
}

type CreditService interface {
	Get(id string) (*domain.Transaction,error)
	Authorize(transaction *domain.Transaction) (*domain.Transaction,error)
}