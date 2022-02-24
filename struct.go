package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

// Calculates checksum for a struct in a manner that order of fields does not
// matter. It's used to calculate checksums for JSON objects.
func calcStructChecksum(obj interface{}) (string, error) {
	// json.Marshal sorts keys lexicographically
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return toSha1(bytes), nil
}

func toSha1(bytes []byte) string {
	h := sha1.New()
	h.Write(bytes)
	return fmt.Sprintf("%x", h.Sum(nil))
}
