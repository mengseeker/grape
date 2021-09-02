package logs

const (
	LogTypeEnvoyAccess = "nxmc.envoyaccess"
	LogTypeTrace       = "nxmc.trace"
)

type Message struct {
	MessageType string
	Val         []byte
}

func (m *Message) GetLog() []byte {
	return GetFluentBitMessageLog(m)
}
