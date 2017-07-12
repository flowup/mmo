package config

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func (s *ConfigSuite) TestMarshalConfig() {

	conf, err := LoadConfig("fixtures/mmo.yaml")

	s.Nil(err)
	s.Equal("example-project-name", conf.Name)
	s.Equal("go", conf.Lang)
	s.Equal("glide", conf.DepManager)
	s.Equal("github.com/flowup/goginn", conf.GoPackage)

	s.Equal("handles authentication and authorization of users", conf.Services["auth"].Description)
	s.True(conf.Services["auth"].WebRPC)

	s.Equal("handles product creation and reading", conf.Services["product"].Description)
	s.False(conf.Services["product"].WebRPC)
	s.Len(conf.Services["product"].Dependencies, 3)

	s.Equal("auth", conf.Services["product"].Dependencies[0].Name)
	s.Equal("grpc", conf.Services["product"].Dependencies[0].Type)
	s.Equal("auth", conf.Services["product"].Dependencies[0].Run.Name)
	s.Equal("local", conf.Services["product"].Dependencies[0].Run.Location)

	s.Equal("auth2", conf.Services["product"].Dependencies[1].Name)
	s.Equal("grpc", conf.Services["product"].Dependencies[1].Type)
	s.Equal("auth", conf.Services["product"].Dependencies[1].Run.Name)
	s.Equal("octo", conf.Services["product"].Dependencies[1].Run.Location)

	s.Equal("db", conf.Services["product"].Dependencies[2].Name)
	s.Equal("postgres", conf.Services["product"].Dependencies[2].Type)
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, &ConfigSuite{})
}
