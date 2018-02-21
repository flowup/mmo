package main

import (
	"bufio"
	"go/build"
	"os"
	"strings"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/generator"
	"github.com/flowup/mmo/kubernetes"
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
			},
			Action: func(c *cli.Context) error {
				bootstrap(c)

				if c.Args().First() == "" {
					return errors.New("Missing mmo name argument")
				}

				m := Mmo{}
				m.Config = &config.Config{}

				m.Config.Name = c.Args().First()

				// if we can't stat the folder, we'll just create new
				if _, err := os.Stat(m.Config.Name); err == nil {

					reader := bufio.NewReader(os.Stdin)
					answer := ""
					log.Info("Initializing mmo in an existing folder. Do you want to proceed? [y/n]: ")
					for answer != "y" && answer != "n" {
						answer, _ = reader.ReadString('\n')
						answer = strings.Trim(answer, "\n")
					}

					if answer == "n" {
						return nil
					}
				}

				err := generator.GenerateProject(
					generator.Project{Name: c.Args().First()},
					c.StringSlice("x"),
					c.String("t"),
					".",
				)

				if err != nil {
					log.Fatal(err)
				}

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
				bootstrap(c)

				m, err := GetMmo()
				if err != nil {
					return utils.ErrNoProject
				}

				if c.IsSet("reset") {
					if err := m.ResetContext(); err != nil {
						log.Fatal(err)
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
				bootstrap(c)

				m, err := GetMmo()

				if err != nil {
					return utils.ErrNoProject
				}

				var services []string

				if m.Context.IsGlobal() {
					services = m.Config.ServiceNames()
					log.Debug("Global context")
				} else {
					log.Debug("Not global context")
					services = m.Context.GetServices()
				}

				for _, service := range services {
					log.Debug("Running service " + service)
					err = m.Plugins.RunGen([]string{service}, m.Config.Services[service].Plugins)
					if err != nil {
						log.Error(err)
					}
				}

				err = m.Plugins.RunGen(m.Config.ServiceNames(), m.Config.Plugins)
				if err != nil {
					log.Error(err)
				}

				return nil
			},
		}, {
			Name:  "template",
			Usage: "shows information about template",
			Action: func(c *cli.Context) error {
				bootstrap(c)

				if c.NArg() == 0 {
					return utils.ErrNoArg
				}

				log.Info(utils.GetTemplateHelp(c.Args().First()))

				return nil
			},
		},
		{
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
						bootstrap(c)

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

						m.Config.Services[c.Args().First()] = config.Service{
							Name:        c.Args().First(),
							Description: c.String("d"),
						}

						err = generator.GenerateService(generator.Service{
							Name:    c.Args().First(),
							Project: m.Config.Name,
						},
							c.StringSlice("x"),
							c.String("t"),
							".",
						)

						if err != nil {
							log.Fatal(err)
							return nil
						}

						if err := config.SaveConfig(m.Config); err != nil {
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

						if c.NArg() == 0 {
							return utils.ErrNoArg
						}

						m, err := GetMmo()
						if err != nil {
							return utils.ErrNoProject
						}

						m.Config.AddPlugin(c.Args().First())
						err = config.SaveConfig(m.Config)
						if err != nil {
							log.Error(err)
						}

						return nil
					},
				},
			},
		},
		{
			Name:  "kube",
			Usage: "manages Kubernetes resources",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "adds kubernetes resource for service",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "wizard, w",
							Usage: "wizard for creating service",
						},
						cli.StringFlag{
							Name:  "template, t",
							Usage: "template folder for service generation",
							Value: build.Default.GOPATH + "/src/github.com/flowup/mmo/templates/kubernetes",
						},
					},
					Action: func(c *cli.Context) error {
						bootstrap(c)

						if c.NArg() == 0 {
							return utils.ErrNoArg
						}

						m, err := GetMmo()
						if err != nil {
							return utils.ErrNoProject
						}

						service, ok := m.Config.Services[c.Args().First()]
						if !ok {
							log.Warnln("Service doesn't exist")
							return nil
						}

						options := make(map[string]interface{})
						options["Name"] = c.Args().First()
						options["Project"] = m.Config.Name
						options["k"] = kubernetes.FromPlugins(service.Plugins)

						err = generator.Generate(options, c.String("t"), "./infrastructure/services")

						if err != nil {
							log.Warnln(err)
							return nil
						}

						return nil
					},
				},
				{
					Name:  "extend",
					Usage: "extends existing resource of the service",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "template, t",
							Usage: "template folder for service generation",
							Value: build.Default.GOPATH + "/src/github.com/flowup/mmo/templates/go-service",
						},
					},
					Action: func(c *cli.Context) error {
						bootstrap(c)

						if c.NArg() == 0 {
							return utils.ErrNoArg
						}

						_, err := GetMmo()
						if err != nil {
							log.Fatal(err)
							return nil
						}

						return nil
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
