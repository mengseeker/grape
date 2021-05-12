package auth

import "errors"

// GET::/api/healthcheck
type Endpoint string

var (
	errUnregisteredEndpoint = errors.New("unregistered endpoint")

	endpints = map[Endpoint]bool{}
)

func GetEndpoint(method, path string) Endpoint {
	return Endpoint(method + "::" + path)
}

func NeedAuth(end Endpoint) bool {
	return endpints[end]
}

func buildAppEndpoints() {
	newEnds := map[Endpoint]bool{}
	tmps := make([]Endpoint, 0, len(endpints))
	for _, app := range apps {
		for _, end := range app.Endpoints {
			newEnds[end] = true
			tmps = append(tmps, end)
		}
	}
	endpints = newEnds
	log.Infof("endpoints %v updated", tmps)
}
