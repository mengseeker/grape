package models

import "time"

// Cluster is an object representing the database table.
type Cluster struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"index;unique;not null;" json:"name"`
	Code       string `gorm:"index;unique;not null;" json:"code"`
	DeployType string `gorm:"not null;" json:"deploy_type"`
	EtcdID     int    `gorm:"not null;" json:"etcd_id"`
	Note       string `gorm:"not null;default:''" json:"note"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}