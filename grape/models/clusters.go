package models

import (
	"errors"
	"grape/pkg/arrays"
	"time"

	"gorm.io/gorm"
)

const (
	DeployTypeK8s = "k8s"
)

var (
	AllDeployTypes = [...]string{DeployTypeK8s}
)

var (
	ErrClusterAttrDeployTypeInvalid = errors.New("ClusterAttrDeployTypeInvalid")
)

// Cluster is an object representing the database table.
type Cluster struct {
	// record
	ID         int64  `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"index;unique;not null;" json:"name"`
	Code       string `gorm:"index;unique;not null;" json:"code"`
	DeployType string `gorm:"not null;" json:"deploy_type"`
	EtcdID     int64  `gorm:"not null;" json:"etcd_id"`
	Note       string `gorm:"not null;default:''" json:"note"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	F_Etcd       *EtcdLink   `gorm:"foreignKey:EtcdID" json:"-"`
	F_Namespaces []Namespace `gorm:"foreignKey:ClusterID" json:"-"`
}

func (r *Cluster) BeferSave(*gorm.DB) error {
	if !arrays.ContainsStr(AllDeployTypes[:], r.DeployType) {
		return ErrClusterAttrDeployTypeInvalid
	}
	return nil
}

func (r *Cluster) EctdLink() *EtcdLink {
	if r.F_Etcd != nil {
		return r.F_Etcd
	}
	PanicErr(db.Model(r).Association("F_Etcd").Find(&r.F_Etcd))
	return r.F_Etcd
}
