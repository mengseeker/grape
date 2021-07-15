package bl

import "grape/grapeapi/models"

func CreateGroup(
	name, code, version, lang, note string,
	namespace_id, service_id, cluster_id int64,
) models.Group {
	record := models.Group{
		Name:        name,
		NamespaceID: namespace_id,
		ServiceID:   service_id,
		ClusterID:   cluster_id,
		Lang:        lang,
		Version:     version,
		Note:        note,
	}
	PanicErr(db().Create(&record).Error)
	return record
}
