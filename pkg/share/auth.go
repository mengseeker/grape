package share

import "strings"

const AuthAppDB = "/GRAPE/:clusterCode:/AUTH/APP/"
const AuthTokenDB = "/GRAPE/:clusterCode:/AUTH/TOKEN/"

func AuthAppDBKey(clusterCode string) string {
	return strings.ReplaceAll(AuthAppDB, ":clusterCode:", clusterCode)
}

func AuthTokenDBKey(clusterCode string) string {
	return strings.ReplaceAll(AuthTokenDB, ":clusterCode:", clusterCode)
}
