package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("AllYourBase")

var (
	Token        string
	RefreshToken string
)

type Claims struct {
	Username string `json:"username"`
	UserCode string `json:"user_code"`
	jwt.StandardClaims
}

func GenerateToken(username, user_code string) (string, error) {
	expireAt := time.Now().Add(time.Hour).Unix()
	claims := Claims{
		username,
		user_code,
		jwt.StandardClaims{
			//NotBefore: expireAt - 60,
			ExpiresAt: expireAt + 60*60*2,
			Issuer:    "xxx-xxx",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
