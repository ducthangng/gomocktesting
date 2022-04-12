package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Type         string `mapstructure:"DATABASE_TYPE"`
	User         string `mapstructure:"DATABASE_USER"`
	Password     string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseHost string `mapstructure:"DATABASE_HOST"`
	DatabasePort string `mapstructure:"DATABASE_PORT"`
	Name         string `mapstructure:"DATABASE_NAME"`
	SSLMode      string `mapstructure:"DATABASE_SSL_MODE"`
	CACERTBASE64 string `mapstructure:"DATABASE_CACERTBASE64"`
}

func ReadConfig(path string) (config Config, err error) {
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
