package docker

import (
	"os/exec"
	"os"
	"fmt"
)

type Client struct {
}

func CreateClient() (*Client, error) {
	c := Client{}
	err := c.Run(CreateRunOptions("", "info"), false)
	return &c, err
}

func (c *Client) Run(options *RunOptions, output bool) error {
	fmt.Println(options.ToArgs())
	cmd := exec.Command("docker", options.ToArgs()...)
	if output {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
