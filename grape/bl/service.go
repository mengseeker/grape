package bl

import (
	"grape/grape/models"
)

func CreateService(name, code string, port, protocol, external int, note string) models.Service {
	record := models.Service{
		Name:     name,
		Code:     code,
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
