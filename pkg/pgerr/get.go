package pgerr

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetPgErr(err error) *pgconn.PgError {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) {
		return pgError
	}
	return nil
}
