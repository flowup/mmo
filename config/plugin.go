package config

import (
	"github.com/flowup/mmo/docker"
	"github.com/sirupsen/logrus"
	"os"
	"github.com/pkg/errors"
)

type Plugins struct {
	Gen    []Plugin
	Client *docker.Client
}

type Plugin struct {
	Image string `yaml:"-"`
}

// NewPlugins is function to get instance of the plugin system
func NewPlugins(pluginNames []string) (Plugins, error) {
	cli, err := docker.CreateClient()
	if err != nil {
		return Plugins{}, err
	}

	plugins := make([]Plugin, len(pluginNames))
	for i, pluginName := range pluginNames {
		plugins[i] = Plugin{Image: pluginName}
	}

	return Plugins{Client: cli, Gen: plugins}, nil

}

// RunHook is method to run all plugins that have specified hook
func (p *Plugins) RunGen(contextServices []string) error {

	for _, plugin := range p.Gen {
		logrus.Infof("Running plugin %s", plugin.Image)
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		opts := docker.CreateRunOptions(plugin.Image, "run")
		opts.AddArguments(contextServices...)
		opts.AutoRemove = true
		opts.MountHostVolume(pwd, "/source")

		err = p.Client.Run(opts, true)
		if err != nil {
			return errors.Wrap(err, "Error running plugin " + plugin.Image)
		}
	}

	return nil
}
