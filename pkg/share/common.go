package share

import "strings"

const ViperEnvPrefix = "GRAPE"
const EnvoyAccessLogPrefix = "envoy_access."

const resourceServicePrefix = "/GRAPE/:clusterCode:/SERVICE/"
const resourceGroupPrefix = "/GRAPE/:clusterCode:/GROUP/"
const resourcePolicyPrefix = "/GRAPE/:clusterCode:/Policy/"

func ResourceServicePrefix(clusterCode string) string {
	return strings.ReplaceAll(resourceServicePrefix, ":clusterCode:", clusterCode)
}

func ResourceGroupPrefix(clusterCode string) string {
	return strings.ReplaceAll(resourceGroupPrefix, ":clusterCode:", clusterCode)
}

func ResourcePolicyPrefix(clusterCode string) string {
	return strings.ReplaceAll(resourcePolicyPrefix, ":clusterCode:", clusterCode)
}
