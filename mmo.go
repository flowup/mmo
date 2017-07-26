package main

import (
	"bufio"
	"github.com/evalphobia/logrus_sentry"
	"github.com/flowup/mmo/commands/project"
	"github.com/flowup/mmo/commands/service"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"strings"
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
	log.SetLevel(log.DebugLevel)

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
			log.Fatalln("Got Nuked :( :", errors.WithStack(err.(error)))
		}
	}()

	app := cli.NewApp()
	app.Name = "mmo"
	app.Usage = ""

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "sets logging level to debug",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{},
			Usage:   "creates new project with a given name",
			Action: func(c *cli.Context) error {
				bootstrap(c)

				if c.Args().First() == "" {
					return errors.New("Missing project name argument")
				}

				mmo := project.Mmo{}
				mmo.Config = &config.Config{}

				mmo.Config.Name = c.Args().First()
				mmo.Config.Lang = "go"
				mmo.Config.DepManager = "glide"

				// if we can't stat the folder, we'll just create new
				if _, err := os.Stat(mmo.Config.Name); err != nil {
					// create project folder
					err := os.Mkdir(mmo.Config.Name, os.ModePerm)
					if err != nil {
						return err
					}
				} else {
					reader := bufio.NewReader(os.Stdin)
					answer := ""
					log.Info("Initializing project in an existing folder. Do you want to proceed? [y/n]: ")
					for answer != "y" && answer != "n" {
						answer, _ = reader.ReadString('\n')
						answer = strings.Trim(answer, "\n")
					}

					if answer == "n" {
						return nil
					}
				}

				if err := mmo.InitProject(); err != nil {
					log.Fatal(err.Error())
				}

				return nil
			},
		},
		{
			Name:    "context",
			Aliases: []string{"ctx"},
			Usage:   "sets context to the service(s) given by the argument(s)",
			Action: func(c *cli.Context) error {
				bootstrap(c)

				mmo, err := project.GetMmo()
				if err != nil {
					return utils.ErrNoProject
				}

				if c.NArg() == 0 {
					log.Infoln("Current context:", mmo.Context.Services)
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
				bootstrap(c)

				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "run",
			Usage: "runs services and their dependencies using docker on your machine",
			Action: func(c *cli.Context) error {
				bootstrap(c)

				mmo, err := project.GetMmo()

				if err != nil {
					return utils.ErrNoProject
				}

				if err := mmo.Run(); err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
		{
			Name:  "build",
			Usage: "builds docker images for all services targeted by the context",
			Action: func(c *cli.Context) error {
				bootstrap(c)

				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "integration",
			Usage: "builds all the services, deploys them to the kubernetes development cluster and starts up the integration tests. ",
			Action: func(c *cli.Context) error {
				bootstrap(c)

				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "test",
			Usage: "runs tests for all services targeted by the context",
			Action: func(c *cli.Context) error {
				bootstrap(c)

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
			Usage: "generates third party resources such as proto stubs or serializers",
			Subcommands: []cli.Command{
				{
					Name:  "proto",
					Usage: "generates API clients and server stubs from proto definition for all services targeted by the context",
					Action: func(c *cli.Context) error {
						bootstrap(c)

						mmo, err := project.GetMmo()

						if err != nil {
							return utils.ErrNoProject
						}

						lang := mmo.Config.Lang

						if c.NArg() == 1 {
							lang = c.Args().Get(0)
						}

						services := mmo.Context.Services
						if len(services) == 0 {
							log.Warnln("No context set, using global")
							services = mmo.Config.ServiceNames()
						}

						if err := mmo.ProtoGen(services, lang); err != nil {
							log.Fatal(err)
						}
						return nil
					},
				}, {
					Name:  "gateway",
					Usage: "generates GRPC gateways from proto definition for all services targeted by the context",
					Action: func(c *cli.Context) error {
						bootstrap(c)

						mmo, err := project.GetMmo()

						if err != nil {
							return utils.ErrNoProject
						}

						services := mmo.Context.Services
						if len(services) == 0 {
							log.Warnln("No context set, using global")
							services = mmo.Config.ServiceNames()
						}

						if err := mmo.ProtoGen(services, "gw"); err != nil {
							log.Fatal(err)
						}
						return nil
					},
				},
			},
		},
		{
			Name:  "add",
			Usage: "scaffolds resources across the project such as services, models, etc.",
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
							Usage: "automatically generates webrpc integration",
						},
						cli.BoolFlag{
							Name:  "sentry",
							Usage: "allows Sentry integration",
						},
					},

					Action: func(c *cli.Context) error {
						bootstrap(c)

						if c.NArg() == 0 {
							return utils.ErrNoArg
						}

						mmo, err := project.GetMmo()
						if err != nil {
							return utils.ErrNoProject
						}

						mmo.Config.Services = make(map[string]config.Service, len(mmo.Config.Services)+1)
						//mmo.Config.Services[c.Args().First()] = service.Wizzar(c.Args().First())
						mmo.Config.Services[c.Args().First()] = service.FromCliContext(c.Args().First(), c)

						if err := config.SaveConfig(mmo.Config, config.FilenameConfig); err != nil {
							log.Fatal(err)
						}

						if err := service.InitService(mmo.Config.Services[c.Args().First()]); err != nil {
							log.Fatal(err)
						}
						return nil
					},
				},
				{
					Name:  "plugin",
					Usage: "adds plugin to the current service",
					Action: func(c *cli.Context) error {
						bootstrap(c)

						return utils.ErrNotImplemented
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

// bootstrap bootstraps the command to global flags
func bootstrap(c *cli.Context) {
	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Bootstrap successful")
}
