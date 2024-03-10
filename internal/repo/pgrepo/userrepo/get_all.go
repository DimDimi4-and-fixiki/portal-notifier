package userrepo

import (
	"context"
	"fmt"
	e "notify/internal/entity"
)

func (r *Repository) GetAll(_ context.Context) (*[]e.UserDB, error) {
	var users []e.UserDB
	res := r.conn.Find(&users)

	if res.Error != nil {
		return &users,
			fmt.Errorf("get all users error %w %w", ErrGetUser, res.Error)
	}
	return &users, nil
}
