package crypt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
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
	for i, tCase := range tCases {
		pass := tCase
		testName := fmt.Sprintf("case %s", strconv.Itoa(i))
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			hashedPass, err := Hash(pass)
			require.NoError(t, err)
			verifyRes := VerifyPassword(pass, hashedPass)
			assert.True(t, verifyRes)
		})
	}
}
