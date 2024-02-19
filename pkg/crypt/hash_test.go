package crypt

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHash(t *testing.T) {
	t.Parallel()
	tCases := []string{
		"some_pass!$",
		"Some_pass___!!!",
		"____!!!___$$@@@**^^",
		";_:*&^#!()-=+",
	}
	for _, tCase := range tCases {
		hashedPass, err := Hash(tCase)
		require.NoError(t, err)
		verifyRes := VerifyPassword(tCase, hashedPass)
		assert.True(t, verifyRes)

	}
}
