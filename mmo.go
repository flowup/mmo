package main

import (
	"errors"
	"github.com/flowup/mmo/utils"
	"github.com/urfave/cli"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/commands/project"
	"github.com/flowup/mmo/commands/service"
	"os"
	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	// Log global settings of logger
	dsn = "https://5554d201ce064e8790792540de39c608:fb93a01a320a486ab40b4fbb5feaf7ac@sentry.io/190135"
)

func init() {
	// Logging format is Text
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	// this should be override-able by some debug flag
	log.SetLevel(log.InfoLevel)

	levels := []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	}

	hook, err := logrus_sentry.NewSentryHook(dsn, levels)
	hook.Timeout = 20 * time.Second
	hook.StacktraceConfiguration.Enable = true

	if err == nil {
		log.AddHook(hook)
	}
}


func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("Got Nuked :( :", err)
		}
	}()

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

				if err := mmo.Init(); err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
		{
			Name:    "context",
			Aliases: []string{"ctx"},
			Usage:   "sets context to the service(s) given by the argument(s)",
			Action: func(c *cli.Context) error {

				mmo, err := project.GetMmo()
				if err != nil {
					return utils.ErrNoProject
				}

				if c.NArg() == 0 {
					log.Println("Current context:", mmo.Context.Services)
					return nil
				}

				services := make([]string, c.NArg())
				for i := 0; i < c.NArg(); i++ {
					services[i] = c.Args().Get(i)
				}

				if err := mmo.SetContext(services); err != nil {
					log.Fatal(err)
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

				mmo, err := project.GetMmo()

				if err != nil {
					return utils.ErrNoProject
				}

				if mmo.Context == nil {
					return utils.ErrContextNotSet
				}

				if err := mmo.RunTests(); err != nil {
					log.Fatal(err)
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

						mmo, err := project.GetMmo()

						if err != nil {
							return utils.ErrNoProject
						}

						services := mmo.Context.Services
						if len(services) == 0 {
							log.Warnln("No context set, using global")
							services = mmo.Config.ServiceNames()
						}

						if err := mmo.ProtoGen(services); err != nil {
							log.Fatal(err)
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
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "description, d",
							Usage: "set service description",
						},
						cli.BoolFlag{
							Name:  "webrpc",
							Usage: "set webrpc to service",
						},
						cli.StringFlag{
							Name:  "sentry-dsn",
							Usage: "set sentry dns",
						},
					},

					Action: func(c *cli.Context) error {
						if c.NArg() == 0 {
							return utils.ErrNoArg
						}

						mmo := project.GetMmo()
						if mmo.Config == nil {
							return utils.ErrNoProject
						}

						mmo.Config.Services = make(map[string]config.Service, len(mmo.Config.Services)+1)
						//mmo.Config.Services[c.Args().First()] = service.Wizzar(c.Args().First())
						mmo.Config.Services[c.Args().First()] = service.Flags(c.Args().First(), c)

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
