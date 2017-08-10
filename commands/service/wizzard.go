package service

import (
	"bufio"
	"fmt"
	"github.com/flowup/mmo/config"
	"github.com/urfave/cli"
	"os"
)

// Wizzard for setup service
func Wizzard(serviceName string) config.Service {
	newService := config.Service{Name: serviceName}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Description : ")
	if text, _ := reader.ReadString('\n'); text != "\n" {
		newService.Description = text[:len(text)-1]
	}

	fmt.Print("WebRPC [y]: ")
	if text, _ := reader.ReadString('\n'); text == "y\n" {
		newService.WebRPC = true
	}

	fmt.Print("Dsn: ")
	if text, _ := reader.ReadString('\n'); text != "\n" {
		newService.Sentry = text == "y\n"
	}

	return newService
}

// FromCliContext Create new service according to flags
func FromCliContext(serviceName string, ctx *cli.Context) config.Service {
	return config.Service{
		Name:        serviceName,
		Description: ctx.String("description"),
		WebRPC:      ctx.Bool("webrpc"),
		Sentry:      ctx.Bool("sentry"),
		Gateway:     ctx.Bool("gateway"),
	}
}
