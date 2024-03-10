package pgrepo

type PgErrorCode string

func (val PgErrorCode) Str() string {
	return string(val)
}

const (
	UNIQUE_CONSTRAINT_VIOLATED PgErrorCode = "23505"
)
