package userrepo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	e "notify/internal/entity"
)

func wrapErr(id uuid.UUID, err error) error {
	return fmt.Errorf("user id=%q %w %w", id, err, ErrUpdateUser)
}

func (r *Repository) UpdateCommonInfoByID(c context.Context, id uuid.UUID, data e.UserCommonInfo) (*e.UserDB, error) {
	user, err := r.GetByID(c, id)
	if err != nil {
		return user, wrapErr(id, err)
	}

	user.Info = data
	res := r.conn.Save(user)

	if res.Error != nil {
		return user, wrapErr(id, err)
	}
	return user, nil
}
