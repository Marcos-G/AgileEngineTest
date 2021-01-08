package dependencies

import (
	"example.com/test_task/adapters/lock"
	"example.com/test_task/adapters/repositories/accountsrepo"
	"example.com/test_task/adapters/repositories/transactionsrepo"
	"example.com/test_task/core/domain"
	"example.com/test_task/core/ports"
	"example.com/test_task/core/services/account"
	"example.com/test_task/core/services/transactions"
	"example.com/test_task/pkg"
)

var (
	AccountService ports.AccountsService
	TransactionsService  ports.TransactionsService
)

func Bind() {
	if true {
		BindUniqueEnv()
	}
}

func BindUniqueEnv() {
	accountsRepo := accountsrepo.New(&pkg.AccountsClient{})
	accountsRepo.Save(&domain.Account{ID: "unique_account",DebitBalance: 100,CreditBalance: 500})
	transactionsRepo := transactionsrepo.New(&pkg.TransactionsClient{})
	lock := lock.New()

	AccountService = account.New(accountsRepo,transactionsRepo,lock)
	TransactionsService = transactions.New(accountsRepo, transactionsRepo,lock)
}
