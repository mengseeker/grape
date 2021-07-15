package mcp

import (
	"grape/api"
	"grape/grapeapi/models"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func buildServiceResource(records []models.Service, groups []models.Group, policies []models.Policy) Resource {
	var msgs []proto.Message
	groupList := map[int64][]string{}
	policyList := map[int64][]string{}
	for _, r := range groups {
		if groupList[r.ServiceID] == nil {
			groupList[r.ServiceID] = []string{r.Name}
		} else {
			groupList[r.ServiceID] = append(groupList[r.ServiceID], r.Name)
		}
	}
	for _, r := range policies {
		if policyList[r.ServiceID] == nil {
			policyList[r.ServiceID] = []string{r.Name}
		} else {
			policyList[r.ServiceID] = append(groupList[r.ServiceID], r.Name)
		}
	}
	for _, r := range records {
		m := api.Service{
			Name:            r.Name,
			Namespace:       r.Namespace().Name,
			ServiceProtocol: api.Service_Protocol(r.Protocol),
			DeployGroup:     groupList[r.ID],
			Policy:          policyList[r.ID],
		}
		msgs = append(msgs, &m)
	}
	return Resource{
		Version:   time.Now().String(),
		Resources: msgs,
	}
}

func buildPolicyResource(records []models.Policy) Resource {
	var msgs []proto.Message
	for _, r := range records {
		m := api.Policy{
			Name:       r.Name,
			PolicyType: api.Policy_Type(r.PolicyType),
			Policy:     buildPolicyMsg(r),
		}
		msgs = append(msgs, &m)
	}
	return Resource{
		Version:   time.Now().String(),
		Resources: msgs,
	}
}

func buildGroupResource(records []models.Group, nodes []models.Node) Resource {
	var msgs []proto.Message
	nodeList := map[int64][]*api.Node{}
	for _, r := range nodes {
		if nodeList[r.GroupID] == nil {
			nodeList[r.ServiceID] = []*api.Node{}
		} else {
			node := &api.Node{Ip: r.IP}
			nodeList[r.ServiceID] = append(nodeList[r.ServiceID], node)
		}
	}
	for _, r := range records {
		m := api.DeployGroup{
			Name:  r.Name,
			Nodes: nodeList[r.ID],
		}
		msgs = append(msgs, &m)
	}
	return Resource{
		Version:   time.Now().String(),
		Resources: msgs,
	}
}

func buildPolicyMsg(p models.Policy) *anypb.Any {
	var msg proto.Message
	switch p.PolicyType {
	case int(api.Policy_HealthCheck):
		msg = &api.HealthCheckPolicy{
			Path: p.OptionsMap()["path"].(string),
		}
	}
	any, err := anypb.New(msg)
	if err != nil {
		panic(err)
	}
	return any
}
