package service

import (
	"github.com/flowup/mmo/config"
	"bufio"
	"os"
	"fmt"
)

// Wizzar for setup service
func Wizzar(serviceName string) config.Service {
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
		newService.Dsn = text[:len(text)-1]
	}

	return newService
}
