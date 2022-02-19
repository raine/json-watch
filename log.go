package main

import (
	"fmt"
	"os"
)

func logStderr(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func die(format string, a ...interface{}) {
	logStderr(format, a...)
	os.Exit(1)
}

func dieWithError(err error) {
	die("Error: %s\n", err.Error())
}
