package config

type CryptConfig struct {
	JWTExpireMinutes uint   `env:"JWT_EXPIRE_MINUTES" env-default:"30"`
	JWTKey           string `env:"JWT_KEY" env-default:"dima_jula_simon"`
	APIKeyLength     int    `env:"API_KEY_LENGTH" env-default:"32"`
}
