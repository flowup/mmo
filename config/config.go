package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
}

func ReadConfig() (*Config, error) {
	viper.SetConfigName("mmo")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	return &Config{config: viper.GetViper()}, err
}

func (c *Config) GetGoPrefix() string {
	return viper.GetString("goPrefix")
}

func (c *Config) GetProjectName() string {
	return viper.GetString("name")
}

func (c *Config) GetLang() string {
	return viper.GetString("lang")
}

func (c *Config) HasWebRPC() bool {
	return viper.GetBool("webRPC")
}
