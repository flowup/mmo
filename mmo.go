package main

import (
	"errors"
	"github.com/flowup/mmo/utils"
	"github.com/urfave/cli"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/commands/project"
	"github.com/flowup/mmo/commands/service"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "mmo"
	app.Usage = ""

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{},
			Usage:   "creates new project with a given name",
			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					return errors.New("Missing project name argument")
				}

				mmo := project.Mmo{}
				mmo.Config = &config.Config{}

				mmo.Config.Name = c.Args().First()
				mmo.Config.Lang = "go"
				mmo.Config.DepManager = "glide"

				if err := mmo.InitProject(); err != nil {
					utils.Log.Fatal(err)
				}
				return nil
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

				mmo := project.GetMmo()
				if mmo.Config == nil {
					return utils.ErrNoProject
				}

				services := make([]string, c.NArg())
				for i := 0; i < c.NArg(); i++ {
					services[i] = c.Args().Get(i)
				}

				if err := mmo.SetContext(services); err != nil {
					utils.Log.Fatal(err)
				}
				return nil
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
			Name:  "run",
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
			Name:  "integration",
			Usage: "builds all the services, deploys them to the kubernetes development cluster and starts up the integration tests. ",
			Action: func(c *cli.Context) error {
				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "test",
			Usage: "runs tests for all services targeted by the context",
			Action: func(c *cli.Context) error {

				mmo := project.GetMmo()

				if mmo.Config == nil {
					return utils.ErrNoProject
				}

				if mmo.Context == nil {
					return utils.ErrContextNotSet
				}

				if err := mmo.RunTests(); err != nil {
					utils.Log.Fatal(err)
				}
				return nil
			},
		},
		{
			Name:  "gen",
			Usage: "is used to generate various components across services",
			Subcommands: []cli.Command{
				{
					Name:  "proto",
					Usage: "generates API clients and server stubs from proto definition for all services targeted by the context",
					Action: func(c *cli.Context) error {

						mmo := project.GetMmo()

						if mmo.Config == nil {
							return utils.ErrNoProject
						}

						if mmo.Context == nil {
							return utils.ErrContextNotSet
						}

						if err := mmo.ProtoGen(); err != nil {
							utils.Log.Fatal(err)
						}
						return nil
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
					Name:  "service",
					Usage: "creates new service within the project",
					Action: func(c *cli.Context) error {
						mmo := project.GetMmo()
						if mmo.Config == nil {
							return utils.ErrNoProject
						}

						mmo.Config.Services = make(map[string]config.Service, len(mmo.Config.Services)+1)
						mmo.Config.Services[c.Args().First()] = service.Wizzar(c.Args().First())

						if err := config.SaveConfig(*mmo.Config, config.FilenameConfig); err != nil {
							utils.Log.Fatal(err)
						}

						if err := service.InitService(mmo.Config.Services[c.Args().First()]); err != nil {
							utils.Log.Fatal(err)
						}
						return nil
					},
				}, {
					Name:  "plugin",
					Usage: "adds plugin to the current service",
					Action: func(c *cli.Context) error {
						return utils.ErrNotImplemented
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
