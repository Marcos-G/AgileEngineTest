package domain

type Transaction struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Amount    float64 `json:"amount"`
	Mode      string `json:"mode"`
	AccountId string `json:"account_id"`
}
var APPROVED_STATUS = "approved"
var REJECTED_STATUS = "rejected"