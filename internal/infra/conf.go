package infra

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUrl  string `mapstructure:"DB_URL"`
	DBName string `mapstructure:"DB_NAME"`

	NatsURL string `mapstructure:"NATS_URL"`

	F1TelServerPort int `mapstructure:"F1_TELEMETRY_SERVER_PORT"`
}

func LoadConfigFromEnv() Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load config from env: %s", err)
	}

	cfg := Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to parse env config: %s", err)
	}

	log.Printf("Loaded config from env: %+v", cfg)

	return cfg
}
