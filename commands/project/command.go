package project

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/commands"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/docker"
	"github.com/flowup/mmo/minikube"
	"github.com/flowup/mmo/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
	"text/template"
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

	// create project folder
	err := os.Mkdir(mmo.Config.Name, os.ModePerm)
	if err != nil {
		return err
	}

	// go through assets and generate them
	for name, assetGetter := range _bindata {
		asset, err := assetGetter()
		if err != nil {
			return err
		}

		// get correct path to the file
		filePath := strings.Replace(asset.info.Name(), "commands/project/template", mmo.Config.Name, 1)
		// create template for the file
		tmpl := template.Must(template.New(name).Parse(string(asset.bytes)))

		// create the file in path
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// execute the template to the file
		err = tmpl.Execute(file, mmo.Config)
		if err != nil {
			return err
		}
	}

	// change to the newly created project and init the dep manager
	err = os.Chdir(mmo.Config.Name)
	if err != nil {
		return err
	}

	err = mmo.InitializeDependencyManager()
	if err != nil {
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
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
		if err := glideInstallCmd.Run(); err != nil {
			return err
		}

		glideInitCmd := exec.Command("glide", "init", "--non-interactive")
		if err := glideInitCmd.Run(); err != nil {
			return err
		}

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

// Run is command to run all services in cluster (minikube) - all services are built as docker images and deployed to cluster
func (mmo *Mmo) Run() error {

	kubeClient, err := minikube.ConnectToCluster()
	if err != nil {
		return err
	}

	err = minikube.IsRegistryRunning(kubeClient)
	if err != nil {
		err = minikube.DeployDockerRegistry(kubeClient)
		if err != nil {
			return err
		}
	}

	portFwdCmd, err := minikube.ForwardRegistryPort()

	builder, err := docker.GetBuilder(mmo.Config.GoPackage)
	if err != nil {
		return err
	}

	for service := range mmo.Config.Services {
		_, err := builder.BuildService(service)
		if err != nil {
			return err
		}
	}

	// TODO: build service

	// TODO: deploy service

	portFwdCmd.Process.Kill()

	return nil
}
