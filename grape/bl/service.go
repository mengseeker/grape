package bl

import (
	"grape/grape/models"
)

func CreateService(name, code string, port, protocol, external int, note string) models.Service {
	if external > 0 {
		external = 1
	}
	service := models.Service{
		Name:     name,
		Code:     code,
		Port:     port,
		Protocol: protocol,
		External: external,
		Note:     note,
	}
	err := db().Create(&service).Error
	PanicErr(err)
	return service
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
