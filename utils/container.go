package utils

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"errors"
	"strconv"
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

	code, err := cli.ContainerWait(context.Background(), containerID)

	if code != 0 {
		return errors.New("Container exited with code " + strconv.Itoa(int(code)))
	}

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

// PushImage is function to push image to registry
func PushImage(cli *client.Client, image string) error {
	cmd := exec.Command("docker", "push", image)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
