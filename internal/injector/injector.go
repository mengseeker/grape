package injector

import (
	"encoding/json"
	"errors"
	"fmt"
	"grape/pkg/etcdcli"
	"grape/pkg/logger"
	"io/ioutil"
	"net/http"

	kubeApiAdmissionv1 "k8s.io/api/admission/v1"
)

var (
	errUnsupportKind = errors.New("unsupport kind")
)

type InjectorConfig struct {
	Cli                    *etcdcli.Client
	Log                    logger.Logger
	EnableConfd            bool
	EnableMesh             bool
	EnableView             bool
	InjectDiscoveryAddress string
	MeshSidecarImage       string
}

func (cf *InjectorConfig) NewjectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			cf.Log.Error(err)
			http.Error(w, err.Error(), 400)
			return
		}
		if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
			cf.Log.Errorf("contentType=%s, expect application/json", contentType)
			http.Error(w, "expect contentType application/json", 400)
			return
		}

		review := &kubeApiAdmissionv1.AdmissionReview{}
		if err := json.Unmarshal(data, &review); err != nil {
			cf.Log.Errorf("Fail to deserialize object: %s with error: %v", string(data), err)
			errAdmissionResponse(review, fmt.Sprintf("Fail to deserialize object: %v", err))
			goto WRITE_RESP
		}

		if err := doInject(cf, review); err != nil {
			cf.Log.Errorf("failed to inject: %v", err)
			errAdmissionResponse(review, fmt.Sprintf("failed to inject: %v", err))
			goto WRITE_RESP
		}

	WRITE_RESP:
		w.Header().Set("Content-Type", "application/json")
		review.Request = nil
		if err := json.NewEncoder(w).Encode(review); err != nil {
			cf.Log.Errorf("Marshal of response failed with error: %v", err)
			http.Error(w, err.Error(), 500)
			return
		}
	}

}

func doInject(cf *InjectorConfig, ar *kubeApiAdmissionv1.AdmissionReview) error {
	req := ar.Request
	if (req.Kind.Group == "apps" || req.Kind.Group == "") && req.Kind.Kind == "Pod" {
		return injectPod(cf, ar)
	} else {
		return errUnsupportKind
	}
}
