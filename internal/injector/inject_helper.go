package injector

import (
	"fmt"

	kubeApiAdmissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SidecarContainerName = "grape-sidecar"
	InitContainerName    = "grape-initProxy"

	DefaultApplicationContainerName = "app"
)

var (
	errApplicationContainerNotfound = fmt.Errorf("container %q not found", DefaultApplicationContainerName)
)

func FindContainer(name string, containers []corev1.Container) *corev1.Container {
	for i := range containers {
		if containers[i].Name == name {
			return &containers[i]
		}
	}
	return nil
}

func FindSidecar(containers []corev1.Container) *corev1.Container {
	return FindContainer(InitContainerName, containers)
}

func FindInitContainer(containers []corev1.Container) *corev1.Container {
	return FindContainer(InitContainerName, containers)
}

func potentialPodName(metadata metav1.ObjectMeta) string {
	if metadata.Name != "" {
		return metadata.Name
	}
	if metadata.GenerateName != "" {
		return metadata.GenerateName + "***** (actual name not yet known)"
	}
	return ""
}

func errAdmissionResponse(ar *kubeApiAdmissionv1.AdmissionReview, errMessage string) {

}

func getAppContatiner(pod *corev1.Pod) (*corev1.Container, error) {
	if len(pod.Spec.Containers) == 1 {
		return &pod.Spec.Containers[0], nil
	} else {
		for i := 0; i < len(pod.Spec.Containers); i++ {
			if pod.Spec.Containers[i].Name == DefaultApplicationContainerName {
				return &pod.Spec.Containers[i], nil
			}
		}
	}
	return nil, errApplicationContainerNotfound
}

func appendContainerEnv(c *corev1.Container, key, value string) {
	for i := 0; i < len(c.Env); i++ {
		if c.Env[i].Name == key {
			c.Env[i].Value = value
			c.Env[i].ValueFrom = nil
			return
		}
	}
	c.Env = append(c.Env, corev1.EnvVar{Name: key, Value: value})
}
