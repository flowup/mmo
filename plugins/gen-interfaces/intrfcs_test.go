package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/flowup/mmo/utils/filecompare"
)

type IntrfcsAccessSuite struct {
	suite.Suite
}

func (s *IntrfcsAccessSuite) SetupSuite() {
}

func (s *IntrfcsAccessSuite) SetupTest() {

}

// Test
type testStruct struct {
	Input    string
	Output   string
	Expected string
}

func (s *IntrfcsAccessSuite) TestParse() {
	candidates := []testStruct{
		{
			Input:    "fixtures/proto1_go",
			Output:   "fixtures/service1_go",
			Expected: "fixtures/service_orig_go",
		},
		{
			Input:    "fixtures/proto1_go",
			Output:   "fixtures/service2_go",
			Expected: "fixtures/service_orig_go",
		},
		{
			Input:    "fixtures/proto1_go",
			Output:   "fixtures/service3_go",
			Expected: "fixtures/service_orig_go",
		},
	}

	for _, candidate := range candidates {
		s.Nil(Parse(candidate.Input, candidate.Output))
		s.Nil(filecompare.CompareFiles(candidate.Output, candidate.Expected))
	}
}

func TestIntrfcsAccessSuite(t *testing.T) {
	suite.Run(t, new(IntrfcsAccessSuite))
}

func (s *IntrfcsAccessSuite) TearDownSuite() {

}
