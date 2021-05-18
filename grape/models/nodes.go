package models

import "time"

// Node is an object representing the database table.
type Node struct {
	// record
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"index;unique;not null;" json:"name"`
	Code        string `gorm:"index;unique;not null;" json:"code"`
	NamespaceID int64  `gorm:"index;not null;" json:"namespace_id"`
	ServiceID   int64  `gorm:"index;not null;" json:"service_id"`
	ClusterID   int64  `gorm:"index;not null;" json:"cluster_id"`
	GroupID     int64  `gorm:"index;not null;" json:"group_id"`
	IP          string `gorm:"not null;" json:"ip"`
	State       int    `gorm:"not null;" json:"state" toml:"state" yaml:"state"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	F_Service *Service `gorm:"foreignKey:ServiceID" json:"-"`
	F_Cluster *Cluster `gorm:"foreignKey:ClusterID" json:"-"`
	F_Group   *Group   `gorm:"foreignKey:GroupID" json:"-"`
}

func (r *Node) Cluster() *Cluster {
	if r.F_Cluster != nil {
		return r.F_Cluster
	}
	PanicErr(db.Model(r).Association("F_Cluster").Find(&r.F_Cluster))
	return r.F_Cluster
}

func (r *Node) Service() *Service {
	if r.F_Service != nil {
		return r.F_Service
	}
	PanicErr(db.Model(r).Association("F_Service").Find(&r.F_Service))
	return r.F_Service
}

func (r *Node) Group() *Group {
	if r.F_Group != nil {
		return r.F_Group
	}
	PanicErr(db.Model(r).Association("F_Group").Find(&r.F_Group))
	return r.F_Group
}
