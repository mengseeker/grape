package injector

import (
	"encoding/json"
	"fmt"
	"grape/api/v1/confd"
	"grape/internal/confdserver"
	"strconv"
	"time"

	"gomodules.xyz/jsonpatch/v3"
	kubeApiAdmissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	ConfdEnableKey = "grape/confd"
	MeshEnableKey  = "grape/mesh"
	ViewEnableKey  = "grape/view"

	ServiceCodeKey = "grape/service-code"
	ServicePortKey = "grape/service-port"
	GroupCodeKey   = "grape/group-code"
)

const (
	ConfdAgentContainerName = "confd-agent"
)

func injectPod(cf *InjectorConfig, ar *kubeApiAdmissionv1.AdmissionReview) error {
	var pod corev1.Pod
	req := ar.Request
	if err := json.Unmarshal(req.Object.Raw, &pod); err != nil {
		cf.Log.Errorf("Could not unmarshal raw object: %v %s", err,
			string(req.Object.Raw))
		return err
	}

	podName := potentialPodName(pod.ObjectMeta)
	if pod.ObjectMeta.Namespace == "" {
		pod.ObjectMeta.Namespace = req.Namespace
	}
	mergedPod := pod.DeepCopy()

	cf.Log.Infof("injection request for %v/%v", req.Namespace, podName)
	cf.Log.Debugf("Object: %v", string(req.Object.Raw))

	ann := pod.Annotations
	if cf.EnableConfd && ann[ConfdEnableKey] != "disable" {
		cf.Log.Debugf("inject confd for %s/%s", pod.Namespace, podName)
		if err := doInjectConfd(cf, &pod, mergedPod); err != nil {
			return err
		}
	}

	if cf.EnableMesh && ann[MeshEnableKey] != "disable" {
		cf.Log.Debugf("inject mesh for %s/%s", pod.Namespace, podName)
		if err := doInjectMesh(cf, &pod, mergedPod); err != nil {
			return err
		}
	}

	if cf.EnableView && ann[ViewEnableKey] != "disable" {
		cf.Log.Debugf("inject view for %s/%s", pod.Namespace, podName)
		if err := doInjectView(cf, &pod, mergedPod); err != nil {
			return err
		}
	}

	return patchPodResponse(ar, &pod, mergedPod)
}

func patchPodResponse(ar *kubeApiAdmissionv1.AdmissionReview, pod, merge *corev1.Pod) error {
	original, err := json.Marshal(pod)
	if err != nil {
		return err
	}
	reinjected, err := json.Marshal(merge)
	if err != nil {
		return err
	}
	patch, err := jsonpatch.CreatePatch(original, reinjected)
	if err != nil {
		return err
	}
	patchBytes, _ := json.Marshal(patch)

	jsonPatch := kubeApiAdmissionv1.PatchTypeJSONPatch

	ar.Response = &kubeApiAdmissionv1.AdmissionResponse{
		Allowed:   true,
		UID:       ar.Request.UID,
		Patch:     patchBytes,
		PatchType: &jsonPatch,
	}

	return nil
}

func doInjectConfd(ijf *InjectorConfig, pod, merge *corev1.Pod) error {
	serviceCode := pod.Annotations[ServiceCodeKey]
	if serviceCode == "" {
		return fmt.Errorf("annotation %s not found", ServiceCodeKey)
	}
	groupCode := pod.Annotations[GroupCodeKey]
	service := pod.Namespace + "/" + serviceCode
	ijf.Log.Debugf("inject serverConfigs %s(%s)", service, groupCode)
	appContainer, err := getAppContatiner(merge)
	if err != nil {
		return err
	}
	serverConfigs, rev, err := confdserver.GetServiceConfigs(ijf.Cli, service, groupCode, 0)
	if err != nil {
		return err
	}
	if serverConfigs == nil {
		return nil
	}
	injectEnv(appContainer, serverConfigs.EnvConfigs)
	injectFiles(ijf, serverConfigs.FileConfigs, appContainer, merge, serviceCode, groupCode, rev)
	return nil
}

func injectFiles(ijf *InjectorConfig, cf []*confd.FileConfig, ac *corev1.Container, pod *corev1.Pod, serviceCode, groupCode string, rev int64) {
	c := corev1.Container{}
	c.Name = ConfdAgentContainerName
	c.Image = ijf.ConfdAgentImage
	c.ImagePullPolicy = corev1.PullIfNotPresent
	c.Args = append(c.Args, "-s", fmt.Sprintf("%s/%s", pod.Namespace, serviceCode))
	c.Args = append(c.Args, "-g", groupCode)
	c.Args = append(c.Args, "-a", ijf.DiscoveryAddress)
	c.Args = append(c.Args, "-l", strconv.Itoa(int(rev)))
	for _, f := range cf {
		hf := fmt.Sprintf("%s%s/%d", ConfdHostPathBaseDir, time.Now().Format("2006-01-02"), time.Now().UnixNano())
		var hp corev1.HostPathType = "FileOrCreate"
		pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{
			Name: f.Name, VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: hf, Type: &hp,
				},
			},
		})
		c.VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{Name: f.Name, MountPath: f.Path, ReadOnly: false})
		ac.VolumeMounts = append(ac.VolumeMounts, corev1.VolumeMount{Name: f.Name, MountPath: f.Path, ReadOnly: true})
	}
	pod.Spec.InitContainers = append(pod.Spec.InitContainers, c)
}

func injectEnv(c *corev1.Container, env []*confd.EnvConfig) {
	for _, e := range env {
		appendContainerEnv(c, e.Key, e.Value)
	}
}

func doInjectMesh(cf *InjectorConfig, pod, merge *corev1.Pod) error {

	return nil
}

func doInjectView(cf *InjectorConfig, pod, merge *corev1.Pod) error {

	return nil
}

// func NewSidecar(serviceCode, servicePort, groupCode string) corev1.Container {
// 	c := corev1.Container{
// 		Name:      SidecarContainerName,
// 		Image:     proxyImage,
// 		Ports:     []corev1.ContainerPort{{Name: "http-admin", Protocol: "TCP", ContainerPort: 9901}},
// 		Env:       []corev1.EnvVar{{Name: "serviceCode", Value: serviceCode}, {Name: "serviceDeployGroupCode", Value: groupCode}, {Name: "POD_IP", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{APIVersion: "v1", FieldPath: "status.podIP"}}}},
// 		Resources: corev1.ResourceRequirements{Limits: corev1.ResourceList{"cpu": resource.MustParse("300m"), "memory": resource.MustParse("512Mi")}, Requests: corev1.ResourceList{"cpu": resource.MustParse("50m"), "memory": resource.MustParse("128Mi")}},
// 	}
// 	if logPvcName != "" {
// 		c.VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{
// 			Name:      "envoylog",
// 			MountPath: "/opt/dataforce/log/envoy",
// 			SubPath:   groupCode,
// 		})
// 	}
// 	return c
// }

// func NewInitContainer(serviceCode, servicePort, groupCode string) corev1.Container {
// 	c := corev1.Container{
// 		Name:            InitContainerName,
// 		Image:           initImage,
// 		Env:             []corev1.EnvVar{{Name: "PROXY_PORT", Value: servicePort}},
// 		SecurityContext: &corev1.SecurityContext{Capabilities: &corev1.Capabilities{Add: []corev1.Capability{"NET_ADMIN"}}},
// 	}
// 	return c
// }
