package utils

import (
	"os"
	"path/filepath"
	"regexp"
)

func SepareteFileNameFromPath(path string) (string, error) {
	golintRegex := regexp.MustCompile(`\/(.[^\/]+)$`)
	filename := golintRegex.FindStringSubmatch(path)
	if len(filename) != 2 {
		return "", ErrParseFileName
	}

	return filename[1], nil
}

func ImportPath() (string, error) {
	golintRegex := regexp.MustCompile(`(github.+)`)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	fp := golintRegex.FindStringSubmatch(dir)
	return fp[0], nil
}
