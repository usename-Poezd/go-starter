package config

import (
	"sync"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Server   string `mapstructure:"DB_SERVER"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type Config struct {
	DB        DBConfig
}

var once sync.Once
var config *Config

var configError error

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init() (*Config, error) {
	once.Do(func() {
		if err := parseConfigFile(); err != nil {
			config = nil
			configError = err
			return
		}

		var cfg Config
		if err := unmarshal(&cfg); err != nil {
			config = nil
			configError = err
			return
		}

		config = &cfg
		configError = nil
	})

	return config, configError
}

func parseConfigFile() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return viper.Unmarshal(&cfg.DB)
}
