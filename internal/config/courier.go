package config

type CourierConfig struct {
	ApiKey     string `env:"COURIER_API_KEY" env-default:""`
	OwnerEmail string `env:"COURIER_OWNER_EMAIL" env-default:"dimdimi4andfixiki@gmail.com"`
}
