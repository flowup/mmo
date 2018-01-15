package docker

import (
	"fmt"
	"os"
	"os/exec"
)

// Client represents Docker client
type Client struct {
}

// CreateClient is constructor for initializing Client
func CreateClient() (*Client, error) {
	c := Client{}
	err := c.Run(CreateRunOptions("", "info"), false)
	return &c, err
}

// Run is method to run Docker image according to passed options
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
