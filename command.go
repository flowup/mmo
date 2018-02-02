package main

import (
	"os"

	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Mmo represents config and context
type Mmo struct {
	Config  *config.Config
	Context *config.Context
	Plugins *config.Plugins
}

// GetMmo Load context and config from config files
func GetMmo() (*Mmo, error) {

	mmo := &Mmo{
		&config.Config{},
		&config.Context{},
		&config.Plugins{},
	}

	mmoContext, err := config.LoadContext()
	if err == nil {
		mmo.Context = mmoContext
	}

	mmoConfig, err := config.LoadConfig(config.FilenameConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load mmo config")
	}

	mmo.Config = mmoConfig

	mmoPlugins, err := config.NewPlugins(mmoConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to init mmo plugins")
	}

	mmo.Plugins = &mmoPlugins

	return mmo, nil
}

// SetContext is cli function to set context of mmo to specified service or services
func (mmo *Mmo) SetContext(services []string) error {
	log.Debugln("Setting context for services:", services)
	for _, service := range services {
		if _, ok := mmo.Config.Services[service]; !ok {
			return utils.ErrServiceNotExists
		}

		if _, err := os.Stat(service); os.IsNotExist(err) {
			return errors.Wrap(utils.ErrServiceNotExists, service)
		}
	}

	serviceContext := &config.Context{
		Services: services,
	}

	err := config.SaveContext(serviceContext)

	return err
}

// ResetContext resets user's context of the MMO
func (mmo *Mmo) ResetContext() error {
	log.Debugln("Resetting context")
	return mmo.SetContext([]string{})
}
