package models

import (
	httpRequest "test/models/http/request"
	"time"
)

// struct is a type which contains named fields
type Users struct {
	ID			int64 		`json:"id" gorm:"primary_key"`
	Email		string 		`json:"email" gorm:"column:email"`
	Password 	string 		`json:"password" gorm:"column:password"`
	Name 		string 		`json:"name" gorm:"column:name"`
	UpdatedAt 	time.Time 	`json:"updated_at" gorm:"column:updated_at"`
	CreatedAt 	time.Time 	`json:"created_at" gorm:"column:created_at"`
}

// UserService represent the users service contract
type IUserService interface {
	Create(data *httpRequest.UserRequest) (Users, error)
	GetUserByID(id int) (Users, error)
}

// UserRepository represent the users repository contract
type IUserRepository interface {
	Create(data *httpRequest.UserRequest) (Users, error)
	GetUserByID(id int) (Users, error)
	GetUserByEmail(email string) (Users, error)
}