package main

import (
	"bufio"
	"go/build"
	"os"
	"strings"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/flowup/mmo/api/server"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/generator"
	"github.com/flowup/mmo/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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
			Usage:   "creates new mmo with a given name",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template, t",
					Usage: "template folder for mmo generation",
					Value: build.Default.GOPATH + "/src/github.com/flowup/mmo/templates/go-project",
				},
				cli.StringSliceFlag{
					Name:  "option, x",
					Usage: "additional options passed to template engine (-x name=value)",
				},
				cli.StringFlag{
					Name:  "prefix, p",
					Usage: "Go prefix of the project",
					Value: "",
				},
			},
			Action: func(c *cli.Context) error {
				debug(c, "Initialization process was started")
				var err error
				arg := c.Args().First()

				if arg == "" {
					return errors.New("Missing mmo name argument")
				}

				prefix := c.String("p")
				if prefix == "" {
					prefix, err = getPrefix()
					if err != nil {
						return err
					}
				}

				m := &Mmo{
					Config: &config.Config{
						Name:   arg,
						Prefix: config.GoPrefix(prefix),
					},
				}

				// if we can't stat the folder, we'll just create new
				if _, err := os.Stat(m.Config.Name); err == nil {

					reader := bufio.NewReader(os.Stdin)
					answer := ""

					log.Warnln("Initializing mmo in an existing folder. Do you want to proceed? [y/n]: ")

					for answer != "y" && answer != "n" {
						answer, _ = reader.ReadString('\n')
						answer = strings.Trim(answer, "\n")
					}

					if answer == "n" {
						return nil
					}
				}

				err = generator.GenerateProject(
					m.Config,
					c.StringSlice("x"),
					c.String("t"),
					".",
				)

				if err != nil {
					return err
				}

				log.Infof("Project %s was created", arg)

				return nil
			},
		},
		{
			Name:    "context",
			Aliases: []string{"ctx"},
			Usage:   "sets context to the service(s) given by the argument(s)",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "reset, r",
					Usage: "reset context back to global",
				},
			},
			Action: func(c *cli.Context) error {
				debug(c, "Set context process was started")

				m, err := GetMmo()
				if err != nil {
					return utils.ErrNoProject
				}

				if c.IsSet("reset") {
					if err := m.ResetContext(); err != nil {
						return err
					}
					return nil
				}

				if c.NArg() == 0 {
					log.Infoln("Current context:", m.Context.Services)
					return nil
				}

				services := make([]string, c.NArg())
				for i := 0; i < c.NArg(); i++ {
					services[i] = c.Args().Get(i)
				}

				if err := m.SetContext(services); err != nil {
					return err
				}

				log.Infoln("Current context:", m.Context.Services)

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

				_, err := GetMmo()

				if err != nil {
					return utils.ErrNoProject
				}

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
				m, err := GetMmo()

				if err != nil {
					return utils.ErrNoProject
				}

				if m.Context == nil {
					return utils.ErrContextNotSet
				}

				return utils.ErrNotImplemented
			},
		},
		{
			Name:  "gen",
			Usage: "generates third party resources such as proto stubs or serializers",
			Action: func(c *cli.Context) error {
				debug(c, "start generating MMO resources")

				m, err := GetMmo()

				if err != nil {
					return utils.ErrNoProject
				}

				var services []string

				if m.Context.IsGlobal() {
					services = m.Config.ServiceNames()
					log.Debugln("Global context")
				} else {
					log.Debugln("Not global context")
					services = m.Context.GetServices()
				}

				for _, service := range services {
					log.Debugln("Running service " + service)
					err = m.Plugins.RunGen([]string{service}, m.Config.Services[service].Plugins)
					if err != nil {
						log.Error(err)
					}
				}

				err = m.Plugins.RunGen(m.Config.ServiceNames(), m.Config.Plugins)
				if err != nil {
					return err
				}

				log.Infoln("Generation was completed")

				return nil
			},
		}, {
			Name:  "template",
			Usage: "shows information about template",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return utils.ErrNoArg
				}

				log.Infoln(utils.GetTemplateHelp(c.Args().First()))

				return nil
			},
		}, {
			Name:  "ui",
			Usage: "Run MMO UI",
			Action: func(c *cli.Context) error {
				server.Serve()
				return nil
			},
		}, {
			Name:  "add",
			Usage: "scaffolds resources across the mmo such as services, models, etc.",
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
					Usage: "creates new service within the mmo",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "template, t",
							Usage: "template folder for service generation",
							Value: build.Default.GOPATH + "/src/github.com/flowup/mmo/templates/go-service",
						},
						cli.StringFlag{
							Name:  "description, d",
							Usage: "description of the service",
						},
						cli.StringSliceFlag{
							Name:  "option, x",
							Usage: "additional options passed to template engine (-x name=value)",
						},
					},
					Action: func(c *cli.Context) error {
						debug(c, "Start creating new MMO service")

						arg := c.Args().First()

						if c.NArg() == 0 {
							return utils.ErrNoArg
						}

						m, err := GetMmo()
						if err != nil {
							log.Fatal(err)
							return nil
						}

						if m.Config.Services == nil {
							m.Config.Services = make(map[string]config.Service)
						}

						m.Config.Services[arg] = config.Service{
							Name:        arg,
							Description: c.String("d"),
						}

						err = generator.GenerateService(generator.Service{
							Name:    arg,
							Project: m.Config.Name,
							Package: string(m.Config.Prefix),
						},
							c.StringSlice("x"),
							c.String("t"),
							".",
						)

						if err != nil {
							return err
						}

						if err := config.SaveConfig(m.Config); err != nil {
							return err
						}

						log.Infof("Service %s was created", arg)

						return nil
					},
				},
				{
					Name:  "plugin",
					Usage: "adds plugin to the current service",
					Action: func(c *cli.Context) error {
						debug(c, "Start adding new MMO plugin")

						arg := c.Args().First()

						if c.NArg() == 0 {
							return utils.ErrNoArg
						}

						m, err := GetMmo()
						if err != nil {
							return utils.ErrNoProject
						}

						m.Config.AddPlugin(arg)
						err = config.SaveConfig(m.Config)
						if err != nil {
							return err
						}

						log.Infof("Plugin %s was created", arg)

						return nil
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}

// debug debugs the command to global flags
func debug(c *cli.Context, message string) {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	log.WithFields(log.Fields{
		"arg":     c.Args(),
		"flags":   c.GlobalFlagNames(),
		"command": c.Command.Name,
	}).Debugln(message)
}

func getPrefix() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(path, os.Getenv("GOPATH")+"/src/"), nil
}
