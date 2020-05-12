package models

import "github.com/taufiqade/gowallet/models/http/request"

// ITransactionService represent transaction service contract
type ITransactionService interface {
	TopUp(email string, payload *request.TransactionRequest) error
	Transfer(obligorID int, beneficiary string, payload *request.TransactionRequest) error
}
