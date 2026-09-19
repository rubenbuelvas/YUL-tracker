package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Load the two variables from config.yaml
	viper.BindEnv("YUL_ARRIVALS_URL", "YUL_ARRIVALS_URL")
	viper.BindEnv("YUL_DEPARTURES_URL", "YUL_DEPARTURES_URL")
}
