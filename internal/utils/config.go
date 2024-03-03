package utils

import (
	"log"
	"todo-list-app/domain"

	"github.com/spf13/viper"
)

func LoadConfig() (*domain.Config, error) {
	viper.AddConfigPath("./..")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("failed to load config.env: ", err)
		return nil, err
	}

	cfg := &domain.Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalln("failed to unmarshal config: ", err)
		return nil, err
	}

	return cfg, nil
}
