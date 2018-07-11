package config

import (
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/flowup/mmo/docker"
	"github.com/pkg/errors"
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

// gen is method to run all plugins that have specified hook
func (p *Plugins) gen(contextServices []string, plugins []string) error {

	for _, plugin := range plugins {
		log.Infof("Running plugin %s", plugin)
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		opts := docker.CreateRunOptions(plugin, "run")
		opts.AddArguments(contextServices...)
		opts.AutoRemove = true
		opts.MountHostVolume(pwd, "/source")
		opts.AddEnvVariable("GO_PREFIX", string(p.Config.Prefix))

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

// RunGen is method to run all gived services with plugins that have specified hook
func RunGen(config *Config, plugins *Plugins, services []string) error {
	var err error
	for _, service := range services {
		log.Debugln("Running service " + service)
		err = plugins.gen([]string{service}, config.Services[service].Plugins)
		if err != nil {
			log.Error(err)
		}
	}

	err = plugins.gen(config.ServiceNames(), config.Plugins)
	if err != nil {
		return err
	}

	return nil
}
