package filecompare

import (
	"io/ioutil"
	"github.com/pkg/errors"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var (
	// ErrFilesNotEqual should be fired any time two compared files
	// are not byte-equal
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

	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(string(src), string(exp), true)

	if len(diffs) > 1 {
		fmt.Println(dmp.DiffPrettyText(diffs))
		return errors.Wrap(ErrFilesNotEqual, "Files: "+source+" / "+expected)
	}

	return nil
}
