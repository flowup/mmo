package kubernetes

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/pkg/util"
)

// RegistryReplicationController is k8s resource for deploying docker registry replication controller
var RegistryReplicationController = apiv1.ReplicationController{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "kube-registry-v0",
		Namespace: "kube-system",
		Labels: map[string]string{
			"k8s-app": "kube-registry",
			"version": "v0",
		},
	},
	Spec: apiv1.ReplicationControllerSpec{
		Replicas: util.Int32Ptr(1),
		Selector: map[string]string{
			"k8s-app": "kube-registry",
			"version": "v0",
		},
		Template: &apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"k8s-app": "kube-registry",
					"version": "v0",
				},
			},
			Spec: apiv1.PodSpec{
				Containers: []apiv1.Container{
					{
						Name:  "registry",
						Image: "registry:2",
						Resources: apiv1.ResourceRequirements{
							Requests: apiv1.ResourceList{
								"cpu":    resource.MustParse("10m"),
								"memory": resource.MustParse("10Mi"),
							},
						},
						Env: []apiv1.EnvVar{
							{Name: "REGISTRY_HTTP_ADDR", Value: ":17465"},
							{Name: "REGISTRY_STORAGE_FILESYSTEM_ROOTDIRECTORY", Value: "/var/lib/registry"},
						},
						VolumeMounts: []apiv1.VolumeMount{
							{Name: "image-store", MountPath: "/var/lib/registry"},
						},
						Ports: []apiv1.ContainerPort{
							{
								ContainerPort: 17465,
								Name:          "registry",
								Protocol:      "TCP",
							},
						},
					},
				},
				Volumes: []apiv1.Volume{
					{
						Name: "image-store",
						VolumeSource: apiv1.VolumeSource{
							HostPath: &apiv1.HostPathVolumeSource{
								Path: "/data/registry/",
							},
						},
					},
				},
			},
		},
	},
}

// RegistryService is k8s resource for deploying docker registry service
var RegistryService = apiv1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "kube-registry",
		Namespace: "kube-system",
		Labels: map[string]string{
			"k8s-app": "kube-registry",
		},
	},
	Spec: apiv1.ServiceSpec{
		Selector: map[string]string{
			"k8s-app": "kube-registry",
		},
		Ports: []apiv1.ServicePort{
			{
				Port:     17465,
				Name:     "registry",
				Protocol: "TCP",
			},
		},
	},
}

// RegistryDaemonSet is k8s resource for deploying docker registry daemon set
var RegistryDaemonSet = v1beta1.DaemonSet{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "kube-registry-proxy",
		Namespace: "kube-system",
		Labels: map[string]string{
			"k8s-app":                       "kube-registry",
			"kubernetes.io/cluster-service": "true",
			"version":                       "v0.4",
		},
	},
	Spec: v1beta1.DaemonSetSpec{
		Template: apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"k8s-app": "kube-registry",
					"version": "v0.4",
				},
			},
			Spec: apiv1.PodSpec{
				Containers: []apiv1.Container{
					{
						Name:  "kube-registry-proxy",
						Image: "gcr.io/google_containers/kube-registry-proxy:0.4",
						Env: []apiv1.EnvVar{
							{Name: "REGISTRY_HOST", Value: "kube-registry.kube-system.svc.cluster.local"},
							{Name: "REGISTRY_PORT", Value: "17465"},
						},
						Ports: []apiv1.ContainerPort{
							{
								ContainerPort: 80,
								HostPort:      17465,
								Name:          "registry",
							},
						},
					},
				},
			},
		},
	},
}
