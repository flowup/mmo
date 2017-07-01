package config

import (
	"os"
	"encoding/json"
	"io/ioutil"
)

const (
	filenameContext = ".mmo.cache"
)

type Context struct {
	Services []string
}

func LoadContext() (Context, error) {
	b, err := ioutil.ReadFile(filenameContext)
	if err != nil {
		return Context{}, err
	}

	var context Context
	err = json.Unmarshal(b, &context)

	return context,err
}

func SaveContext(context Context) error {
	b, err := json.Marshal(context)

	if err != nil {
		return nil
	}

	f, err := os.Create(filenameContext)
	_, err = f.Write(b)

	return err
}