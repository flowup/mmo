package config

import (
	"archive/tar"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Plugins struct {
	Loaded  bool
	Plugins []Plugin
	Client  *client.Client
}

type Plugin struct {
	Name   string            `yaml:"name"`
	Image  string            `yaml:"-"`
	Hooks  map[string]string `yaml:"hooks"`
	Global bool              `yaml:"global"`
}

func NewPlugins(pluginNames []string) (Plugins, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return Plugins{}, err
	}

	plugins := make([]Plugin, len(pluginNames))
	for i, pluginName := range pluginNames {
		plugins[i] = Plugin{Image: pluginName}
	}

	return Plugins{Client: cli, Plugins: plugins, Loaded: false}, nil

}
func (p *Plugins) Load() error {
	for i, plugin := range p.Plugins {

		err := utils.PullImage(p.Client, plugin.Image)
		if err != nil {
			logrus.Warn(errors.Wrap(err, "Failed to pull image "+plugin.Image))
		}

		cont, err := p.Client.ContainerCreate(context.Background(), &container.Config{
			Image: plugin.Image,
		}, &container.HostConfig{AutoRemove: true}, nil, "")

		if err != nil {
			return errors.Wrap(err, "Failed to create container from image "+plugin.Image)
		}

		content, stat, err := p.Client.CopyFromContainer(context.Background(), cont.ID, "/plugin/mmo-plugin.yaml")
		if err != nil {
			return err
		}

		if stat.Mode.IsDir() {
			return errors.New("/plugin/mmo-plugin.yaml is dir")
		}

		tr := tar.NewReader(content)

		if _, err = tr.Next(); err != nil {
			return err
		}

		pluginYaml, err := ioutil.ReadAll(tr)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(pluginYaml, &p.Plugins[i])
		if err != nil {
			return err
		}

		content.Close()
	}

	p.Loaded = true
	return nil
}

func (p *Plugins) GetByHook() ([]Plugin, error) {
	if !p.Loaded {
		err := p.Load()
		if err != nil {
			return nil, errors.Wrap(err, "Error loading plugin")
		}
	}
	return nil, nil
}

func (p *Plugins) RunHook(hook string, projectServices []string, contextServices []string) error {
	if !p.Loaded {
		err := p.Load()
		if err != nil {
			return errors.Wrap(err, "Error loading plugin")
		}
	}

	for _, plugin := range p.Plugins {
		if val, ok := plugin.Hooks[hook]; ok {
			logrus.Debugf("Running plugin %s and hook %s", plugin.Name, hook)
			pwd, err := os.Getwd()
			if err != nil {
				return err
			}

			var cmd []string
			if plugin.Global {
				cmd = append([]string{"/hooks/" + val}, projectServices...)
			} else {
				cmd = append([]string{"/hooks/" + val}, contextServices...)
			}

			cont, err := p.Client.ContainerCreate(context.Background(), &container.Config{
				Image: plugin.Image,
				Cmd:   cmd,
			}, &container.HostConfig{
				AutoRemove: true,
				Mounts: []mount.Mount{
					{
						Type:   mount.TypeBind,
						Source: pwd,
						Target: "/source",
					},
				},
			}, nil, "")

			if err != nil {
				return errors.Wrap(err, "Failed to create container from image "+plugin.Image)
			}

			err = utils.ContainerRunStdout(p.Client, cont.ID)
			if err != nil {
				return errors.Wrap(err, "Failed to run container from image "+plugin.Image)
			}
		}
	}

	return nil
}
