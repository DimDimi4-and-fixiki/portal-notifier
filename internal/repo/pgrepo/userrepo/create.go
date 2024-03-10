package userrepo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	e "notify/internal/entity"
	"notify/internal/repo/pgrepo"
	"notify/pkg/logger"
	"notify/pkg/pgerr"
)

func (r *Repository) Create(_ context.Context, user *e.User) (*e.UserDB, error) {
	var hashedUser e.UserDB
	hashedUser, err := hashUser(*user)
	if err != nil {
		return nil, err
	}

	res := r.conn.Create(&hashedUser)
	if res.Error != nil {
		pgErr := pgerr.GetPgErr(res.Error)
		if pgErr != nil {
			switch pgErr.Code {
			case pgrepo.UNIQUE_CONSTRAINT_VIOLATED.Str():
				return nil, fmt.Errorf("user login=%s %w",
					user.Cred.Login, ErrUserAlreadyExists)
			}
		}
		logger.NewSugar().Infof("Error Create %s", res.Error.(*pgconn.PgError).Code)
		return nil, fmt.Errorf("user login=%s %w %w",
			user.Cred.Login, res.Error, ErrCreateUser)
	}
	return &hashedUser, nil
}
