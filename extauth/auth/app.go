package auth

import (
	"encoding/json"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

type App struct {
	ID        int        `json:"id"`
	APPID     string     `json:"app_id"`
	APPSec    string     `json:"app_sec"`
	Username  string     `json:"username"`
	Endpoints []Endpoint `json:"endpoints"`
	// Tokens      []string     `json:"tokens"`
	idxEndpoint map[Endpoint]bool
}

// 读多写少，且每次更新时全量更新，不用加锁.
var (
	apps    = map[int]*App{}
	appKeys = map[string]int{}
)

func (a *App) Auth(endpoint Endpoint) error {
	if !a.HasEndpoint(endpoint) {
		return errUnregisteredEndpoint
	}
	return nil
}

func (a *App) Headers() map[string]string {
	return map[string]string{
		"X-Username": a.Username,
	}
}

func (a *App) HasEndpoint(endpoint Endpoint) bool {
	return a.idxEndpoint[endpoint]
}

func (a *App) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

func UnmarshalApp(raw []byte) (*App, error) {
	var a App
	err := json.Unmarshal(raw, &a)
	return &a, err
}

func SetupApp(kv *mvccpb.KeyValue) error {
	a, err := UnmarshalApp(kv.Value)
	if err != nil {
		return err
	}
	a.idxEndpoint = map[Endpoint]bool{}
	for _, end := range a.Endpoints {
		a.idxEndpoint[end] = true
		endpints[end] = true
	}
	apps[a.ID] = a
	appKeys[string(kv.Key)] = a.ID
	buildAppEndpoints()
	log.Infof("app %s added", a.APPID)
	return nil
}

// 删除app、删除app、重新build endpints
// 对于tokens，由grape server负责删除
func RemoveApp(kv *mvccpb.KeyValue) error {
	aid := appKeys[string(kv.Key)]
	appid := apps[aid].APPID
	delete(apps, aid)
	delete(appKeys, string(kv.Key))
	buildAppEndpoints()
	log.Infof("app %s removed", appid)
	return nil
}

func UpdateApp(kv *mvccpb.KeyValue) error {
	a, err := UnmarshalApp(kv.Value)
	if err != nil {
		return err
	}
	apps[a.ID] = a
	buildAppEndpoints()
	log.Infof("app %s updated", a.APPID)
	return nil
}
