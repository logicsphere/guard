package test

import (
  "github.com/logicsphere/guard/adapter"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSoma(t *testing.T) {
  got := adapter.Soma(1, 4)
  expected := 5

  assert.Equal(t, expected, got, "Os valores estão diferentes")
}
