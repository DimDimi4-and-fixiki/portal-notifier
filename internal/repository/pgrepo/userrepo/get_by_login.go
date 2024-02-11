package userrepo

import (
	e "notify/internal/entity"
)

func (r *Repository) GetByLogin(login string) (e.User, error) {
	var user e.User
	query := r.conn.Model(e.User{Cred: e.UserCred{Login: login}})
	res := r.conn.Model(query).First(&user)
	return user, res.Error
}
