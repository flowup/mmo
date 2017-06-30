package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
}

func ReadConfig() *Config {
	viper.SetConfigName("mmo")
	viper.AddConfigPath(".")

	viper.ReadInConfig()

	return &Config{config: viper.GetViper()}
}

func (c *Config) GetGoPrefix() string {
	return viper.GetString("goPrefix")
}

func (c *Config) GetProjectName() string {
	return viper.GetString("name")
}