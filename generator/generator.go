package generator

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"cloud.google.com/go/storage"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/utils"
	"google.golang.org/api/option"
)

const bucketOutFolder = "/tmp/bucket/"

var output = flag.String("output", ".", "Defines output path")

// Service represents data needed for service generation
type Service struct {
	Name    string
	Package string
	Project string
}

// GenerateProject is function to initialize project from template
func GenerateProject(config *config.Config, opts []string, tmpl string, out string) error {
	options, err := ParseOptions(opts)
	if err != nil {
		return err
	}

	options["Name"] = config.Name
	options["Package"] = config.Prefix

	return Generate(options, tmpl, out)
}

// GenerateService is function for service generation from template
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

// Generate is function for general generation of files from template (cookiecutter)
func Generate(options map[string]interface{}, tmpl string, out string) error {

	var err error
	templateFromGS := false
	// Download template from google cloud storage
	if strings.HasPrefix(tmpl, "gs://") {
		client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(os.Getenv("HOME")+"/cloudkey.json"))
		if err != nil {
			return err
		}
		splitted := strings.SplitN(strings.TrimPrefix(tmpl, "gs://"), "/", 2)
		if len(splitted) != 2 {
			return errors.New("invalid gs uri, should be gs://<bucket_name>/<path>")
		}
		err = utils.RecursiveCopy(client, splitted[0], splitted[1], bucketOutFolder)
		if err != nil {
			return err
		}
		// Modify template path to newly extracted
		tmpl = bucketOutFolder
		templateFromGS = true

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
	skipFiles["__help.txt"] = true

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
		if pathBytes.String() == filepath.Dir(filename)+"/" {
			return nil
		}

		pathTemplated := filepath.Join(out, pathBytes.String())
		if info.IsDir() {
			if pathBytes.Len() == 0 {
				skipFiles[filename+"/"] = true
				return nil
			}

			dir, err := os.Stat(pathTemplated)
			if err == nil && dir.IsDir() {
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
	if err != nil {
		return err
	}
	// Remove temporary folder used for template
	if templateFromGS {
		err = os.RemoveAll(bucketOutFolder)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

// ParseOptions takes slice of options in format key=value and creates map of these options
func ParseOptions(opts []string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	for _, opt := range opts {
		keyval := strings.Split(opt, "=")
		if len(keyval) == 1 {
			return nil, errors.New("additional option " + opt + " is not key value pair")
		}

		out[keyval[0]] = keyval[1]
	}

	return out, nil
}
