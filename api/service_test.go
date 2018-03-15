package api

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ApiServiceSuite struct {
	suite.Suite

	service *Service
}

func (s *ApiServiceSuite) SetupSuite() {
}

func (s *ApiServiceSuite) SetupTest() {

}

// Test
func (s *ApiServiceSuite) Test() {

}

func TestApiServiceSuite(t *testing.T) {
	suite.Run(t, new(ApiServiceSuite))
}

func (s *ApiServiceSuite) TearDownSuite() {

}
