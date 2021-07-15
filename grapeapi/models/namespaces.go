/*
命名空间为逻辑概念，一个集群划分多个命名空间，多个集群相同的命名空间合并抽象为一个命名空间。
*/
package models

import "time"

// Namespace is an object representing the database table.
type Namespace struct {
	// record
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"index;unique;not null;" json:"name"`
	Note      string `gorm:"not null;default:'';" json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time

	F_Groups  []Group  `gorm:"foreignKey:NamespaceID" json:"-"`
	F_Nodes   []Node   `gorm:"foreignKey:NamespaceID" json:"-"`
}
