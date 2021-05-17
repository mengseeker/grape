package models

import "time"

// Node is an object representing the database table.
type Node struct {
	// record
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"index;unique;not null;" json:"name"`
	Code        string `gorm:"index;unique;not null;" json:"code"`
	NamespaceID int    `gorm:"index;not null;" json:"namespace_id"`
	ServiceID   int    `gorm:"index;not null;" json:"service_id"`
	ClusterID   int    `gorm:"index;not null;" json:"cluster_id"`
	GroupID     int    `gorm:"index;not null;" json:"group_id"`
	IP          string `gorm:"not null;" json:"ip"`
	State       int    `gorm:"not null;" json:"state" toml:"state" yaml:"state"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *Node) Cluster() *Cluster {
	clu := Cluster{}
	err := db.First(&clu, r.ClusterID).Error
	if err != nil {
		panic(err)
	}
	return &clu
}

func (r *Node) Service() *Service {
	srv := Service{}
	err := db.First(&srv, r.ServiceID).Error
	if err != nil {
		panic(err)
	}
	return &srv
}

func (r *Node) Group() *Group {
	g := Group{}
	err := db.First(&g, r.GroupID).Error
	if err != nil {
		panic(err)
	}
	return &g
}
