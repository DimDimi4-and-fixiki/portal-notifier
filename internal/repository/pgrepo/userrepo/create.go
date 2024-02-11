package userrepo

import e "notify/internal/entity"

func (r *Repository) Create(user *e.User) (e.User, error) {
	res := r.conn.Create(user)
	return *user, res.Error
}

func (r *Repository) CreateWithDetails(info e.UserCommonInfo, cred e.UserCred) (e.User, error) {
	user := e.User{Cred: cred, Info: info}
	res := r.conn.Create(&user)

	return user, res.Error
}
