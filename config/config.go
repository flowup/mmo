package config

import (
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
)

const (
	FilenameConfig = "mmo.yaml"
)

type Config struct {
	Name       string `yaml:"name"`
	Lang       string `yaml:"lang"`
	DepManager string `yaml:"dependencyManager"`
	GoPackage  string `yaml:"goPackage"`
	Services   map[string]Service `yaml:"services"`
}

type Service struct {
	Description  string `yaml:"description"`
	WebRPC       bool `yaml:"webRPC"`
	Dependencies []Dependency `yaml:"dependencies"`
}

type Dependency struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Run  DependencyRun `yaml:"run"`
}

type DependencyRun struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
}

// LoadContext loads project context from the given directory
func LoadConfig(filenameConfig string) (Config, error) {
	b, err := ioutil.ReadFile(filenameConfig)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = yaml.Unmarshal(b, &cfg)

	return cfg, err
}

// SaveContext saves given context to the current path
func SaveConfig(cfg Config, filenameConfig string) error {
	b, err := yaml.Marshal(cfg)

	if err != nil {
		return nil
	}

	f, err := os.Create(filenameConfig)
	_, err = f.Write(b)

	return err
}
