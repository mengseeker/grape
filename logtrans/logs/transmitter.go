package logs

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
