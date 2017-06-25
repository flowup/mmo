package utils

import "regexp"

func SepareteFileNameFromPath(path string) (string, error) {
	golintRegex := regexp.MustCompile(`\/(.[^\/]+)$`)
	filename := golintRegex.FindStringSubmatch(path)
	if len(filename) != 2 {
		return "", ErrParseFileName
	}

	return filename[1], nil
}
