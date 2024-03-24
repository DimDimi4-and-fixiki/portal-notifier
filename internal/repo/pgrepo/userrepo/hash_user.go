package userrepo

import (
	"fmt"
	e "notify/internal/entity"
	"notify/pkg/crypt"
)

// hashUser hashes user sensitive fields e.g. password
func hashUser(user e.User) (e.UserDB, error) {
	hashedPassword, err := crypt.Hash(user.Cred.Password)
	if err != nil {
		err = fmt.Errorf("user_login = %q %w %w",
			user.Cred.Login, err, ErrHashUser)
	}
	return e.UserDB{
		Info: user.Info,
		Cred: e.HashedUserCred{
			Login:          user.Cred.Login,
			HashedPassword: hashedPassword,
		},
	}, err

}
