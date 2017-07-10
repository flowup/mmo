package main

import (
	"errors"
	"github.com/flowup/mmo/commands/project"
	"github.com/flowup/mmo/utils"
	"github.com/urfave/cli"
	"os"
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
					return errors.New("Missing project name argument")
				}

				return project.Create(project.ProjectOptions{
					Name:              c.Args().First(),
					Language:          "go",
					DependencyManager: "glide",
				})
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

				if c.NArg() == 0 {
					return utils.ErrSetContextNoArg
				}

				services := make([]string, c.NArg())
				for i := 0; i < c.NArg(); i++ {
					services[i] = c.Args().Get(i)
				}

				return project.SetContext(services)
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
			Name: "run",
			Usage: "runs services and their dependencies using docker on your machine",
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
			Name:  "test",
			Usage: "runs tests for all services targeted by the context",
			Action: func(c *cli.Context) error {
				return project.RunTests()
			},
		},
		{
			Name:  "gen",
			Usage: "is used to generate various components across services",
			Subcommands: []cli.Command{
				{
					Name: "proto",
					Usage: "generates API clients and server stubs from proto definition for all services targeted by the context",
					Action: func(c *cli.Context) error {
						return project.ProtoGen()
					},
				},
			},
		},
		{
			Name:  "add",
			Usage: "adds selected resource to the given service",
			Subcommands: []cli.Command{
				{
					Name:  "model",
					Usage: "adds model with a given name to the service",
					Action: func(c *cli.Context) error {
						return utils.ErrNotImplemented
					},
				},
				{
					Name:  "dep",
					Usage: "adds dependency with the given name to the service",
					Action: func(c *cli.Context) error {
						return utils.ErrNotImplemented
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
