package utils

import "os"

// CreateDir Create directory in path
func CreateDir(path string) error{
	return os.Mkdir(path, 0777)
}
