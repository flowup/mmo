package main

import (
	"github.com/flowup/mmo/config"
	"gopkg.in/yaml.v2"
	"fmt"
)

func main() {
	cfg := config.Config{Name: "example", Lang: "go", DepManager: "glide", GoPackage: "github.com/example",
	Services:                  make([]config.Service, 0)}

	exampleService := config.Service{Name: "auth", Description: "description", WebRPC: true, Dependencies: make([]config.Dependency, 0)}

	exampleDep := config.Dependency{Name: "base", Type: "grpc", Run: config.DependencyRun{Name: "base", Location: "local"}}

	exampleService.Dependencies = append(exampleService.Dependencies, exampleDep)

	cfg.Services = append(cfg.Services, exampleService)

	b,err := yaml.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
