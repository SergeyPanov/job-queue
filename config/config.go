package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbPassword       string `mapstructure:"POSTGRES_PASSWORD"`
	DbUser           string `mapstructure:"POSTGRES_USER"`
	DbName           string `mapstructure:"POSTGRES_DB"`
	DbHost           string `mapstructure:"POSTGRES_HOST"`
	DbPort           int    `mapstructure:"POSTGRES_PORT"`
	MaxJobsToProcess int    `mapstructure:"MAX_JOBS_TO_PROCESS"`
	JobsSupplyRate   int    `mapstructure:"JOBS_SUPPLY_RATE"`
}

func NewConfig(path string) (config Config, err error) {
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
