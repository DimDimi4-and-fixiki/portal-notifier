package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	UserLogin string   `json:"user_login"`
	UserRole  UserRole `json:"user_role"`
	jwt.RegisteredClaims
}

func JWTFromUserDB(user UserDB, expireMinutes uint) *JWT {
	minutes := time.Duration(expireMinutes)
	expirationTime := time.Now().Add(minutes * time.Minute)
	data := &JWT{
		UserLogin: user.Cred.Login,
		UserRole:  user.Info.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	return data
}
