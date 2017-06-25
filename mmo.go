package main

import (
	"os"
	"github.com/urfave/cli"
	"fmt"
	"github.com/flowup/mmo/services"
	"github.com/flowup/mmo/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "mmo"
	app.Usage = ""

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 1 {
			return utils.ErrNoArg
		}

		switch (c.Args().Get(0)) {

		case "project":
			if (c.NArg() == 2) {
				if err := services.Project(c.Args().Get(1)); err != nil {
					fmt.Println(err)
					return err
				}
				return nil
			}
			fmt.Println("project not implemented yet")
			return utils.ErrNotImplemented

		case "service":
			fmt.Println("service not implemented yet")
			return utils.ErrNotImplemented

		case "set-context":
			fmt.Println("set-context not implemented yet")
			return utils.ErrNotImplemented

		case "dev":
			fmt.Println("dev not implemented yet")
			return utils.ErrNotImplemented

		case "build":
			fmt.Println("build not implemented yet")
			return utils.ErrNotImplemented

		case "e2e":
			fmt.Println("e2e not implemented yet")
			return utils.ErrNotImplemented

		case "deploy":
			fmt.Println("deploy not implemented yet")
			return utils.ErrNotImplemented

		case "add":
			fmt.Println("add not implemented yet")
			return utils.ErrNotImplemented

		}
		return nil
	}
	app.Run(os.Args)
}
