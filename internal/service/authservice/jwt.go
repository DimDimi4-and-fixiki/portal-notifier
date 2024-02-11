package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"notify/internal/entity"
)

func (s *Service) GenJWT(login string, expireMinutes uint, key string) (string, error) {
	user, err := s.userRepo.GetByLogin(login)
	if err != nil {
		return "", err
	}

	data := entity.JWTFromUser(user, expireMinutes)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err := token.SignedString(key)
	return tokenString, err

}
