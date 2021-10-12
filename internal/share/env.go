package share

import "os"

const (
	EnvServiceCode = "Grape_ServiceCode"
	EnvNamespace   = "Grape_Namespace"
	EnvGroupCode   = "Grape_Group"
)

func GetServiceCode() string {
	return os.Getenv(EnvServiceCode)
}

func GetGroupCode() string {
	return os.Getenv(EnvGroupCode)
}

func GetNamespace() string {
	return os.Getenv(EnvNamespace)
}

func GetService() string {
	ns := GetNamespace()
	svc := GetServiceCode()
	if ns != "" && svc != "" {
		return ns + "/" + svc
	}
	return ""
}
