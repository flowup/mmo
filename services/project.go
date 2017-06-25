package services

import (
	"github.com/flowup/mmo/utils"
)

func Project(name string) error {
	definition := utils.Definition{Name: name}

	if err := utils.CreateDir(name); err != nil {
		return err
	}

	// Create README.md
	if err := utils.CreateFileFromTemplate(definition, "template/README.md"); err != nil {
		return err
	}

	//Create mmo.json
	if err := utils.CreateFileFromTemplate(definition, "template/mmo.json"); err != nil {
		return err
	}

	//TODO glide
	//TODO infra folde
	//TODO contributing
	//TODO issue template

	//TODO wercker.yml
	/*if err := utils.CreateFileFromTemplate(definition, "template/wercker.yml"); err != nil {
		return err
	}*/

	//TODO gitignore
	/*if err := utils.CreateFileFromTemplate(definition, "template/gitignore"); err != nil {
		return err
	}*/

	return nil
}
