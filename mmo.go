package main

import (
	"os"
	"github.com/urfave/cli"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "mmo"
	app.Usage = ""

	app.Action = func(c *cli.Context) error {
		if c.NArg() < 0 {
			return ErrNoArg
		}

		switch (c.Args().Get(0)) {

		case "project":
			fmt.Println("project not implemented yet")
			return ErrNotImplemented

		case "service":
			fmt.Println("service not implemented yet")
			return ErrNotImplemented

		case "set-context":
			fmt.Println("set-context not implemented yet")
			return ErrNotImplemented

		case "dev":
			fmt.Println("dev not implemented yet")
			return ErrNotImplemented

		case "build":
			fmt.Println("build not implemented yet")
			return ErrNotImplemented

		case "e2e":
			fmt.Println("e2e not implemented yet")
			return ErrNotImplemented

		case "deploy":
			fmt.Println("deploy not implemented yet")
			return ErrNotImplemented

		case "add":
			fmt.Println("add not implemented yet")
			return ErrNotImplemented

		}
		return nil
	}
	app.Run(os.Args)
}
