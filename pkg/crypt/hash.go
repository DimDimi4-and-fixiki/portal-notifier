package crypt

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const COST int = 7

var ErrHash = errors.New("hashing error")
var ErrVerifyHash = errors.New("verifying hash error")

func Hash(text string) (string, error) {
	btext := []byte(text)
	result, err := bcrypt.GenerateFromPassword(btext, COST)
	if err != nil {
		err = fmt.Errorf("%w, %w", ErrHash, err)
	}
	return string(result), err
}

func VerifyPassword(plain string, hashed string) bool {
	bplain, bhashed := []byte(plain), []byte(hashed)
	err := bcrypt.CompareHashAndPassword(bhashed, bplain)
	return err == nil
}
