package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// UserJWT data to be encoded in jwt token
type UserJWT struct {
	UserLogin string   `json:"user_login"`
	UserRole  UserRole `json:"user_role"`
	jwt.RegisteredClaims
}

type JWTString struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func JWTFromUserDB(user UserDB, expireMinutes uint) *UserJWT {
	minutes := time.Duration(expireMinutes)
	expirationTime := time.Now().Add(minutes * time.Minute)
	data := &UserJWT{
		UserLogin: user.Cred.Login,
		UserRole:  user.Info.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	return data
}
