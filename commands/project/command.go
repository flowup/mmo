package project

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/commands"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/utils"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"strings"
	"text/template"
	log "github.com/sirupsen/logrus"
	"bufio"
)

// Mmo represents config and context
type Mmo struct {
	Config  *config.Config
	Context *config.Context
}

// GetMmo Load context and config from config files
func GetMmo() (*Mmo, error) {

	mmo := &Mmo{
		&config.Config{},
		&config.Context{},
	}

	mmoContext, err := config.LoadContext()
	if err == nil {
		mmo.Context = mmoContext
	}

	mmoConfig, err := config.LoadConfig(config.FilenameConfig)
	if err != nil {
		return nil, err
	}
	mmo.Config = mmoConfig

	return mmo, nil
}

// InitProject extends all assets using project options passed by the caller
// This automatically creates a project folder with all files
func (mmo *Mmo) InitProject() error {
	rewriteAll := false

	// go through assets and generate them
	for name, assetGetter := range _bindata {
		asset, err := assetGetter()
		if err != nil {
			return err
		}

		// get correct path to the handle
		filePath := strings.Replace(asset.info.Name(), "commands/project/template", mmo.Config.Name, 1)
		// create template for the handle
		tmpl := template.Must(template.New(name).Parse(string(asset.bytes)))

		var handle *os.File

		if _, err := os.Stat(filePath); err == nil {
			// ask user if he want's to rewrite if not rewritten
			if !rewriteAll {
				log.Info("File ", filePath, " already exists. Do you want to rewrite it? [y/yy/n]:")

				reader := bufio.NewReader(os.Stdin)
				answer := ""
				for answer != "y" && answer != "n" && answer != "yy" {
					answer, _ = reader.ReadString('\n')
					answer = strings.Trim(answer, "\n")
				}

				// rewrites all files automatically
				if answer == "yy" {
					rewriteAll = true
				}
			}

			log.Infoln("Rewriting file:", filePath)
			err = os.Remove(filePath)
			if err != nil {
				log.Info("ERROGING")
				return err
			}
		}

		log.Debugln("Creating file", filePath)
		// create the handle in path
		handle, err = os.Create(filePath)
		if err != nil {
			return err
		}

		// execute the template to the handle
		err = tmpl.Execute(handle, mmo.Config)
		if err != nil {
			return err
		}

		err = handle.Close()
		if err != nil {
			return err
		}
	}

	// change to the newly created project and init the dep manager
	err := os.Chdir(mmo.Config.Name)
	if err != nil {
		return errors.Wrap(err, "chdir failed with an error")
	}

	err = mmo.ClearDependencyManager()
	if err != nil {
		return err
	}

	err = mmo.InitializeDependencyManager()
	if err != nil {
		return errors.Wrap(err, "dependency manager initialization failed")
	}

	err = os.Chdir("..")
	if err != nil {
		return errors.Wrap(err, "chdir back failed with an error")
	}

	return nil
}

// InitializeDependencyManager initializes given dependency manager
// within the current project.
// It will also automatically update the dependency manager
func (mmo *Mmo) InitializeDependencyManager() error {
	switch mmo.Config.DepManager {
	case "glide":
		glideInstallCmd := exec.Command("go", "get", "github.com/Masterminds/glide")
		glideInstallCmd.Stdout = os.Stdout
		glideInstallCmd.Stderr = os.Stdout
		if err := glideInstallCmd.Run(); err != nil {
			return errors.Wrap(err, "glide installation failed")
		}

		glideInitCmd := exec.Command("glide", "init", "--non-interactive")
		glideInitCmd.Stdout = os.Stdout
		glideInitCmd.Stderr = os.Stdout
		if err := glideInitCmd.Run(); err != nil {
			return errors.Wrap(err, "glide initialization failed")
		}

	default:
		return errors.New("Unrecognized dependency manager: " + mmo.Config.DepManager)
	}

	return nil
}

// ClearDependencyManager clears contents of the content manager
func (mmo *Mmo) ClearDependencyManager() error {
	switch mmo.Config.DepManager {
	case "glide":
		// remove mail glide file
		os.Remove("glide.yaml")
		// cache removal
		os.RemoveAll(".glide/")

	default:
		return errors.New("Unrecognized dependency manager: " + mmo.Config.DepManager)
	}

	return nil
}

// RunTests is cli function to run tests of application in docker
func (mmo *Mmo) RunTests() error {

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, serviceName := range mmo.Context.Services {

		log.Infoln("Running tests for service \"" + serviceName + "\":")

		testContainer, err := cli.ContainerCreate(context.Background(), &container.Config{
			Image:      "flowup/mmo-webrpc",
			Cmd:        []string{"bash", "-c", "go test $(glide novendor)"},
			WorkingDir: "/go/src/" + mmo.Config.GoPackage + "/" + serviceName,
		}, &container.HostConfig{
			AutoRemove: true,
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: pwd,
					Target: "/go/src/" + mmo.Config.GoPackage,
				},
			},
		}, nil, "")

		if err != nil {
			return err
		}

		err = utils.ContainerRunStdout(cli, testContainer.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// SetContext is cli function to set context of mmo to specified service or services
func (mmo *Mmo) SetContext(services []string) error {
	for _, service := range services {
		if _, ok := mmo.Config.Services[service]; !ok {
			return utils.ErrServiceNotExists
		}

		if _, err := os.Stat(service); os.IsNotExist(err) {
			return errors.Wrap(utils.ErrServiceNotExists, service)
		}
	}

	serviceContext := &config.Context{
		Services: services,
	}

	err := config.SaveContext(serviceContext)

	return err
}

// ProtoGen is cli function to generate API clients and server stubs of specified service or services
func (mmo *Mmo) ProtoGen(services []string) error {

	for _, serviceName := range services {
		log.Infoln("Generating protobuf for:", serviceName)

		if _, err := os.Stat(serviceName + "/protobuf"); os.IsNotExist(err) {
			log.Warnln("No protobuf files found for service:", serviceName, " -> Skipping")
			continue
		}

		if _, err := os.Stat(serviceName + "/sdk"); os.IsNotExist(err) {
			os.Mkdir(serviceName+"/sdk", os.ModePerm)
		}

		err := commands.GenerateProto(mmo.Config.Lang, serviceName)
		if err != nil {
			return err
		}

		if mmo.Config.Services[serviceName].WebRPC {
			err = commands.GenerateProto(commands.TypeScript, serviceName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
