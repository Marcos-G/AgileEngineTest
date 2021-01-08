package account

import (
	"example.com/test_task/core/domain"
)

func (s *service) Get(id string) (*domain.Account, error) {

	account, err := s.accountsRepo.Get(id)
	if err != nil {
		return nil,err
	}

	return account, nil
}

func (s *service) SetBalance(id, mode string, balance float64) (*domain.Account, error) {

	account, err := s.accountsRepo.Get(id)
	if err != nil {
		return nil,err
	}

	switch mode {
	case "credit":
		account.CreditBalance = balance
	case "debit":
		account.DebitBalance = balance
	}

	err = s.accountsRepo.Save(account)
	if err != nil {
		return nil,err
	}

	return account, nil
}
