package project

import (
	"os"
	"strings"
	"text/template"
	"github.com/pkg/errors"
	"os/exec"
	"github.com/docker/docker/client"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/flowup/mmo/config"
	"github.com/docker/docker/api/types"
	"io"
	"github.com/flowup/mmo/utils"
	"fmt"
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
		filePath := strings.Replace(asset.info.Name(), "template", "new", 1)

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

	err = InitializeDependencyManager(opts.DependencyManager)
	if err != nil {
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}

func InitializeDependencyManager(man string) error {
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

func RunTests() error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	pConfig := config.ReadConfig()

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	testContainer, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "flowup/mmo-webrpc",
		Cmd:   []string{"bash", "-c", "go test $(glide novendor)"},
		WorkingDir: "/go/src/" + pConfig.GetGoPrefix(),
	}, &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: pwd,
				Target: "/go/src/" + pConfig.GetGoPrefix(),
			},
		},
	}, nil, pConfig.GetProjectName())

	if err != nil {
		return err
	}

	resp, err := cli.ContainerAttach(ctx, testContainer.ID, types.ContainerAttachOptions{
		Stdout: true,
		Stderr: true,
		Stream: true,
	})

	if err != nil {
		return err
	}

	go func() { io.Copy(os.Stdout, resp.Reader) }()

	err = cli.ContainerStart(ctx, testContainer.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	_, err = cli.ContainerWait(context.Background(), testContainer.ID)
	if err != nil {
		return err
	}

	return nil
}

func SetContext(services []string) error {
	for _,service := range services {
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