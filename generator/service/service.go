package service

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

var templateSource = flag.String("template", "template", "Defines path to directory with service templates")
var output = flag.String("output", ".", "Defines output path")

type Service struct {
	Name    string
	Package string
	Project string
}

func Generate(service Service, tmpl string, out string) error {

	if (tmpl)[len(tmpl)-1] != '/' {
		tmpl += "/"
	}

	if (out)[len(out)-1] != '/' {
		out += "/"
	}

	err := filepath.Walk(tmpl, func(path string, info os.FileInfo, err error) error {
		filename := path[len(*templateSource):]

		pathBytes := &bytes.Buffer{}
		fileBytes := &bytes.Buffer{}

		template.Must(template.New("").Parse(filename)).Execute(pathBytes, &service)

		pathTemplated := filepath.Join(*output, service.Name, pathBytes.String())
		if info.IsDir() {
			os.Mkdir(pathTemplated, 0755)
			fmt.Println(filepath.Join(*output, service.Name, pathBytes.String()))
			return nil
		}

		fileTemplate, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		template.Must(template.New("").Parse(string(fileTemplate))).Execute(fileBytes, &service)

		err = ioutil.WriteFile(pathTemplated, fileBytes.Bytes(), 0755)

		return err
	})

	return err
}

func main() {
	flag.Parse()
}
