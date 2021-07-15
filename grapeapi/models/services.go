/*
服务必须在命名空间之下，可以部署在多个集群中，所以多个集群部署服务时，命名空间相同。
*/
package models

import (
	"errors"
	"grape/api"
	"time"

	"gorm.io/gorm"
)

// Service is an object representing the database table.

var (
	ErrServiceAttrExternalInvalid = errors.New("ServiceAttrExternalInvalid")
)

type Service struct {
	// record
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"index;unique;not null;" json:"name"`
	Port        int    `gorm:"not null;" json:"port"`
	Protocol    int    `gorm:"not null" json:"protocol"`
	External    int    `gorm:"index;not null;default:0" json:"external"`
	Note        string `gorm:"not null;default:'';" json:"note"`
	NamespaceID int64  `gorm:"index;not null;" json:"namespace_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	F_Namespace *Namespace `gorm:"foreignKey:NamespaceID" json:"-"`
	F_Groups    []Group    `gorm:"foreignKey:ServiceID" json:"-"`
	F_Nodes     []Node     `gorm:"foreignKey:ServiceID" json:"-"`
	F_Policies  []Policy   `gorm:"foreignKey:ServiceID" json:"-"`
}

type ServiceProtocol int

//go:generate stringer -type ServiceProtocol -trimprefix PROTOCOL_ -output service_protocol_str.go
const (
	PROTOCOL_HTTP  ServiceProtocol = ServiceProtocol(api.Service_HTTP)
	PROTOCOL_HTTPS                 = ServiceProtocol(api.Service_HTTPS)
	PROTOCOL_HTTP2                 = ServiceProtocol(api.Service_HTTP2)
	PROTOCOL_TCP                   = ServiceProtocol(api.Service_TCP)
	PROTOCOL_UDP                   = ServiceProtocol(api.Service_UDP)
	PROTOCOL_TLS                   = ServiceProtocol(api.Service_TLS)
	PROTOCOL_GRPC                  = ServiceProtocol(api.Service_GRPC)
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

func (r *Service) Namespace() *Namespace {
	if r.F_Namespace != nil {
		return r.F_Namespace
	}
	PanicErr(db.Model(r).Association("F_Namespace").Find(&r.F_Namespace))
	return r.F_Namespace
}

func (r *Service) Groups() []Group {
	if r.F_Groups != nil {
		return r.F_Groups
	}
	PanicErr(db.Model(r).Association("F_Groups").Find(&r.F_Groups))
	return r.F_Groups
}

func (r *Service) Policies() []Policy {
	if r.F_Policies != nil {
		return r.F_Policies
	}
	PanicErr(db.Model(r).Association("F_Policies").Find(&r.F_Policies))
	return r.F_Policies
}
