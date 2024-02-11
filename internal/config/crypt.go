package config

type CryptConfig struct {
	JWTExpireMinutes uint   `env:"JWT_EXPIRE_MINUTES" env-default:"10"`
	JWTKey           string `env:"JWT_KEY" env-default:"dima_jula_simon"`
}
