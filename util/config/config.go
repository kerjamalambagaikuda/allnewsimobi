package config

import (
	"time"

	"github.com/kerjamalambagaikuda/allnewsimobi/util/logger"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	LOG_MAXSIZE         int           `mapstructure:"LOG_MAXSIZE"`
	LOG_MAXBACKUP       int           `mapstructure:"LOG_MAXBACKUP"`
	LOG_MAXAGE          int           `mapstructure:"LOG_MAXAGE"`
	LOG_COMPRESS        bool          `mapstructure:"LOG_COMPRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadLogger(path string, config Config) error {
	return logger.LoadLogger(path, logger.LogConfig{
		LOG_MAXSIZE:   config.LOG_MAXSIZE,
		LOG_MAXBACKUP: config.LOG_MAXBACKUP,
		LOG_MAXAGE:    config.LOG_MAXAGE,
		LOG_COMPRESS:  config.LOG_COMPRESS,
	})
}
