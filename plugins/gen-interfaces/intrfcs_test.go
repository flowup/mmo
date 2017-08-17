package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type IntrfcsAccessSuite struct {
	suite.Suite
}

func (s *IntrfcsAccessSuite) SetupSuite() {
}

func (s *IntrfcsAccessSuite) SetupTest() {

}

// Test
func (s *IntrfcsAccessSuite) TestParse() {
	err := Parse("fixtures/proto1_go")
	s.Nil(err)
}

func TestIntrfcsAccessSuite(t *testing.T) {
	suite.Run(t, new(IntrfcsAccessSuite))
}

func (s *IntrfcsAccessSuite) TearDownSuite() {

}
