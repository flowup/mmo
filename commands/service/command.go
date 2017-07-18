package service

import (
	"github.com/flowup/mmo/commands"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/utils"
	"github.com/flowup/mmo/utils/cookiecutter"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InitService is cli function to generate service with given name
func InitService(configService config.Service) error {
	// go through assets and generate them
	for name, assetGetter := range _bindata {

		asset, err := assetGetter()
		if err != nil {
			return err
		}

		log.Debugln("Restoring file:", name)
		err = cookiecutter.Restore(
			configService.Name,
			strings.Replace(name, "_go", ".go", 1),
			string(asset.bytes),
			configService,
			utils.DefaultFuncMap,
		)

		if err != nil {
			return err
		}
	}

	if err := addGoImportManager("./"); err != nil {
		return errors.Wrap(err, "Failed to run import manager")
	}

	if err := commands.GenerateProto(commands.Go, configService.Name); err != nil {
		return errors.Wrap(err, "Failed to generate API servers and stubs from proto")
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
