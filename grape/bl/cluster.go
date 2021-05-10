package bl

import (
	"grape/grape/models"
)

func CreateCluster(name, code, note, deployType string, etcdLinkID int) models.Cluster {
	record := models.Cluster{
		Name:       name,
		Code:       code,
		DeployType: deployType,
		EtcdID:     etcdLinkID,
		Note:       note,
	}
	PanicErr(db().Create(&record).Error)
	return record
}
