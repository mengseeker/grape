package models

import (
	"time"
)

// User is an object representing the database table.
type User struct {
	// record
	ID        int64  `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"not null;index" json:"username"`
	Password  string `gorm:"not null;" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
