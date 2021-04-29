// Code generated by "stringer -type ServiceProtocol -trimprefix PROTOCOL_ -output service_protocol_str.go"; DO NOT EDIT.

package models

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PROTOCOL_TCP-1]
	_ = x[PROTOCOL_UDP-2]
	_ = x[PROTOCOL_GRPC-3]
	_ = x[PROTOCOL_HTTP-4]
	_ = x[PROTOCOL_HTTP2-5]
	_ = x[PROTOCOL_HTTPS-6]
	_ = x[PROTOCOL_TLS-7]
	_ = x[PROTOCOL_MONGO-8]
	_ = x[PROTOCOL_REDIS-9]
}

const _ServiceProtocol_name = "TCPUDPGRPCHTTPHTTP2HTTPSTLSMONGOREDIS"

var _ServiceProtocol_index = [...]uint8{0, 3, 6, 10, 14, 19, 24, 27, 32, 37}

func (i ServiceProtocol) String() string {
	i -= 1
	if i < 0 || i >= ServiceProtocol(len(_ServiceProtocol_index)-1) {
		return "ServiceProtocol(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ServiceProtocol_name[_ServiceProtocol_index[i]:_ServiceProtocol_index[i+1]]
}
