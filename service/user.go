package service

import (
	dbEntity "test/models"
	httpRequest "test/models/http/request"
)

type userService struct {
	userRepo 		dbEntity.IUserRepository
}

func NewUserService(u dbEntity.IUserRepository) *userService {
	return &userService{
		userRepo: u,
	}
}

func (u *userService) GetUserByID(id int) (dbEntity.Users, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return dbEntity.Users{}, err
	}
	return user, err
}

func (u *userService) Create(data *httpRequest.UserRequest) (dbEntity.Users, error) {
	user, err := u.userRepo.Create(data)
	if err != nil {
		return dbEntity.Users{}, err
	}
	return user, err
}