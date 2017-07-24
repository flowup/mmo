package project

import (
	"bufio"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/flowup/mmo/commands"
	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/docker"
	"github.com/flowup/mmo/kubernetes"
	"github.com/flowup/mmo/utils"
	"github.com/flowup/mmo/utils/cookiecutter"
	"github.com/flowup/mmo/utils/dockercmd"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
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

		if _, err := os.Stat(name); err == nil {
			// ask user if he want's to rewrite if not rewritten
			if !rewriteAll {
				log.Info("File ", name, " already exists. Do you want to rewrite it? [y/yy/n]:")

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

			log.Infoln("Rewriting file:", name)
			err = os.Remove(name)
			if err != nil {
				return err
			}
		}

		log.Debugln("Creating file:", name)
		err = cookiecutter.Restore(
			mmo.Config.Name,
			name,
			string(asset.bytes),
			mmo.Config,
			utils.DefaultFuncMap,
		)
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
	log.Debugln("Initializing dep. manager:", mmo.Config.DepManager)
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
	log.Debugln("Clearning dep. manager:", mmo.Config.DepManager)
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
	log.Debugln("Trying to set context for services:", services)
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
func (mmo *Mmo) ProtoGen(services []string, lang string) error {

	for _, serviceName := range services {
		log.Infoln("Generating protobuf for:", serviceName)

		if _, err := os.Stat(serviceName + "/protobuf"); os.IsNotExist(err) {
			log.Warnln("No protobuf files found for service:", serviceName, " -> Skipping")
			continue
		}

		if mmo.Config.Services[serviceName].WebRPC {
			if _, err := os.Stat(serviceName + "/sdk"); os.IsNotExist(err) {
				err := os.Mkdir(serviceName+"/sdk", os.ModePerm)
				if err != nil {
					return err
				}
			}
		}

		err := commands.GenerateProto(lang, serviceName)
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

	log.Infoln("Connecting to the kubernetes cluster")
	kubeClient, err := kubernetes.ConnectToCluster()
	if err != nil {
		return err
	}

	log.Debugln("Checking container registry status")
	err = kubernetes.IsRegistryRunning(kubeClient)
	if err != nil {
		log.Infoln("Registry not running, deploying now")
		err = kubernetes.DeployDockerRegistry(kubeClient)
		if err != nil {
			return err
		}
	}

	log.Infoln("Forwarding registry port to localhost")
	portFwdCmd, err := kubernetes.ForwardRegistryPort()
	defer func() {
		log.Debugln("Cleaning port forwarding for docker registry")
		err := portFwdCmd.Process.Kill()
		if err != nil {
			log.Errorln(err)
		}
	}()

	log.Debugln("Getting builder for the repository:", mmo.Config.Name)
	builder, err := docker.GetBuilder(mmo.Config.Name)
	if err != nil {
		return err
	}
	defer func() {
		log.Debugln("Cleaning docker builder images")
		err := builder.Clean()
		if err != nil {
			log.Errorln(err)
		}
	}()

	var env = make(kubernetes.DeployEnvironment)
	env["DOCKER_REGISTRY"] = dockercmd.MinikubeRegistry
	env["PROJECT_NAME"] = mmo.Config.Name

	for service := range mmo.Config.Services {

		log.Infoln("Building service:", service)
		image, err := builder.BuildService(service)
		if err != nil {
			return err
		}

		log.Infoln("Pushing image:", image.GetFullname())
		err = builder.PushService(image)
		if err != nil {
			return err
		}

		env["SERVICE"] = service
		env["WERCKER_GIT_COMMIT"] = image.Tag

		log.Debugln("Expanding templates for service:", service)
		err = kubernetes.DeployService(kubeClient, env)
		if err != nil {
			return err
		}
	}

	return nil
}
