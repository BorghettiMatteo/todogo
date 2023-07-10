package auth

import (
	"fmt"
	"main/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaim struct {
	jwt.RegisteredClaims
	user       string
	expiration string
}

var SecretKey = []byte("password")

//ritorna la stringa che rappresenta l'utenza

func GenerateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       user.Username,
		"expiration": time.Now().Local().Add(time.Minute * time.Duration(1)).Unix(),
	})
	signedKey, err := token.SignedString(SecretKey)
	return signedKey, err
}

func VerifyToken(stringifyToken string) bool {
	// io uso hmac come metodo di sign
	token, err := jwt.Parse(stringifyToken, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return false
	}
	if token.Valid {
		fmt.Printf("\"valid\": %v\n", "valid")
	}
	return true
}
