package main

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/flowup/mmo/utils"
	"os"
	"github.com/flowup/mmo/services"
)

func main() {
	app := cli.NewApp()
	app.Name = "mmo"
	app.Usage = ""

	app.Commands = []cli.Command{
		{
			Name:    "project",
			Aliases: []string{},
			Usage:   "creates new project with a given name",
			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					fmt.Println("Missing project name")
					return utils.ErrNoArg
				}

				if err := services.Project(c.Args().First()); err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "service",
			Usage: "creates new service within the project",
			Action: func(c *cli.Context) error {
				return utils.ErrNotImplemented
			},
		},
		{
			Name:    "set-context",
			Aliases: []string{"ctx"},
			Usage:   "sets context to the service(s) given by the argument(s)",
			Action: func(c *cli.Context) error {
				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "dev",
			Usage: "starts up development environment for all services targeted by the context",
			Action: func(c *cli.Context) error {
				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "build",
			Usage: "builds docker images for all services targeted by the context",
			Action: func(c *cli.Context) error {
				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "e2e",
			Usage: "spins up e2e tests for all services targeted by the context. Make sure you are targeting all dependencies",
			Action: func(c *cli.Context) error {
				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "deploy",
			Usage: "performs clean build and applies all configurations to the current kubectl context",
			Action: func(c *cli.Context) error {
				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "add",
			Usage: "adds selected resource to the given service",
			Subcommands: []cli.Command{
				{
					Name: "model",
					Usage: "adds model with a given name to the service",
					Action: func(c *cli.Context) error {
						return utils.ErrNotImplemented
					},
				},
				{
					Name: "plugin",
					Usage: "adds plugin with the given name to the service",
					Action: func(c *cli.Context) error {
						return utils.ErrNotImplemented
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
