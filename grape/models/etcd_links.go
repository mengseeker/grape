package models

import "time"


// EtcdLink is an object representing the database table.
type EtcdLink struct {
	// record
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"index;unique;not null;" json:"name"`
	Address   string `gorm:"not null;" json:"address"`
	Note      string `gorm:"not null;default:'';" json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
