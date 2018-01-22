package filecompare

import (
	"os"
	"path/filepath"
	"strings"
)

//CompareFolder take source as source path and
// expected as expected path. This Function compare
// two folders recursively with all files.
// Detect also all missing and extra files.es.
func CompareFolder(source, expected string) error {
	if err := filepath.Walk(expected, func(expectedPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		sourcePath := strings.Replace(expectedPath, expected[2:], source, 1)

		return CompareFiles(expectedPath, sourcePath)

	}); err != nil {
		return err
	}

	return filepath.Walk(source, func(sourcePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		expectedPath := strings.Replace(sourcePath, source[2:], expected, 1)

		return CompareFiles(sourcePath, expectedPath)
	})
}
