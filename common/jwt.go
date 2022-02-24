package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Myclaims struct {
	Username string `json:"username"`
	UserId   int    `json:"user_id"`
	jwt.StandardClaims
}

var Secret = []byte("secret")

const TokenExpireDuration = time.Hour * 24

// GenToken 生成Token
func GenToken(username string, userId int) (string, error) {
	c := Myclaims{
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "gin-demo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Myclaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Myclaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
