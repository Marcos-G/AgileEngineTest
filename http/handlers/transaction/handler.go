package transaction

import (
	"encoding/json"
	"errors"
	"example.com/test_task/core/domain"
	"example.com/test_task/core/ports"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type handler struct {
	transactionService ports.TransactionsService
}

func New(service ports.TransactionsService) *handler {
	return &handler{service}
}

func (h *handler)Post(w http.ResponseWriter, r *http.Request) {
	data := &postTransactionBody{}
	if err := render.Bind(r, data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transaction, err := h.transactionService.Authorize(data.Transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response,err := json.Marshal(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)

}

func (h *handler)Get(w http.ResponseWriter, r *http.Request) {
	transactionId := chi.URLParam(r, "id")

	transaction, err := h.transactionService.Get(transactionId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if transaction == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response,err := json.Marshal(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)

}

type postTransactionBody struct {
	*domain.Transaction
}
func (tb *postTransactionBody) Bind(r *http.Request) error {
	if tb.Transaction == nil{
		return errors.New("Invalid Transaction")
	}
	tb.AccountId = "unique_account"
	return nil
}