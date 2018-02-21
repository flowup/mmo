package api

import (
	"github.com/flowup/mmo/config"
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
		result[i] = &Plugin{Name: plugin}
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
		result[i] = &Plugin{Name: plugin}
		i++
	}

	return &Plugins{
		Plugins: result,
	}, nil
}
