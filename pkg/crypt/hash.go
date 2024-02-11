package crypt

import "golang.org/x/crypto/bcrypt"

const COST int = 7

func Hash(text string) (string, error) {
	btext := []byte(text)
	result, err := bcrypt.GenerateFromPassword(btext, COST)
	return string(result), err
}

func VerifyPassword(plain string, hashed string) bool {
	bplain, bhashed := []byte(plain), []byte(hashed)
	err := bcrypt.CompareHashAndPassword(bhashed, bplain)
	return err != nil
}
