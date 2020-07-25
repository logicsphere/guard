package config

import (
  "github.com/aws/aws-sdk-go/aws/session"
)

type EnvironemtType string

const (
  AwsEnv EnvironemtType = "AWS"
  GcpEnv EnvironemtType = "GCP"
  K8sEnv EnvironemtType = "K8S"
  StdEnv EnvironemtType = "Standard"
)

func DetectEnvironment() EnvironemtType {

  _, err := session.NewSession()

  if err != nil {
    return AwsEnv
  }

  return StdEnv
}
