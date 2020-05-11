package helper

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashString godoc
func HashString(password string) string {
	str, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(str)
}

// CompareHash godoc
func CompareHash(plainText string, chipperText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(plainText), []byte(chipperText))
	return err == nil
}

var jwtKey = []byte("miniwallet")

// Claim godoc
type Claim struct {
	UserID int    `json:"user_id"`
	Type   string `json:"type"`
	jwt.StandardClaims
}

// CreateToken godoc
func CreateToken(userID int, Type string, exp int64) (string, error) {
	claim := &Claim{
		UserID: userID,
		Type:   Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(jwtKey)
}

// VerifyToken godoc
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

	_, err = Get(tokenString)
	if err != nil {
		fmt.Println("redis token not found/expired")
	}
	return *claims, status, err
}
