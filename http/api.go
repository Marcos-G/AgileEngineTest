package http

import (
	"example.com/test_task/AgileEngineTest/http/handlers/account"
	"example.com/test_task/AgileEngineTest/http/handlers/transaction"
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

	r.Get("/account/{id}", account.Get)

	r.Post("/transaction/debit", transaction.Post)
	r.Get("/transaction/debit/{id}", transaction.Get)

	r.Post("/transaction/credit", transaction.Post)
	r.Get("/transaction/credit/{id}", transaction.Get)

	return r
}
