package authservice

import e "notify/internal/entity"

func (s *Service) CreateUserWithDetails(info e.UserCommonInfo, cred e.UserCred) (e.User, error) {
	return s.userRepo.CreateWithDetails(info, cred)
}
