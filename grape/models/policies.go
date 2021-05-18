package models

import (
	"encoding/json"
	"time"
)

// Policy is an object representing the database table.
type Policy struct {
	// record
	ID        int64  `gorm:"primaryKey" json:"id"`
	Code      string `gorm:"index;unique;not null;" json:"code"`
	Category  int    `gorm:"index;not null;" json:"category"`
	ServiceID int64  `gorm:"index;not null;" json:"service_id"`
	Active    bool   `gorm:"index;not null;default:true" json:"active"`
	Options   string `gorm:"not null;default:'{}'" json:"options"`
	CreatedAt time.Time
	UpdatedAt time.Time

	F_Service *Service `gorm:"foreignKey:ServiceID" json:"-"`

	options map[string]interface{}
}

func (r *Policy) Service() *Service {
	if r.F_Service != nil {
		return r.F_Service
	}
	PanicErr(db.Model(r).Association("F_Group").Find(&r.F_Service))
	return r.F_Service
}

func (r *Policy) OptionsMap() map[string]interface{} {
	if r.options != nil {
		return r.options
	}
	err := json.Unmarshal([]byte(r.Options), &r.options)
	if err != nil {
		panic(err)
	}
	return r.options
}
