package pgerr

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPgErr(t *testing.T) {
	t.Parallel()
	pgErr := &pgconn.PgError{Code: "1", Message: "pg error"}

	tCases := []error{
		pgErr,
		fmt.Errorf("wrapper pg err %w", pgErr),
	}

	for _, tCase := range tCases {
		res := GetPgErr(tCase)
		require.NotNil(t, res)
		require.ErrorAs(t, res, &tCase)
	}
}

func TestGetPgErrNil(t *testing.T) {
	simpleErr := errors.New("some error")

	tCases := []error{
		nil,
		simpleErr,
		fmt.Errorf("wrapped error %w", simpleErr),
	}

	for _, tCase := range tCases {
		res := GetPgErr(tCase)
		require.Nil(t, res)
	}
}
