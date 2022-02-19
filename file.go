package main

import (
	"bufio"
	"os"
)

func readLinesMap(file *os.File) (map[string]bool, error) {
	lines := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines[scanner.Text()] = true
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	} else {
		return lines, nil
	}
}
