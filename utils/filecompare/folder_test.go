package filecompare

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/pkg/errors"
	"os"
	"syscall"
)

type FolderCompareSuite struct {
	suite.Suite
}

type folderCompareTest struct {
	source  string
	compare string
	err     error
}

func (s *FolderCompareSuite) TestCompareFiles() {
	candidates := []fileCompareTest{
		{
			source:  "./fixtures/fold_orig/",
			compare: "./fixtures/fold_copy/",
			err:     nil,
		},
		{
			source:  "./fixtures/fold_orig/",
			compare: "./fixtures/fold_fail/",
			err:     ErrFilesNotEqual,
		},
		{
			source:  "./fixtures/fold_orig/",
			compare: "./fixtures/fold_fail2/",
			err:     &os.PathError{"open", "./fixtures/fold_fail2/fold/f2.txt", syscall.ENOENT},
		},

		{
			source:  "./fixtures/fold_orig/",
			compare: "./fixtures/fold_fail3/",
			err:     &os.PathError{"open", "./fixtures/fold_orig/fold/f3.txt", syscall.ENOENT},
		},
	}

	for _, can := range candidates {
		err := CompareFolder(can.source, can.compare)
		s.Equal(can.err, errors.Cause(err))
	}
}

func TestFolderCompareSuite(t *testing.T) {
	suite.Run(t, &FolderCompareSuite{})
}
