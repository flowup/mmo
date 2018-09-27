package generator

import (
	"bytes"
	"context"
	"flag"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

	// skipped files by the templating engine
	skipFiles := []string{
		"__help.txt",
	}

	// files that should be written as is
	skipTemplating := []string{
		"_helpers.tpl",
		"chart/templates/.*yaml",
	}

	err = filepath.Walk(tmpl, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		filename := path[len(tmpl):]

		if matchInStringSlice(path, skipFiles) {
			return nil
		}

		pathBytes := &bytes.Buffer{}
		fileBytes := &bytes.Buffer{}

		// template the directory
		tpl, err := template.New("").Parse(filename)
		if err != nil {
			return WrapFileError(err, path)
		}

		err = tpl.Execute(pathBytes, options)
		if err != nil {
			return WrapFileError(err, path)
		}

		pathTemplated := filepath.Join(out, pathBytes.String())
		if info.IsDir() {
			if pathBytes.Len() == 0 {
				return nil
			}

			dir, err := os.Stat(pathTemplated)
			if err == nil && dir.IsDir() {
				return nil
			}

			return os.Mkdir(pathTemplated, 0755)
		}

		// read the template from the file
		fileTemplate, err := ioutil.ReadFile(path)
		if err != nil {
			return WrapFileError(err, path)
		}

		if matchInStringSlice(path, skipTemplating) {
			// skipping the templating only writes original contents
			fileBytes.Write(fileTemplate)
		} else {
			contentTpl, err := template.New("").
				Funcs(utils.DefaultFuncMap).
				Parse(string(fileTemplate))
			if err != nil {
				return WrapFileError(err, path)
			}

			err = contentTpl.Execute(fileBytes, options)
			if err != nil {
				return WrapFileError(err, path)
			}
		}

		// write the contents of the template (either not templated or templated)
		err = ioutil.WriteFile(pathTemplated, fileBytes.Bytes(), 0755)
		if err != nil {
			return WrapFileError(err, path)
		}

		return nil
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

func WrapFileError(err error, file string) error {
	return errors.Wrap(err, "An error occurred during templating: "+file)
}

func matchInStringSlice(key string, opts []string) bool {
	for _, rxp := range opts {
		rxp, err := regexp.Compile(rxp)
		if err != nil {
			panic(err)
		}

		if rxp.MatchString(key) {
			return true
		}
	}

	return false
}
