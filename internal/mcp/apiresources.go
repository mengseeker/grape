package mcp

import (
	"grape/api"
	"grape/grapeapi/models"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type ResourceHash string

func Hash(cluster string, rt api.MCPResourceType) ResourceHash {
	return ResourceHash(rt.String() + "::" + cluster)
}

func HashBack(hash ResourceHash) (cluster string, rt api.MCPResourceType) {
	strs := strings.SplitN(string(hash), "::", 2)
	cluster = strs[1]
	it, _ := strconv.Atoi(strs[0])
	rt = api.MCPResourceType(it)
	return
}

type Resource struct {
	Version   string
	Resources []proto.Message
}

func (r Resource) toAny() []*anypb.Any {
	var ansy []*anypb.Any
	for _, r := range r.Resources {
		anr, err := anypb.New(r)
		if err != nil {
			println("================ resource to anypb err =========================", err)
		}
		ansy = append(ansy, anr)
	}
	return ansy
}

type ApiResources struct {
	Resources      map[ResourceHash]Resource
	recordVersions map[ResourceHash]string
}

func NewApiResources() *ApiResources {
	return &ApiResources{
		Resources:      make(map[ResourceHash]Resource),
		recordVersions: make(map[ResourceHash]string),
	}
}

func (r *ApiResources) GetResponse(hash ResourceHash) *api.DiscoveryResponse {
	rts := strings.Split(string(hash), "::")[0]
	typeID, err := strconv.Atoi(rts)
	if err != nil {
		println("=================== resource hash to resourceType fail =================", err)
	}
	res := r.Resources[hash]
	return &api.DiscoveryResponse{
		VersionInfo:  res.Version,
		ResourceType: api.MCPResourceType(typeID),
		Resource:     res.toAny(),
	}
}

func (r *ApiResources) GetVersion(hash ResourceHash) string {
	return r.Resources[hash].Version
}

// db paniced
func (r *ApiResources) CheckUpdate() error {
	var err error
	var clusters []models.Cluster
	err = models.GetDB().Model(&clusters).Order("id desc").Find(&clusters).Error
	if err != nil {
		return err
	}
	for _, cluster := range clusters {
		err = r.updateClusterResources(cluster)
		if err != nil {
			return err
		}
	}
	return r.removeOutofClusterResources(clusters)
}

func (r *ApiResources) LoadAll() error {
	return r.CheckUpdate()
}

func (r *ApiResources) updateClusterResources(cluster models.Cluster) error {
	var err error
	var groups []models.Group
	var clusterServices []models.Service
	var policies []models.Policy
	var nodes []models.Node

	err = models.GetDB().Model(&groups).Where("cluster_id = ?", cluster.ID).Order("id desc").Find(&groups).Error
	if err != nil {
		return err
	}

	groupServiceIDs := GetServiceIDsFromGroups(groups)
	err = models.GetDB().Model(&clusterServices).Where("id in (?)", groupServiceIDs).Order("id desc").Find(&clusterServices).Error
	if err != nil {
		return err
	}

	err = models.GetDB().Model(&policies).Where("service_id in (?) and active = 1", groupServiceIDs).Order("id desc").Find(&policies).Error
	if err != nil {
		return err
	}

	groupIDs := []int64{}
	for _, g := range groups {
		groupIDs = append(groupIDs, g.ID)
	}
	err = models.GetDB().Model(&nodes).Where("group_id in (?)", groupIDs).Order("id desc").Find(&nodes).Error
	if err != nil {
		return err
	}

	r.updateServiceResource(cluster, clusterServices, groups, policies)
	r.updatePolicyResource(cluster, policies)
	r.updateGroupResource(cluster, groups, nodes)
	return nil
}

func (r *ApiResources) removeOutofClusterResources(clusters []models.Cluster) error {
	clusterNames := map[string]bool{}
	for _, clu := range clusters {
		clusterNames[clu.Name] = true
	}
	for hash := range r.Resources {
		resourceClu, _ := HashBack(hash)
		if !clusterNames[resourceClu] {
			delete(r.Resources, hash)
			delete(r.recordVersions, hash)
		}
	}
	return nil
}

func (r *ApiResources) updateServiceResource(cluster models.Cluster, records []models.Service, groups []models.Group, policies []models.Policy) {
	hash := Hash(cluster.Name, api.MCPResourceType_SERVICE)
	var recordVersion string
	for _, rec := range records {
		recordVersion = recordVersion + rec.UpdatedAt.String()
	}
	if r.recordVersions[hash] != recordVersion {
		r.Resources[hash] = buildServiceResource(records, groups, policies)
		r.recordVersions[hash] = recordVersion
	}
}

func (r *ApiResources) updatePolicyResource(cluster models.Cluster, records []models.Policy) {
	hash := Hash(cluster.Name, api.MCPResourceType_POLICY)
	var recordVersion string
	for _, rec := range records {
		recordVersion = recordVersion + rec.UpdatedAt.String()
	}
	if r.recordVersions[hash] != recordVersion {
		r.Resources[hash] = buildPolicyResource(records)
		r.recordVersions[hash] = recordVersion
	}
}

func (r *ApiResources) updateGroupResource(cluster models.Cluster, records []models.Group, nodes []models.Node) {
	hash := Hash(cluster.Name, api.MCPResourceType_DEPLOY_GROUP)
	var recordVersion string
	for _, rec := range records {
		recordVersion = recordVersion + rec.UpdatedAt.String()
	}
	for _, rec := range nodes {
		recordVersion = recordVersion + rec.UpdatedAt.String()
	}
	if r.recordVersions[hash] != recordVersion {
		r.Resources[hash] = buildGroupResource(records, nodes)
		r.recordVersions[hash] = recordVersion
	}
}
