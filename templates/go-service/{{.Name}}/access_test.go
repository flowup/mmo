package {{.Name}}

import (
	"github.com/stretchr/testify/suite"
    "testing"
)

type {{.Name | Title}}AccessSuite struct {
	suite.Suite

	access *Access
}

func (s *{{.Name | Title}}AccessSuite) SetupSuite() {
	s.access = NewAccess()
}

func (s *{{.Name | Title}}AccessSuite) SetupTest() {

}

// Test
func (s *{{.Name | Title}}AccessSuite) Test() {

}

func Test{{.Name | Title}}AccessSuite(t *testing.T) {
	suite.Run(t, new({{.Name | Title}}AccessSuite))
}

func (s *{{.Name | Title}}AccessSuite) TearDownSuite() {

}
