package service

import (
	"bytes"
	"errors"
	"flag"
	"github.com/flowup/mmo/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var templateSource = flag.String("template", "template", "Defines path to directory with service templates")
var output = flag.String("output", ".", "Defines output path")

type Service struct {
	Name    string
	Package string
	Project string
	Opts    map[string]string
}

func Generate(service Service, opts []string, tmpl string, out string) error {

	if tmpl == "" {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			return errors.New("error: set your $GOPATH or explicitly define template with -t option")
		}

		tmpl = filepath.Join(gopath, "src/github.com/flowup/mmo/templates/go-service")
	}

	options, err := ParseOptions(opts)
	if err != nil {
		return err
	}

	service.Opts = options

	if (tmpl)[len(tmpl)-1] != '/' {
		tmpl += "/"
	}

	if (out)[len(out)-1] != '/' {
		out += "/"
	}

	err = filepath.Walk(tmpl, func(path string, info os.FileInfo, err error) error {
		filename := path[len(tmpl):]

		pathBytes := &bytes.Buffer{}
		fileBytes := &bytes.Buffer{}

		template.Must(template.New("").Parse(filename)).Execute(pathBytes, &service)

		pathTemplated := filepath.Join(*output, service.Name, pathBytes.String())
		if info.IsDir() {
			os.Mkdir(pathTemplated, 0755)
			return nil
		}

		fileTemplate, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		template.Must(
			template.New("").
				Funcs(utils.DefaultFuncMap).
				Parse(string(fileTemplate)),
		).Execute(fileBytes, &service)

		err = ioutil.WriteFile(pathTemplated, fileBytes.Bytes(), 0755)

		return err
	})

	return err
}

func ParseOptions(opts []string) (map[string]string, error) {
	out := make(map[string]string)
	for _, opt := range opts {
		keyval := strings.Split(opt, "=")
		if len(keyval) == 1 {
			return nil, errors.New("additional option " + opt + " is not key value pair")
		}

		out[keyval[0]] = keyval[1]
	}

	return out, nil
}
