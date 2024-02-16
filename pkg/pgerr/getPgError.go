package pgerr

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetPgErr(err error) *pgconn.PgError {
	if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
		return pgError
	}
	return nil
}
