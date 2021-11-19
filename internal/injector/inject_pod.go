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
	if cf.EnableConfd && ann[share.Annotation_DisableConfd] != "true" {
		cf.Log.Debugf("inject confd for %s/%s", pod.Namespace, podName)
		if err := doInjectConfd(cf, &pod, mergedPod); err != nil {
			return err
		}
	}

	if cf.EnableMesh && ann[share.Annotation_DisableMesh] != "true" {
		cf.Log.Debugf("inject mesh for %s/%s", pod.Namespace, podName)
		if err := doInjectMesh(cf, &pod, mergedPod); err != nil {
			return err
		}
	}

	return patchPodResponse(cf, ar, &pod, mergedPod)
}

func patchPodResponse(cf *InjectorConfig, ar *kubeApiAdmissionv1.AdmissionReview, pod, merge *corev1.Pod) error {
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
	cf.Log.Debugf("patchs: %s", patchBytes)

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
	projectName := pod.Labels[share.Label_ProjectName]
	if projectName == "" {
		return fmt.Errorf("label %s undefined", share.Label_ProjectName)
	}
	groupCode := pod.Labels[share.Label_Group]
	if groupCode == "" {
		// get form name
		groupCode = getDeploymentName(pod.GenerateName)
	}
	ijf.Log.Debugf("inject serverConfigs %s:%s", projectName, groupCode)

	appContainer, err := getAppContatiner(merge)
	if err != nil {
		return err
	}
	args := []string{"-a", ijf.InjectDiscoveryAddress}
	args = append(args, "-p", projectName)
	args = append(args, "-g", groupCode)
	args = append(args, "-d")

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
