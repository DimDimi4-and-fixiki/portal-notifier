package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT struct {
	UserID    uint     `json:"user_id"`
	UserLogin string   `json:"user_login"`
	UserRole  UserRole `json:"user_role"`
	jwt.RegisteredClaims
}

func JWTFromUser(user User, expireMinutes uint) *JWT {
	minutes := time.Duration(expireMinutes)
	expirationTime := time.Now().Add(minutes * time.Minute)
	data := &JWT{
		UserID:    user.Model.ID,
		UserLogin: user.Cred.Login,
		UserRole:  user.Info.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	return data

}
