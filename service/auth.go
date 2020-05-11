package service

import (
	dbEntity "github.com/taufiqade/gowallet/models"
	helper "github.com/taufiqade/gowallet/utils/helper"
)

// AuthService godoc
type AuthService struct {
	userRepo dbEntity.IUserRepository
	// redisRepo dbEntity.IRedisAuthRepository
}

// NewAuthService initialize new auth service
func NewAuthService(u dbEntity.IUserRepository) *AuthService {
	return &AuthService{
		userRepo: u,
		// redisRepo: r,
	}
}

// CreateToken godoc
func (a *AuthService) CreateToken(email string, password string) (string, error) {
	user, err := a.userRepo.GetUserByEmail(email)
	if err != nil {
		return "user not found", err
	}
	check := helper.CompareHash(user.Password, password)
	if check != true {
		return "password doesn't match", err
	}

	// it should be store to redis

	token, err := helper.CreateToken(int(user.ID), user.Type)
	return token, err
}
