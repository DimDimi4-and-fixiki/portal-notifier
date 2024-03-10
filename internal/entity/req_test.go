package entity

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"notify/internal/fixture"
	"testing"
)

func TestCreateServiceReq_ToUser(t *testing.T) {
	t.Parallel()
	cfg := fixture.NewMockCryptConfig()

	t.Run("success", func(t *testing.T) {
		// arrange
		req := CreateServiceReq{
			ServiceName: "test service",
			Login:       "test login",
		}
		expected := &User{
			Info: UserCommonInfo{
				Role: Service,
				Name: "test service",
			},
			Cred: UserCred{
				Login:    "test login",
				Password: mock.Anything,
			},
		}

		// act
		actual, err := req.ToUser(cfg)

		// assert
		require.Equal(t, len(actual.Cred.Password), cfg.APIKeyLength)
		require.NoError(t, err)
		actual.Cred.Password = mock.Anything
		require.Equal(t, expected, actual)
	})
}
