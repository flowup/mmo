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
			Usage:   "",
			Action: func(c *cli.Context) error {
				if c.Args().First() == ""{
					fmt.Println("No arguments")
					return utils.ErrNoArg
				}

				if err := services.Project(c.Args().First()); err != nil {
					fmt.Println(err)
					return err
				}
				return nil
			},

		},

		{
			Name:    "service",
			Aliases: []string{},
			Usage:   "",
			Action: func(c *cli.Context) error {
				fmt.Println("add not implemented yet")
				return utils.ErrNotImplemented

			},

		},

		{
			Name:    "set-contex",
			Aliases: []string{},
			Usage:   "",
			Action: func(c *cli.Context) error {
				fmt.Println("add not implemented yet")
				return utils.ErrNotImplemented

			},

		},

		{
			Name:    "dev",
			Aliases: []string{},
			Usage:   "",
			Action: func(c *cli.Context) error {
				fmt.Println("add not implemented yet")
				return utils.ErrNotImplemented

			},

		},

		{
			Name:    "build",
			Aliases: []string{},
			Usage:   "",
			Action: func(c *cli.Context) error {
				fmt.Println("add not implemented yet")
				return utils.ErrNotImplemented

			},

		},

		{
			Name:    "e2e",
			Aliases: []string{},
			Usage:   "",
			Action: func(c *cli.Context) error {
				fmt.Println("add not implemented yet")
				return utils.ErrNotImplemented

			},

		},


		{
			Name:    "deploy",
			Aliases: []string{},
			Usage:   "",
			Action: func(c *cli.Context) error {
				fmt.Println("add not implemented yet")
				return utils.ErrNotImplemented

			},

		},


		{
			Name:    "add",
			Aliases: []string{},
			Usage:   "",
			Action: func(c *cli.Context) error {
				fmt.Println("add not implemented yet")
				return utils.ErrNotImplemented

			},

		},


	}

	app.Run(os.Args)
}
