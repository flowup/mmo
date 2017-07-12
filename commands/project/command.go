package project

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/commands"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/utils"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type Mmo struct {
	Config  *config.Config
	Context *config.Context
}

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

// Create extends all assets using project options passed by the caller
// This automatically creates a project folder with all files
func (mmo *Mmo) Create() error {

	// create project folder
	err := os.Mkdir(mmo.Config.Name, os.ModePerm)
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
		filePath := strings.Replace(asset.info.Name(), "template", mmo.Config.Name, 1)
		// create template for the file
		tmpl := template.Must(template.New(name).Parse(string(asset.bytes)))

		// create the file in path
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}

		// execute the template to the file
		err = tmpl.Execute(file, mmo.Config)
		if err != nil {
			return err
		}

		file.Close()
	}

	// change to the newly created project and init the dep manager
	err = os.Chdir(mmo.Config.Name)
	if err != nil {
		return err
	}

	err = mmo.InitializeDependencyManager()
	if err != nil {
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}

// InitializeDependencyManager initializes given dependency manager
// within the current project.
// It will also automatically update the dependency manager
func (mmo *Mmo) InitializeDependencyManager() error {
	switch mmo.Config.DepManager {
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
		return errors.New("Unrecognized dependency manager: " + mmo.Config.DepManager)
	}

	return nil
}

// RunTests is cli function to run tests of application in docker
func (mmo *Mmo) RunTests() error {

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, serviceName := range mmo.Context.Services {

		fmt.Println("Running tests for service \"" + serviceName + "\":")

		testContainer, err := cli.ContainerCreate(context.Background(), &container.Config{
			Image:      "flowup/mmo-webrpc",
			Cmd:        []string{"bash", "-c", "go test $(glide novendor)"},
			WorkingDir: "/go/src/" + mmo.Config.GoPackage + "/" + serviceName,
		}, &container.HostConfig{
			AutoRemove: true,
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: pwd,
					Target: "/go/src/" + mmo.Config.GoPackage,
				},
			},
		}, nil, "")

		if err != nil {
			return err
		}

		err = utils.ContainerRunStdout(cli, testContainer.ID)
		if err != nil {
			return err
		}

		fmt.Println()
	}

	return nil
}

// SetContext is cli function to set context of mmo to specified service or services
func (mmo *Mmo) SetContext(services []string) error {
	for _, service := range services {
		if _, ok := mmo.Config.Services[service]; !ok {
			return utils.ErrServiceNotExists
		}

		if _, err := os.Stat(service); os.IsNotExist(err) {
			return errors.Wrap(utils.ErrServiceNotExists, service)
		}
	}

	serviceContext := config.Context{
		Services: services,
	}

	err := config.SaveContext(serviceContext)

	return err
}

// ProtoGen is cli function to generate API clients and server stubs of specified service or services
func (mmo *Mmo) ProtoGen() error {

	for _, serviceName := range mmo.Context.Services {
		if _, err := os.Stat(serviceName + "/sdk"); os.IsNotExist(err) {
			os.Mkdir(serviceName+"/sdk", os.ModePerm)
		}

		err := commands.GenerateProto(mmo.Config.Lang, serviceName)
		if err != nil {
			return err
		}

		if mmo.Config.Services[serviceName].WebRPC {
			err = commands.GenerateProto(commands.TypeScript, serviceName)
			if err != nil {
				return err
			}
		}

		fmt.Println()
	}

	return nil
}
