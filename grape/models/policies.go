package models

import "time"

// Policy is an object representing the database table.
type Policy struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Category  int    `gorm:"index;not null;" json:"category"`
	ServiceID int    `gorm:"index;not null;" json:"service_id"`
	Active    int    `gorm:"index;not null;default:1" json:"active"`
	Options   string `gorm:"not null;default:'{}'" json:"options"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
