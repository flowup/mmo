package filecompare

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ByteCompareSuite struct {
	suite.Suite
}

type fileCompareTest struct {
	source  string
	compare string
	err     error
}

func (s *ByteCompareSuite) TestCompareFiles() {
	candidates := []fileCompareTest{
		{
			source:  "./fixtures/f1.txt",
			compare: "./fixtures/f1_copy.txt",
			err:     nil,
		},
		{
			source:  "./fixtures/f1.txt",
			compare: "./fixtures/f1_fail.txt",
			err:     ErrFilesNotEqual,
		},
	}

	for _, can := range candidates {
		err := CompareFiles(can.source, can.compare)
		s.Equal(can.err, errors.Cause(err))
	}
}

func TestByteCompareSuite(t *testing.T) {
	suite.Run(t, &ByteCompareSuite{})
}
