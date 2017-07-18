package kubernetes

import (
	"fmt"
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

func ExpandTemplate(env DeployEnvironment) error {
	err := SetupDeployEnvironment(env)
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(env["SERVICE"] + "deployment/deployment.yaml.template")
	if err != nil {
		return err
	}

	s := string(file)
	expanded := os.ExpandEnv(s)

	fmt.Println(expanded)
	return nil
}
