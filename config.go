package main

import (
	"os"
	"path/filepath"
)

var myConfigDir string

func init() {
	myConfigDir = getMyConfigDir()
	os.MkdirAll(filepath.Join(myConfigDir, "watches"), os.ModePerm)
}

func getMyConfigDir() string {
	// Allow overriding the config directory for tests
	if configDirOverride, ok := os.LookupEnv("JSON_WATCH_CONFIG_DIR"); ok {
		return configDirOverride
	}

	home, err := os.UserHomeDir()
	if err != nil {
		dieWithError(err)
	}
	return filepath.Join(home, ".config", "json-watch")
}

func formatWatchPath(name string) string {
	return filepath.Join(myConfigDir, "watches", name)
}
