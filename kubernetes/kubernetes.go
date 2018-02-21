package kubernetes

import (
	"github.com/flowup/mmo/docker"
)

type Kubernetes struct {
	Ports     []Port
	Variables []EnvVar
}

type EnvVar struct {
	Name  string
	Value string
}

type Port struct {
	Name   string
	Number string
}

func New() *Kubernetes {
	return &Kubernetes{
		Ports:     make([]Port, 0),
		Variables: make([]EnvVar, 0),
	}
}

func FromPlugins(plugins []string) *Kubernetes {
	k := New()

	for _, plugin := range plugins {
		image := docker.ImageFromString(plugin)
		if image.Name == "flowup/mmo-gen-go-grpc" {
			k.Ports = append(k.Ports, Port{Name: "grpc", Number: "50051"})
		}
		if image.Name == "flowup/mmo-gen-grpc-gateway" {
			k.Ports = append(k.Ports, Port{Name: "http-gateway", Number: "50080"})
		}
	}

	return k
}
