package project

import (
	"bufio"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/utils"
	"github.com/flowup/mmo/utils/cookiecutter"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
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

	mmoPlugins, err := config.NewPlugins(mmoConfig.Plugins)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to init mmo plugins")
	}

	mmo.Plugins = &mmoPlugins

	return mmo, nil
}

// InitProject extends all assets using project options passed by the caller
// This automatically creates a project folder with all files
func (mmo *Mmo) InitProject() error {
	rewriteAll := false

	// go through assets and generate them
	for name, assetGetter := range _bindata {
		asset, err := assetGetter()
		if err != nil {
			return err
		}

		if _, err := os.Stat(name); err == nil {
			// ask user if he want's to rewrite if not rewritten
			if !rewriteAll {
				log.Info("File ", name, " already exists. Do you want to rewrite it? [y/yy/n]:")

				reader := bufio.NewReader(os.Stdin)
				answer := ""
				for answer != "y" && answer != "n" && answer != "yy" {
					answer, _ = reader.ReadString('\n')
					answer = strings.Trim(answer, "\n")
				}

				// rewrites all files automatically
				if answer == "yy" {
					rewriteAll = true
				}
			}

			log.Infoln("Rewriting file:", name)
			err = os.Remove(name)
			if err != nil {
				return err
			}
		}

		log.Debugln("Creating file:", name)
		err = cookiecutter.Restore(
			mmo.Config.Name,
			name,
			string(asset.bytes),
			mmo.Config,
			utils.DefaultFuncMap,
		)
		if err != nil {
			return err
		}
	}

	// change to the newly created project and init the dep manager
	err := os.Chdir(mmo.Config.Name)
	if err != nil {
		return errors.Wrap(err, "chdir failed with an error")
	}

	err = os.Chdir("..")
	if err != nil {
		return errors.Wrap(err, "chdir back failed with an error")
	}

	return nil
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

// ResetContext
func (mmo *Mmo) ResetContext() error {
	log.Debugln("Resetting context")
	return mmo.SetContext([]string{})
}
