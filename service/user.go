package service

import (
	dbEntity "github.com/taufiqade/gowallet/models"
	httpRequest "github.com/taufiqade/gowallet/models/http/request"
)

// UserService struct
type UserService struct {
	userRepo dbEntity.IUserRepository
}

// NewUserService initialize new user service
func NewUserService(u dbEntity.IUserRepository) *UserService {
	return &UserService{
		userRepo: u,
	}
}

// GetUserByID godoc
func (u *UserService) GetUserByID(id int) (dbEntity.Users, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return dbEntity.Users{}, err
	}
	return user, err
}

// Create godoc
func (u *UserService) Create(data *httpRequest.UserRequest) (dbEntity.Users, error) {
	user, err := u.userRepo.Create(data)
	if err != nil {
		return dbEntity.Users{}, err
	}
	return user, err
}
