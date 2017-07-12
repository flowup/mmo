package commands

import "github.com/flowup/mmo/config"

// Mmo represents config and context
type Mmo struct {
	Config  *config.Config
	Context *config.Context
}

// GetMmo Load context and config from config files
func GetMmo() *Mmo {

	var mmo Mmo
	mmoContext, err := config.LoadContext()
	if err != nil {
		mmo.Context = nil
	} else {
		mmo.Context = &mmoContext
	}

	mmoConfig, err := config.LoadConfig(config.FilenameConfig)
	if err != nil {
		mmo.Config = nil
	} else {
		mmo.Config = &mmoConfig
	}

	return &mmo
}