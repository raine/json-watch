package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructChecksum(t *testing.T) {
	str := `
{
  "a": 1,
  "b": {
    "a": [
      "b",
      "c"
    ],
    "b": 1
  }
}`

	var parsed interface{}
	json.Unmarshal([]byte(str), &parsed)
	expected, _ := calcStructChecksum(parsed)
	actual := "86014b6f1e4d01600cdf98794fe5a289da9cda48"
	assert.Equal(t, expected, actual)
}

func TestStableStructChecksum(t *testing.T) {
	str1 := `
{
  "b": {
    "b": 1,
    "a": [
      "b",
      "c"
    ]
  },
  "a": 1
}`

	str2 := `
{
  "a": 1,
  "b": {
    "a": [
      "b",
      "c"
    ],
    "b": 1
  }
}`

	var parsed1 interface{}
	var parsed2 interface{}
	json.Unmarshal([]byte(str1), &parsed1)
	json.Unmarshal([]byte(str2), &parsed2)
	checksum1, _ := calcStructChecksum(parsed1)
	checksum2, _ := calcStructChecksum(parsed2)
	assert.Equal(t, checksum1, checksum2)
}
