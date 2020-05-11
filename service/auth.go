package service

import (
	"strconv"
	"time"

	dbEntity "github.com/taufiqade/gowallet/models"
	helper "github.com/taufiqade/gowallet/utils/helper"
)

// AuthService godoc
type AuthService struct {
	userRepo  dbEntity.IUserRepository
	redisRepo dbEntity.IRedisAuthRepository
}

// NewAuthService initialize new auth service
func NewAuthService(u dbEntity.IUserRepository, r dbEntity.IRedisAuthRepository) *AuthService {
	return &AuthService{
		userRepo:  u,
		redisRepo: r,
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

	exp := time.Now().Add(60 * time.Minute).Unix()
	token, err := helper.CreateToken(int(user.ID), user.Type, exp)
	if err != nil {
		return "", err
	}
	// it should be store to redis
	go a.redisRepo.Set(token, strconv.Itoa(int(user.ID)), exp)
	return token, err
}
