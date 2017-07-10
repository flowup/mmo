package config

import (
	"github.com/spf13/viper"
	"github.com/flowup/mmo/commands"
)

// Config represents a configuration manager
type Config struct {
	config *viper.Viper
}

// ReadConfig returns the new configuration
// TODO: return struct instead of viper configuration
func ReadConfig() (*Config, error) {
	viper.SetConfigName("mmo")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	return &Config{config: viper.GetViper()}, err
}

// GetGoPrefix returns the golang directory prefix
// TODO: deprecate this
func (c *Config) GetGoPrefix() string {
	return viper.GetString("goPrefix")
}

// GetProjectName returns name of the project from the
// configuration
func (c *Config) GetProjectName() string {
	return viper.GetString("name")
}

// GetLang returns language of the current project
func (c *Config) GetLang() commands.Language {
	return commands.Language(viper.GetString("lang"))
}

// HasWebRPC resolves if the project has activated the web-rpc
// TODO: deprecate this. Each service should have this flag
func (c *Config) HasWebRPC() bool {
	return viper.GetBool("webRPC")
}
