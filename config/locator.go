package config

import "net/http"

type EnvironemtType string

const (
	AwsEnv EnvironemtType = "AWS"
	GcpEnv EnvironemtType = "GCP"
	K8sEnv EnvironemtType = "K8S"
	StdEnv EnvironemtType = "Standard"
)

func DetectEnvironment() EnvironemtType {

	if checkAwsEnvironment() {
		return AwsEnv
	}

	return StdEnv
}

func checkAwsEnvironment() bool {
	_, err := http.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	return err == nil
}
