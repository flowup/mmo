package config

import (
	"github.com/flowup/mmo/docker"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

// Plugins is structure that represents plugin system of the MMO
type Plugins struct {
	Client *docker.Client
	Config *Config
}

// NewPlugins is function to get instance of the plugin system
func NewPlugins(config *Config) (Plugins, error) {
	cli, err := docker.CreateClient()
	if err != nil {
		return Plugins{}, err
	}

	return Plugins{Client: cli, Config: config}, nil
}

// RunGen is method to run all plugins that have specified hook
func (p *Plugins) RunGen(contextServices []string, plugins []string) error {

	for _, plugin := range plugins {
		logrus.Infof("Running plugin %s", plugin)
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		opts := docker.CreateRunOptions(plugin, "run")
		opts.AddArguments(contextServices...)
		opts.AutoRemove = true
		opts.MountHostVolume(pwd, "/source")
		opts.AddEnvVariable("GO_PREFIX", p.Config.Prefix)

		if len(contextServices) == 1 {
			opts.WorkingDir = "/source/" + contextServices[0]
		}

		err = p.Client.Run(opts, true)
		if err != nil {
			return errors.Wrap(err, "Error running plugin "+plugin)
		}
	}

	return nil
}
