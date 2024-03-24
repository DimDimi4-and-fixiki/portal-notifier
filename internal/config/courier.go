package config

type CourierConfig struct {
	ApiKey     string `env:"COURIER_API_KEY" env-default:""`
	OwnerEmail string `env:"OWNER_EMAIL" env-default:""`
}
