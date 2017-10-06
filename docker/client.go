package docker

import (
	"os/exec"
)

type Client struct {
}

func CreateClient() (*Client, error) {
	c := Client{}
	err := c.GetInfo()
	return &c, err
}

func (c *Client) GetInfo() error {
	cmd := exec.Command("docker", "info")
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
