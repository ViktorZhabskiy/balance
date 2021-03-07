package config

import (
	"github.com/spf13/viper"
)

func Init() error {

	viper.AddConfigPath("internal/config/")
	viper.SetConfigName("config")

	// skip defaults configs

	return viper.ReadInConfig()
}
