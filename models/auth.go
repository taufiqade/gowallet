package models

// transaction service contract
type IAuthService interface {
	Login(email string, password string) (string, error)
}