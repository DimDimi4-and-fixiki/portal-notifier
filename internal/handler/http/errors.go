package http

import (
	"errors"
	"fmt"
)

// Err is basic struct for all http errors
// Each error should have custom code
type Err struct {
	code string
	err  error
}

func (e Err) Error() string {
	return e.err.Error()
}

func ValidationError(err error) Err {
	var ErrValidation = errors.New("validation error")
	return Err{
		code: "4000",
		err:  fmt.Errorf("%w: %w", ErrValidation, err),
	}
}

func AuthError(err error) Err {
	var ErrAuth = errors.New("error while authentication")
	return Err{
		code: "5000",
		err:  fmt.Errorf("%w %w", ErrAuth, err),
	}
}

func UnauthorizedError() Err {
	return Err{
		code: "3000",
		err:  errors.New("unauthorized error, wrong credentials"),
	}
}
