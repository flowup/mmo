package commands

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/utils"
	"github.com/flowup/mmo/utils/dockercmd"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// Supported languages
const (
	Go          = "go"
	Python      = "python"
	TypeScript  = "ts"
	GRPCGateway = "gw"
	GRPCSwagger = "swagger"
)

// GenerateProto generates proto files in a given language for
// the given service.
// Current Support: go, python, typescript(ts)
func GenerateProto(lang string, serviceName string) error {

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	log.Infoln("Generating "+string(lang)+" API clients and server stubs for service:", serviceName)

	var inputMount string
	var outputMount string
	var cmd string
	var image string

	inputMount = pwd + "/" + serviceName + "/protobuf"
	outputMount = pwd + "/" + serviceName

	switch lang {
	case Go:
		cmd = dockercmd.GoGen
		image = dockercmd.ImageGo
	case Python:
		cmd = dockercmd.PyGen
		image = dockercmd.ImagePy
	case TypeScript:
		outputMount = pwd + "/" + serviceName + "/sdk"
		cmd = dockercmd.TsGen
		image = dockercmd.ImageTs
	case GRPCGateway:
		cmd = dockercmd.GGwGen
		image = dockercmd.ImageGo
	case GRPCSwagger:
		cmd = dockercmd.SwaggerGen
		image = dockercmd.ImageGo
	default:
		return errors.New("Invalid generation language: " + lang)
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	err = utils.PullImage(cli, image)
	if err != nil {
		return errors.Wrap(err, "Failed to pull image "+image)
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
		return errors.Wrap(err, "Failed to create container from image "+image)
	}

	err = utils.ContainerRunStdout(cli, cont.ID)
	return errors.Wrap(err, "Failed to run container from image "+image)
}
