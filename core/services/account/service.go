package account

import "example.com/test_task/core/ports"

type service struct {
	accountsRepo ports.AccountsRepository
}

func New(accountsRepository ports.AccountsRepository) ports.AccountsService {
	return &service{accountsRepo: accountsRepository}
}
