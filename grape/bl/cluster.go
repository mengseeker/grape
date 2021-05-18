package bl

import (
	"grape/grape/models"
)

func CreateCluster(name, code, note, deployType string, etcdLinkID int64) models.Cluster {
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

func ListClusters() ([]models.Cluster, *SearchInfo) {
	var total int64
	var list []models.Cluster
	order := "id desc"
	PanicErr(db().Model(&list).Count(&total).Order(order).Find(&list).Error)
	return list, &SearchInfo{
		Total: total,
		Order: order,
	}
}
