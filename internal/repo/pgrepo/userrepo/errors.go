package userrepo

import "errors"

var ErrCreateUser = errors.New("create user error")
var ErrGetUser = errors.New("get user error")
var ErrUpdateUser = errors.New("update user error")
var ErrHashUser = errors.New("hashing user error")
var ErrUserAlreadyExists = errors.New("user with this details already exists")
