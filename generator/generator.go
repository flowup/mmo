package generator

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
	"google.golang.org/api/option"
	"cloud.google.com/go/storage"
	"context"
)

const bucketOutFolder = "/tmp/bucket/"
var output = flag.String("output", ".", "Defines output path")

type Service struct {
	Name    string
	Package string
	Project string
}

type Project struct {
	Name string
}

func GenerateProject(project Project, opts []string, tmpl string, out string) error {
	options, err := ParseOptions(opts)
	if err != nil {
		return err
	}

	options["Name"] = project.Name

	return Generate(options, tmpl, out)
}

func GenerateService(service Service, opts []string, tmpl string, out string) error {
	options, err := ParseOptions(opts)
	if err != nil {
		return err
	}

	options["Name"] = service.Name
	options["Package"] = service.Package
	options["Project"] = service.Project

	return Generate(options, tmpl, out)
}

func Generate(options map[string]string, tmpl string, out string) error {

	var err error
	// Download template from google cloud storage
	if strings.HasPrefix(tmpl, "gs://") {
		client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(os.Getenv("HOME") + "/cloudkey.json"))
		if err != nil {
			return err
		}
		splitted := strings.SplitN(strings.TrimPrefix(tmpl,"gs://"), "/", 2)
		if len(splitted) != 2 {
			return errors.New("invalid gs uri, should be gs://<bucket_name>/<path>")
		}
		err = utils.RecursiveCopy(client, splitted[0], splitted[1], bucketOutFolder)
		if err != nil {
			return err
		}
		// Modify template path to newly extracted
		tmpl = bucketOutFolder

	}
	if (tmpl)[len(tmpl)-1] != '/' {
		tmpl += "/"
	}

	if out == "" {
		out = *output
	}

	if (out)[len(out)-1] != '/' {
		out += "/"
	}

	skipFiles := make(map[string]bool)

	err = filepath.Walk(tmpl, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		filename := path[len(tmpl):]

		for key := range skipFiles {
			if strings.HasPrefix(filename, key) {
				return nil
			}
		}

		pathBytes := &bytes.Buffer{}
		fileBytes := &bytes.Buffer{}

		template.Must(template.New("").Parse(filename)).Execute(pathBytes, options)

		pathTemplated := filepath.Join(*output, pathBytes.String())
		if info.IsDir() {
			if pathBytes.Len() == 0 {
				skipFiles[filename+"/"] = true
				return nil
			}
			return os.Mkdir(pathTemplated, 0755)
		}

		fileTemplate, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		template.Must(
			template.New("").
				Funcs(utils.DefaultFuncMap).
				Parse(string(fileTemplate)),
		).Execute(fileBytes, options)

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
