package bl

import (
	"grape/grapeapi/models"
)

func CreateService(name string, port, protocol, external int, note string) models.Service {
	record := models.Service{
		Name:     name,
		Port:     port,
		Protocol: protocol,
		External: external,
		Note:     note,
	}
	PanicErr(db().Create(&record).Error)
	return record
}

func SearchService(limit, offset int, params gromParam) ([]models.Service, *SearchInfo) {
	var total int64
	var records []models.Service
	order := "id desc"
	err := db().Model(&records).Where(params).
		Count(&total).
		Limit(limit).Offset(offset).
		Order(order).
		Find(&records).Error
	PanicErr(err)
	searchInfo := &SearchInfo{
		Total: total,
		Order: order,
	}
	return records, searchInfo
}

func GetServiceClusters(svc *models.Service) []models.Cluster {
	var cls []models.Cluster
	idsq := db().Select("cluster_id").Where("service_id = ?", svc.ID).Table(models.TableNameGroup)
	PanicErr(db().Model(&cls).Where("id in (?)", idsq).Find(&cls).Error)
	return cls
}
