package models

type ServiceProtocol int

//go:generate stringer -type ServiceProtocol -trimprefix PROTOCOL_ -output service_protocol_str.go
const (
	PROTOCOL_TCP ServiceProtocol = iota + 1
	PROTOCOL_UDP
	PROTOCOL_GRPC
	PROTOCOL_HTTP
	PROTOCOL_HTTP2
	PROTOCOL_HTTPS
	PROTOCOL_TLS
	PROTOCOL_MONGO
	PROTOCOL_REDIS
)

func (o *Service) ProtocolString() string {
	return ServiceProtocol(o.Protocol).String()
}