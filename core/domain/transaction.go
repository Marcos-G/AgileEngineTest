package domain

type Transaction struct {
	ID        string
	Status    string
	Amount    float64
	Currency  string
	Mode      string
	AccountId string
}
var APPROVED_STATUS = "approved"
var REJECTED_STATUS = "rejected"