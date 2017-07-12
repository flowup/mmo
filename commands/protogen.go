package commands

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/utils"
	"github.com/flowup/mmo/utils/dockercmd"
	"os"
)

const (
	// Go represents supported languages
	Go  = "go"
	Python = "python"
	TypeScript = "ts"
)

// GenerateProto generates proto files in a given language for
// the given service.
// Current Support: go, python, typescript(ts)
func GenerateProto(lang string, serviceName string) error {

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println("Generating " + string(lang) + " API clients and server stubs for service \"" + serviceName + "\"...")

	var inputMount string
	var outputMount string
	var cmd string
	var image string

	inputMount = pwd + "/" + serviceName + "/protobuf"

	switch lang {
	case Go:
		outputMount = pwd + "/" + serviceName
		cmd = dockercmd.GoGen
		image = dockercmd.ImageGo
	case Python:
		outputMount = pwd + "/" + serviceName
		cmd = dockercmd.PyGen
		image = dockercmd.ImagePy
	case TypeScript:
		outputMount = pwd + "/" + serviceName + "/sdk"
		cmd = dockercmd.TsGen
		image = dockercmd.ImageTs
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	err = utils.PullImage(cli, image)
	if err != nil {
		return err
	}

	cont, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: image,
		Cmd:   []string{"bash", "-c", cmd},
	}, &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: inputMount,
				Target: "/in",
			}, {
				Type:   mount.TypeBind,
				Source: outputMount,
				Target: "/out",
			},
		},
	}, nil, "")

	if err != nil {
		return err
	}

	err = utils.ContainerRunStdout(cli, cont.ID)
	return err
}
