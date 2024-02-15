package userrepo

import (
	"context"
	"fmt"
	e "notify/internal/entity"
)

func (r *Repository) Create(_ context.Context, user *e.User) (*e.UserDB, error) {
	var hashedUser e.UserDB
	hashedUser, err := hashUser(*user)
	if err != nil {
		return nil, err
	}

	res := r.conn.Create(&hashedUser)
	if res.Error != nil {
		return nil, fmt.Errorf("user login=%s %w %w",
			user.Cred.Login, res.Error, ErrCreateUser)
	}
	return &hashedUser, nil
}
