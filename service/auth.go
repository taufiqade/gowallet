package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	dbEntity "github.com/taufiqade/gowallet/models"
)

// AuthService godoc
type AuthService struct {
	userRepo dbEntity.IUserRepository
}

// NewAuthService initialize new auth service
func NewAuthService(u dbEntity.IUserRepository) *AuthService {
	return &AuthService{
		userRepo: u,
	}
}

// Login godoc
func (u *AuthService) Login(email string, password string) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "test123") //this should be in an env file
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return "user not found", err
	}

	if user.Password != password {
		return "password doesn't match", err
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
