package config

type DbConfig struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	Password string `env:"DB_PASSWORD" env-default:"notify"`
	Name     string `env:"DB_NAME" env-default:"notify"`
	User     string `env:"DB_USER" env-default:"notify"`
}
