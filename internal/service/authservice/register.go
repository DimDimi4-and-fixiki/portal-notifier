package authservice

import (
	"context"
	e "notify/internal/entity"
)

func (s *Service) Register(_ context.Context, info e.UserCommonInfo, cred e.UserCred) (e.User, error) {
	user := e.User{Info: info, Cred: cred}
	res, err := s.userRepo.Create(&user)
	return res, err
}
