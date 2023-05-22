package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
