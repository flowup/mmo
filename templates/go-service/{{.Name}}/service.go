package {{.Name}}

import (
    "golang.org/x/net/context"
    google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

// {{.Name | Title}} represents an implementation of the service interface
type Service struct {
}

// New{{.Name | Title}} creates a new service object
func NewService() *Service {
    return &Service{}
}

func (s *Service) GetVersion(ctx context.Context, in *google_protobuf.Empty) (*Version, error) {
    return &Version{
        Name: "1.0.0"
    }, nil
}
