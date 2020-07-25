package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestDetectEnvironment(t *testing.T) {
  got := DetectEnvironment()
  expected := StdEnv
  assert.Equal(t, got, expected)
}
