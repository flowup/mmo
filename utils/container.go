package utils

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"os"
)

// ContainerRunStdout is function to run created docker container and attach output of container to stdout of mmo
func ContainerRunStdout(cli *client.Client, containerID string) error {

	resp, err := cli.ContainerAttach(context.Background(), containerID, types.ContainerAttachOptions{
		Stdout: true,
		Stderr: true,
		Stream: true,
	})

	if err != nil {
		return err
	}

	go func() { io.Copy(os.Stdout, resp.Reader) }()

	err = cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	_, err = cli.ContainerWait(context.Background(), containerID)
	return err
}

// PullImage is function to pull image
func PullImage(cli *client.Client, image string) error {
	out, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(ioutil.Discard, out)
	return err
}
