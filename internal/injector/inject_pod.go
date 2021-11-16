package injector

import (
	"encoding/json"
	"fmt"
	"grape/internal/share"

	"gomodules.xyz/jsonpatch/v3"
	kubeApiAdmissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
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
	if cf.EnableConfd && ann[share.Annotation_ConfdEnableKey] != "false" {
		cf.Log.Debugf("inject confd for %s/%s", pod.Namespace, podName)
		if err := doInjectConfd(cf, &pod, mergedPod); err != nil {
			return err
		}
	}

	if cf.EnableMesh && ann[share.Annotation_MeshEnableKey] != "false" {
		cf.Log.Debugf("inject mesh for %s/%s", pod.Namespace, podName)
		if err := doInjectMesh(cf, &pod, mergedPod); err != nil {
			return err
		}
	}

	if cf.EnableView && ann[share.Annotation_ViewEnableKey] != "false" {
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
	serviceCode := pod.Annotations[share.Annotation_ServiceCodeKey]
	if serviceCode == "" {
		return fmt.Errorf("annotation %s not found", share.Annotation_ServiceCodeKey)
	}
	groupCode := pod.Annotations[share.Annotation_GroupCodeKey]
	ijf.Log.Debugf("inject serverConfigs %s/%s(%s)", pod.Namespace, serviceCode, groupCode)
	appContainer, err := getAppContatiner(merge)
	if err != nil {
		return err
	}
	args := []string{"-a", ijf.InjectDiscoveryAddress}
	args = append(args, "-s", serviceCode)
	args = append(args, "-n", pod.Namespace)
	args = append(args, "-g", groupCode)

	// disable confd agent discovery, just download configs
	if pod.Annotations[share.Annotation_Confd_Discovery] != "true" {
		args = append(args, "-d")
	}

	// Cover the runCmd defined by conf-server force
	if pod.Annotations[share.Annotation_Confd_RunCmd] != "" {
		args = append(args, "-r", pod.Annotations[share.Annotation_Confd_RunCmd])
	}

	appContainer.Command = []string{"confd"}
	appContainer.Args = args
	return nil
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
