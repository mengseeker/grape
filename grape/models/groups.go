package models

import "time"

// Group is an object representing the database table.
type Group struct {
	// record
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"index;unique;not null;" json:"name"`
	Code        string `gorm:"index;unique;not null;" json:"code"`
	Version     string `gorm:"not null;default:'';" json:"version"`
	Lang        string `gorm:"not null;default:'';" json:"lang"`
	NamespaceID int    `gorm:"index;not null;" json:"namespace_id"`
	ServiceID   int    `gorm:"index;not null;" json:"service_id"`
	ClusterID   int    `gorm:"index;not null;" json:"cluster_id"`
	Replicas    int    `gorm:"not null;default:1" json:"replicas"`
	DeployType  int    `gorm:"not null;" json:"deploy_type"`
	Note        string `gorm:"not null;default:'';" json:"note"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
