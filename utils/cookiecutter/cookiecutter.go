package cookiecutter

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Restore restores contents of a given text to the given file, replacing
// all paths and file data by the data interface
func Restore(dir string, path string, text string, data interface{}, funcs template.FuncMap) error {
	pathBytes := &bytes.Buffer{}
	template.Must(template.New(path).Parse(path)).Execute(pathBytes, data)

	err := os.MkdirAll(_filePath(dir, filepath.Dir(pathBytes.String())), os.FileMode(0755))
	if err != nil {
		return err
	}

	fileBytes := &bytes.Buffer{}
	template.Must(template.New("file:"+path).Funcs(funcs).Parse(text)).Execute(fileBytes, data)

	err = ioutil.WriteFile(_filePath(dir, pathBytes.String()), fileBytes.Bytes(), os.FileMode(0755))
	if err != nil {
		return err
	}

	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
