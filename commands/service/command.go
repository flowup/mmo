package service

import (
	"github.com/flowup/mmo/utils"
	"strings"
	"os"
	"github.com/flowup/mmo/commands"
	"text/template"
	"github.com/flowup/mmo/config"
)

// InitService is cli function to generate service with given name
func InitService(configService config.Service) error {
	// create proto dir
	if err := utils.CreateDir(configService.Name + "/protobuf"); err != nil {
		return err
	}

	// create main dir
	if err := utils.CreateDir(configService.Name + "/cmd/" + configService.Name); err != nil {
		return err
	}

	// go through assets and generate them
	for name, assetGetter := range _bindata {

		asset, err := assetGetter()
		if err != nil {
			return err
		}

		// get correct path to the file
		filePath := ""

		switch asset.info.Name() {
		case "commands/service/template/main_go":
			filePath = strings.Replace(asset.info.Name(), "commands/service/template", configService.Name+"/cmd/"+configService.Name, 1)
		case "commands/service/template/proto.proto":
			filePath = strings.Replace(asset.info.Name(), "commands/service/template", configService.Name+"/protobuf", 1)
		default:
			filePath = strings.Replace(asset.info.Name(), "commands/service/template", configService.Name, 1)
		}

		filePath = strings.Replace(filePath, "_go", ".go", 1)

		// create template for the file
		tmpl := template.Must(template.New(name).Parse(string(asset.bytes)))

		// create the file in path
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// execute the template to the file
		err = tmpl.Execute(file, configService)
		if err != nil {
			return err
		}
	}

	if err := commands.GenerateProto(commands.Go, configService.Name); err != nil {
		return err
	}

	return nil
}
