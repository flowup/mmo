package config

import (
	"os"
	"encoding/json"
	"io/ioutil"
)

const (
	filenameContext = ".mmo.cache"
)

// Context represents a cached configuration for the project
type Context struct {
	Services []string
}

// LoadContext loads project context from the given directory
func LoadContext() (*Context, error) {
	b, err := ioutil.ReadFile(filenameContext)
	if err != nil {
		return &Context{}, err
	}

	context := &Context{}
	return context, json.Unmarshal(b, &context)
}

// SaveContext saves given context to the current path
func SaveContext(context *Context) error {
	b, err := json.Marshal(context)

	if err != nil {
		return nil
	}

	f, err := os.Create(filenameContext)
	_, err = f.Write(b)

	return err
}