package worker

type InfClient struct {
	Addr string
}

func NewInfClient(addr string) (*InfClient, error) {
	return nil, nil
}

func (e *InfClient) NewRunner() Runner {
	return nil
}
