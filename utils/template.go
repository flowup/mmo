package utils

import (
	"bytes"
	"text/template"
)

func templateReader(definition Definition, path string) (string, error) {
	templateString := new(bytes.Buffer)

	filename, err := SepareteFileNameFromPath(path)
	if err != nil {
		return "", err
	}

	t, err := template.New(filename).Funcs(template.FuncMap{}).ParseFiles(path)
	if err != nil {
		return "", err
	}

	t.Execute(templateString, definition)

	return templateString.String(), nil
}

func CreateFileFromTemplate(definition Definition, path string) error {
	filename, err := SepareteFileNameFromPath(path)
	if err != nil {
		return err
	}

	if err := CreateFile(definition.Name + "/" + filename); err != nil {
		return err
	}

	resultString, err := templateReader(definition, path)
	if err != nil {
		return err
	}

	WriteFile(definition.Name+"/"+filename, resultString)
	return nil
}
