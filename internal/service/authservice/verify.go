package authservice

import "notify/pkg/crypt"

func (s *Service) VerifyUser(login string, password string) (bool, error) {
	user, err := s.userRepo.GetByLogin(login)
	if err != nil {
		return false, err
	}

	return crypt.VerifyPassword(password, user.Cred.HashedPassword), nil
}
