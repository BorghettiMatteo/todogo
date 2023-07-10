package auth

import (
	"main/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("password")

//ritorna la stringa che rappresenta l'utenza

func GenerateToken(user model.User) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * 5)),
		Issuer:    "todo",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey, err := token.SignedString(SecretKey)
	return signedKey, err
}
func VerifyToken(stringifyToken string) bool {
	// io uso hmac come metodo dir tokenClaim TokenClaim
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(stringifyToken, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
