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

	if checkGcpEnvironment() {
		return GcpEnv
	}

	return StdEnv
}

func checkAwsEnvironment() bool {
	res, err := http.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	return err == nil && res.StatusCode == 200
}

func checkGcpEnvironment() bool {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/id", nil)
	req.Header.Set("Metadata-Flavor", "Google")
	res, err := client.Do(req)
	return err == nil && res.StatusCode == 200
}
