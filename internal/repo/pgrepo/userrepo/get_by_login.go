package userrepo

import (
	"context"
	"fmt"
	e "notify/internal/entity"
)

func (r *Repository) GetByLogin(_ context.Context, login string) (*e.UserDB, error) {
	var user = e.UserDB{}
	res := r.conn.Where("login = ?", login).First(&user)
	if res.Error != nil {
		return &user, fmt.Errorf("user login=%v %w %w",
			login, ErrGetUser, res.Error)
	}
	return &user, nil
}
