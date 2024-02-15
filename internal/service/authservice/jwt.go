package authservice

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"notify/internal/entity"
)

func (s *Service) GenJWT(ctx context.Context, login string, expireMinutes uint, key string) (string, error) {
	user, err := s.userRepo.GetByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	data := entity.JWTFromUserDB(*user, expireMinutes)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err := token.SignedString(key)
	return tokenString, err

}
