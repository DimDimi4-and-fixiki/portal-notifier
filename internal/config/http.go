package config

type HttpServerConfig struct {
	Host               string `env:"SERVER_HOST" env-default:"localhost"`
	Port               string `env:"SERVER_PORT" env-default:"8080"`
	LogTimeFormat      string `env:"SERVER_LOG_TIME_FORMAT" env-default:"RFC3339"`
	GracefulTimeoutSec int    `env:"SERVER_GRACEFUL_TIMEOUT_SEC" env-default:"5"`
}
