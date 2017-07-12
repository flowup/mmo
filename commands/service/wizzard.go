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
	fmt.Print("WebRPC [y]: ")
	text, _ := reader.ReadString('\n')
	if text == "y\n" {
		newService.WebRPC = true
	}

	fmt.Print("Dsn: ")
	text, _ = reader.ReadString('\n')
	newService.Dsn = text

	return newService
}
