package kubernetes

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os/exec"
	"os/user"
)

const (
	minikubeUsername = "minikube"
	minikubeHost     = "192.168.99.100:8443"
	minikubeCAFile   = "/.minikube/ca.crt"
	minikubeKeyFile  = "/.minikube/apiserver.key"
	minikubeCertFile = "/.minikube/apiserver.crt"
)

var (
	errNotInstalled = errors.New("minikube is not installed. Please install minikube: https://github.com/kubernetes/minikube")
	errNotRunning   = errors.New("minikube is not running. Run minikube using command 'minikube start'")
	errPwFailed     = errors.New("Forwarding port of the minikube's docker registry failed")
)

// IsRunning is function to find out minikube is running or not
func IsRunning() error {

	_, err := exec.LookPath("minikube")
	if err != nil {
		return errNotInstalled
	}

	cmdCube := exec.Command("minikube", "ip")

	err = cmdCube.Run()
	if err != nil {
		return errNotRunning
	}

	return nil
}

// ConnectToCluster is function to connect to local minikube cluster
func ConnectToCluster() (*kubernetes.Clientset, error) {

	usr, err := user.Current()
	if err != nil {
		return &kubernetes.Clientset{}, err
	}

	config := &rest.Config{}
	config.Username = minikubeUsername
	config.Host = minikubeHost
	config.TLSClientConfig.CAFile = usr.HomeDir + minikubeCAFile
	config.TLSClientConfig.KeyFile = usr.HomeDir + minikubeKeyFile
	config.TLSClientConfig.CertFile = usr.HomeDir + minikubeCertFile

	return kubernetes.NewForConfig(config)
}

// IsRegistryRunning is function to check if all parts of docker registry are deployed
func IsRegistryRunning(client *kubernetes.Clientset) error {
	rplInterface := client.CoreV1Client.ReplicationControllers(RegistryReplicationController.ObjectMeta.Namespace)
	_, err := rplInterface.Get(RegistryReplicationController.ObjectMeta.Name, v1.GetOptions{})
	if err != nil {
		return err
	}

	svcInterface := client.CoreV1Client.Services(RegistryService.ObjectMeta.Namespace)
	_, err = svcInterface.Get(RegistryService.ObjectMeta.Name, v1.GetOptions{})
	if err != nil {
		return err
	}

	daemonInterface := client.ExtensionsV1beta1Client.DaemonSets(RegistryDaemonSet.ObjectMeta.Namespace)
	_, err = daemonInterface.Get(RegistryDaemonSet.ObjectMeta.Name, v1.GetOptions{})
	return err
}

// IsRegistryAccessible is function to check if docker registry running in minikube is accessible
func IsRegistryAccessible() error {
	return nil
}

// DeployDockerRegistry is function to deploy docker registry to connected k8s cluster
func DeployDockerRegistry(client *kubernetes.Clientset) error {
	rplInterface := client.CoreV1Client.ReplicationControllers(RegistryReplicationController.ObjectMeta.Namespace)
	_, err := rplInterface.Create(&RegistryReplicationController)
	if err != nil {
		return err
	}

	svcInterface := client.CoreV1Client.Services(RegistryService.ObjectMeta.Namespace)
	_, err = svcInterface.Create(&RegistryService)
	if err != nil {
		return err
	}

	daemonInterface := client.ExtensionsV1beta1Client.DaemonSets(RegistryDaemonSet.ObjectMeta.Namespace)
	_, err = daemonInterface.Create(&RegistryDaemonSet)
	return err
}

// ForwardRegistryPort is function to forward minikube's registry port, returned Cmd should be killed when forwarding is not needed
func ForwardRegistryPort() (*exec.Cmd, error) {
	cmdCube := exec.Command("bash", "-c", "kubectl port-forward --namespace kube-system "+
		"$(kubectl get po -n kube-system | grep kube-registry-v0 | awk '{print $1;}') 17465:17465")

	err := cmdCube.Start()
	if err != nil {
		return cmdCube, errors.Wrap(errPwFailed, err.Error())
	}
	return cmdCube, nil
}
