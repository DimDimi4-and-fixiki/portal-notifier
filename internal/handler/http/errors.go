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

func InternalError(err error) Err {
	var ErrInternal = errors.New("internal error")
	return Err{
		code: "5000",
		err:  fmt.Errorf("%w: %w", ErrInternal, err),
	}
}
