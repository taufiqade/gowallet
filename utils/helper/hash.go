package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func HashString(payload string) string {
	str, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.MinCost)
	if err == nil {
		log.Println(err)
		return ""
	}
	return string(str)
}

func CompareHash(plainText string, chipperText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(plainText), []byte(chipperText))
	return err == nil
}

var jwtKey = []byte("miniwallet")

type Claim struct {
	UserID	int	`json:"user_id"`
	UserType	string	`json:"user_type"`
	jwt.StandardClaims
}

func CreateToken(userId int, userType string) (string, error) {
	claim := &Claim{
		UserID:         userId,
		UserType:       userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (Claim, int, error) {
	claims := &Claim{}
	var status = 0
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// unauthorized
			fmt.Println("signature invalid")
			status = 1
		}
		// bad request
		status = 2
	}
	if !tkn.Valid {
		// unauthorized
		fmt.Println("token invalid")
		status = 1
	}
	return *claims, status, err
}