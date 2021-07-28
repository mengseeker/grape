package logs

const (
	LogTypeEnvoyAccess = "nxmc.envoy_access"
	LogTypeTrace       = "nxmc.trace"
)

type Message struct {
	MessageType string
	Val         []byte
}

type Receiver interface {
	Receive(msg Message)
}

type Distributor interface {
	Distribute(types []string, rec Receiver)
}

type Transmitter interface {
	Receiver
	Distributor
}
