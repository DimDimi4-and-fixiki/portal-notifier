package config

import (
	"github.com/golang-jwt/jwt/v5"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"notify/pkg/logger"
)

type Config struct {
	Server  HttpServerConfig
	Db      DbConfig
	Crypt   CryptConfig
	Courier CourierConfig
}

var JwtSigningMethod = jwt.SigningMethodES256

var cfg Config
var once sync.Once
var log = logger.NewSugar()

func (c *Config) isEmpty() bool {
	return *c == (Config{})
}

func Read() Config {
	var read = func() {
		err := cleanenv.ReadConfig(".env", &cfg)
		if err != nil {
			log.Panicf("Error while reading config: %s", err)
		}
	}
	once.Do(read)
	return cfg
}

func Get() Config {
	if cfg.isEmpty() {
		cfg = Read()
	}

	return cfg
}
