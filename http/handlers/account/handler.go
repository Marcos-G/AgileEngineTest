package account

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
	accountService ports.AccountsService
}

func New(service ports.AccountsService) *handler {
	return &handler{service}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

	account, transactions, err := h.accountService.Get("unique_account")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if account == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response := buildAccountResponse(account, transactions)
	w.Write(response)
}

func (h *handler) PutBalance(w http.ResponseWriter, r *http.Request) {
	mode := chi.URLParam(r, "mode")

	data := &putBalanceBody{}
	if err := render.Bind(r, data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	account, err := h.accountService.SetBalance("unique_account", mode, data.Balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if account == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response,err := json.Marshal(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

type putBalanceBody struct {
	Balance float64 `json:"balance"`
}
func (bb *putBalanceBody) Bind(r *http.Request) error {
	if bb.Balance == 0{
		return errors.New("invalid balance")
	}
	return nil
}

type accountResponse struct {
	domain.Account
	Transactions []domain.Transaction  `json:"transactions"`
}
func buildAccountResponse(account *domain.Account,transactions []domain.Transaction)[]byte{
	data,_ := json.Marshal(accountResponse{
		Account:      *account,
		Transactions: transactions,
	})
	return data
}