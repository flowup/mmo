package {{.Name}}

import (
	"github.com/stretchr/testify/suite"
    "testing"
)

type {{.Name | Title}}ServiceSuite struct {
	suite.Suite

	service *Service
}

func (s *{{.Name | Title}}ServiceSuite) SetupSuite() {
}

func (s *{{.Name | Title}}ServiceSuite) SetupTest() {

}

// Test
func (s *{{.Name | Title}}ServiceSuite) Test() {

}

func Test{{.Name | Title}}ServiceSuite(t *testing.T) {
	suite.Run(t, new({{.Name | Title}}ServiceSuite))
}

func (s *{{.Name | Title}}ServiceSuite) TearDownSuite() {

}
