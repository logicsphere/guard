package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDetectEnvironment(t *testing.T) {
	got := DetectEnvironment()

	var expected EnvironemtType

	switch os.Getenv("TEST_RUNTIME") {
	case "AWS":
		expected = AwsEnv
	default:
		expected = StdEnv
	}

	assert.Equal(t, expected, got)
}
