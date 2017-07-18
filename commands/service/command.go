package service

import (
	"github.com/flowup/mmo/utils"
	"strings"
	"os"
	"github.com/flowup/mmo/commands"
	"text/template"
	"github.com/flowup/mmo/config"
	"os/exec"
	"path/filepath"
	"bytes"
	log "github.com/sirupsen/logrus"
)

// InitService is cli function to generate service with given name
func InitService(configService config.Service) error {
	// create proto dir
	if err := utils.CreateDir(configService.Name + "/protobuf"); err != nil {
		return err
	}

	// create main dir
	if err := utils.CreateDir(configService.Name + "/cmd/" + configService.Name); err != nil {
		return err
	}

	// go through assets and generate them
	for name, assetGetter := range _bindata {

		asset, err := assetGetter()
		if err != nil {
			return err
		}

		// get the file path to the directory
		filePath := strings.Replace(name, "commands/service/template/", "", 1)
		// add the config service name to the asset path
		filePath = configService.Name + "/" + strings.Replace(filePath, "_go", ".go", 1)

		// template the path
		assetTmpl := template.Must(template.New("assetPath").Parse(filePath))
		buf := &bytes.Buffer{}
		assetTmpl.Execute(buf, configService)

		assetPath := buf.String()
		log.Debugln("New asset path was templated:", assetPath)

		// create the directory if it doesn't exist
		if stat, err := os.Stat(filepath.Dir(assetPath)); err != nil || !stat.IsDir() {
			err := os.MkdirAll(filepath.Dir(assetPath), os.ModePerm)
			if err != nil {
				return err
			}
		}

		// create template for the file
		tmpl := template.Must(template.New(name).Parse(string(asset.bytes)))

		log.Debugln("Creating new file:", assetPath)
		// create the file in path
		file, err := os.Create(assetPath)
		if err != nil {
			return err
		}
		defer file.Close()

		// execute the template to the file
		err = tmpl.Execute(file, configService)
		if err != nil {
			return err
		}
	}

	if err := addGoImportManager("./"); err != nil {
		return err
	}

	if err := commands.GenerateProto(commands.Go, configService.Name); err != nil {
		return err
	}

	return nil
}

func addGoImportManager(root string) error {
	goImportsInstallCmd := exec.Command("go", "get", "golang.org/x/tools/cmd/goimports")
	goImportsInstallCmd.Stdout = os.Stdout
	goImportsInstallCmd.Stderr = os.Stdout
	if err := goImportsInstallCmd.Run(); err != nil {
		return err
	}

	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if !strings.HasSuffix(info.Name(), ".go") || info.IsDir() {
			return nil
		}

		/*out := exec.Command("goimports", path)

		if err := out.Run(); err != nil {
			return errors.Wrap(err, "goimports failed")
		}

		if err := utils.WriteFile(path, string(out)); err != nil {
			return err
		}*/

		return nil
	}); err != nil {
		return err
	}

	return nil

}
