package userrepo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func (r *Repository) GetByLogin(_ context.Context, login string) (*e.UserDB, error) {
	var user = e.UserDB{}
	res := r.conn.Where("login = ?", login).First(&user)
	if res.Error != nil {
		return &user, fmt.Errorf("user login=%v %w %w",
			login, ErrGetUser, res.Error)
	}
	return &user, nil
}

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

func (r *Repository) GetByEmail(_ context.Context, email string) (*e.UserDB, error) {
	var user = e.UserDB{}
	res := r.conn.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return &user, fmt.Errorf("user email=%v %w %w",
			email, ErrGetUser, res.Error)
	}
	return &user, nil
}
