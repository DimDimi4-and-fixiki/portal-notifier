package userrepo

import (
	"context"
	"fmt"
	e "notify/internal/entity"
)

func (r *Repository) GetByLogin(_ context.Context, login string) (*e.UserDB, error) {
	var user e.UserDB
	query := r.conn.Model(
		e.UserDB{
			Cred: e.HashedUserCred{Login: login},
		},
	)
	res := r.conn.Model(query).First(&user)
	if res.Error != nil {
		return &user, fmt.Errorf("user login=%q %w %w",
			login, res.Error, ErrGetUser)
	}
	return &user, nil
}
