/*
一个服务有多个部署组，一个部署组只允许在一个集群中，即：多个集群中属于多个部署组。
*/
package models

import (
	"time"
)

// Group is an object representing the database table.
type Group struct {
	// record
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"index;unique;not null;" json:"name"`
	Version     string `gorm:"not null;default:'';" json:"version"`
	Lang        string `gorm:"not null;default:'';" json:"lang"`
	NamespaceID int64  `gorm:"index;not null;" json:"namespace_id"`
	ServiceID   int64  `gorm:"index;not null;" json:"service_id"`
	ClusterID   int64  `gorm:"index;not null;" json:"cluster_id"`
	Note        string `gorm:"not null;default:'';" json:"note"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	F_Namespace *Namespace `gorm:"foreignKey:NamespaceID" json:"-"`
	F_Service   *Service   `gorm:"foreignKey:ServiceID" json:"-"`
	F_Cluster   *Cluster   `gorm:"foreignKey:ClusterID" json:"-"`
	F_Nodes     []Node     `gorm:"foreignKey:GroupID" json:"-"`
}

func (r *Group) Namespace() *Namespace {
	if r.F_Namespace != nil {
		return r.F_Namespace
	}
	PanicErr(db.Model(r).Association("F_Namespace").Find(&r.F_Namespace))
	return r.F_Namespace
}

func (r *Group) Cluster() *Cluster {
	if r.F_Cluster != nil {
		return r.F_Cluster
	}
	PanicErr(db.Model(r).Association("F_Cluster").Find(&r.F_Cluster))
	return r.F_Cluster
}

func (r *Group) Service() *Service {
	if r.F_Service != nil {
		return r.F_Service
	}
	PanicErr(db.Model(r).Association("F_Service").Find(&r.F_Service))
	return r.F_Service
}

func (r *Group) Nodes() []Node {
	if r.F_Nodes != nil {
		return r.F_Nodes
	}
	PanicErr(db.Model(r).Association("F_Nodes").Find(&r.F_Nodes))
	return r.F_Nodes
}

func (r *Group) NodeIPs() []string {
	ips := []string{}
	for _, n := range r.Nodes() {
		ips = append(ips, n.IP)
	}
	return ips
}
