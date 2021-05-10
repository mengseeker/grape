package models

import "time"

// Namespace is an object representing the database table.
type Namespace struct {
	// record
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"index;unique;not null;" json:"name"`
	Code      string `gorm:"index;unique;not null;" json:"code"`
	Note      string `gorm:"not null;default:'';" json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
