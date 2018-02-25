package api

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/flowup/mmo/config"
	"github.com/flowup/mmo/docker"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
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

	result := make([]*KubernetesConfig, 0)

	err := filepath.Walk("./infrastructure", func(path string, info os.FileInfo, err error) error {
		logrus.Debugln("Walking file", info.Name(), "in path", path)
		if info.IsDir() {
			return nil
		}

		if !(strings.HasSuffix(info.Name(), ".yaml") ||
			strings.HasSuffix(info.Name(), ".yml")) {
			return nil
		}

		if !strings.HasPrefix(info.Name(), in.Name+"-") {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrap(err, "Failed to read kubernetes config")
		}

		kType := struct {
			Kind string `yaml:"Kind"`
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

func (s *APIService) KubernetesConfigFromForm(ctx context.Context, in *KubernetesServiceForm) (*KubernetesConfigs, error) {
	logrus.Debugln("Generating kubernetes configs... ", in)
	return nil, nil
}

//func (s *APIService) GetPlugins(ctx context.Context, in *Service) (*Plugins, error) {}
