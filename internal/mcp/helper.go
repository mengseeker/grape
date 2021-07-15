package mcp

import "grape/grapeapi/models"

func GetServiceIDsFromGroups(gs []models.Group) []int64 {
	m := make(map[int64]bool)
	for _, g := range gs {
		m[g.ServiceID] = true
	}
	ids := []int64{}
	for id := range m {
		ids = append(ids, id)
	}
	return ids
}
