package bl

import "grape/grapeapi/models"

func CreateNs(name, code, note string, clusterID int64) models.Namespace {
	record := models.Namespace{
		Name: name,
		Note: note,
	}
	PanicErr(db().Create(&record).Error)
	return record
}
