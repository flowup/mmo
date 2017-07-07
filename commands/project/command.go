package project

import (
	"os"
	"strings"
	"text/template"
	"errors"
	"os/exec"
)

// ProjectOptions encapsulates options that can be passed to the
// project creator
type ProjectOptions struct {
	Name              string
	Language          string
	Path              string
	DependencyManager string
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
		filePath := strings.Replace(asset.info.Name(), "template", opts.Name, 1)
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

	// change to the newly created project and init the dep manager
	err = os.Chdir(opts.Name)
	if err != nil {
		return err
	}

	err = InitializeDependnecyManager(opts.DependencyManager)
	if err != nil {
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}

func InitializeDependnecyManager(man string) error {
	switch man {
	case "glide":
		glideInstallCmd := exec.Command("go", "get", "github.com/Masterminds/glide")
		if err := glideInstallCmd.Run(); err != nil {
			return err
		}

		glideInitCmd := exec.Command("glide", "init", "--non-interactive")
		if err := glideInitCmd.Run(); err != nil {
			return err
		}

	default:
		return errors.New("Unrecognized dependency manager: " + man)
	}

	return nil
}