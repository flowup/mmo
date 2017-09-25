package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"io/ioutil"
	"bytes"
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
		output, err := ioutil.ReadFile(candidate.Output)
		s.Nil(err)

		newInterfaces, err := Parse(candidate.Input, candidate.Output, false)
		s.Nil(err)

		for _, newInterface := range newInterfaces {
			output = append(output, []byte(newInterface)...)
		}

		expected, err := ioutil.ReadFile(candidate.Expected)
		s.Nil(err)
		s.True(bytes.Equal(output, expected))
	}
}

func TestIntrfcsAccessSuite(t *testing.T) {
	suite.Run(t, new(IntrfcsAccessSuite))
}

func (s *IntrfcsAccessSuite) TearDownSuite() {

}
