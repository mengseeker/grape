package models

import "time"

// Namespace is an object representing the database table.
type Namespace struct {
	// record
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"index;unique;not null;" json:"name"`
	Code      string `gorm:"index;unique;not null;" json:"code"`
	ClusterID int64  `gorm:"index;not null;" json:"cluster_id"`
	Note      string `gorm:"not null;default:'';" json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time

	F_Cluster *Cluster `gorm:"foreignKey:ClusterID" json:"-"`
	F_Groups  []Group  `gorm:"foreignKey:NamespaceID" json:"-"`
	F_Nodes   []Node   `gorm:"foreignKey:NamespaceID" json:"-"`
}
