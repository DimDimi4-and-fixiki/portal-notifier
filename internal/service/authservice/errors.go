package authservice

import "errors"

var ErrEmptyUserInput = errors.New("user input should consist of user_id or user_login")
