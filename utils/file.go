package utils

import (
	"os"
)

// CreateFile create file in path
func CreateFile(path string) error {
	_, err := os.Create(path)
	return err
}

// WriteFile put content into file with path
func WriteFile(path string, content string) error {
	var file, err = os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return file.Sync()
}
