package main

import (
	"github.com/flowup/mmo/docker"
	"log"
)

func main() {
	_, err := docker.CreateClient()
	if err != nil {
		log.Fatal(err)
	}
}
