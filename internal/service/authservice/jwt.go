package authservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"notify/internal/config"
	e "notify/internal/entity"
)

var SigningMethod = jwt.SigningMethodHS256

func (s *Service) JWTFromLogin(ctx context.Context, login string) (*e.JWTString, error) {
	cfg := config.Get()
	user, err := s.userRepo.GetByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	data := e.JWTFromUserDB(*user, cfg.Crypt.JWTExpireMinutes)
	token := jwt.NewWithClaims(SigningMethod, data)
	tokenString, err := token.SignedString([]byte(cfg.Crypt.JWTKey))
	if err != nil {
		return nil, fmt.Errorf("error generating jwt: %v", err)
	}
	expTime, _ := data.GetExpirationTime()
	return &e.JWTString{Token: tokenString, ExpiresAt: expTime.Time}, err
}

func (s *Service) JWTFromUser(_ context.Context, u *e.UserDB) (*e.JWTString, error) {
	cfg := config.Get()
	data := e.JWTFromUserDB(*u, cfg.Crypt.JWTExpireMinutes)
	token := jwt.NewWithClaims(SigningMethod, data)
	tokenString, err := token.SignedString([]byte(cfg.Crypt.JWTKey))
	if err != nil {
		return nil, fmt.Errorf("error generating jwt: %v", err)
	}
	expTime, _ := data.GetExpirationTime()
	return &e.JWTString{Token: tokenString, ExpiresAt: expTime.Time}, err
}

func (s *Service) DecodeJWTToUser(_ context.Context, token string) (*e.UserJWT, error) {
	parsed, err := jwt.ParseWithClaims(token, &e.UserJWT{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := errors.New("unexpected signing method for JWT")
			s.log.Error("JWT decoding error", zap.Error(err))
			return nil, err
		}
		return []byte(config.Get().Crypt.JWTKey), nil
	})

	if err != nil {
		err := fmt.Errorf("error parsing jwt: %w", err)
		return nil, err
	}

	user := parsed.Claims.(*e.UserJWT)
	return user, nil
}
