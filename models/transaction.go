package models

import "github.com/taufiqade/gowallet/models/http/request"

// transaction service contract
type ITransactionService interface {
	TopUp(beneficiaryID int, payload *request.TopUpRequest) error
	//Transfer(obligorID int, beneficiaryID int, ip string, agent string, payload string) error
}