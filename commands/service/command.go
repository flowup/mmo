package service

import (
	"github.com/flowup/mmo/utils"
	"strings"
	"text/template"
	"os"
	"github.com/flowup/mmo/commands"
)

// SetviceOptions encapsulates options that can be passed to the
// service creator
type SetviceOptions struct {
	Name        string
	WebGrpc     bool
	Dsn         string
	ProjectName string
}

// Init is cli function to generate service with given name
func Init(serv SetviceOptions) error {

	if err := utils.CreateDir(serv.Name); err != nil {
		return err
	}

	// create proto dir
	if err := utils.CreateDir(serv.Name + "/protobuf"); err != nil {
		return err
	}

	// create main dir
	if err := utils.CreateDir(serv.Name + "/cmd/" + serv.Name); err != nil {
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
		case "template/main.go":
			filePath = strings.Replace(asset.info.Name(), "template", serv.Name+"/cmd/"+serv.Name, 1)
		case "template/proto.proto":
			filePath = strings.Replace(asset.info.Name(), "template", serv.Name+"/protobuf", 1)
		default:
			filePath = strings.Replace(asset.info.Name(), "template", serv.Name, 1)
		}

		// create template for the file
		tmpl := template.Must(template.New(name).Parse(string(asset.bytes)))

		// create the file in path
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// execute the template to the file
		err = tmpl.Execute(file, serv)
		if err != nil {
			return err
		}
	}

	if err := commands.GenerateProto(commands.Go, serv.Name); err != nil {
		return err
	}

	return nil
}
