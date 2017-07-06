package filecompare

import (
	"io/ioutil"
	"bytes"
	"github.com/pkg/errors"
)

var (
	ErrFilesNotEqual = errors.New("Given files are not equal")
)

// CompareFiles compare two file paths
func CompareFiles(source, expected string) error {
	src, srcErr := ioutil.ReadFile(source)

	if srcErr != nil {
		return srcErr
	}

	exp, expErr := ioutil.ReadFile(expected)

	if expErr != nil {
		return expErr
	}

	if !bytes.Equal(src, exp) {
		return errors.Wrap(ErrFilesNotEqual, "Files: " + source + " / " + expected)
	}

	return nil
}