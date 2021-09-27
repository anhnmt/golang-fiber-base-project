package config

import (
	"context"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type DefaultConfig struct {
	// DEFAULT
	Ctx context.Context

	// APP
	AppName    string `mapstructure:"APP_NAME"`
	AppPort    int    `mapstructure:"APP_PORT"`
	AppPrefork bool   `mapstructure:"APP_PREFORK"`
	AppNode    int64  `mapstructure:"APP_NODE"`

	// LOGGER
	LogEnabled bool `mapstructure:"LOG_ENABLED"`

	// DATABASE
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBUrl    string `mapstructure:"DB_URL"`

	// JWT
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

var (
	once          *sync.Once
	defaultConfig *DefaultConfig
)

func Config() *DefaultConfig {
	if defaultConfig == nil {
		once = &sync.Once{}

		once.Do(func() {
			if defaultConfig == nil {
				// SetConfigFile explicitly defines the path, name and extension of the config file.
				// Viper will use this and not check any of the config paths.
				// .env - It will search for the .env file in the current directory
				viper.SetConfigFile(".env")
				viper.AddConfigPath(".")
				viper.AutomaticEnv()

				// Find and read the config file
				if err := viper.ReadInConfig(); err != nil {
					log.Fatalf("Error reading config file, %s", err)
				}

				defaultConfig = new(DefaultConfig)
				if err := viper.Unmarshal(defaultConfig); err != nil {
					log.Fatalf("unable to decode into struct, %v", err)
				}

				defaultConfig.Ctx = context.Background()

				// log.Println(defaultConfig)
			}
		})
	}

	return defaultConfig
}
