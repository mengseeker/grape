package logs

const (
	LogTypeEnvoyAccess = "nxmc.envoy_access"
	LogTypeTrace       = "nxmc.trace"
)

type Message struct {
	MessageType string
	Val         []byte
}

func (m *Message) GetLog() []byte {
	return GetFluentBitMessageLog(m)
}
