/*
集群划分多个命名空间，和命名空间多对多关系，由部署组关联。
注意不同集群必须名字不同。
*/
package models

import (
	"errors"
	"grape/pkg/arrays"
	"time"

	"gorm.io/gorm"
)

const (
	ClusterTypeK8s = "k8s"
	ClusterTypeVm  = "vm"
)

var (
	AllTypes = [...]string{ClusterTypeK8s, ClusterTypeVm}
)

var (
	ErrClusterAttrDeployTypeInvalid = errors.New("ClusterAttrDeployTypeInvalid")
)

// Cluster is an object representing the database table.
type Cluster struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"index;unique;not null;" json:"name"`
	ClusterType string `gorm:"not null;" json:"deploy_type"`
	Note        string `gorm:"not null;default:''" json:"note"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *Cluster) BeferSave(*gorm.DB) error {
	if !arrays.ContainsStr(AllTypes[:], r.ClusterType) {
		return ErrClusterAttrDeployTypeInvalid
	}
	return nil
}
