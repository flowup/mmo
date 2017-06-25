package services

import (
	"github.com/flowup/mmo/utils"
)

func Project(name string) error {
	path, err := utils.ImportPath()
	if err != nil {
		return err
	}

	definition := utils.Definition{Name: name, Path: path + "/" + name}

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

	//Create glide.yaml
	if err := utils.CreateFileFromTemplate(definition, "template/glide.yaml"); err != nil {
		return err
	}

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
