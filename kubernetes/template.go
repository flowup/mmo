package kubernetes

import (
	"bytes"
	"io/ioutil"
	"os"
)

// DeployEnvironment is type to hold all needed env variables to export
type DeployEnvironment map[string]string

// SetupDeployEnvironment is function to export variables needed for expanding k8s templates
func SetupDeployEnvironment(env DeployEnvironment) error {
	for k := range env {
		err := os.Setenv(k, env[k])
		if err != nil {
			return err
		}
	}

	return nil
}

// ExpandTemplate is function to expand default service's template with env variables
func ExpandTemplate(env DeployEnvironment) ([]byte, error) {
	err := SetupDeployEnvironment(env)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(env["SERVICE"] + "/deployment/deployment.yaml.template")
	if err != nil {
		return nil, err
	}

	s := string(file)
	return []byte(os.ExpandEnv(s)), nil
}

func splitYamlDocument(contents []byte) [][]byte {
	return bytes.Split(contents, []byte("---\n"))
}
