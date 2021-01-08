package debit

import (
	"errors"
	"example.com/test_task/core/domain"
)

func (s *service) Get(id string) (*domain.Transaction, error) {
	transaction,err:=s.transactionsRepo.Get(id)
	return transaction,err
}

func (s *service) Authorize(transaction *domain.Transaction) (*domain.Transaction, error) {
	lockID := transaction.AccountId+transaction.Mode
	err := s.lock.Acquire(lockID)
	if err != nil {
		return nil,err
	}
	defer s.lock.Release(lockID)

	account,err := s.accountsRepo.Get(transaction.AccountId)
	if err != nil{
		return nil,err
	}

	if account == nil {
		return nil, errors.New("invalid account")
	}

	if account.CreditBalance < transaction.Amount{
		transaction.Status = domain.REJECTED_STATUS

	}else{
		account.CreditBalance -= transaction.Amount
		transaction.Status = domain.APPROVED_STATUS
		err = s.accountsRepo.Save(account)
		if err!= nil{
			return nil,errors.New("internal error saving account")
		}
	}

	err = s.transactionsRepo.Save(transaction)
	if err!= nil{
		return nil,errors.New("internal error saving transaction")
	}
	return transaction,nil
}
