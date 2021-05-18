package bl

import (
	"grape/grape/models"
	"grape/grape/models/etlinks"
	"grape/pkg/resr"
)

func SyncServiceConf(r *models.Service, del bool) {
	groups := r.Groups()
	ids := []int64{}
	for _, g := range groups {
		ids = append(ids, g.ID)
	}
	PanicErr(db().Model(groups).Preload("F_Nodes").Find(&groups, ids).Error)

	ps := r.Policies()
	rs := resr.NewService(r)
	for _, clu := range GetServiceClusters(r) {
		syncConf(&clu, rs, del)
		syncGroupConf(&clu, rs, groups, del)
		syncPolicyConf(&clu, rs, ps, del)
	}
}

func syncGroupConf(clu *models.Cluster, svc *resr.Service, rs []models.Group, del bool) {
	for _, r := range rs {
		syncConf(clu, resr.NewGroup(svc, &r), del)
	}
}

func syncPolicyConf(clu *models.Cluster, svc *resr.Service, rs []models.Policy, del bool) {
	for _, r := range rs {
		syncConf(clu, resr.NewPolicy(svc, &r), del)
	}
}

func syncConf(clu *models.Cluster, r resr.Res, del bool) {
	if del {
		PanicErr(resr.Delete(etlinks.GetCli(clu), r))
	} else {
		PanicErr(resr.Update(etlinks.GetCli(clu), r))
	}
}
