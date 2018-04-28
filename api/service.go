package api

import (
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/docker"
	"github.com/flowup/mmo/generator"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
)

var (
	kubernetesTemplate = build.Default.GOPATH + "/src/github.com/flowup/mmo/templates/kubernetes"
)

// Api represents an implementation of the service interface
type APIService struct {
	Config       *config.Config
	GithubClient *github.Client
}

// NewApi creates a new service object
func NewAPIService(config *config.Config, githubClient *github.Client) *APIService {
	return &APIService{Config: config, GithubClient: githubClient}
}

func (s *APIService) GetVersion(ctx context.Context, in *google_protobuf.Empty) (*Version, error) {
	return &Version{
		Name: "1.0.0",
	}, nil
}

func (s *APIService) GetServices(ctx context.Context, in *google_protobuf.Empty) (*Services, error) {
	result := make([]*Service, len(s.Config.Services))
	i := 0
	for key, val := range s.Config.Services {
		result[i] = &Service{Name: key, Description: val.Description}
		i++
	}

	return &Services{
		Services: result,
	}, nil
}

func (s *APIService) GetGlobalPlugins(ctx context.Context, in *google_protobuf.Empty) (*Plugins, error) {
	result := make([]*Plugin, len(s.Config.Plugins))
	i := 0
	for _, plugin := range s.Config.Plugins {
		image := docker.ImageFromString(plugin)
		result[i] = &Plugin{Name: image.Registry + image.Name, Version: image.Tag}
		i++
	}

	return &Plugins{
		Plugins: result,
	}, nil
}

func (s *APIService) GetPlugins(ctx context.Context, in *Service) (*Plugins, error) {
	service, ok := s.Config.Services[in.Name]
	if !ok {
		return &Plugins{}, errors.New("Service doesn't exist")
	}

	result := make([]*Plugin, len(service.Plugins))
	i := 0
	for _, plugin := range service.Plugins {
		image := docker.ImageFromString(plugin)
		result[i] = &Plugin{Name: image.Registry + image.Name, Version: image.Tag}
		i++
	}

	return &Plugins{
		Plugins: result,
	}, nil
}

func (s *APIService) GetKubernetesConfigs(ctx context.Context, in *Service) (*KubernetesConfigs, error) {

	result := make([]*KubernetesConfig, 0)

	err := filepath.Walk("./infrastructure", func(path string, info os.FileInfo, err error) error {
		// logrus.Debugln("Walking file", info.Name(), "in path", path)
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if !(strings.HasSuffix(info.Name(), ".yaml") ||
			strings.HasSuffix(info.Name(), ".yml")) {
			return nil
		}

		if !strings.HasPrefix(info.Name(), in.Name) {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrap(err, "Failed to read kubernetes config")
		}

		kType := struct {
			Kind string `yaml:"kind"`
		}{}

		yaml.Unmarshal(data, &kType)
		if kType.Kind == "" {
			kType.Kind = "Invalid Kubernetes config"
		}

		k := &KubernetesConfig{}
		k.Name = info.Name()
		k.Path = path
		k.Data = string(data)
		k.Type = kType.Kind

		result = append(result, k)

		return nil
	})

	return &KubernetesConfigs{Configs: result}, err
}

func (s *APIService) SaveKuberentesConfig(ctx context.Context, in *KubernetesConfig) (*google_protobuf.Empty, error) {
	info, err := os.Stat(in.Path)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(in.Path, []byte(in.Data), info.Mode())
	return &google_protobuf.Empty{}, err
}

func (s *APIService) RemoveKubernetesConfig(ctx context.Context, in *KubernetesConfig) (*google_protobuf.Empty, error) {
	os.Remove(in.Path)
	return &google_protobuf.Empty{}, nil
}

func (s *APIService) KubernetesFormFromPlugins(ctx context.Context, in *Service) (*KubernetesServiceForm, error) {

	mmoService, ok := s.Config.Services[in.Name]
	if !ok {
		return nil, errors.New("Service doesn't exist")
	}

	k := KubernetesServiceForm{
		ServiceName: in.Name,
		Ports:       make([]*KubernetesPort, 0),
	}

	for _, plugin := range mmoService.Plugins {
		image := docker.ImageFromString(plugin)
		if image.Name == "flowup/mmo-gen-go-grpc" {
			k.Ports = append(k.Ports, &KubernetesPort{Name: "grpc", Port: "50051"})
		}
		if image.Name == "flowup/mmo-gen-grpc-gateway" {
			k.Ports = append(k.Ports, &KubernetesPort{Name: "http-gateway", Port: "50080"})
		}
	}

	return &k, nil

	// options := make(map[string]interface{})
	// options["Name"] = in.Name
	// options["Project"] = s.Config.Name
	// options["k"] = kubernetes.FromPlugins(mmoService.Plugins)

	// err := generator.Generate(options, os.Getenv("GOPATH")+"/src/github.com/flowup/mmo/templates/kubernetes", ".")
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Failed to generate kubernetes configs")
	// }

	// return &KubernetesServiceForm{}, nil
}

func (s *APIService) KubernetesConfigFromForm(ctx context.Context, in *KubernetesServiceForm) (*KubernetesServiceForm, error) {
	logrus.Debugln("Generating kubernetes configs... ")
	logrus.Debugln(in.ServiceName)
	logrus.Debugln(in.Ports)
	logrus.Debugln(in.Variables)
	logrus.Debugln(in.Volumes)

	input := make(map[string]interface{})
	input["ServiceName"] = in.ServiceName
	input["ProjectName"] = s.Config.Name
	input["EnvMap"] = in.ConfigEnvConfigmap

	if len(in.Ports) > 0 {
		input["Ports"] = in.Ports
	}
	if len(in.Variables) > 0 {
		input["Variables"] = in.Variables
	}
	if len(in.Volumes) > 0 {
		input["Volumes"] = in.Volumes
	}

	err := generator.Generate(input, kubernetesTemplate, ".")
	if err != nil {
		logrus.Warnln(err)
	}
	return in, err
}

func (s *APIService) GithubDeploy(ctx context.Context, in *GithubDeployRequest) (*google_protobuf.Empty, error) {
	request := &github.DeploymentRequest{
		Environment: github.String(in.Environment),
		Ref:         github.String(in.Ref),
		Description: github.String(in.Message),
	}

	_, _, err := s.GithubClient.Repositories.CreateDeployment(
		context.Background(),
		s.Config.Prefix.GetOwner(),
		s.Config.Prefix.GetRepository(),
		request,
	)

	return &google_protobuf.Empty{}, err
}

func (s *APIService) GetKubernetesClusters(ctx context.Context, in *google_protobuf.Empty) (*KubernetesClusters, error) {
	out, err := exec.Command("kubectl", "config", "get-clusters").CombinedOutput()
	if err != nil {
		return nil, err
	}

	clusters := strings.Split(string(out), "\n")

	result := &KubernetesClusters{
		Clusters:     clusters[1 : len(clusters)-1],
		Environments: []string{},
	}

	files, err := ioutil.ReadDir("./infrastructure")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			result.Environments = append(result.Environments, filepath.Base(f.Name()))
		}
	}

	return result, nil
}

func (s *APIService) KubernetesDeploy(ctx context.Context, in *KubernetesDeployRequest) (*KubernetesConfigs, error) {
	configs := KubernetesConfigs{
		Configs: []*KubernetesConfig{},
	}

	dirs := []string{"shared", in.Environment}

	for _, dir := range dirs {
		err := filepath.Walk("infrastructure/"+dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				logrus.Warnln("error", err)
				return nil
			}

			if info.IsDir() {
				return nil
			}

			k := &KubernetesConfig{}
			k.Name = info.Name()
			k.Path = path

			logrus.Debugln("Appending", k)
			configs.Configs = append(configs.Configs, k)

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return &configs, nil
}

func (s *APIService) ConfirmKubernetesDeploy(ctx context.Context, in *KubernetesDeployRequest) (*ConsoleOutput, error) {
	out, err := s.KubernetesDeploy(ctx, in)
	if err != nil {
		return nil, err
	}

	log := "Switching to selected to cluster...\n"

	currentContext, err := exec.Command("kubectl", "config", "current-context").CombinedOutput()
	if err != nil {
		logrus.Warnln(err)
	}
	console, err := exec.Command("kubectl", "config", "use-context", in.Cluster).CombinedOutput()
	if err != nil {
		logrus.Warnln(err)
	}
	log += string(console) + "\n"

	for _, config := range out.Configs {
		out, err := exec.Command("kubectl", "apply", "-n", in.Namespace, "-f", config.Path).CombinedOutput()
		if err != nil {
			logrus.Warnln("Error deploying", err)
		}

		log += string(out)
	}

	log += "\nSwitching back to previously selected cluster...\n"
	console, err = exec.Command("kubectl", "config", "use-context", strings.Split(string(currentContext), "\n")[0]).CombinedOutput()
	if err != nil {
		logrus.Warnln(err)
	}
	log += string(console)

	return &ConsoleOutput{Output: log}, nil
}

//func (s *APIService) GetPlugins(ctx context.Context, in *Service) (*Plugins, error) {}
