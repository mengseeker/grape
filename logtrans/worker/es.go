package worker

import "log"

type EsClient struct {
	Addr string
}

func NewEsClient(addr string) (*EsClient, error) {

	return nil, nil
}

func (e *EsClient) NewRunner() Runner {
	return func(m *Message) error {
		log.Println(string(m.Value))
		return nil
	}
}
