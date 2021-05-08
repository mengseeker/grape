package models

import "time"

// Node is an object representing the database table.
type Node struct {
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
