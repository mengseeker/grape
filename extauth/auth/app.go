package auth

import (
	"errors"
)

type AppID int
type EndpointID int

type App struct {
	ID          AppID        `json:"id"`
	APPID       string       `json:"app_id"`
	APPSec      string       `json:"app_sec"`
	Username    string       `json:"username"`
	Endpoints   []EndpointID `json:"endpoints"`
	Tokens      []string     `json:"tokens"`
	idxEndpoint map[EndpointID]bool
}

// 读多写少，且每次更新时全量更新，不用加锁.
var (
	endpints = map[string]EndpointID{}
	apps     = map[AppID]*App{}
	tokens   = map[string]AppID{}
)

var (
	errUnauthorizedToken    = errors.New("unauthorized token")
	errUnregisteredEndpoint = errors.New("unregistered endpoint")
)

// check token then return app
func GetAppByToken(token string) (*App, error) {
	id, ok := tokens[token]
	if !ok {
		return nil, errUnauthorizedToken
	}
	return apps[id], nil
}

func GetEndpoint(method, path string) EndpointID {
	return endpints[method+"::"+path]
}

func (a *App) Auth(endpoint EndpointID) error {
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

func (a *App) HasEndpoint(endpoint EndpointID) bool {
	return a.idxEndpoint[endpoint]
}