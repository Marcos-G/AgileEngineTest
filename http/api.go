package http

import (
	"example.com/test_task/dependencies"
	"example.com/test_task/http/handlers/account"
	"example.com/test_task/http/handlers/transaction"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func Start() {
	err := http.ListenAndServe(":8080", Router())
	if err != nil {
		fmt.Println("Error in creating server ", err)
	}
}

func Router() *chi.Mux {
	r := chi.NewRouter()

	accountHandler := account.New(dependencies.AccountService)
	transactionHandler := transaction.New(dependencies.TransactionsService)

	r.Get("/account", accountHandler.Get)
	r.Put("/account/{mode}/balance", accountHandler.PutBalance)

	r.Post("/transaction", transactionHandler.Post)
	r.Get("/transaction/{id}", transactionHandler.Get)

	return r
}
