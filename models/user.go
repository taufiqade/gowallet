package models

import (
	"time"

	httpRequest "github.com/taufiqade/gowallet/models/http/request"
)

// Users struct
type Users struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	Name      string    `json:"name" gorm:"column:name"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

// IUserService represent user service contract
type IUserService interface {
	Create(data *httpRequest.UserRequest) (Users, error)
	GetUserByID(id int) (Users, error)
}

// IUserRepository represent user interface contract
type IUserRepository interface {
	Create(data *httpRequest.UserRequest) (Users, error)
	GetUserByID(id int) (Users, error)
	GetUserByEmail(email string) (Users, error)
}
