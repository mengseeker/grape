package api

import "grape/grape/pkg/postgresdb"

var DB = postgresdb.GetDB
