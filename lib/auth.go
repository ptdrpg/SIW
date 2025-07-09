package lib

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccesClaims struct {
	UserId string `json:"Id"`
	jwt.StandardClaims
}

var jwtKey = []byte("chucknoristokenkey")

func GenerateToken(id string) (string, error) {
	validity := time.Now().Add(24 * time.Hour)
	claims := &AccesClaims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: validity.Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifToken(tokenString string) (bool, error) {
	tokenres, err := jwt.Parse(tokenString, func (token *jwt.Token)(any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("mauvaise methode de signature: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return false, err
	}
	if !tokenres.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}
