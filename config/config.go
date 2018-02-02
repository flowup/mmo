package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sort"
)

const (
	// FilenameConfig name of default mmo.yaml file
	FilenameConfig = "mmo.yaml"
)

// Config represents projects configuration
type Config struct {
	Name     string             `yaml:"name"`
	Plugins  []string           `yaml:"plugins"`
	Prefix   string             `yaml:prefix`
	Services map[string]Service `yaml:"services"`
}

// ServiceNames returns an array of all services registered within the config
func (c *Config) ServiceNames() []string {
	names := make([]string, 0, len(c.Services))
	for key := range c.Services {
		names = append(names, key)
	}

	sort.Strings(names)

	return names
}

// AddPlugin is method for adding global plugin
func (c *Config) AddPlugin(name string) {
	for _, plugin := range c.Plugins {
		if plugin == name {
			return
		}
	}

	c.Plugins = append(c.Plugins, name)
}

// Service represents service configuration from Config
type Service struct {
	Name        string   `yaml:"-"`
	Description string   `yaml:"description"`
	Plugins     []string `yaml:"plugins"`
}

// Dependency represents service dependency configuration from Config
type Dependency struct {
	Name string        `yaml:"name"`
	Type string        `yaml:"type"`
	Run  DependencyRun `yaml:"run"`
}

// DependencyRun represents what should dependency run
type DependencyRun struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
}

// LoadConfig loads project config from the given directory
func LoadConfig(filenameConfig string) (*Config, error) {
	b, err := ioutil.ReadFile(filenameConfig)
	if err != nil {
		return &Config{}, err
	}

	cfg := &Config{}
	return cfg, yaml.Unmarshal(b, &cfg)
}

// SaveConfig saves given config to  the given directory
func SaveConfig(cfg *Config) error {
	b, err := yaml.Marshal(cfg)

	if err != nil {
		return nil
	}

	f, err := os.Create(FilenameConfig)
	_, err = f.Write(b)

	return err
}
