package common

import (
	"github.com/dgrijalva/jwt-go"
	"mwx563796/ginessential/model"
	"time"
)

var jwtkey = []byte("mwx563796")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string,error) {
	expirationTime := time.Now().Add(7*24*time.Hour)
	claims := &Claims{
		UserId:          user.ID,
		StandardClaims:  jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer: "mwx563796",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err := token.SignedString(jwtkey)

	if err != nil{
		return "", err
	}
	return tokenString,nil
}