package bl

import "grape/grape/models"

func CreateNs(name, code, note, deployType string, etcdLinkID int) models.Namespace {
	record := models.Namespace{
		Name: name,
		Code: code,
		Note: note,
	}
	PanicErr(db().Create(&record).Error)
	return record
}
