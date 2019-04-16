package config

type Config struct {
	Home         string `env:"HOME"`
	Port         string `env:"PORT" envDefault:"3000"`
	IsProduction bool   `env:"PRODUCTION"`
}
