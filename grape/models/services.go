package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Service is an object representing the database table.

var (
	ErrServiceAttrExternalInvalid = errors.New("ServiceAttrExternalInvalid")
)

type Service struct {
	// record
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"index;unique;not null;" json:"name"`
	Code      string `gorm:"index;unique;not null;" json:"code"`
	Port      int    `gorm:"not null;" json:"port"`
	Protocol  int    `gorm:"not null" json:"protocol"`
	External  int    `gorm:"index;not null;default:0" json:"external"`
	Note      string `gorm:"not null;default:'';" json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

func (r *Service) BeferSave(*gorm.DB) error {
	if r.External < 0 || r.External > 1 {
		return ErrServiceAttrExternalInvalid
	}
	return nil
}
