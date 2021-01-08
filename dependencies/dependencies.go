package dependencies

import (
	"example.com/test_task/AgileEngineTest/adapters/repositories/accountsrepo"
	"example.com/test_task/AgileEngineTest/adapters/repositories/transactionsrepo"
	"example.com/test_task/AgileEngineTest/core/ports"
	"example.com/test_task/AgileEngineTest/core/services/account"
	"example.com/test_task/AgileEngineTest/core/services/credit"
	"example.com/test_task/AgileEngineTest/core/services/debit"
	"example.com/test_task/AgileEngineTest/pkg"
)

var (
	AccountService ports.AccountsService
	DebitService   ports.DebitService
	CreditService  ports.CreditService
)

func Bind(){
	if true{
		BindUniqueEnv()
	}
}


func BindUniqueEnv() {
	accountsRepo := accountsrepo.New(&pkg.AccountsClient{})
	transactionsRepo := transactionsrepo.New(&pkg.TransactionsClient{})

	AccountService = account.New(accountsRepo)
	CreditService = credit.New(accountsRepo,transactionsRepo)
	DebitService = debit.New(accountsRepo,transactionsRepo)
}
