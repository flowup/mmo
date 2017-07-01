package commands

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/utils"
)

func GenerateProto(in string, out string, cmd string) error {

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	cont, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: "flowup/mmo-webrpc",
		Cmd:   []string{"bash", "-c", cmd},
	}, &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: in,
				Target: "/in",
			}, {
				Type:   mount.TypeBind,
				Source: out,
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
