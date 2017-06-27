package project

import (
	"os"
	"strings"
	"text/template"
)

// ProjectOptions encapsulates options that can be passed to the
// project creator
type ProjectOptions struct {
	Name string
	Language string
	Path string
}

// Create extends all assets using project options passed by the caller
// This automatically creates a project folder with all files
func Create(opts ProjectOptions) error {

	// create project folder
	err := os.Mkdir(opts.Name, os.ModePerm)
	if err != nil {
		return err
	}

	// go through assets and generate them
	for name, assetGetter := range _bindata {
		asset, err := assetGetter()
		if err != nil {
			return err
		}

		// get correct path to the file
		filePath := strings.Replace(asset.info.Name(), name, "template", 1)

		// create template for the file
		tmpl := template.Must(template.New(name).Parse(string(asset.bytes)))

		// create the file in path
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// execute the template to the file
		err = tmpl.Execute(file, opts)
		if err != nil {
			return err
		}
	}

	return nil
}