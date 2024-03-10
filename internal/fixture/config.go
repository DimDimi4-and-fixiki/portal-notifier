package fixture

import "notify/internal/config"

func NewMockCryptConfig() config.CryptConfig {
	return config.CryptConfig{
		APIKeyLength:     32,
		JWTExpireMinutes: 10,
		JWTKey:           "test_salt",
	}
}
