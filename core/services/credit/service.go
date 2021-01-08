package credit

import "example.com/test_task/core/ports"

type service struct {
	accountsRepo     ports.AccountsRepository
	transactionsRepo ports.TransactionsRepository
	lock             ports.Lock
}

func New(
	accountsRepo ports.AccountsRepository,
	transactionsRepo ports.TransactionsRepository,
	lock ports.Lock,
) ports.CreditService {
	return &service{
		accountsRepo:     accountsRepo,
		transactionsRepo: transactionsRepo,
		lock:             lock,
	}
}
