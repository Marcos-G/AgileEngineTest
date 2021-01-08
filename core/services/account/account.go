package account

import (
	"example.com/test_task/core/domain"
)

func (s *service) Get(id string) (*domain.Account, []domain.Transaction, error) {

	account, err := s.accountsRepo.Get(id)
	if err != nil {
		return nil, nil, err
	}
	transactions, err := s.transactionsRepo.GetByAccountID(id)
	if err != nil {
		return account, nil, err
	}

	return account, transactions,nil
}

func (s *service) SetBalance(id, mode string, balance float64) (*domain.Account, error) {

	account, err := s.accountsRepo.Get(id)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, nil
	}

	switch mode {
	case "transactions":
		account.CreditBalance = balance
	case "debit":
		account.DebitBalance = balance
	}

	err = s.accountsRepo.Save(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
