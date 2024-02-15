package userrepo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	e "notify/internal/entity"
)

func (r *Repository) GetByID(_ context.Context, id uuid.UUID) (*e.UserDB, error) {
	var user = &e.UserDB{}
	var err error

	query := r.conn.Model(&e.UserDB{ID: id})
	res := query.First(user)

	if res.Error != nil {
		err = fmt.Errorf("user with id=%q: %w; %w", id, res.Error, ErrGetUser)
	}
	return user, err

}
