package config

import (
	"scoreboard/internal/models/config"
	"scoreboard/logger"

	"github.com/spf13/viper"
)

var (
	AppConfig config.Config
	err       error
	Database  config.Database
	Server    config.Server
)

func init() {
	LoadConfig()

	Database = AppConfig.Database
	Server = AppConfig.Server
}

func loadConfig(path string) (config config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadConfig() {
	AppConfig, err = loadConfig(".")
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("config loaded")
}
