package bl

import (
	"grape/grape/models"
	"grape/grape/models/etlinks"
	"grape/pkg/resr"
)

func SyncServiceConf(r *models.Service, del bool) {
	for _, clu := range GetServiceClusters(r) {
		syncConf(&clu, resr.NewService(r), del)
	}
}

func SyncGroupConf(r *models.Group, del bool) {
	syncConf(r.Cluster(), resr.NewGroup(r), del)
}

func SyncPolicyConf(r *models.Policy, del bool) {
	svc := r.Service()
	for _, clu := range GetServiceClusters(svc) {
		syncConf(&clu, resr.NewPolicy(r), del)
	}
}

func syncConf(clu *models.Cluster, r resr.Res, del bool) {
	if del {
		PanicErr(resr.Delete(etlinks.GetCli(clu), r))
	} else {
		PanicErr(resr.Update(etlinks.GetCli(clu), r))
	}
}
