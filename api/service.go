package api

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/docker"
	"github.com/flowup/mmo/generator"
	"github.com/flowup/mmo/kubernetes"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// Api represents an implementation of the service interface
type APIService struct {
	Config *config.Config
}

// NewApi creates a new service object
func NewAPIService(config *config.Config) *APIService {
	return &APIService{Config: config}
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
	services, err := ioutil.ReadDir("./infrastructure/services/")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read kubernetes configs")
	}

	deployments, err := ioutil.ReadDir("./infrastructure/deployments/")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read kubernetes configs")
	}

	result := make([]*KubernetesConfig, 0)

	for _, file := range deployments {
		if strings.HasPrefix(file.Name(), in.Name+"-") {
			data, err := ioutil.ReadFile("./infrastructure/deployments/" + file.Name())
			if err != nil {
				return nil, errors.Wrap(err, "Failed to read kubernetes config")
			}

			result = append(result, &KubernetesConfig{Name: file.Name(), Data: string(data)})
		}
	}

	for _, file := range services {
		if strings.HasPrefix(file.Name(), in.Name+"-") {
			data, err := ioutil.ReadFile("./infrastructure/services/" + file.Name())
			if err != nil {
				return nil, errors.Wrap(err, "Failed to read kubernetes config")
			}

			result = append(result, &KubernetesConfig{Name: file.Name(), Data: string(data)})
		}
	}

	return &KubernetesConfigs{Configs: result}, nil
}

func (s *APIService) KubernetesConfigFromPlugins(ctx context.Context, in *Service) (*KubernetesConfigs, error) {

	mmoService, ok := s.Config.Services[in.Name]
	if !ok {
		return nil, errors.New("Service doesn't exist")
	}

	options := make(map[string]interface{})
	options["Name"] = in.Name
	options["Project"] = s.Config.Name
	options["k"] = kubernetes.FromPlugins(mmoService.Plugins)

	err := generator.Generate(options, os.Getenv("GOPATH")+"/src/github.com/flowup/mmo/templates/kubernetes", ".")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to generate kubernetes configs")
	}

	return &KubernetesConfigs{}, nil
}

func (s *APIService) KubernetesConfigFromForm(ctx context.Context, in *KubernetesServiceForm) (*KubernetesConfigs, error) {
	return nil, nil
}

//func (s *APIService) GetPlugins(ctx context.Context, in *Service) (*Plugins, error) {}
