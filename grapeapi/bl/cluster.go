package bl

import (
	"grape/grapeapi/models"
)

func CreateCluster(name, clusterType, note string) models.Cluster {
	record := models.Cluster{
		Name:        name,
		ClusterType: clusterType,
		Note:        note,
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
